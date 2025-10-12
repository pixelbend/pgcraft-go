package clause

import (
	"github.com/arkamfahry/pgcraft-go"
	"io"
)

type Set struct {
	Set []any
}

func (s *Set) AppendSet(exprs ...any) {
	s.Set = append(s.Set, exprs...)
}

func (s Set) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.ExpressSlice(w, start, s.Set, "", ",\n", "")
}
