package core

import "sort"

type Pos struct {
	Key          int
	StartIndex   int
	EndIndexChar string
}

func GetMinPos(array []Pos) Pos {
	var valueArray []int
	var min int
	for _, v := range array {
		valueArray = append(valueArray, v.StartIndex)
	}
	sort.Ints(valueArray)
	if len(valueArray) > 0 {
		min = valueArray[0]
	}
	for _, v := range array {
		if v.StartIndex == min {
			return v
		}
	}
	return array[0]
}
