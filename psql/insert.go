package psql

import (
	"github.com/teapartydev/pgcraft-go/query"
)

func Insert(mods ...pgcraft.Mod[*query.Insert]) pgcraft.BaseQuery[*query.Insert] {
	q := &query.Insert{}
	for _, mod := range mods {
		mod.Apply(q)
	}

	return pgcraft.BaseQuery[*query.Insert]{
		Expression: q,
	}
}
