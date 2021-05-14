package util

import "fmt"

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
