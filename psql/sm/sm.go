package sm

import (
	"github.com/driftdev/pgcraft"
	"github.com/driftdev/pgcraft/clause"
	"github.com/driftdev/pgcraft/mods"
	"github.com/driftdev/pgcraft/query"
)

func With(name string, columns ...string) query.CTEChain[*query.Select] {
	return query.With[*query.Select](name, columns...)
}

func Recursive(r bool) pgcraft.Mod[*query.Select] {
	return mods.Recursive[*query.Select](r)
}

func Distinct(on ...any) pgcraft.Mod[*query.Select] {
	if on == nil {
		on = []any{} // nil means no distinct
	}

	return mods.QueryModFunc[*query.Select](func(q *query.Select) {
		q.Distinct.On = on
	})
}

func Columns(clauses ...any) pgcraft.Mod[*query.Select] {
	return mods.Select[*query.Select](clauses)
}

func From(table any) query.FromChain[*query.Select] {
	return query.From[*query.Select](table)
}

func FromFunction(funcs ...*query.Function) query.FromChain[*query.Select] {
	var table any

	if len(funcs) == 1 {
		table = funcs[0]
	}

	if len(funcs) > 1 {
		table = query.Functions(funcs)
	}

	return query.From[*query.Select](table)
}

func InnerJoin(e any) query.JoinChain[*query.Select] {
	return query.InnerJoin[*query.Select](e)
}

func LeftJoin(e any) query.JoinChain[*query.Select] {
	return query.LeftJoin[*query.Select](e)
}

func RightJoin(e any) query.JoinChain[*query.Select] {
	return query.RightJoin[*query.Select](e)
}

func FullJoin(e any) query.JoinChain[*query.Select] {
	return query.FullJoin[*query.Select](e)
}

func CrossJoin(e any) pgcraft.Mod[*query.Select] {
	return query.CrossJoin[*query.Select](e)
}

func Where(e pgcraft.Expression) mods.Where[*query.Select] {
	return mods.Where[*query.Select]{E: e}
}

func Having(e any) pgcraft.Mod[*query.Select] {
	return mods.Having[*query.Select]{e}
}

func GroupBy(e any) pgcraft.Mod[*query.Select] {
	return mods.GroupBy[*query.Select]{
		E: e,
	}
}

func GroupByDistinct(distinct bool) pgcraft.Mod[*query.Select] {
	return mods.GroupByDistinct[*query.Select](distinct)
}

func Window(name string) query.WindowsMod[*query.Select] {
	m := query.WindowsMod[*query.Select]{
		Name: name,
	}

	m.WindowChain = &query.WindowChain[*query.WindowsMod[*query.Select]]{
		Wrap: &m,
	}
	return m
}

func OrderBy(e any) query.OrderBy[*query.Select] {
	return query.OrderBy[*query.Select](func() clause.OrderDef {
		return clause.OrderDef{
			Expression: e,
		}
	})
}

func Limit(count any) pgcraft.Mod[*query.Select] {
	return mods.Limit[*query.Select]{
		Count: count,
	}
}

func Offset(count any) pgcraft.Mod[*query.Select] {
	return mods.Offset[*query.Select]{
		Count: count,
	}
}

func Fetch(count int64, withTies bool) pgcraft.Mod[*query.Select] {
	return mods.Fetch[*query.Select]{
		Count:    &count,
		WithTies: withTies,
	}
}

func Union(q pgcraft.Query) pgcraft.Mod[*query.Select] {
	return mods.Combine[*query.Select]{
		Strategy: clause.Union,
		Query:    q,
		All:      false,
	}
}

func UnionAll(q pgcraft.Query) pgcraft.Mod[*query.Select] {
	return mods.Combine[*query.Select]{
		Strategy: clause.Union,
		Query:    q,
		All:      true,
	}
}

func Intersect(q pgcraft.Query) pgcraft.Mod[*query.Select] {
	return mods.Combine[*query.Select]{
		Strategy: clause.Intersect,
		Query:    q,
		All:      false,
	}
}

func IntersectAll(q pgcraft.Query) pgcraft.Mod[*query.Select] {
	return mods.Combine[*query.Select]{
		Strategy: clause.Intersect,
		Query:    q,
		All:      true,
	}
}

func Except(q pgcraft.Query) pgcraft.Mod[*query.Select] {
	return mods.Combine[*query.Select]{
		Strategy: clause.Except,
		Query:    q,
		All:      false,
	}
}

func ExceptAll(q pgcraft.Query) pgcraft.Mod[*query.Select] {
	return mods.Combine[*query.Select]{
		Strategy: clause.Except,
		Query:    q,
		All:      true,
	}
}

func ForUpdate(tables ...string) query.LockChain[*query.Select] {
	return query.LockChain[*query.Select](func() clause.For {
		return clause.For{
			Strength: clause.LockStrengthUpdate,
			Tables:   tables,
		}
	})
}

func ForNoKeyUpdate(tables ...string) query.LockChain[*query.Select] {
	return query.LockChain[*query.Select](func() clause.For {
		return clause.For{
			Strength: clause.LockStrengthNoKeyUpdate,
			Tables:   tables,
		}
	})
}

func ForShare(tables ...string) query.LockChain[*query.Select] {
	return query.LockChain[*query.Select](func() clause.For {
		return clause.For{
			Strength: clause.LockStrengthShare,
			Tables:   tables,
		}
	})
}

func ForKeyShare(tables ...string) query.LockChain[*query.Select] {
	return query.LockChain[*query.Select](func() clause.For {
		return clause.For{
			Strength: clause.LockStrengthKeyShare,
			Tables:   tables,
		}
	})
}
