package main

import (
	"go-sqlmap/core"
	"go-sqlmap/log"
	"go-sqlmap/util"
)

const (
	version = "0.1"
	author  = "4ra1n"
	url     = "https://github.com/EmYiQing/go-sqlmap"
)

func main() {
	core.PrintLogo(version, author, url)
	params := core.ParseInput()
	target := util.CheckUrl(params.Url)
	log.Info("target is " + target)
	core.DetectAlive(target)
	core.DetectWaf(target)
	core.DetectSqlInject(target)
	suffix := core.GetSuffix(params.Url)
	key := core.GetOrderByNum(suffix, target)
	cleanUrl := util.GetCleanUrl(target)
	pos := core.GetUnionSelectPos(suffix, cleanUrl, key)
	core.GetVersion(pos, suffix, cleanUrl, key)
	core.GetDatabase(pos, suffix, cleanUrl, key)
	if params.Database != "" {
		core.GetAllTables(pos, suffix, cleanUrl, key)
	}
	if params.Table != "" {
		core.GetColumns(pos, suffix, cleanUrl, key, params.Table)
	}
	if len(params.Columns) > 0 {
		core.GetData(pos, suffix, cleanUrl, key, params.Table, params.Columns)
	}
}
