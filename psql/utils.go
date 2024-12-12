package psql

import (
	"github.com/codefrantic/pgcraft-go"
	"github.com/codefrantic/pgcraft-go/expr"
	"github.com/codefrantic/pgcraft-go/mods"
	"github.com/codefrantic/pgcraft-go/query"
)

type Expression = query.Expression

//nolint:gochecknoglobals
var bmod = expr.Builder[Expression, Expression]{}

// F creates a function expression with the given name and args
//
//	SQL: generate_series(1, 3)
//	Go: psql.F("generate_series", 1, 3)
func F(name string, args ...any) mods.Moddable[*query.Function] {
	f := query.NewFunction(name, args...)

	return mods.Moddable[*query.Function](func(mods ...pgcraft.Mod[*query.Function]) *query.Function {
		for _, mod := range mods {
			mod.Apply(f)
		}

		return f
	})
}

// S creates a string literal
// SQL: 'a string'
// Go: psql.S("a string")
func S(s string) Expression {
	return bmod.S(s)
}

// SQL: NOT true
// Go: psql.Not("true")
func Not(exp pgcraft.Expression) Expression {
	return bmod.Not(exp)
}

// SQL: a OR b OR c
// Go: psql.Or("a", "b", "c")
func Or(args ...pgcraft.Expression) Expression {
	return bmod.Or(args...)
}

// SQL: a AND b AND c
// Go: psql.And("a", "b", "c")
func And(args ...pgcraft.Expression) Expression {
	return bmod.And(args...)
}

// SQL: a || b || c
// Go: psql.Concat("a", "b", "c")
func Concat(args ...pgcraft.Expression) Expression {
	return expr.X[Expression, Expression](expr.Join{Exprs: args, Sep: " || "})
}

// SQL: $1, $2, $3
// Go: psql.Args("a", "b", "c")
func Arg(args ...any) Expression {
	return bmod.Arg(args...)
}

// SQL: ($1, $2, $3)
// Go: psql.ArgGroup("a", "b", "c")
func ArgGroup(args ...any) Expression {
	return bmod.ArgGroup(args...)
}

// SQL: $1, $2, $3
// Go: psql.Placeholder(3)
func Placeholder(n uint) Expression {
	return bmod.Placeholder(n)
}

// SQL: (a, b)
// Go: psql.Group("a", "b")
func Group(exps ...pgcraft.Expression) Expression {
	return bmod.Group(exps...)
}

// SQL: "table"."column"
// Go: psql.Quote("table", "column")
func Quote(ss ...string) Expression {
	return bmod.Quote(ss...)
}

// SQL: where a = $1
// Go: psql.Raw("where a = ?", "something")
func Raw(query string, args ...any) Expression {
	return bmod.Raw(query, args...)
}

// SQL: CAST(a AS int)
// Go: psql.Cast("a", "int")
func Cast(exp pgcraft.Expression, typname string) Expression {
	return bmod.Cast(exp, typname)
}
