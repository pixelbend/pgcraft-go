package query

import (
	"github.com/teapartydev/pgcraft-go"
	"github.com/teapartydev/pgcraft-go/clause"
	"io"
)

type Delete struct {
	clause.With
	Only bool
	clause.Table
	clause.From
	clause.Where
	clause.Returning
}

func (d Delete) WriteSQL(w io.Writer, start int) ([]any, error) {
	var args []any

	withArgs, err := pgcraft.ExpressIf(w, start+len(args), d.With,
		len(d.With.CTEs) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, withArgs...)

	w.Write([]byte("DELETE FROM "))

	if d.Only {
		w.Write([]byte("ONLY "))
	}

	tableArgs, err := pgcraft.ExpressIf(w, start+len(args), d.Table, true, "", "")
	if err != nil {
		return nil, err
	}
	args = append(args, tableArgs...)

	usingArgs, err := pgcraft.ExpressIf(w, start+len(args), d.From,
		d.From.Table != nil, "\nUSING ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, usingArgs...)

	whereArgs, err := pgcraft.ExpressIf(w, start+len(args), d.Where,
		len(d.Where.Conditions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, whereArgs...)

	retArgs, err := pgcraft.ExpressIf(w, start+len(args), d.Returning,
		len(d.Returning.Expressions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, retArgs...)

	return args, nil
}
