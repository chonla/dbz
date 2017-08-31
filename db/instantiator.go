package db

import "errors"

// NewDbz to create a new database instance
func NewDbz(f string) (Interface, error) {
	c, e := NewConfig(f)
	if e != nil {
		return nil, e
	}

	instance, e := createInstance(c)
	if e != nil {
		return nil, e
	}

	return instance, nil
}

func createInstance(c *Config) (Interface, error) {
	info := parse(c.Database)

	switch info.Type {
	case "sqlite":
		o := Sqlite{
			db: info.Database,
		}
		o.Parse(c)
		return &o, nil
	}
	return nil, errors.New("unsupport database driver or invalid dsn declaration")
}
