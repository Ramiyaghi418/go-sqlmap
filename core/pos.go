package core

import "sort"

type Pos struct {
	Key          int
	StartIndex   int
	EndIndexChar string
}

// GetMinPos 获取位置中开始索引最小的
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
	if len(array) > 0 {
		return array[0]
	}
	return Pos{}
}
