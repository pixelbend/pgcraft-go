package clause

import (
	"github.com/codefrantic/pgcraft-go"
	"github.com/codefrantic/pgcraft-go/expr"
	"io"
)

type Table struct {
	Expression any
	Alias      string
	Columns    []string
}

func (t Table) As(alias string, columns ...string) Table {
	t.Alias = alias
	t.Columns = append(t.Columns, columns...)

	return t
}

func (t Table) WriteSQL(w io.Writer, start int) ([]any, error) {
	args, err := pgcraft.Express(w, start, t.Expression)
	if err != nil {
		return nil, err
	}

	if t.Alias != "" {
		w.Write([]byte(" AS "))
		expr.WriteQuote(w, t.Alias)
	}

	if len(t.Columns) > 0 {
		w.Write([]byte(" ("))
		for k, cAlias := range t.Columns {
			if k != 0 {
				w.Write([]byte(", "))
			}

			expr.WriteQuote(w, cAlias)
		}
		w.Write([]byte(")"))
	}

	return args, nil
}
