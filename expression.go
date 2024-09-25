package pgcraft

import (
	"fmt"
	"io"
)

type Expression interface {
	WriteSQL(w io.Writer, start int) (args []any, err error)
}

type ExpressionFunc func(w io.Writer, start int) ([]any, error)

func (e ExpressionFunc) WriteSQL(w io.Writer, start int) ([]any, error) {
	return e(w, start)
}

func Express(w io.Writer, start int, e any) ([]any, error) {
	switch v := e.(type) {
	case string:
		w.Write([]byte(v))
		return nil, nil
	case []byte:
		w.Write(v)
		return nil, nil
	case Expression:
		return v.WriteSQL(w, start)
	default:
		fmt.Fprint(w, e)
		return nil, nil
	}
}

// ExpressIf expands an express if the condition evaluates to true
// it can also add a prefix and suffix
func ExpressIf(w io.Writer, start int, e any, cond bool, prefix, suffix string) ([]any, error) {
	if !cond {
		return nil, nil
	}

	w.Write([]byte(prefix))
	args, err := Express(w, start, e)
	if err != nil {
		return nil, err
	}
	w.Write([]byte(suffix))

	return args, nil
}

// ExpressSlice is used to express a slice of expressions along with a prefix and suffix
func ExpressSlice[T any](w io.Writer, start int, expressions []T, prefix, sep, suffix string) ([]any, error) {
	if len(expressions) == 0 {
		return nil, nil
	}

	var args []any
	w.Write([]byte(prefix))

	for k, e := range expressions {
		if k != 0 {
			w.Write([]byte(sep))
		}

		newArgs, err := Express(w, start+len(args), e)
		if err != nil {
			return args, err
		}

		args = append(args, newArgs...)
	}
	w.Write([]byte(suffix))

	return args, nil
}
