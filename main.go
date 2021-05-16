package main

import (
	"fmt"
	"go-sqlmap/constant"
	"go-sqlmap/log"
	"go-sqlmap/start"
	"go-sqlmap/util"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	start.PrintLogo(constant.Version, constant.Author, constant.Url)

	params := start.ParseInput()

	target := util.CheckUrl(params.Url)
	log.Info("target is " + target)

	if !start.DetectAlive(target) {
		os.Exit(-1)
	}

	if start.DetectSafeDogWaf(target) {
		os.Exit(-1)
	}

	start.NewStarter(target, params)

	wait()
}

// 使用信号优雅退出
func wait() {
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
