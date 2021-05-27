package start

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/core"
	"github.com/EmYiQing/go-sqlmap/log"
)

// NewStarter 核心启动函数
func NewStarter(target string, params Input) {
	success, _ := core.DetectSqlInject(target, constant.DefaultMethod)
	if !success {
		return
	}
	success, suffixList := core.GetSuffixList(params.Url)
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
