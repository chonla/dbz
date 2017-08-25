package db

import (
	"regexp"
)

// Info is database engine info
type Info struct {
	Type     string
	Database string
}

var re = regexp.MustCompile("(?P<type>[a-zA-Z0-9]+):///(?P<database>.+)")

func parse(engine string) Info {
	matches := re.FindStringSubmatch(engine)
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = matches[i]
		}
	}

	return Info{
		Type:     result["type"],
		Database: result["database"],
	}
}
