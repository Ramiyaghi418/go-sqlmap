package util

import (
	"fmt"
)

func PrintDatabases(databases []interface{}) {
	fmt.Println("|------------------------------|")
	fmt.Printf("|%-30s|\n", "All Databases")
	fmt.Println("|------------------------------|")
	for _, v := range databases {
		fmt.Printf("|%-30s|\n", v)
	}
	fmt.Println("|------------------------------|")
}

func PrintTables(tables []interface{}) {
	fmt.Println("|------------------------------|")
	fmt.Printf("|%-30s|\n", "All Tables")
	fmt.Println("|------------------------------|")
	for _, v := range tables {
		fmt.Printf("|%-30s|\n", v)
	}
	fmt.Println("|------------------------------|")
}

func PrintColumns(columns []interface{}) {
	fmt.Println("|------------------------------|")
	fmt.Printf("|%-30s|\n", "All Columns")
	fmt.Println("|------------------------------|")
	for _, v := range columns {
		fmt.Printf("|%-30s|\n", v)
	}
	fmt.Println("|------------------------------|")
}

func PrintData(columns []interface{}, data [][]interface{}) {
	fmt.Print("|")
	for i := 0; i < len(columns); i++ {
		if i == 0 {
			fmt.Print("----------")
		} else {
			fmt.Print("-----------")
		}
	}
	fmt.Print("|\n")
	fmt.Printf("|%-10s|%-10s|%-10s|\n", columns...)
	fmt.Print("|")
	for i := 0; i < len(columns); i++ {
		if i == 0 {
			fmt.Print("----------")
		} else {
			fmt.Print("-----------")
		}
	}
	fmt.Print("|\n")
	for _, v := range data {
		fmt.Printf("|%-10s|%-10s|%-10s|\n", v...)
	}
	fmt.Print("|")
	for i := 0; i < len(columns); i++ {
		if i == 0 {
			fmt.Print("----------")
		} else {
			fmt.Print("-----------")
		}
	}
	fmt.Print("|")
}
