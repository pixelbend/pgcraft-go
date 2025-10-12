package clause

import (
	"fmt"
	"github.com/arkamfahry/pgcraft-go"
	"io"
)

var ErrNoLockStrength = fmt.Errorf("no lock strength specified")

const (
	LockStrengthUpdate      = "UPDATE"
	LockStrengthNoKeyUpdate = "NO KEY UPDATE"
	LockStrengthShare       = "SHARE"
	LockStrengthKeyShare    = "KEY SHARE"
)

const (
	LockWaitNoWait     = "NOWAIT"
	LockWaitSkipLocked = "SKIP LOCKED"
)

type For struct {
	Strength string
	Tables   []string
	Wait     string
}

func (f *For) SetFor(lock For) {
	*f = lock
}

func (f For) WriteSQL(w io.Writer, start int) ([]any, error) {
	if f.Strength == "" {
		return nil, nil
	}

	w.Write([]byte("FOR "))
	if f.Strength != "" {
		fmt.Fprintf(w, "%s ", f.Strength)
	}

	args, err := pgcraft.ExpressSlice(w, start, f.Tables, "OF ", ", ", "")
	if err != nil {
		return nil, err
	}

	if f.Wait != "" {
		w.Write([]byte(" "))
		w.Write([]byte(f.Wait))
	}

	return args, nil
}
