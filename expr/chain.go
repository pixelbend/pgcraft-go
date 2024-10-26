package expr

import (
	"io"
)

type Chain[T pgcraft.Expression, B builder[T]] struct {
	Base pgcraft.Expression
}

// WriteSQL satisfies the pgcraft.Expression interface
func (x Chain[T, B]) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.Express(w, start, x.Base)
}

// IS DISTINCT FROM
func (x Chain[T, B]) IsDistinctFrom(exp pgcraft.Expression) T {
	return X[T, B](Join{Exprs: []pgcraft.Expression{x.Base, isDistinctFrom, exp}})
}

// IS NOT DISTINCT FROM
func (x Chain[T, B]) IsNotDistinctFrom(exp pgcraft.Expression) T {
	return X[T, B](Join{Exprs: []pgcraft.Expression{x.Base, isNotDistinctFrom, exp}})
}

// IS NUll
func (x Chain[T, B]) IsNull() T {
	return X[T, B](Join{Exprs: []pgcraft.Expression{x.Base, isNull}})
}

// IS NOT NUll
func (x Chain[T, B]) IsNotNull() T {
	return X[T, B](Join{Exprs: []pgcraft.Expression{x.Base, isNotNull}})
}

// Generic Operator
func (x Chain[T, B]) OP(op string, target pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: target, operator: op})
}

// Equal
func (x Chain[T, B]) EQ(target pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: target, operator: "="})
}

// Not Equal
func (x Chain[T, B]) NE(target pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: target, operator: "<>"})
}

// Less than
func (x Chain[T, B]) LT(target pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: target, operator: "<"})
}

// Less than or equal to
func (x Chain[T, B]) LTE(target pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: target, operator: "<="})
}

// Greater than
func (x Chain[T, B]) GT(target pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: target, operator: ">"})
}

// Greater than or equal to
func (x Chain[T, B]) GTE(target pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: target, operator: ">="})
}

// IN
func (x Chain[T, B]) In(vals ...pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: group(vals), operator: "IN"})
}

// NOT IN
func (x Chain[T, B]) NotIn(vals ...pgcraft.Expression) T {
	return X[T, B](leftRight{left: x.Base, right: group(vals), operator: "NOT IN"})
}

// OR
func (x Chain[T, B]) Or(targets ...pgcraft.Expression) T {
	return X[T, B](Join{Exprs: append([]pgcraft.Expression{x.Base}, targets...), Sep: " OR "})
}

// AND
func (x Chain[T, B]) And(targets ...pgcraft.Expression) T {
	return X[T, B](Join{Exprs: append([]pgcraft.Expression{x.Base}, targets...), Sep: " AND "})
}

// Concatenate: ||
func (x Chain[T, B]) Concat(targets ...pgcraft.Expression) T {
	return X[T, B](Join{Exprs: append([]pgcraft.Expression{x.Base}, targets...), Sep: " || "})
}

// BETWEEN a AND b
func (x Chain[T, B]) Between(a, b pgcraft.Expression) T {
	return X[T, B](Join{Exprs: []pgcraft.Expression{x.Base, between, a, and, b}})
}

// NOT BETWEEN a AND b
func (x Chain[T, B]) NotBetween(a, b pgcraft.Expression) T {
	return X[T, B](Join{Exprs: []pgcraft.Expression{
		x.Base, notBetween, a, and, b,
	}})
}

// Subtract
func (x Chain[T, B]) Minus(target pgcraft.Expression) T {
	return X[T, B](leftRight{operator: "-", left: x.Base, right: target})
}

// Like operator
func (x Chain[T, B]) Like(target pgcraft.Expression) T {
	return X[T, B](leftRight{operator: "LIKE", left: x.Base, right: target})
}

// As does not return a new chain. Should be used at the end of an expression
// useful for columns
func (x Chain[T, B]) As(alias string) pgcraft.Expression {
	return leftRight{left: x.Base, operator: "AS", right: quoted{alias}}
}
