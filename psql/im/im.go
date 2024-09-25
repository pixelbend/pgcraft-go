package im

import (
	"github.com/driftdev/pgcraft"
	"github.com/driftdev/pgcraft/clause"
	"github.com/driftdev/pgcraft/expr"
	"github.com/driftdev/pgcraft/internal"
	"github.com/driftdev/pgcraft/mods"
	"github.com/driftdev/pgcraft/query"
)

func With(name string, columns ...string) query.CTEChain[*query.Insert] {
	return query.With[*query.Insert](name, columns...)
}

func Recursive(r bool) pgcraft.Mod[*query.Insert] {
	return mods.Recursive[*query.Insert](r)
}

func Into(name any, columns ...string) pgcraft.Mod[*query.Insert] {
	return mods.QueryModFunc[*query.Insert](func(i *query.Insert) {
		i.Table = clause.Table{
			Expression: name,
			Columns:    columns,
		}
	})
}

func IntoAs(name any, alias string, columns ...string) pgcraft.Mod[*query.Insert] {
	return mods.QueryModFunc[*query.Insert](func(i *query.Insert) {
		i.Table = clause.Table{
			Expression: name,
			Alias:      alias,
			Columns:    columns,
		}
	})
}

func OverridingSystem() pgcraft.Mod[*query.Insert] {
	return mods.QueryModFunc[*query.Insert](func(i *query.Insert) {
		i.Overriding = "SYSTEM"
	})
}

func OverridingUser() pgcraft.Mod[*query.Insert] {
	return mods.QueryModFunc[*query.Insert](func(i *query.Insert) {
		i.Overriding = "USER"
	})
}

func Values(clauses ...pgcraft.Expression) pgcraft.Mod[*query.Insert] {
	return mods.Values[*query.Insert](clauses)
}

func Rows(rows ...[]pgcraft.Expression) pgcraft.Mod[*query.Insert] {
	return mods.Rows[*query.Insert](rows)
}

// Insert from a query
func Query(q pgcraft.Query) pgcraft.Mod[*query.Insert] {
	return mods.QueryModFunc[*query.Insert](func(i *query.Insert) {
		i.Query = q
	})
}

// The column to target. Will auto add brackets
func OnConflict(columns ...any) mods.Conflict[*query.Insert] {
	return mods.Conflict[*query.Insert](func() clause.Conflict {
		return clause.Conflict{
			Target: clause.ConflictTarget{
				Columns: columns,
			},
		}
	})
}

func OnConflictOnConstraint(constraint string) mods.Conflict[*query.Insert] {
	return mods.Conflict[*query.Insert](func() clause.Conflict {
		return clause.Conflict{
			Target: clause.ConflictTarget{
				Constraint: constraint,
			},
		}
	})
}

func Returning(clauses ...any) pgcraft.Mod[*query.Insert] {
	return mods.Returning[*query.Insert](clauses)
}

//========================================
// For use in ON CONFLICT DO UPDATE SET
//========================================

func Set(sets ...pgcraft.Expression) pgcraft.Mod[*clause.Conflict] {
	return mods.QueryModFunc[*clause.Conflict](func(c *clause.Conflict) {
		c.Set.Set = append(c.Set.Set, internal.ToAnySlice(sets)...)
	})
}

func SetCol(from string) mods.Set[*clause.Conflict] {
	return mods.Set[*clause.Conflict]{from}
}

func SetExcluded(cols ...string) pgcraft.Mod[*clause.Conflict] {
	exprs := make([]any, 0, len(cols))
	for _, col := range cols {
		if col == "" {
			continue
		}
		exprs = append(exprs,
			expr.Join{Exprs: []pgcraft.Expression{
				expr.Quote(col), expr.Raw("= EXCLUDED."), expr.Quote(col),
			}},
		)
	}

	return mods.QueryModFunc[*clause.Conflict](func(c *clause.Conflict) {
		c.Set.Set = append(c.Set.Set, exprs...)
	})
}

func Where(e pgcraft.Expression) pgcraft.Mod[*clause.Conflict] {
	return mods.QueryModFunc[*clause.Conflict](func(c *clause.Conflict) {
		c.Where.Conditions = append(c.Where.Conditions, e)
	})
}
