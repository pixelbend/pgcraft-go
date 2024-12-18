package psql

import (
	"github.com/pixelbend/pgcraft-go"
	"github.com/pixelbend/pgcraft-go/query"
)

func Delete(mods ...pgcraft.Mod[*query.Delete]) pgcraft.BaseQuery[*query.Delete] {
	q := &query.Delete{}
	for _, mod := range mods {
		mod.Apply(q)
	}

	return pgcraft.BaseQuery[*query.Delete]{
		Expression: q,
	}
}
