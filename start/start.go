package start

import (
	"go-sqlmap/constant"
	"go-sqlmap/core"
	"go-sqlmap/util"
)

func RunUnionSelect(target string, params Input) bool {
	success, _ := core.DetectUnionSelectSqlInject(target, constant.DefaultMethod)
	if !success {
		return false
	}
	success, suffixList := core.GetSuffix(params.Url)
	if !success {
		return false
	}
	suffix, key := core.GetOrderByNum(suffixList, target)
	if key == 0 {
		return false
	}
	cleanUrl := util.GetCleanUrl(target)
	pos := core.GetUnionSelectPos(suffix, cleanUrl, key)
	core.GetVersion(pos, suffix, cleanUrl, key)
	core.GetCurrentDatabase(pos, suffix, cleanUrl, key)
	core.GetAllDatabases(pos, suffix, cleanUrl, key)
	if params.Database != "" {
		core.GetAllTables(pos, suffix, cleanUrl, key, params.Database)
	}
	if params.Database != "" && params.Table != "" {
		core.GetColumns(pos, suffix, cleanUrl, key, params.Database, params.Table)
	}
	if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
		core.GetData(pos, suffix, cleanUrl, key, params.Database, params.Table, params.Columns)
	}
	return true
}

func RunErrorBased(target string, params Input) bool {
	// TODO
	return true
}

func RunBoolBlind(target string, params Input) bool {
	// TODO
	return true
}

func RunTimeBlind(target string, params Input) bool {
	// TODO
	return true
}
