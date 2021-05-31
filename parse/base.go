package parse

type BaseResponse struct {
	Code    int
	Headers map[string]string
	Body    []byte
}
