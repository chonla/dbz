package db

import (
	"errors"
	"fmt"
	"strings"
)

// Sqlite uses sqlite engine
type Sqlite struct {
	sql []string
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
	return nil
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
	return strings.TrimSpace(fmt.Sprintf("%s %s %s %s", f.Name, f.Type, nullable, pk))
}

// SQL return generated sql statements
func (d *Sqlite) SQL() []string {
	return d.sql
}

// Execute to execute statements from SQL()
func (d *Sqlite) Execute() error {
	return nil
}
