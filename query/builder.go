package query

import (
	"github.com/teapartydev/pgcraft-go"
	"github.com/teapartydev/pgcraft-go/expr"
	"strings"
)

var (
	and                 = expr.Raw("AND")
	betweenSymmetric    = expr.Raw("BETWEEN")
	notBetweenSymmetric = expr.Raw("NOT BETWEEN")
	iLike               = expr.Raw("ILIKE")
)

type Expression struct {
	expr.Chain[Expression, Expression]
}

func (Expression) New(exp pgcraft.Expression) Expression {
	var b Expression
	b.Base = exp
	return b
}

// Implements fmt.Stringer()
func (x Expression) String() string {
	w := strings.Builder{}
	x.WriteSQL(&w, 1) //nolint:errcheck
	return w.String()
}

// BETWEEN SYMMETRIC a AND b
func (x Expression) BetweenSymmetric(a, e pgcraft.Expression) Expression {
	return expr.X[Expression, Expression](expr.Join{Exprs: []pgcraft.Expression{
		x.Base, betweenSymmetric, a, and, e,
	}})
}

// NOT BETWEEN SYMMETRIC a AND b
func (x Expression) NotBetweenSymmetric(a, e pgcraft.Expression) Expression {
	return expr.X[Expression, Expression](expr.Join{Exprs: []pgcraft.Expression{
		x.Base, notBetweenSymmetric, a, and, e,
	}})
}

// ILIKE val
func (x Expression) ILike(val pgcraft.Expression) Expression {
	return expr.X[Expression, Expression](expr.Join{Exprs: []pgcraft.Expression{
		x.Base, iLike, val,
	}})
}
