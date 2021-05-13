package core

import (
	"go-sqlmap/log"
	"go-sqlmap/util"
	"os"
	"strings"
)

func DetectWaf(url string) {
	waf := false
	payload := "'%20or%201=1--+"
	url = url + payload
	code, headers, body := util.Request("GET", url, nil, nil)
	if code != -1 {
		keyword := "www.safedog.cn"
		if strings.Contains(string(body), keyword) {
			if strings.Contains(headers["X-Powered-By"], "WAF") {
				waf = true
			}
		}
	}
	if waf {
		log.Info("there is a waf")
		os.Exit(-1)
	}
}
