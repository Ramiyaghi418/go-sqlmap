package start

import (
	"go-sqlmap/constant"
	"go-sqlmap/log"
	"go-sqlmap/util"
	"time"
)

func DetectAlive(url string) bool {
	code, _, _ := util.Request(constant.DefaultMethod, url, nil, nil)
	if code != -1 {
		log.Info("connect success...")
		return true
	} else {
		log.Info("connect error and try again...")
		time.Sleep(time.Second * 3)
		innerCode, _, _ := util.Request(constant.DefaultMethod, url, nil, nil)
		if innerCode == -1 {
			log.Error("connect error!")
			return false
		} else {
			return true
		}
	}
}
