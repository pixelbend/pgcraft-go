package clause

import (
	"fmt"
	"io"
)

var ErrNoCombinationStrategy = fmt.Errorf("combination strategy must be set")

const (
	Union     = "UNION"
	Intersect = "INTERSECT"
	Except    = "EXCEPT"
)

type Combine struct {
	Strategy string
	Query    pgcraft.Query
	All      bool
}

func (s *Combine) SetCombine(c Combine) {
	*s = c
}

func (s Combine) WriteSQL(w io.Writer, start int) ([]any, error) {
	if s.Strategy == "" {
		return nil, ErrNoCombinationStrategy
	}

	w.Write([]byte(s.Strategy))

	if s.All {
		w.Write([]byte(" ALL "))
	} else {
		w.Write([]byte(" "))
	}

	args, err := pgcraft.Express(w, start, s.Query)
	if err != nil {
		return nil, err
	}

	return args, nil
}
