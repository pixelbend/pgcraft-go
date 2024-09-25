package clause

import (
	"github.com/driftdev/pgcraft"
	"io"
)

type Where struct {
	Conditions []any
}

func (wh *Where) AppendWhere(e ...any) {
	wh.Conditions = append(wh.Conditions, e...)
}

func (wh Where) WriteSQL(w io.Writer, start int) ([]any, error) {
	args, err := pgcraft.ExpressSlice(w, start, wh.Conditions, "WHERE ", " AND ", "")
	if err != nil {
		return nil, err
	}

	return args, nil
}
