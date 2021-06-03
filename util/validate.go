package util

import (
	"github.com/EmYiQing/go-sqlmap/log"
	"os"
	"strings"
)

// CheckUrl 校验URL
func CheckUrl(url string) string {
	var result string
	if strings.TrimSpace(url) == "" {
		log.Error("Need Url!")
		os.Exit(-1)
	}
	if !strings.HasPrefix(url, "http") {
		result = "http://" + url
	} else {
		result = url
	}
	if strings.HasSuffix(url, "/") {
		r := []rune(result)
		result = string(r[:len(r)-1])
	}
	return result
}
