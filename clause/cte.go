package clause

import (
	"fmt"
	"github.com/teapartydev/pgcraft-go"
	"io"
)

type CTE struct {
	Query        pgcraft.Query
	Name         string
	Columns      []string
	Materialized *bool
	Search       CTESearch
	Cycle        CTECycle
}

func (c CTE) WriteSQL(w io.Writer, start int) ([]any, error) {
	w.Write([]byte(c.Name))
	_, err := pgcraft.ExpressSlice(w, start, c.Columns, "(", ", ", ")")
	if err != nil {
		return nil, err
	}

	w.Write([]byte(" AS "))

	switch {
	case c.Materialized == nil:
		// do nothing
		break
	case *c.Materialized:
		w.Write([]byte("MATERIALIZED "))
	case !*c.Materialized:
		w.Write([]byte("NOT MATERIALIZED "))
	}

	w.Write([]byte("("))
	args, err := c.Query.WriteQuery(w, start)
	if err != nil {
		return nil, err
	}
	w.Write([]byte(")"))

	searchArgs, err := pgcraft.ExpressIf(w, start+len(args), c.Search,
		len(c.Search.Columns) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, searchArgs...)

	cycleArgs, err := pgcraft.ExpressIf(w, start+len(args), c.Cycle,
		len(c.Cycle.Columns) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, cycleArgs...)

	return args, nil
}

const (
	SearchBreadth = "BREADTH"
	SearchDepth   = "DEPTH"
)

type CTESearch struct {
	Order   string
	Columns []string
	Set     string
}

func (c CTESearch) WriteSQL(w io.Writer, start int) ([]any, error) {
	// [ SEARCH { BREADTH | DEPTH } FIRST BY column_name [, ...] SET search_seq_col_name ]
	fmt.Fprintf(w, "SEARCH %s FIRST BY ", c.Order)

	args, err := pgcraft.ExpressSlice(w, start, c.Columns, "", ", ", "")
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(w, " SET %s", c.Set)

	return args, nil
}

type CTECycle struct {
	Columns    []string
	Set        string
	Using      string
	SetVal     any
	DefaultVal any
}

func (c CTECycle) WriteSQL(w io.Writer, start int) ([]any, error) {
	//[ CYCLE column_name [, ...] SET cycle_mark_col_name [ TO cycle_mark_value DEFAULT cycle_mark_default ] USING cycle_path_col_name ]
	w.Write([]byte("CYCLE "))

	args, err := pgcraft.ExpressSlice(w, start, c.Columns, "", ", ", "")
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(w, " SET %s", c.Set)

	markArgs, err := pgcraft.ExpressIf(w, start+len(args), c.SetVal,
		c.SetVal != nil, " TO ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, markArgs...)

	defaultArgs, err := pgcraft.ExpressIf(w, start+len(args), c.DefaultVal,
		c.DefaultVal != nil, " DEFAULT ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, defaultArgs...)

	fmt.Fprintf(w, " USING %s", c.Using)

	return args, nil
}
