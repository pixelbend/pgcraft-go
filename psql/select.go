package psql

import (
	"github.com/arkamfahry/pgcraft-go"
	"github.com/arkamfahry/pgcraft-go/query"
)

func Select(mods ...pgcraft.Mod[*query.Select]) pgcraft.BaseQuery[*query.Select] {
	q := &query.Select{}
	for _, mod := range mods {
		mod.Apply(q)
	}

	return pgcraft.BaseQuery[*query.Select]{
		Expression: q,
	}
}
