package start

import "go-sqlmap/constant"

// NewStarter 核心启动函数
func NewStarter(target string, params Input) {
	for _, v := range params.Technique {
		if v == constant.UnionSelectTech {
			RunUnionSelect(target, params)
		}
		if v == constant.ErrorBasedTech {
			RunErrorBased(target, params)
		}
		if v == constant.BoolBlindTech {
			RunBoolBlind(target, params)
		}
		if v == constant.TimeBlindTech {
			RunTimeBlind(target, params)
		}
	}
}
