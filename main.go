package main

import (
	"fmt"
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/input"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/start"
	"github.com/EmYiQing/go-sqlmap/util"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	start.PrintLogo(constant.Version, constant.Author, constant.Url)
	params := input.ParseInput()
	if params.Url != "" {
		target := doPre(params.Url)
		start.NewStarter(target, params)
		wait()
	}
}

// 预处理
func doPre(url string) string {
	target := util.CheckUrl(url)
	log.Info("target is " + target)
	if !start.DetectAlive(target) {
		os.Exit(-1)
	}
	if start.DetectSafeDogWaf(target) {
		os.Exit(-1)
	}
	return target
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
