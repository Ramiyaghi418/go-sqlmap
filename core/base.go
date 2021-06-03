package core

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
	"strings"
)

// DetectSqlInject 检测是否存在注入
func DetectSqlInject(fixUrl parse.BaseUrl, paramKey string) bool {
	temp := fixUrl.Params[paramKey]
	for _, v := range constant.SuffixList {
		fixUrl.SetParam(paramKey, temp+v)
		response := fixUrl.SendRequestByBaseUrl()
		if response.Code != -1 {
			// 关键字判断回显型注入
			if strings.Contains(strings.ToLower(string(response.Body)),
				strings.ToLower(constant.DetectedKeyword)) {
				log.Info("detected sql injection!")
				fixUrl.Params[paramKey] = temp
				return true
			}
		}
		// Bool盲注检测
		fixUrl.SetParam(paramKey, temp+v+constant.BlindDetectTruePayload)
		trueResp := fixUrl.SendRequestByBaseUrl()
		fixUrl.SetParam(paramKey, temp+v+constant.BlindDetectFalsePayload)
		falseResp := fixUrl.SendRequestByBaseUrl()
		if len(trueResp.Body) == len(falseResp.Body) {
			continue
		}
		if string(trueResp.Body) != string(falseResp.Body) {
			fixUrl.Params[paramKey] = temp
			return true
		}
	}
	log.Info("not detected sql injection!")
	fixUrl.Params[paramKey] = temp
	return false
}

// GetSuffixList 获取可能的闭合符号列表
func GetSuffixList(fixUrl parse.BaseUrl, key string) (bool, []string) {
	defaultResp := fixUrl.SendRequestByBaseUrl()
	defaultBody := string(defaultResp.Body)
	temp := fixUrl.Params[key]
	var suffixList []string
	for _, v := range constant.SuffixList {
		// 尝试闭合后直接注释
		fixUrl.SetParam(key, temp+v+constant.SuffixCondition)
		conditionResp := fixUrl.SendRequestByBaseUrl()
		conditionBody := conditionResp.Body
		// 与True条件应正常返回
		fixUrl.SetParam(key, temp+v+constant.SuffixTruePayload)
		trueResp := fixUrl.SendRequestByBaseUrl()
		trueBody := trueResp.Body
		// 与False条件应返回有误
		fixUrl.SetParam(key, temp+v+constant.SuffixFalsePayload)
		falseResp := fixUrl.SendRequestByBaseUrl()
		falseBody := falseResp.Body
		// 双重验证只能尽量保证闭合符号正确，得出有最有可能的闭合符号
		if defaultBody == string(conditionBody) &&
			string(trueBody) != string(falseBody) {
			suffixList = append(suffixList, v)
		}
	}
	fixUrl.SetParam(key, temp)
	if len(suffixList) > 0 {
		return true, suffixList
	}
	return false, suffixList
}
