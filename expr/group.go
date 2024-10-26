package expr

import (
	"io"
)

type group []pgcraft.Expression

func (g group) WriteSQL(w io.Writer, start int) ([]any, error) {
	if len(g) == 0 {
		return pgcraft.ExpressIf(w, start, null, true, openPar, closePar)
	}

	return pgcraft.ExpressSlice(w, start, g, openPar, commaSpace, closePar)
}
