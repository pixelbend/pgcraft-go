package clause

import (
	"io"
)

type IWindow interface {
	SetFrom(string)
	AddPartitionBy(...any)
	AddOrderBy(...any)
	SetMode(string)
	SetStart(any)
	SetEnd(any)
	SetExclusion(string)
}

type Window struct {
	From        string // an existing window name
	orderBy     []any
	partitionBy []any
	Frame
}

func (wi *Window) SetFrom(from string) {
	wi.From = from
}

func (wi *Window) AddPartitionBy(condition ...any) {
	wi.partitionBy = append(wi.partitionBy, condition...)
}

func (wi *Window) AddOrderBy(order ...any) {
	wi.orderBy = append(wi.orderBy, order...)
}

func (wi Window) WriteSQL(w io.Writer, start int) ([]any, error) {
	if wi.From != "" {
		w.Write([]byte(wi.From))
		w.Write([]byte(" "))
	}

	args, err := pgcraft.ExpressSlice(w, start, wi.partitionBy, "PARTITION BY ", ", ", " ")
	if err != nil {
		return nil, err
	}

	orderArgs, err := pgcraft.ExpressSlice(w, start, wi.orderBy, "ORDER BY ", ", ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, orderArgs...)

	frameArgs, err := pgcraft.ExpressIf(w, start, wi.Frame, wi.Frame.Defined, " ", "")
	if err != nil {
		return nil, err
	}
	args = append(args, frameArgs...)

	return args, nil
}

type NamedWindow struct {
	Name       string
	Definition any
}

func (n NamedWindow) WriteSQL(w io.Writer, start int) ([]any, error) {
	w.Write([]byte(n.Name))
	w.Write([]byte(" AS ("))
	args, err := pgcraft.Express(w, start, n.Definition)
	w.Write([]byte(")"))

	return args, err
}

type Windows struct {
	Windows []NamedWindow
}

func (wi *Windows) AppendWindow(w NamedWindow) {
	wi.Windows = append(wi.Windows, w)
}

func (wi Windows) WriteSQL(w io.Writer, start int) ([]any, error) {
	return pgcraft.ExpressSlice(w, start, wi.Windows, "WINDOW ", ", ", "")
}
