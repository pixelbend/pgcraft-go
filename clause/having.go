package clause

import (
	"github.com/pixelbend/pgcraft-go"
	"io"
)

type Having struct {
	Conditions []any
}

func (h *Having) AppendHaving(e ...any) {
	h.Conditions = append(h.Conditions, e...)
}

func (h Having) WriteSQL(w io.Writer, start int) ([]any, error) {
	args, err := pgcraft.ExpressSlice(w, start, h.Conditions, "HAVING ", " AND ", "")
	if err != nil {
		return nil, err
	}

	return args, nil
}
