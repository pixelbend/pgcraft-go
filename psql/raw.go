package psql

import (
	"github.com/arkamfahry/pgcraft-go"
	"github.com/arkamfahry/pgcraft-go/expr"
)

func RawQuery(query string, args ...any) pgcraft.BaseQuery[expr.Clause] {
	return expr.RawQuery(query, args...)
}
