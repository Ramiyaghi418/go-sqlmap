package util

import (
	"bytes"
	"github.com/EmYiQing/go-sqlmap/log"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request 发送http请求
func Request(method string, url string, data interface{},
	headers interface{}) (int, map[string]string, []byte) {
	var innerData map[string]string
	var innerHeaders map[string]string
	if data == nil {
		innerData = make(map[string]string)
	} else {
		innerData = data.(map[string]string)
	}
	if headers == nil {
		innerHeaders = make(map[string]string)
	} else {
		innerHeaders = data.(map[string]string)
	}
	client := http.Client{}
	var (
		req       *http.Request
		finalData []byte
	)
	if value, ok := innerHeaders["Content-Type"]; ok {
		if value == "application/json" {
			resolveJson(innerData)
		}
	} else {
		finalData = resolveForm(innerData)
	}
	if len(finalData) != 0 {
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(finalData))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	for k, v := range innerHeaders {
		req.Header.Set(k, v)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")
	resp, err := client.Do(req)
	if err == nil {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		defer func() {
			if resp != nil {
				_ = resp.Body.Close()
			}
		}()
		respHeader := make(map[string]string)
		for k, v := range resp.Header {
			respHeader[k] = strings.Join(v, "")
		}
		return resp.StatusCode, respHeader, body
	} else {
		log.Error("request error:" + err.Error() + "->" + url)
	}
	return -1, nil, nil
}

// 解决表单方式的提交
func resolveForm(data map[string]string) []byte {
	var temp bytes.Buffer
	var finalData []byte
	for k, v := range data {
		temp.WriteString(k + "=" + v + "&")
	}
	finalData = []byte(strings.TrimRight(temp.String(), "&"))
	return finalData
}

// 解决json方式的提交
func resolveJson(data map[string]string) []byte {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytesData, _ := json.Marshal(data)
	return bytesData
}
