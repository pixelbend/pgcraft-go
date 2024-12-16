package mods

import (
	"github.com/pixelbend/pgcraft-go"
	"github.com/pixelbend/pgcraft-go/clause"
)

type Conflict[Q interface{ SetConflict(clause.Conflict) }] func() clause.Conflict

func (s Conflict[Q]) Apply(q Q) {
	q.SetConflict(s())
}

func (c Conflict[Q]) Where(where ...any) Conflict[Q] {
	conflict := c()
	conflict.Target.Where = append(conflict.Target.Where, where...)

	return func() clause.Conflict {
		return conflict
	}
}

func (c Conflict[Q]) DoNothing() pgcraft.Mod[Q] {
	conflict := c()
	conflict.Do = "NOTHING"

	return Conflict[Q](func() clause.Conflict {
		return conflict
	})
}

func (c Conflict[Q]) DoUpdate(sets ...pgcraft.Mod[*clause.Conflict]) pgcraft.Mod[Q] {
	conflict := c()
	conflict.Do = "UPDATE"

	for _, set := range sets {
		set.Apply(&conflict)
	}

	return Conflict[Q](func() clause.Conflict {
		return conflict
	})
}
