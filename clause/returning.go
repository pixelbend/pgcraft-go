package clause

import (
	"github.com/arkamfahry/pgcraft-go"
	"io"
)

type Returning struct {
	Expressions []any
}

func (r *Returning) HasReturning() bool {
	return len(r.Expressions) > 0
}

func (r *Returning) AppendReturning(columns ...any) {
	r.Expressions = append(r.Expressions, columns...)
}

func (r Returning) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.ExpressSlice(w, start, r.Expressions, "RETURNING ", ", ", "")
}
