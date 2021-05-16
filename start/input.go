package start

import (
	"flag"
	"go-sqlmap/constant"
	"os"
	"strings"
)

type Input struct {
	Url       string
	Database  string
	Table     string
	Shell     bool
	Columns   []string
	Technique []string
}

// ParseInput 处理输入参数
func ParseInput() Input {
	var url string
	var database string
	var table string
	var columns string
	var help bool
	var technique string
	flag.StringVar(&url, "u", "", "Input Target Url")
	flag.StringVar(&database, "D", "", "Get All Databases")
	flag.StringVar(&table, "T", "", "Get All Tables")
	flag.StringVar(&columns, "C", "", "Get All Columns")
	flag.StringVar(&technique, "technique", "BTUE",
		"Set Technique(B:bool-blind,T:time-blind,U:union-select,E:error-based)")
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
	if strings.Contains(technique, constant.TimeBlindTech) {
		finalTech = append(finalTech, constant.TimeBlindTech)
	}
	if strings.Contains(technique, constant.BoolBlindTech) {
		finalTech = append(finalTech, constant.BoolBlindTech)
	}
	if strings.Contains(technique, constant.ErrorBasedTech) {
		finalTech = append(finalTech, constant.ErrorBasedTech)
	}
	if strings.Contains(technique, constant.UnionSelectTech) {
		finalTech = append(finalTech, constant.UnionSelectTech)
	}

	result := Input{
		Url:       url,
		Database:  database,
		Table:     table,
		Columns:   finalColumns,
		Technique: finalTech,
	}

	return result
}
