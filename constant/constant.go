package constant

var (
	SuffixList = []string{"%20", "'", "\"", ")", "')", "\")"}
)

const (
	DefaultMethod    = "GET"
	DetectWafPayload = "'%20or%201=1--+"
	DetectedKeyword  = "You have an error in your SQL syntax"
	OrderKeyword     = "Unknown column"
	Annotator        = "--+"
	Space            = "%20"
)

const (
	SafeDogKeyword       = "www.safedog.cn"
	SafeDogHeaderKey     = ""
	SafeDogHeaderKeyword = ""
)
