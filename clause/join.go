package clause

import (
	"github.com/teapartydev/pgcraft-go"
	"github.com/teapartydev/pgcraft-go/expr"
	"io"
)

const (
	InnerJoin = "INNER JOIN"
	LeftJoin  = "LEFT JOIN"
	RightJoin = "RIGHT JOIN"
	FullJoin  = "FULL JOIN"
	CrossJoin = "CROSS JOIN"
)

type Join struct {
	Type string
	To   From

	Natural bool
	On      []pgcraft.Expression
	Using   []string
}

func (j Join) WriteSQL(w io.Writer, start int) ([]any, error) {
	if j.Natural {
		w.Write([]byte("NATURAL "))
	}

	w.Write([]byte(j.Type))
	w.Write([]byte(" "))

	args, err := pgcraft.Express(w, start, j.To)
	if err != nil {
		return nil, err
	}

	onArgs, err := pgcraft.ExpressSlice(w, start+len(args), j.On, " ON ", " AND ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, onArgs...)

	for k, col := range j.Using {
		if k == 0 {
			w.Write([]byte(" USING("))
		} else {
			w.Write([]byte(", "))
		}

		_, err = expr.Quote(col).WriteSQL(w, 1) // start does not matter
		if err != nil {
			return nil, err
		}

		if k == len(j.Using)-1 {
			w.Write([]byte(") "))
		}

	}

	return args, nil
}
