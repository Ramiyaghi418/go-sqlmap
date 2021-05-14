package util

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
