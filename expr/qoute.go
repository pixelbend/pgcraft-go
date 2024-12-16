package expr

import (
	"github.com/pixelbend/pgcraft-go"
	"io"
)

var doubleQuote = []byte(`"`)

func WriteQuote(w io.Writer, s string) {
	w.Write(doubleQuote)
	w.Write([]byte(s))
	w.Write(doubleQuote)
}

func Quote(aa ...string) pgcraft.Expression {
	ss := make([]string, 0, len(aa))
	for _, v := range aa {
		if v == "" {
			continue
		}
		ss = append(ss, v)
	}

	return quoted(ss)
}

// quoted and joined... something like "users"."id"
type quoted []string

func (q quoted) WriteSQL(w io.Writer, start int) ([]any, error) {
	if len(q) == 0 {
		return nil, nil
	}

	// wrap in parenthesis and join with comma
	k := 0 // not using the loop index to avoid empty strings
	for _, a := range q {
		if a == "" {
			continue
		}

		if k != 0 {
			w.Write([]byte("."))
		}
		k++

		WriteQuote(w, a)
	}

	return nil, nil
}
