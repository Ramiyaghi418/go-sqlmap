package parse

import (
	"io/ioutil"
	"runtime"
	"strings"
)

type BaseRequest struct {
	Method  string
	Path    string
	Headers map[string]string
	Cookie  map[string]string
	Data    string
}

// RequestParse 解析HTTP协议
func RequestParse(filename string) (req *BaseRequest) {
	req = &BaseRequest{}
	sysType := runtime.GOOS
	lineSep := "\n"
	if sysType == "windows" {
		lineSep = "\r\n"
	}
	requestByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	request := string(requestByte)
	temp := strings.Split(request, lineSep)
	if len(temp) < 1 {
		return
	}
	firstLine := temp[0]
	firstTemp := strings.Split(firstLine, " ")
	if len(firstTemp) < 3 {
		return
	}
	requestMethod := firstTemp[0]
	path := firstTemp[1]

	var cookieIndex int
	headers := make(map[string]string)
	for i := 1; i < len(temp); i++ {
		if strings.TrimSpace(temp[i]) == "" {
			break
		}
		key := strings.Split(temp[i], ": ")[0]
		value := strings.Split(temp[i], ": ")[1]
		if strings.ToLower(key) == "cookie" {
			cookieIndex = i
			continue
		}
		headers[key] = value
	}
	cookies := make(map[string]string)
	tempCookie := strings.Split(temp[cookieIndex], ": ")[1]
	if !strings.Contains(tempCookie, "; ") {
		key := strings.Split(tempCookie, "=")[0]
		value := strings.Split(tempCookie, "=")[1]
		cookies[key] = value
	} else {
		for _, v := range strings.Split(tempCookie, "; ") {
			key := strings.Split(v, "=")[0]
			value := strings.Split(v, "=")[1]
			cookies[key] = value
		}
	}

	dataTemp := strings.Split(request, lineSep+lineSep)
	data := ""
	if len(dataTemp) > 1 {
		data = dataTemp[1]
	}
	req.Data = data
	req.Cookie = cookies
	req.Method = requestMethod
	req.Path = path
	req.Headers = headers
	return
}
