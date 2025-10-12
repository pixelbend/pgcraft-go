package clause

import (
	"github.com/arkamfahry/pgcraft-go"
	"io"
)

type Limit struct {
	Count any
}

func (l *Limit) SetLimit(limit any) {
	l.Count = limit
}

func (l Limit) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.ExpressIf(w, start, l.Count, l.Count != nil, "LIMIT ", "")
}
