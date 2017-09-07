package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Sqlite uses sqlite engine
type Sqlite struct {
	sql []string
	db  string
}

// Parse to initialize Sqlite object
func (d *Sqlite) Parse(c *Config) error {
	d.sql = []string{}
	if len(c.Tables) == 0 {
		return errors.New("no table")
	}
	for i := range c.Tables {
		t := c.Tables[i]
		sql := createTable(t)
		d.sql = append(d.sql, sql)
	}
	for i := range c.Indices {
		t := c.Indices[i]
		sql := createIndex(t)
		d.sql = append(d.sql, sql)
	}
	for i := range c.Data {
		t := c.Data[i]
		for j := range t.Rows {
			sql := createData(t, t.Rows[j])
			d.sql = append(d.sql, sql)
		}
	}
	return nil
}

func createData(t Datum, row []string) string {
	fields := fmt.Sprintf("`%s`", strings.Join(t.Fields, "`, `"))
	values := fmt.Sprintf("\"%s\"", strings.Join(row, "\", \""))
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", t.Table, fields, values)
}

func createIndex(t Index) string {
	unique := ""
	if t.Unique {
		unique = " UNIQUE"
	}
	return fmt.Sprintf("CREATE%s INDEX %s ON %s (%s);", unique, t.Name, t.Table, strings.Join(t.Columns, ", "))
}

func createTable(t Table) string {
	fields := []string{}
	for i := range t.Fields {
		fields = append(fields, createField(t.Fields[i]))
	}
	return fmt.Sprintf("CREATE TABLE %s (%s);", t.Name, strings.Join(fields, ", "))
}

func createField(f Field) string {
	nullable := "NULL"
	if !f.Nullable {
		nullable = "NOT NULL"
	}
	pk := ""
	if f.PrimaryKey {
		pk = "PRIMARY KEY"
	}
	autoinc := ""
	if f.Autoincrement {
		autoinc = "AUTOINCREMENT"
	}
	return strings.TrimSpace(fmt.Sprintf("%s %s %s %s %s", f.Name, f.Type, nullable, pk, autoinc))
}

// SQL return generated sql statements
func (d *Sqlite) SQL() []string {
	return d.sql
}

// Execute to execute statements from SQL()
func (d *Sqlite) Execute(overwrite bool) error {
	if _, e := os.Stat(d.db); e == nil {
		if overwrite {
			os.Remove(d.db)
		} else {
			return errors.New("database is already exist. use -overwrite to overwrite the existing")
		}
	}
	s, e := sql.Open("sqlite3", d.db)
	if e != nil {
		return e
	}
	defer s.Close()

	tx, e := s.Begin()
	if e != nil {
		return e
	}
	for i := range d.sql {
		_, e := tx.Exec(d.sql[i])
		if e != nil {
			return e
		}
	}
	tx.Commit()

	return nil
}
