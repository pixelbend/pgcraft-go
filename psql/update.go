package psql

import (
	"github.com/teapartydev/pgcraft-go/query"
)

func Update(queryMods ...pgcraft.Mod[*query.Update]) pgcraft.BaseQuery[*query.Update] {
	q := &query.Update{}
	for _, mod := range queryMods {
		mod.Apply(q)
	}

	return pgcraft.BaseQuery[*query.Update]{
		Expression: q,
	}
}
