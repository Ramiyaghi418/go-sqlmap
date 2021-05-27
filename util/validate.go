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

// GetCleanUrl 获得需要的URL格式
func GetCleanUrl(url string) string {
	path := strings.Split(url, "?")[0]
	param := strings.Split(url, "?")[1]
	key := strings.Split(param, "=")[0]
	return path + "?" + key + "="
}
