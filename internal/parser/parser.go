package parser

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/engelsjk/faadb/internal/db"
)

type Parser struct {
	Name              string
	path              string
	expectedNumFields int
	delim             string
	decoder           func([]string) (string, string, error)
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

func (p *Parser) SetExpectedNumFields(n int) {
	p.expectedNumFields = n
}

func (p *Parser) LoadLinesToDB(db *db.DB) (int, error) {

	file, err := os.Open(p.path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
		return n, nil
	}
	return n, nil
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
