package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/engelsjk/faadb/internal/db"
	"github.com/engelsjk/faadb/internal/ds3"
)

type Parser struct {
	Name              string
	Source            string
	path              string
	expectedNumFields int
	delim             string
	decoder           func([]string) (string, string, error)
	r                 io.ReadCloser
}

func NewParser(path string, n int, decoder func([]string) (string, string, error)) Parser {
	return Parser{
		Name:              "parser",
		path:              path,
		expectedNumFields: n,
		delim:             ",",
		decoder:           decoder,
	}
}

func (p *Parser) SetReaderFromPath() error {

	// no path provided
	if p.path == "" {
		p.Source = "none"
		return fmt.Errorf("no data path provided")
	}

	// s3
	_, _, err := ds3.ParsePath(p.path)
	if err == nil {
		d, err := ds3.Init(p.path)
		if err != nil {
			return err
		}
		if ok := d.BucketKeyExists(); !ok {
			return fmt.Errorf("bucket/key does not exist")
		}
		r, err := d.Reader()
		if err != nil {
			return err
		}
		p.r = r
		p.Source = "s3"
		return nil
	}

	// on disk

	if _, err := os.Stat(p.path); os.IsNotExist(err) {
		return fmt.Errorf("data path does not exist")
	}

	file, err := os.Open(p.path)
	if err != nil {
		return err
	}

	p.r = file
	p.Source = "disk"
	return nil
}

func (p *Parser) SetExpectedNumFields(n int) {
	p.expectedNumFields = n
}

func (p *Parser) LoadLinesToDB(db *db.DB) (int, error) {

	scanner := bufio.NewScanner(p.r)

	n := 0
	var line string
	for scanner.Scan() {
		line = scanner.Text()

		if n == 0 { // ignore header row
			n++
			continue
		}

		fields := strings.Split(line, p.delim)

		TrimFields(fields)

		if ok := IsNumFieldsExpected(fields, p.ExpectedNumFields()); !ok {
			log.Printf("%s : expected %d fields, got %d...skipping line\n", p.Name, p.ExpectedNumFields(), NumFields(fields))
			continue
		}

		key, val, err := p.decoder(fields)
		if err != nil {
			log.Printf("%s : unable to marshal json of record...skipping line\n", p.Name)
			continue
		}

		key = strings.ToUpper(key)

		if err := db.Set(key, val); err != nil {
			log.Printf("%s : unable to set record key (%s)...skipping line\n", p.Name, key)
			continue
		}

		n++
	}

	if err := scanner.Err(); err != nil {
		p.r.Close()
		return n, nil
	}

	p.r.Close()
	return n, nil
}

func (p *Parser) Close() {
	if p.r != nil {
		p.r.Close()
	}
}

func (p *Parser) ExpectedNumFields() int {
	return p.expectedNumFields
}

func NumFields(fields []string) int {
	return len(fields)
}

func TrimFields(f []string) {
	for i := range f {
		f[i] = strings.TrimSpace(f[i])
	}
}

func IsNumFieldsExpected(fields []string, n int) bool {
	a := len(fields)
	if a != n {
		return false
	}
	return true
}
