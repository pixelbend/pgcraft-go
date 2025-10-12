package dm

import (
	"github.com/arkamfahry/pgcraft-go"
	"github.com/arkamfahry/pgcraft-go/clause"
	"github.com/arkamfahry/pgcraft-go/mods"
	"github.com/arkamfahry/pgcraft-go/query"
)

func With(name string, columns ...string) query.CTEChain[*query.Delete] {
	return query.With[*query.Delete](name, columns...)
}

func Recursive(r bool) pgcraft.Mod[*query.Delete] {
	return mods.Recursive[*query.Delete](r)
}

func Only() pgcraft.Mod[*query.Delete] {
	return mods.QueryModFunc[*query.Delete](func(d *query.Delete) {
		d.Only = true
	})
}

func From(name any) pgcraft.Mod[*query.Delete] {
	return mods.QueryModFunc[*query.Delete](func(u *query.Delete) {
		u.Table = clause.Table{
			Expression: name,
		}
	})
}

func FromAs(name any, alias string) pgcraft.Mod[*query.Delete] {
	return mods.QueryModFunc[*query.Delete](func(u *query.Delete) {
		u.Table = clause.Table{
			Expression: name,
			Alias:      alias,
		}
	})
}

func Using(table any) query.FromChain[*query.Delete] {
	return query.From[*query.Delete](table)
}

func InnerJoin(e any) query.JoinChain[*query.Delete] {
	return query.InnerJoin[*query.Delete](e)
}

func LeftJoin(e any) query.JoinChain[*query.Delete] {
	return query.LeftJoin[*query.Delete](e)
}

func RightJoin(e any) query.JoinChain[*query.Delete] {
	return query.RightJoin[*query.Delete](e)
}

func FullJoin(e any) query.JoinChain[*query.Delete] {
	return query.FullJoin[*query.Delete](e)
}

func CrossJoin(e any) pgcraft.Mod[*query.Delete] {
	return query.CrossJoin[*query.Delete](e)
}

func Where(e pgcraft.Expression) mods.Where[*query.Delete] {
	return mods.Where[*query.Delete]{E: e}
}

func Returning(clauses ...any) pgcraft.Mod[*query.Delete] {
	return mods.Returning[*query.Delete](clauses)
}
