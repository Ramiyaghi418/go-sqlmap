package start

import (
	"github.com/EmYiQing/go-sqlmap/line"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
	"github.com/EmYiQing/go-sqlmap/str"
	"strings"
)

// NewSimpleStarter 启动简单方式的函数
func NewSimpleStarter(target string, params Input) {
	success, _ := line.DetectSqlInject(target, str.RequestMethod)
	if !success {
		return
	}
	success, suffixList := line.GetSuffixList(params.Url)
	if !success {
		return
	}
	for _, v := range params.Technique {
		if v == str.UnionSelectTech {
			if RunUnionSelect(target, params, suffixList) {
				return
			}
			log.Info("finish union select injection")
		}
		if v == str.ErrorBasedTech {
			if RunErrorBased(target, params, suffixList) {
				return
			}
			log.Info("finish error based injection")
		}
		if v == str.BoolBlindTech {
			if RunBoolBlind(target, params, suffixList) {
				return
			}
			log.Info("finish bool blind injection")
		}
	}
}

// NewStarter 启动请求文件形式的函数
func NewStarter(request parse.BaseRequest, params Input) {
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
