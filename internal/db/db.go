package db

import (
	"os"
	"strings"

	"github.com/tidwall/buntdb"
	"github.com/tidwall/gjson"
)

type DB struct {
	Path string
	db   *buntdb.DB
}

func InitDB(path string) (*DB, error) {
	if path == "" {
		path = ":memory:"
	}
	db, err := buntdb.Open(path)
	if err != nil {
		return nil, err
	}
	return &DB{db: db, Path: path}, nil
}

func Exists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (d *DB) CreateIndexString(index, pattern string) error {
	return d.db.CreateIndex(index, pattern, buntdb.IndexString)
}

func (d *DB) CreateIndexJSON(index, pattern string, paths ...string) error {
	less := make([]func(a string, b string) bool, len(paths))
	for i, p := range paths {
		less[i] = buntdb.IndexJSON(p)
	}
	return d.db.CreateIndex(index, pattern, less...)
}

func (d *DB) Set(key, value string) error {
	return d.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
}

func (d *DB) Get(key string) ([][]byte, error) {
	var vals [][]byte
	err := d.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		vals = append(vals, []byte(val))
		return nil
	})
	return vals, err
}

func (d *DB) ListExact(index, match, pattern string, filters map[string]string) ([][]byte, error) {
	var vals [][]byte
	err := d.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend(index, func(key, val string) bool {
			if match == gjson.Get(val, pattern).String() {
				if !passFilters(filters, val) {
					return true
				}
				vals = append(vals, []byte(val))
			}
			return true
		})
		return err
	})
	return vals, err
}

func (d *DB) ListWildcard(index, match, pattern string, filters map[string]string) ([][]byte, error) {
	var vals [][]byte
	err := d.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend(index, func(key, val string) bool {
			item := gjson.Get(val, pattern).String()
			if strings.Contains(item, match) {
				if !passFilters(filters, val) {
					return true
				}
				vals = append(vals, []byte(val))
			}
			return true
		})
		return err
	})
	return vals, err
}

func (d *DB) StartsWith(index, starts_with string) ([][]byte, error) {
	var vals [][]byte
	err := d.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend(index, func(key, val string) bool {
			keysplit := strings.Split(key, "_")
			if starts_with == keysplit[0] {
				vals = append(vals, []byte(val))
			}
			return true
		})
		return err
	})
	return vals, err
}

func passFilters(filters map[string]string, val string) bool {
	if filters == nil {
		return true
	}
	for k, v := range filters {
		if v == "" {
			continue
		}
		item := gjson.Get(val, k)
		// todo: parse listed items on delim ; then do something better than strings.Contains?
		if !strings.Contains(item.String(), v) {
			return false
		}
	}
	return true
}
