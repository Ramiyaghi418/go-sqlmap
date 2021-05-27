package start

import (
	"flag"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/str"
	"os"
	"strings"
)

type Input struct {
	Beta      bool
	Url       string
	Database  string
	Table     string
	Filename  string
	Columns   []string
	Technique []string
}

// ParseInput 处理输入参数
func ParseInput() Input {
	var url string
	var beta bool
	var database string
	var table string
	var columns string
	var help bool
	var technique string
	var filename string
	flag.BoolVar(&beta, "beta", false, "Use Beta Technique")
	flag.StringVar(&url, "u", "", "Input Target Url")
	flag.StringVar(&database, "D", "", "Get All Databases")
	flag.StringVar(&table, "T", "", "Get All Tables")
	flag.StringVar(&columns, "C", "", "Get All Columns")
	flag.StringVar(&technique, "technique", "BUE",
		"Set Technique(B:bool-blind,U:union-select,E:error-based)")
	flag.StringVar(&filename, "r", "", "Use Request Filename")
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

	var finalTech []string
	if strings.Contains(technique, str.BoolBlindTech) {
		finalTech = append(finalTech, str.BoolBlindTech)
	}
	if strings.Contains(technique, str.ErrorBasedTech) {
		finalTech = append(finalTech, str.ErrorBasedTech)
	}
	if strings.Contains(technique, str.UnionSelectTech) {
		finalTech = append(finalTech, str.UnionSelectTech)
	}

	if strings.TrimSpace(filename) != "" {
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				log.Info("request file not exist!")
				if url == "" {
					log.Error("need url or request file!")
					os.Exit(-1)
				}
			}
		}
	}

	result := Input{
		Beta:      beta,
		Url:       url,
		Database:  database,
		Table:     table,
		Columns:   finalColumns,
		Technique: finalTech,
		Filename:  filename,
	}

	return result
}
