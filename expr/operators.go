package expr

import (
	"fmt"
	"github.com/teapartydev/pgcraft-go"
	"io"
)

type leftRight struct {
	operator string
	right    any
	left     any
}

func (lr leftRight) WriteSQL(w io.Writer, start int) ([]any, error) {
	largs, err := pgcraft.Express(w, start, lr.left)
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(w, " %s ", lr.operator)

	rargs, err := pgcraft.Express(w, start+len(largs), lr.right)
	if err != nil {
		return nil, err
	}

	return append(largs, rargs...), nil
}

// Generic operator between a left and right val
func OP(operator string, left, right any) pgcraft.Expression {
	return leftRight{
		right:    right,
		left:     left,
		operator: operator,
	}
}

// If no separator, a space is used
type Join struct {
	Exprs []pgcraft.Expression
	Sep   string
}

func (s Join) WriteSQL(w io.Writer, start int) ([]any, error) {
	sep := s.Sep
	if sep == "" {
		sep = " "
	}

	return pgcraft.ExpressSlice(w, start, s.Exprs, "", sep, "")
}
