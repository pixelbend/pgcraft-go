package psql

import (
	"github.com/driftdev/pgcraft"
	"github.com/driftdev/pgcraft/expr"
)

func RawQuery(query string, args ...any) pgcraft.BaseQuery[expr.Clause] {
	return expr.RawQuery(query, args...)
}
