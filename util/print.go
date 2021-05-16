package util

import (
	"fmt"
)

const (
	Databases = 1
	Tables    = 2
	Columns   = 3
)

func printFunc(printType int, data []interface{}) {
	fmt.Println("|------------------------------|")
	if printType == Databases {
		fmt.Printf("|%-30s|\n", "All Databases")
	} else if printType == Tables {
		fmt.Printf("|%-30s|\n", "All Tables")
	} else if printType == Columns {
		fmt.Printf("|%-30s|\n", "All Columns")
	}
	fmt.Println("|------------------------------|")
	for _, v := range data {
		fmt.Printf("|%-30s|\n", v)
	}
	fmt.Println("|------------------------------|")
}

func PrintDatabases(databases []interface{}) {
	printFunc(Databases, databases)
}

func PrintTables(tables []interface{}) {
	printFunc(Tables, tables)
}

func PrintColumns(columns []interface{}) {
	printFunc(Columns, columns)
}

func PrintData(columns []interface{}, data [][]interface{}) {
	fmt.Print("|")
	for i := 0; i < len(columns); i++ {
		if i == 0 {
			fmt.Print("---------------")
		} else {
			fmt.Print("----------------")
		}
	}
	fmt.Print("|\n")
	formatStr := "|"
	for i := 0; i < len(columns); i++ {
		formatStr += "%-15s|"
	}
	fmt.Printf(formatStr+"\n", columns...)
	fmt.Print("|")
	for i := 0; i < len(columns); i++ {
		if i == 0 {
			fmt.Print("---------------")
		} else {
			fmt.Print("----------------")
		}
	}
	fmt.Print("|\n")
	for _, v := range data {
		fmt.Printf(formatStr+"\n", v...)
	}
	fmt.Print("|")
	for i := 0; i < len(columns); i++ {
		if i == 0 {
			fmt.Print("---------------")
		} else {
			fmt.Print("----------------")
		}
	}
	fmt.Print("|\n")
}
