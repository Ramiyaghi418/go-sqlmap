package core

import (
	"flag"
	"os"
)

type Input struct {
	Url      string
	Database string
	Table    string
	Column   string
	Dump     bool
}

func ParseInput() Input {
	var url string
	var database string
	var table string
	var column string
	var help bool
	var dump bool
	flag.StringVar(&url, "u", "", "Input Target Url")
	flag.StringVar(&database, "D", "", "Get All Databases")
	flag.StringVar(&table, "T", "", "Get All Tables")
	flag.StringVar(&column, "C", "", "Get All Columns")
	flag.BoolVar(&dump, "dump", false, "Get All Data")
	flag.BoolVar(&help, "h", false, "Help Information")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	result := Input{
		Url:      url,
		Dump:     dump,
		Database: database,
		Table:    table,
		Column:   column,
	}

	return result
}
