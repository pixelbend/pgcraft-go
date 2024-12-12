package psql

import (
	"github.com/codefrantic/pgcraft-go"
	"github.com/codefrantic/pgcraft-go/expr"
)

func RawQuery(query string, args ...any) pgcraft.BaseQuery[expr.Clause] {
	return expr.RawQuery(query, args...)
}
