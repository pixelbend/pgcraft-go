package clause

import (
	"github.com/codefrantic/pgcraft-go"
	"github.com/codefrantic/pgcraft-go/expr"
	"io"
)

type From struct {
	Table any

	// Aliases
	Alias   string
	Columns []string

	Only           bool
	Lateral        bool
	WithOrdinality bool

	// Joins
	Joins []Join
}

func (f *From) SetTable(table any) {
	f.Table = table
}

func (f *From) SetTableAlias(alias string, columns ...string) {
	f.Alias = alias
	f.Columns = columns
}

func (f *From) SetOnly(only bool) {
	f.Only = only
}

func (f *From) SetLateral(lateral bool) {
	f.Lateral = lateral
}

func (f *From) SetWithOrdinality(to bool) {
	f.WithOrdinality = to
}

func (f *From) AppendJoin(j Join) {
	f.Joins = append(f.Joins, j)
}

func (f From) WriteSQL(w io.Writer, start int) ([]any, error) {
	if f.Table == nil {
		return nil, nil
	}

	if f.Only {
		w.Write([]byte("ONLY "))
	}

	if f.Lateral {
		w.Write([]byte("LATERAL "))
	}

	args, err := pgcraft.Express(w, start, f.Table)
	if err != nil {
		return nil, err
	}

	if f.WithOrdinality {
		w.Write([]byte(" WITH ORDINALITY"))
	}

	if f.Alias != "" {
		w.Write([]byte(" AS "))
		expr.WriteQuote(w, f.Alias)
	}

	if len(f.Columns) > 0 {
		w.Write([]byte("("))
		for k, cAlias := range f.Columns {
			if k != 0 {
				w.Write([]byte(", "))
			}

			expr.WriteQuote(w, cAlias)
		}
		w.Write([]byte(")"))
	}

	joinArgs, err := pgcraft.ExpressSlice(w, start+len(args), f.Joins, "\n", "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, joinArgs...)

	return args, nil
}
