package pgcraft

import "bytes"

// MustBuild builds a query form start and panics on error
// useful for initializing queries that need to be reused
func MustBuild(q Query) (string, []any) {
	return MustBuildN(q, 1)
}

// MustBuildN builds a query from a point and panics on error
// useful for initializing queries that need to be reused
func MustBuildN(q Query, start int) (string, []any) {
	sql, args, err := BuildN(q, start)
	if err != nil {
		panic(err)
	}

	return sql, args
}

// Build Convenient function to build query from start
func Build(q Query) (string, []any, error) {
	return BuildN(q, 1)
}

// BuildN Convenient function to build query from a point
func BuildN(q Query, start int) (string, []any, error) {
	b := &bytes.Buffer{}
	args, err := q.WriteQuery(b, start)

	return b.String(), args, err
}
