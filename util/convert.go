package util

import "strings"

// ConvertString 将字符串分割并转为接口数组
func ConvertString(result string) []interface{} {
	var final []interface{}
	for _, v := range strings.Split(result, ",") {
		final = append(final, v)
	}
	return final
}

// ConvertInterfaceArray 用于特殊的转换
func ConvertInterfaceArray(columns []string, output [][]string) ([]interface{}, [][]interface{}) {
	var outputHeaderArray []interface{}
	var outputDataArray [][]interface{}
	for _, arg := range columns {
		outputHeaderArray = append(outputHeaderArray, arg)
	}
	for _, arg := range output {
		var temp []interface{}
		for _, v := range arg {
			temp = append(temp, v)
		}
		outputDataArray = append(outputDataArray, temp)
	}
	return outputHeaderArray, outputDataArray
}
