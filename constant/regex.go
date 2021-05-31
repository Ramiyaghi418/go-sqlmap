package constant

const (
	// UpdatexmlRegex Updatexml函数获得数据的正则
	UpdatexmlRegex = "XPATH syntax error: '~(.*?)~'"
)

const (
	// PolygonVersionRegex 报错获得版本的正则
	PolygonVersionRegex = "Illegal.*version.*select\\s'(.*?)' AS.*?"
	// PolygonDataRegex 报错获得数据的正则
	PolygonDataRegex = "Illegal.*group_concat.*select\\s'(.*?)' AS.*?"
	// PolygonFinalDataRegex 报错获得最终数据的正则
	PolygonFinalDataRegex = "Illegal.*?concat.*?from\\s\\(select\\s'(.*?)'\\sAS\\s`"
	// PolygonDatabaseRegex 报错获得当前数据库的正则
	PolygonDatabaseRegex = "Illegal.*database.*select\\s'(.*?)' AS.*?"
)
