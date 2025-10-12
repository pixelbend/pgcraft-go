package clause

import (
	"github.com/arkamfahry/pgcraft-go"
	"io"
)

type Conflict struct {
	Do     string // DO NOTHING | DO UPDATE
	Target ConflictTarget
	Set
	Where
}

func (c *Conflict) SetConflict(conflict Conflict) {
	*c = conflict
}

func (c Conflict) WriteSQL(w io.Writer, start int) ([]any, error) {
	w.Write([]byte("ON CONFLICT"))

	args, err := pgcraft.ExpressIf(w, start, c.Target, true, "", "")
	if err != nil {
		return nil, err
	}

	w.Write([]byte(" DO "))
	w.Write([]byte(c.Do))

	setArgs, err := pgcraft.ExpressIf(w, start+len(args), c.Set, len(c.Set.Set) > 0, " SET\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, setArgs...)

	whereArgs, err := pgcraft.ExpressIf(w, start+len(args), c.Where,
		len(c.Where.Conditions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, whereArgs...)

	return args, nil
}

type ConflictTarget struct {
	Constraint string
	Columns    []any
	Where      []any
}

func (c ConflictTarget) WriteSQL(w io.Writer, start int) ([]any, error) {
	if c.Constraint != "" {
		return pgcraft.ExpressIf(w, start, c.Constraint, true, " ON CONSTRAINT ", "")
	}

	args, err := pgcraft.ExpressSlice(w, start, c.Columns, " (", ", ", ")")
	if err != nil {
		return nil, err
	}

	whereArgs, err := pgcraft.ExpressSlice(w, start+len(args), c.Where, " WHERE ", " AND ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, whereArgs...)

	return args, nil
}
