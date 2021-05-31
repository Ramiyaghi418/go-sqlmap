package start

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/util"
	"time"
)

// DetectAlive 检测目标是否存活
func DetectAlive(url string) bool {
	code, _, _ := util.Request(constant.RequestMethod, url, nil, nil)
	if code != -1 {
		log.Info("connect success...")
		return true
	} else {
		log.Info("connect error and try again...")
		time.Sleep(time.Second * 3)
		innerCode, _, _ := util.Request(constant.RequestMethod, url, nil, nil)
		if innerCode == -1 {
			log.Error("connect error!")
			return false
		} else {
			return true
		}
	}
}
