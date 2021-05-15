package main

import (
	"go-sqlmap/constant"
	"go-sqlmap/core"
	"go-sqlmap/log"
	"go-sqlmap/start"
	"go-sqlmap/util"
	"os"
)

func main() {
	core.PrintLogo(constant.Version, constant.Author, constant.Url)
	params := core.ParseInput()
	target := util.CheckUrl(params.Url)
	log.Info("target is " + target)

	if !core.DetectAlive(target) {
		os.Exit(-1)
	}

	if core.DetectSafeDogWaf(target) {
		os.Exit(-1)
	}

	// Error Based Injection
	start.RunErrorBased(target, params)
}
