package query

import (
	"github.com/driftdev/pgcraft-go"
	"github.com/driftdev/pgcraft-go/clause"
	"io"
)

type Insert struct {
	clause.With
	Overriding string
	clause.Table
	clause.Values
	clause.Conflict
	clause.Returning
}

func (i Insert) WriteSQL(w io.Writer, start int) ([]any, error) {
	var args []any

	withArgs, err := pgcraft.ExpressIf(w, start+len(args), i.With,
		len(i.With.CTEs) > 0, "", "\n")
	if err != nil {
		return nil, err
	}
	args = append(args, withArgs...)

	tableArgs, err := pgcraft.ExpressIf(w, start+len(args), i.Table,
		true, "INSERT INTO ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, tableArgs...)

	_, err = pgcraft.ExpressIf(w, start+len(args), i.Overriding,
		i.Overriding != "", "\nOVERRIDING ", " VALUE")
	if err != nil {
		return nil, err
	}

	valArgs, err := pgcraft.ExpressIf(w, start+len(args), i.Values, true, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, valArgs...)

	conflictArgs, err := pgcraft.ExpressIf(w, start+len(args), i.Conflict,
		i.Conflict.Do != "", "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, conflictArgs...)

	retArgs, err := pgcraft.ExpressIf(w, start+len(args), i.Returning,
		len(i.Returning.Expressions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, retArgs...)

	w.Write([]byte("\n"))
	return args, nil
}
