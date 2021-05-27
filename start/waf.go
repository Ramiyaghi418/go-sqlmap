package start

import (
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/str"
	"github.com/EmYiQing/go-sqlmap/util"
	"strings"
)

// DetectSafeDogWaf 检测安全狗的WAF
func DetectSafeDogWaf(url string) bool {
	url = url + str.DetectWafPayload
	code, headers, body := util.Request(str.RequestMethod, url, nil, nil)
	if code != -1 {
		if strings.Contains(string(body), str.SafeDogKeyword) {
			if strings.Contains(headers[str.SafeDogHeaderKey],
				str.SafeDogHeaderKeyword) {
				log.Error("there is a waf!")
				return true
			}
		}
	}
	return false
}
