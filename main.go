package main

import (
	"fmt"
	"go-sqlmap/constant"
	"go-sqlmap/core"
	"go-sqlmap/log"
	"go-sqlmap/start"
	"go-sqlmap/util"
	"os"
	"os/signal"
	"syscall"
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
	go start.RunErrorBased(target, params)

	exit()
}

func exit() {
	sign := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sign
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	<-done
}
