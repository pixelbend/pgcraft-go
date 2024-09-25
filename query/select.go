package query

import (
	"github.com/driftdev/pgcraft"
	"github.com/driftdev/pgcraft/clause"
	"io"
)

type Select struct {
	clause.With
	clause.SelectList
	Distinct
	clause.From
	clause.Where
	clause.GroupBy
	clause.Having
	clause.Windows
	clause.Combine
	clause.OrderBy
	clause.Limit
	clause.Offset
	clause.Fetch
	clause.For
}

func (s Select) WriteSQL(w io.Writer, start int) ([]any, error) {
	var args []any

	withArgs, err := pgcraft.ExpressIf(w, start+len(args), s.With,
		len(s.With.CTEs) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, withArgs...)

	w.Write([]byte("SELECT "))

	distinctArgs, err := pgcraft.ExpressIf(w, start+len(args), s.Distinct,
		s.Distinct.On != nil, "", " ")
	if err != nil {
		return nil, err
	}
	args = append(args, distinctArgs...)

	selArgs, err := pgcraft.ExpressIf(w, start+len(args), s.SelectList, true, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, selArgs...)

	fromArgs, err := pgcraft.ExpressIf(w, start+len(args), s.From, s.From.Table != nil, "\nFROM ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, fromArgs...)

	whereArgs, err := pgcraft.ExpressIf(w, start+len(args), s.Where,
		len(s.Where.Conditions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, whereArgs...)

	groupByArgs, err := pgcraft.ExpressIf(w, start+len(args), s.GroupBy,
		len(s.GroupBy.Groups) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, groupByArgs...)

	havingArgs, err := pgcraft.ExpressIf(w, start+len(args), s.Having,
		len(s.Having.Conditions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, havingArgs...)

	windowArgs, err := pgcraft.ExpressIf(w, start+len(args), s.Windows,
		len(s.Windows.Windows) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, windowArgs...)

	combineArgs, err := pgcraft.ExpressIf(w, start+len(args), s.Combine,
		s.Combine.Query != nil, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, combineArgs...)

	orderArgs, err := pgcraft.ExpressIf(w, start+len(args), s.OrderBy,
		len(s.OrderBy.Expressions) > 0, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, orderArgs...)

	limitArgs, err := pgcraft.ExpressIf(w, start+len(args), s.Limit,
		s.Limit.Count != nil, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, limitArgs...)

	offsetArgs, err := pgcraft.ExpressIf(w, start+len(args), s.Offset,
		s.Offset.Count != nil, "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, offsetArgs...)

	_, err = pgcraft.ExpressIf(w, start+len(args), s.Fetch,
		s.Fetch.Count != nil, "\n", "")
	if err != nil {
		return nil, err
	}

	forArgs, err := pgcraft.ExpressIf(w, start+len(args), s.For,
		s.For.Strength != "", "\n", "")
	if err != nil {
		return nil, err
	}
	args = append(args, forArgs...)

	w.Write([]byte("\n"))
	return args, nil
}
