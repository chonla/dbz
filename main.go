package main

import (
	"flag"
	"strings"

	"bitbucket.org/bayolab/dbz/db"
	"github.com/chonla/console"
)

func main() {
	var confFile string
	var dumpOnly bool
	flag.StringVar(&confFile, "conf", "", "configuration file")
	flag.BoolVar(&dumpOnly, "dump", false, "do not create database, just dump sql command")
	flag.Parse()

	dbz, err := db.NewDbz(confFile)
	if err != nil {
		console.Printfln("cannot initialize dbz: %v", err, console.ColorRed)
		return
	}
	if dumpOnly {
		console.Printfln("%v", strings.Join(dbz.SQL(), "\n"), console.ColorWhite)
	} else {
		dbz.Execute()
	}
}
