package start

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/core"
	"github.com/EmYiQing/go-sqlmap/input"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
)

// NewStarter 启动简单方式的函数
func NewStarter(target string, params input.Input) {
	// 处理URL得到结构体
	fixUrl := parse.NewUrl(target)
	// 检测输入的-p参数是否是url的参数
	checkParams(&params, fixUrl)
	// 检测是否存在SQL注入
	success := core.DetectSqlInject(fixUrl, params.Param)
	if !success {
		return
	}
	// 尝试得到可能的闭合列表
	success, suffixList := core.GetSuffixList(fixUrl, params.Param)
	if !success {
		return
	}
	// 根据输入选择不同的技术进行注入
	for _, v := range params.Technique {
		if v == constant.UnionSelectTech {
			if RunUnionSelect(fixUrl, params, suffixList) {
				return
			}
			log.Info("finish union select injection")
		}
		if v == constant.ErrorBasedTech {
			if RunErrorBased(fixUrl, params, suffixList) {
				return
			}
			log.Info("finish error based injection")
		}
		if v == constant.BoolBlindTech {
			if RunBoolBlind(fixUrl, params, suffixList) {
				return
			}
			log.Info("finish bool blind injection")
		}
	}
}

// 检查P参数是否合法
func checkParams(param *input.Input, fixUrl parse.BaseUrl) {
	for k := range fixUrl.Params {
		if param.Param == k {
			return
		}
	}
	for k := range fixUrl.Params {
		param.Param = k
	}
}
