package parse

import (
	"github.com/EmYiQing/go-sqlmap/log"
	"os"
)

type BaseResponse struct {
	Code    int
	Headers map[string]string
	Body    []byte
}

// GetUrl 从解析HTTP文件中获取URL
func GetUrl(req BaseRequest) string {
	host, ok := req.Headers["Host"]
	if !ok {
		log.Error("must have host header!")
		os.Exit(-1)
	}
	url := "http://" + host + req.Path
	return url
}
