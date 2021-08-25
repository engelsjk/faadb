package ds3

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type DS3 struct {
	Path   string
	bucket string
	key    string
	region string
	svc    *s3.S3
}

func Init(path string) (*DS3, error) {

	// todo:
	// * aws permissions config?
	// * default/specified aws region

	ds3 := &DS3{
		Path:   path,
		region: "us-east-1",
	}

	var err error
	ds3.bucket, ds3.key, err = ParsePath(path)
	if err != nil {
		return nil, err
	}

	err = ds3.connectToAWS()
	if err != nil {
		return nil, err
	}

	return ds3, nil
}

func (ds3 *DS3) connectToAWS() error {
	config := aws.Config{Region: aws.String(ds3.region)}
	sess := session.Must(session.NewSession(&config))
	if sess == nil {
		return fmt.Errorf("problems with connection to AWS")
	}
	ds3.svc = s3.New(sess)
	return nil
}

func ParsePath(path string) (string, string, error) {

	u, err := url.Parse(path)
	if err != nil {
		return "", "", fmt.Errorf("unable to parse s3 url path")
	}

	if u.Scheme != "s3" {
		return "", "", fmt.Errorf("path is not a valid s3 url")
	}

	u.Path = strings.Trim(u.Path, "/")

	bucket := u.Host
	key := u.Path

	return bucket, key, nil
}

func (ds3 *DS3) BucketKeyExists() bool {

	input := s3.HeadObjectInput{
		Bucket: aws.String(ds3.bucket),
		Key:    aws.String(ds3.key),
	}

	_, err := ds3.svc.HeadObject(&input)
	if err != nil {
		return false
	}

	return true
}

func (ds3 *DS3) Reader() (io.ReadCloser, error) {

	params := &s3.GetObjectInput{
		Bucket: aws.String(ds3.bucket),
		Key:    aws.String(ds3.key),
	}
	resp, err := ds3.svc.GetObject(params)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
