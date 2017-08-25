package db

// Interface is database creator interface
type Interface interface {
	Parse(c *Config) error
	SQL() []string
	Execute() error
}
