package psql

import (
	"github.com/driftdev/pgcraft-go"
	"github.com/driftdev/pgcraft-go/expr"
)

func RawQuery(query string, args ...any) pgcraft.BaseQuery[expr.Clause] {
	return expr.RawQuery(query, args...)
}
