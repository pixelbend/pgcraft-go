package expr

import (
	"io"
	"strconv"
)

var dollar = []byte("$")

func WriteArg(w io.Writer, position int) {
	w.Write(dollar)
	w.Write([]byte(strconv.Itoa(position)))
}

func Arg(vals ...any) pgcraft.Expression {
	return args{vals: vals}
}

func ArgGroup(vals ...any) pgcraft.Expression {
	return args{vals: vals, grouped: true}
}

type args struct {
	vals    []any
	grouped bool
}

func (a args) WriteSQL(w io.Writer, start int) ([]any, error) {
	if len(a.vals) == 0 {
		return nil, nil
	}

	if a.grouped {
		w.Write([]byte(openPar))
	}

	for k := range a.vals {
		if k > 0 {
			w.Write([]byte(commaSpace))
		}

		WriteArg(w, start+k)
	}

	if a.grouped {
		w.Write([]byte(closePar))
	}

	return a.vals, nil
}
