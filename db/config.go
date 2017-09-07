package db

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var readFileFn = ioutil.ReadFile

// Config is configuration
type Config struct {
	Database string
	Tables   []Table
	Indices  []Index
	Data     []Datum
}

// Index is database index
type Index struct {
	Name    string
	Table   string
	Unique  bool `yaml:",omitempty"`
	Columns []string
}

// Table is database table
type Table struct {
	Name   string
	Fields []Field
}

// Field is field
type Field struct {
	Name          string
	Type          string
	Nullable      bool `yaml:",omitempty"`
	PrimaryKey    bool `yaml:"primary_key,omitempty"`
	Autoincrement bool `yaml:",omitempty"`
}

// Datum is preload data
type Datum struct {
	Table  string
	Fields []string
	Rows   []Row
}

// Row is data row
type Row []string

// NewConfig to load config
func NewConfig(f string) (*Config, error) {
	conf := Config{}
	confData, e := readFileFn(f)
	if e != nil {
		return nil, errors.New("cannot load configuration file")
	}
	e = yaml.Unmarshal(confData, &conf)
	if e != nil {
		return nil, errors.New("cannot parse configuration file")
	}
	return &conf, nil
}
