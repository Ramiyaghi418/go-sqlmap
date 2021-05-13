package core

import (
	"go-sqlmap/log"
	"go-sqlmap/util"
	"os"
	"time"
)

func DetectAlive(url string) bool {
	code, _, _ := util.Request("GET", url, nil, nil)
	if code != -1 {
		log.Info("connect success...")
		return true
	} else {
		log.Info("connect error and try again...")
		time.Sleep(time.Second * 3)
		code, _, _ := util.Request("GET", url, nil, nil)
		if code == -1 {
			log.Error("connect error")
			os.Exit(-1)
			return false
		} else {
			return true
		}
	}
}
