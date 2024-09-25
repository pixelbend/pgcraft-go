package main

import (
	"fmt"
	"github.com/driftdev/pgcraft/psql"
	"github.com/driftdev/pgcraft/psql/im"
	"github.com/driftdev/pgcraft/psql/sm"
)

func main() {
	stmt, args, err := psql.Insert(
		im.Into("users"),
		im.Values(psql.Select(
			sm.From("tmp_films"),
			sm.Where(psql.Quote("date_prod").LT(psql.Arg("1971-07-13"))),
		)),
	).Build()

	if err != nil {
		panic(err)
	}

	fmt.Println(stmt)
	fmt.Println(args)
}
