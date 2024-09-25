package pgcraft

import (
	"github.com/qdm12/reprint"
	"io"
)

const (
	openPar  = "("
	closePar = ")"
)

type Query interface {
	Expression
	WriteQuery(writer io.Writer, start int) (args []any, err error)
}

type Mod[T any] interface {
	Apply(T)
}

type BaseQuery[E Expression] struct {
	Expression E
}

func (b BaseQuery[E]) Clone() BaseQuery[E] {
	if c, ok := any(b.Expression).(interface{ Clone() E }); ok {
		return BaseQuery[E]{
			Expression: c.Clone(),
		}
	}

	return BaseQuery[E]{
		Expression: reprint.This(b.Expression).(E),
	}
}

func (b BaseQuery[E]) Apply(mods ...Mod[E]) {
	for _, mod := range mods {
		mod.Apply(b.Expression)
	}
}

func (b BaseQuery[E]) WriteQuery(w io.Writer, start int) ([]any, error) {
	return b.Expression.WriteSQL(w, start)
}

func (b BaseQuery[E]) WriteSQL(w io.Writer, start int) ([]any, error) {
	w.Write([]byte(openPar))
	args, err := b.Expression.WriteSQL(w, start)
	w.Write([]byte(closePar))

	return args, err
}

// MustBuild builds a query form start and panics on error
// useful for initializing queries that need to be reused
func (q BaseQuery[E]) MustBuild() (string, []any) {
	return MustBuildN(q, 1)
}

// MustBuildN builds a query from a point and panics on error
// useful for initializing queries that need to be reused
func (q BaseQuery[E]) MustBuildN(start int) (string, []any) {
	return MustBuildN(q, start)
}

// Build Convenient function to build query from start
func (q BaseQuery[E]) Build() (string, []any, error) {
	return BuildN(q, 1)
}

// BuildN Convenient function to build query from a point
func (q BaseQuery[E]) BuildN(start int) (string, []any, error) {
	return BuildN(q, start)
}

// Cache Convenient function to cache a query from start
func (q BaseQuery[E]) Cache() (BaseQuery[*cached], error) {
	return CacheN(q, 1)
}

// CacheN Convenient function to cache a query from a point
func (q BaseQuery[E]) CacheN(start int) (BaseQuery[*cached], error) {
	return CacheN(q, start)
}
