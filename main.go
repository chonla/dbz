package main

import (
	"flag"
	"strings"

	"github.com/chonla/console"
	"github.com/chonla/dbz/db"
)

func main() {
	var confFile string
	var dumpOnly bool
	var overwrite bool
	flag.StringVar(&confFile, "conf", "", "configuration file")
	flag.BoolVar(&dumpOnly, "dump", false, "do not create database, just dump sql command")
	flag.BoolVar(&overwrite, "overwrite", false, "overwrite output")
	flag.Parse()

	dbz, err := db.NewDbz(confFile)
	if err != nil {
		console.Printfln("cannot initialize dbz: %v", err, console.ColorRed)
		return
	}
	if dumpOnly {
		console.Printfln("%v", strings.Join(dbz.SQL(), "\n"), console.ColorWhite)
	} else {
		e := dbz.Execute(overwrite)
		if e != nil {
			console.Println(e, console.ColorRed)
		} else {
			console.Printfln("done", console.ColorGreen)
		}
	}
}
