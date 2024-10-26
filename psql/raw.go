package psql

import (
	"github.com/teapartydev/pgcraft-go"
	"github.com/teapartydev/pgcraft-go/expr"
)

func RawQuery(query string, args ...any) pgcraft.BaseQuery[expr.Clause] {
	return expr.RawQuery(query, args...)
}
