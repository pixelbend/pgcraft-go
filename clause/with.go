package clause

import (
	"github.com/arkamfahry/pgcraft-go"
	"io"
)

type With struct {
	Recursive bool
	CTEs      []CTE
}

func (w *With) AppendWith(cte CTE) {
	w.CTEs = append(w.CTEs, cte)
}

func (w *With) SetRecursive(r bool) {
	w.Recursive = r
}

func (w With) WriteSQL(wr io.Writer, start int) ([]any, error) {
	prefix := "WITH\n"
	if w.Recursive {
		prefix = "WITH RECURSIVE\n"
	}
	return pgcraft.ExpressSlice(wr, start, w.CTEs, prefix, ",\n", "")
}
