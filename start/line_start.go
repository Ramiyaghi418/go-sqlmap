package start

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/core"
	"github.com/EmYiQing/go-sqlmap/input"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
)

// RunUnionSelect UnionSelect注入
func RunUnionSelect(fixUrl parse.BaseUrl, params input.Input, suffixList []string) bool {
	log.Info("start union select injection...")
	suffix, key := core.GetOrderByNum(fixUrl, params.Param, suffixList)
	if key == 0 {
		return false
	}
	paramKey := params.Param
	fixUrl.Params[paramKey] = constant.UnionSelectUnionCondition
	pos := core.GetUnionSelectPos(suffix, fixUrl, paramKey, key)
	if pos.StartIndex == 0 {
		return false
	}
	core.GetVersion(pos, suffix, fixUrl, paramKey, key)
	core.GetCurrentDatabase(pos, suffix, fixUrl, paramKey, key)
	core.GetAllDatabases(pos, suffix, fixUrl, paramKey, key)
	if params.Database != "" {
		core.GetAllTables(pos, suffix, fixUrl, paramKey, key, params.Database)
	}
	if params.Database != "" && params.Table != "" {
		core.GetColumns(pos, suffix, fixUrl, paramKey, key, params.Database, params.Table)
	}
	if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
		core.GetData(pos, suffix, fixUrl, paramKey, key, params.Database, params.Table, params.Columns)
	}
	return true
}

// RunErrorBased ErrorBased注入
func RunErrorBased(fixUrl parse.BaseUrl, params input.Input, suffixList []string) bool {
	log.Info("start error based injection...")
	success, suffix := core.DetectErrorBased(fixUrl, params.Param, suffixList)
	if !success {
		return false
	}
	if params.Beta == true {
		core.GetVersionByErrorBasedPolygon(fixUrl, params.Param, suffix)
		core.GetCurrentDatabaseByErrorBasedPolygon(fixUrl, params.Param, suffix)
		core.GetAllDatabasesByErrorBasedPolygon(fixUrl, params.Param, suffix)
		if params.Database != "" {
			core.GetAllTablesByErrorBasedPolygon(fixUrl, params.Param, suffix, params.Database)
		}
		if params.Database != "" && params.Table != "" {
			core.GetAllColumnsByErrorBasedPolygon(fixUrl, params.Param, suffix, params.Database, params.Table)
		}
		if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
			core.GetAllDataByErrorBasedPolygon(fixUrl, params.Param, suffix, params.Database, params.Table, params.Columns)
		}
	} else {
		core.GetVersionByErrorBased(fixUrl, params.Param, suffix)
		core.GetCurrentDatabaseByErrorBased(fixUrl, params.Param, suffix)
		core.GetAllDatabasesByErrorBased(fixUrl, params.Param, suffix)
		if params.Database != "" {
			core.GetAllTablesByErrorBased(fixUrl, params.Param, suffix, params.Database)
		}
		if params.Database != "" && params.Table != "" {
			core.GetAllColumnsByErrorBased(fixUrl, params.Param, suffix, params.Database, params.Table)
		}
		if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
			core.GetAllDataByErrorBased(fixUrl, params.Param, suffix, params.Database, params.Table, params.Columns)
		}
	}
	return true
}

// RunBoolBlind BoolBlind注入
func RunBoolBlind(fixUrl parse.BaseUrl, params input.Input, suffixList []string) bool {
	log.Info("start bool blind injection...")
	success, suffix := core.GetBoolBlindSuffix(fixUrl, params.Param, suffixList)
	if !success {
		return false
	}
	core.GetVersionByBoolBlind(fixUrl, params.Param, suffix)
	core.GetCurrentDatabaseByBoolBlind(fixUrl, params.Param, suffix)
	core.GetAllDatabasesByBoolBlind(fixUrl, params.Param, suffix)
	if params.Database != "" {
		core.GetAllTablesByBoolBlind(fixUrl, params.Param, suffix, params.Database)
	}
	if params.Database != "" && params.Table != "" {
		core.GetAllColumnsByBoolBlind(fixUrl, params.Param, suffix, params.Database, params.Table)
	}
	if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
		core.GetAllDataByBoolBlind(fixUrl, params.Param, suffix, params.Database, params.Table, params.Columns)
	}
	return true
}
