package clause

import (
	"fmt"
	"github.com/pixelbend/pgcraft-go"
	"io"
)

type OrderBy struct {
	Expressions []OrderDef
}

func (o *OrderBy) SetOrderBy(orders ...OrderDef) {
	o.Expressions = orders
}

func (o *OrderBy) AppendOrder(order OrderDef) {
	o.Expressions = append(o.Expressions, order)
}

func (o OrderBy) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.ExpressSlice(w, start, o.Expressions, "ORDER BY ", ", ", "")
}

type OrderDef struct {
	Expression any
	Direction  string // ASC | DESC | USING operator
	Nulls      string // FIRST | LAST
	Collation  pgcraft.Expression
}

func (o OrderDef) WriteSQL(w io.Writer, start int) ([]any, error) {
	args, err := pgcraft.Express(w, start, o.Expression)
	if err != nil {
		return nil, err
	}

	if o.Collation != nil {
		_, err = o.Collation.WriteSQL(w, start)
		if err != nil {
			return nil, err
		}
	}

	if o.Direction != "" {
		w.Write([]byte(" "))
		w.Write([]byte(o.Direction))
	}

	if o.Nulls != "" {
		fmt.Fprintf(w, " NULLS %s", o.Nulls)
	}

	return args, nil
}
