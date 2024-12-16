package expr

import (
	"github.com/pixelbend/pgcraft-go"
	"io"
)

func Cast(e pgcraft.Expression, typname string) pgcraft.Expression {
	return cast{e: e, typname: typname}
}

type cast struct {
	e       pgcraft.Expression
	typname string
}

func (c cast) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.ExpressIf(w, start, c.e, c.e != nil, "CAST(", " AS "+c.typname+")")
}
