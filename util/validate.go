package util

import (
	"go-sqlmap/log"
	"os"
	"strings"
)

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

func GetCleanUrl(url string) string {
	path := strings.Split(url, "?")[0]
	param := strings.Split(url, "?")[1]
	key := strings.Split(param, "=")[0]
	return path + "?" + key + "="
}
