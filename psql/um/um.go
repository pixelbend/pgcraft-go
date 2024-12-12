package um

import (
	"github.com/codefrantic/pgcraft-go"
	"github.com/codefrantic/pgcraft-go/clause"
	"github.com/codefrantic/pgcraft-go/internal"
	"github.com/codefrantic/pgcraft-go/mods"
	"github.com/codefrantic/pgcraft-go/query"
)

func With(name string, columns ...string) query.CTEChain[*query.Update] {
	return query.With[*query.Update](name, columns...)
}

func Recursive(r bool) pgcraft.Mod[*query.Update] {
	return mods.Recursive[*query.Update](r)
}

func Only() pgcraft.Mod[*query.Update] {
	return mods.QueryModFunc[*query.Update](func(u *query.Update) {
		u.Only = true
	})
}

func Table(name any) pgcraft.Mod[*query.Update] {
	return mods.QueryModFunc[*query.Update](func(u *query.Update) {
		u.Table = clause.Table{
			Expression: name,
		}
	})
}

func TableAs(name any, alias string) pgcraft.Mod[*query.Update] {
	return mods.QueryModFunc[*query.Update](func(u *query.Update) {
		u.Table = clause.Table{
			Expression: name,
			Alias:      alias,
		}
	})
}

func Set(sets ...pgcraft.Expression) pgcraft.Mod[*query.Update] {
	return mods.QueryModFunc[*query.Update](func(q *query.Update) {
		q.Set.Set = append(q.Set.Set, internal.ToAnySlice(sets)...)
	})
}

func SetCol(from string) mods.Set[*query.Update] {
	return mods.Set[*query.Update]([]string{from})
}

func From(table any) query.FromChain[*query.Update] {
	return query.From[*query.Update](table)
}

func FromFunction(funcs ...*query.Function) query.FromChain[*query.Update] {
	var table any

	if len(funcs) == 1 {
		table = funcs[0]
	}

	if len(funcs) > 1 {
		table = query.Functions(funcs)
	}

	return query.From[*query.Update](table)
}

func InnerJoin(e any) query.JoinChain[*query.Update] {
	return query.InnerJoin[*query.Update](e)
}

func LeftJoin(e any) query.JoinChain[*query.Update] {
	return query.LeftJoin[*query.Update](e)
}

func RightJoin(e any) query.JoinChain[*query.Update] {
	return query.RightJoin[*query.Update](e)
}

func FullJoin(e any) query.JoinChain[*query.Update] {
	return query.FullJoin[*query.Update](e)
}

func CrossJoin(e any) pgcraft.Mod[*query.Update] {
	return query.CrossJoin[*query.Update](e)
}

func Where(e pgcraft.Expression) mods.Where[*query.Update] {
	return mods.Where[*query.Update]{E: e}
}

func Returning(clauses ...any) pgcraft.Mod[*query.Update] {
	return mods.Returning[*query.Update](clauses)
}
