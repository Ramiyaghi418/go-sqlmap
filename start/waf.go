package start

import (
	"github.com/EmYiQing/go-sqlmap/constant"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/util"
	"strings"
)

// DetectSafeDogWaf 检测安全狗的WAF
func DetectSafeDogWaf(url string) bool {
	url = url + constant.DetectWafPayload
	code, headers, body := util.Request(constant.RequestMethod, url, nil, nil)
	if code != -1 {
		if strings.Contains(string(body), constant.SafeDogKeyword) {
			if strings.Contains(headers[constant.SafeDogHeaderKey],
				constant.SafeDogHeaderKeyword) {
				log.Error("there is a waf!")
				return true
			}
		}
	}
	return false
}
