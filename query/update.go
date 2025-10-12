package query

import (
	"github.com/arkamfahry/pgcraft-go"
	"github.com/arkamfahry/pgcraft-go/clause"
	"io"
)

type Update struct {
	clause.With
	Only bool
	clause.Table
	clause.Set
	clause.From
	clause.Where
	clause.Returning
}

func (u Update) WriteSQL(w io.Writer, start int) ([]any, error) {
	var args []any

	withArgs, err := pgcraft.ExpressIf(w, start+len(args), u.With,
		len(u.With.CTEs) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, withArgs...)

	w.Write([]byte("UPDATE "))

	if u.Only {
		w.Write([]byte("ONLY "))
	}

	tableArgs, err := pgcraft.ExpressIf(w, start+len(args), u.Table, true, "", "")
	if err != nil {
		return nil, err
	}
	args = append(args, tableArgs...)

	setArgs, err := pgcraft.ExpressIf(w, start+len(args), u.Set, true, " SET\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, setArgs...)

	fromArgs, err := pgcraft.ExpressIf(w, start+len(args), u.From,
		u.From.Table != nil, "\nFROM ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, fromArgs...)

	whereArgs, err := pgcraft.ExpressIf(w, start+len(args), u.Where,
		len(u.Where.Conditions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, whereArgs...)

	retArgs, err := pgcraft.ExpressIf(w, start+len(args), u.Returning,
		len(u.Returning.Expressions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, retArgs...)

	return args, nil
}
