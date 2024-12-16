package clause

import (
	"github.com/pixelbend/pgcraft-go"
	"io"
)

type Offset struct {
	Count any
}

func (o *Offset) SetOffset(offset any) {
	o.Count = offset
}

func (o Offset) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.ExpressIf(w, start, o.Count, o.Count != nil, "OFFSET ", "")
}
