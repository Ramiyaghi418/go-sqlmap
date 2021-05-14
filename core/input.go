package core

import (
	"flag"
	"os"
	"strings"
)

type Input struct {
	Url      string
	Database string
	Table    string
	Columns  []string
	Dump     bool
}

func ParseInput() Input {
	var url string
	var database string
	var table string
	var columns string
	var help bool
	var dump bool
	flag.StringVar(&url, "u", "", "Input Target Url")
	flag.StringVar(&database, "D", "", "Get All Databases")
	flag.StringVar(&table, "T", "", "Get All Tables")
	flag.StringVar(&columns, "C", "", "Get All Columns")
	flag.BoolVar(&dump, "dump", false, "Get All Data")
	flag.BoolVar(&help, "h", false, "Help Information")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var finalColumns []string
	for _, v := range strings.Split(columns, ",") {
		if strings.TrimSpace(v) != "" {
			finalColumns = append(finalColumns, v)
		}
	}

	result := Input{
		Url:      url,
		Dump:     dump,
		Database: database,
		Table:    table,
		Columns:  finalColumns,
	}

	return result
}
