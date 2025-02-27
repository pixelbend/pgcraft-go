package expr

const (
	openPar    = "("
	closePar   = ")"
	commaSpace = ", "
)

var (
	and               = Raw("AND")
	not               = Raw("NOT")
	null              = Raw("NULL")
	isNull            = Raw("IS NULL")
	isNotNull         = Raw("IS NOT NULL")
	between           = Raw("BETWEEN")
	notBetween        = Raw("NOT BETWEEN")
	isDistinctFrom    = Raw("IS DISTINCT FROM")
	isNotDistinctFrom = Raw("IS NOT DISTINCT FROM")
)
