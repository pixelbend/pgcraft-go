package clause

import (
	"github.com/driftdev/pgcraft"
	"io"
)

type Values struct {
	// Query takes the highest priority
	// If present, will attempt to insert from this query
	Query pgcraft.Query

	// for multiple inserts
	// each sub-slice is one set of values
	Vals []value
}

type value []pgcraft.Expression

func (v value) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.ExpressSlice(w, start, v, "(", ", ", ")")
}

func (v *Values) AppendValues(vals ...pgcraft.Expression) {
	if len(vals) == 0 {
		return
	}

	v.Vals = append(v.Vals, vals)
}

func (v Values) WriteSQL(w io.Writer, start int) ([]any, error) {
	// If a query is present, use it
	if v.Query != nil {
		return v.Query.WriteQuery(w, start)
	}

	// If values are present, use them
	if len(v.Vals) > 0 {
		return pgcraft.ExpressSlice(w, start, v.Vals, "VALUES ", ", ", "")
	}

	// If no value was present, use default value
	w.Write([]byte("DEFAULT VALUES"))
	return nil, nil
}
