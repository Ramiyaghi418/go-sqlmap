package parse

import (
	"bytes"
	"github.com/EmYiQing/go-sqlmap/util"
	"regexp"
	"strconv"
	"strings"
)

type BaseUrl struct {
	Protocol string
	Host     string
	Port     int
	Path     string
	Params   map[string]string
}

// NewUrl 将字符串URL转为URL结构体
func NewUrl(url string) BaseUrl {
	ret := BaseUrl{}
	protocolRe := regexp.MustCompile("(http.*?)://")
	protocol := protocolRe.FindAllStringSubmatch(url, -1)[0][1]
	ret.Protocol = protocol
	hostRe := regexp.MustCompile("http.*?://(.*?)/")
	host := hostRe.FindAllStringSubmatch(url, -1)[0][1]
	port := 80
	ret.Host = host
	if strings.Contains(host, ":") {
		port, _ = strconv.Atoi(strings.Split(host, ":")[1])
		ret.Host = strings.Split(host, ":")[0]
	}
	ret.Port = port
	pathRe := regexp.MustCompile("http.*?://.*?(/.*)\\?")
	path := pathRe.FindAllStringSubmatch(url, -1)[0][1]
	ret.Path = path
	paramsRe := regexp.MustCompile("http.*?://.*?\\?(.*)")
	paramsStr := paramsRe.FindAllStringSubmatch(url, -1)[0][1]
	if strings.TrimSpace(paramsStr) != "" {
		params := strings.Split(paramsStr, "&")
		paramsMap := make(map[string]string)
		for _, v := range params {
			temp := strings.Split(v, "=")
			key := temp[0]
			value := temp[1]
			paramsMap[key] = value
		}
		ret.Params = paramsMap
	}
	return ret
}

// SendRequestByBaseUrl 使用BaseUrl发GET请求
func (u BaseUrl) SendRequestByBaseUrl() BaseResponse {
	finalUrl := bytes.Buffer{}
	finalUrl.WriteString(u.Protocol)
	finalUrl.WriteString("://")
	finalUrl.WriteString(u.Host)
	finalUrl.WriteString(":")
	finalUrl.WriteString(strconv.Itoa(u.Port))
	finalUrl.WriteString(u.Path)
	finalUrl.WriteString("?")
	var temp string
	for k, v := range u.Params {
		temp += k
		temp += "="
		temp += v
		temp += "&"
		finalUrl.WriteString(temp)
	}
	finalUrlStr := finalUrl.String()
	finalUrlStr = util.DeleteLastChar(finalUrlStr)
	code, resHeaders, bodyByte := util.Request("GET", finalUrlStr, nil, nil)
	return BaseResponse{Code: code, Headers: resHeaders, Body: bodyByte}
}

func (u BaseUrl) SetParam(key string, value string) string {
	temp := u.Params[key]
	u.Params[key] = value
	return temp
}
