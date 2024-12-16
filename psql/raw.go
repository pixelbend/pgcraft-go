package psql

import (
	"github.com/pixelbend/pgcraft-go"
	"github.com/pixelbend/pgcraft-go/expr"
)

func RawQuery(query string, args ...any) pgcraft.BaseQuery[expr.Clause] {
	return expr.RawQuery(query, args...)
}
