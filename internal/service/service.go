package service

import (
	"fmt"
	"log"

	"github.com/engelsjk/faadb/internal/db"
	"github.com/engelsjk/faadb/internal/parser"
	"github.com/engelsjk/faadb/internal/utils"
)

type Settings struct {
	Name      string
	DataPath  string
	DBPath    string
	NumFields int
	Reload    bool
}

type Service struct {
	Name string
	db   *db.DB
}

func NewService(settings Settings, decoder func([]string) (string, string, error)) (*Service, error) {

	s := &Service{
		Name: settings.Name,
	}

	log.Printf("%s : new service\n", s.Name)
	log.Printf("%s : initializing db\n", s.Name)

	dbExists := db.Exists(settings.DBPath)

	var err error
	s.db, err = db.InitDB(settings.DBPath)
	if err != nil {
		return nil, err
	}

	if dbExists && !settings.Reload {
		log.Printf("%s : db already exists, skip loading\n", s.Name)
		return s, nil
	}

	log.Printf("%s : parsing data file\n", s.Name)

	p := parser.NewParser(settings.DataPath, settings.NumFields, decoder)

	log.Printf("%s : loading data to db\n", s.Name)

	if err := s.loadToDB(p); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Service) CreateIndexJSON(index, pattern string, path ...string) error {
	log.Printf("%s : creating index '%s'\n", s.Name, index)
	return s.db.CreateIndexJSON(index, pattern, path...)
}

func (s *Service) Get(key string) ([][]byte, error) {
	key = utils.ToUpper(key)
	return s.db.Get(key)
}

func (s *Service) List(index, match, pattern string, exact bool) ([][]byte, error) {
	match = utils.ToUpper(match)
	if exact {
		return s.db.ListExact(index, match, pattern)
	}
	return s.db.ListWildcard(index, match, pattern)
}

func (s *Service) StartsWith(index, starts_with string) ([][]byte, error) {
	starts_with = utils.ToUpper(starts_with)
	return s.db.StartsWith(index, starts_with)
}

type Record interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}

func (s *Service) UnmarshalRecords(bs [][]byte, rs interface{}) error {
	var err error
	records, ok := rs.([]*Record)
	if !ok {
		return fmt.Errorf("not records")
	}
	if len(bs) != len(records) {
		return fmt.Errorf("length mismatch")
	}
	for i, b := range bs {
		r := records[i]
		err = (*r).UnmarshalJSON(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) loadToDB(parser parser.Parser) error {
	count, err := parser.LoadLinesToDB(s.db)
	if err != nil {
		return err
	}
	log.Printf("%s : num. lines loaded to db: %d\n", s.Name, count)
	return nil
}
