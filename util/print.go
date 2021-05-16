package util

import (
	"fmt"
)

// 三种打印
const (
	Databases = 1
	Tables    = 2
	Columns   = 3
)

// 打印提取函数
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

// PrintDatabases 打印数据库
func PrintDatabases(databases []interface{}) {
	printFunc(Databases, databases)
}

// PrintTables 打印所有表
func PrintTables(tables []interface{}) {
	printFunc(Tables, tables)
}

// PrintColumns 打印所有列
func PrintColumns(columns []interface{}) {
	printFunc(Columns, columns)
}

// PrintData 打印所有数据
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
