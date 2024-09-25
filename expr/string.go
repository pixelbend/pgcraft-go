package expr

import (
	"io"
)

type rawString string

func (s rawString) WriteSQL(w io.Writer, start int) ([]any, error) {
	w.Write([]byte("'"))
	w.Write([]byte(s))
	w.Write([]byte("'"))

	return nil, nil
}
