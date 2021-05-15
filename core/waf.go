package core

import (
	"go-sqlmap/constant"
	"go-sqlmap/util"
	"strings"
)

func DetectSafeDogWaf(url string) bool {
	url = url + constant.DetectWafPayload
	code, headers, body := util.Request(constant.DefaultMethod, url, nil, nil)
	if code != -1 {
		if strings.Contains(string(body), constant.SafeDogKeyword) {
			if strings.Contains(headers[constant.SafeDogHeaderKey],
				constant.SafeDogHeaderKeyword) {
				return true
			}
		}
	}
	return false
}
