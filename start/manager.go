package start

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/input"
	"github.com/EmYiQing/go-sqlmap/line"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
	"strings"
)

// NewSimpleStarter 启动简单方式的函数
func NewSimpleStarter(target string, params input.Input) {
	fixUrl := parse.NewUrl(target)
	checkParams(params, fixUrl)
	success := line.DetectSqlInject(fixUrl, params.Param)
	if !success {
		return
	}
	success, suffixList := line.GetSuffixList(fixUrl, params.Param)
	if !success {
		return
	}
	for _, v := range params.Technique {
		if v == constant.UnionSelectTech {
			if RunUnionSelect(target, params, suffixList) {
				return
			}
			log.Info("finish union select injection")
		}
		if v == constant.ErrorBasedTech {
			if RunErrorBased(target, params, suffixList) {
				return
			}
			log.Info("finish error based injection")
		}
		if v == constant.BoolBlindTech {
			if RunBoolBlind(target, params, suffixList) {
				return
			}
			log.Info("finish bool blind injection")
		}
	}
}

// NewStarter 启动请求文件形式的函数
func NewStarter(request parse.BaseRequest, params input.Input) {
	var headerInjectKeys []string
	var dataInjectKeys []string
	for k, v := range request.Headers {
		if strings.HasSuffix(v, "*") {
			headerInjectKeys = append(headerInjectKeys, k)
		}
	}
	for k, v := range request.Data {
		if strings.HasSuffix(v, "*") {
			dataInjectKeys = append(dataInjectKeys, k)
		}
	}
	// TODO http request file injection
}

// 检查P参数是否合法
func checkParams(param input.Input, fixUrl parse.BaseUrl) {
	for k, _ := range fixUrl.Params {
		if param.Param == k {
			return
		}
	}
	for k, _ := range fixUrl.Params {
		param.Param = k
	}
}
