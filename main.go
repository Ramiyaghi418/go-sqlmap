package main

import (
	"fmt"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/parse"
	"github.com/EmYiQing/go-sqlmap/start"
	"github.com/EmYiQing/go-sqlmap/str"
	"github.com/EmYiQing/go-sqlmap/util"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	start.PrintLogo(str.Version, str.Author, str.Url)
	params := start.ParseInput()
	var target string
	if params.Url != "" {
		target = util.CheckUrl(params.Url)
		log.Info("target is " + target)
		if !start.DetectAlive(target) {
			os.Exit(-1)
		}
		if start.DetectSafeDogWaf(target) {
			os.Exit(-1)
		}
		start.NewSimpleStarter(target, params)
		wait()
	} else {
		req := parse.RequestParse(params.Filename)
		host, ok := req.Headers["Host"]
		if !ok {
			log.Error("must have host header!")
			os.Exit(-1)
		}
		url := "http://" + host + req.Path
		target = util.CheckUrl(url)
		log.Info("target is " + target)
		if !start.DetectAlive(target) {
			os.Exit(-1)
		}
		if start.DetectSafeDogWaf(target) {
			os.Exit(-1)
		}
		start.NewStarter(*req, params)
		wait()
	}
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
