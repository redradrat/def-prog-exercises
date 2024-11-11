package safesql

import "github.com/empijei/def-prog-exercises/safesql/internal/raw"

// Not exported on purpose, so this can never be used by a different package
// to treat user input.
type compileTimeConstant string

type TrustedSQL struct {
	s string
}

func New(text compileTimeConstant) TrustedSQL {
	return TrustedSQL{string(text)}
}

func init() {
	raw.TrustedSQLCtor = func(unsafe string) TrustedSQL {
		return TrustedSQL{s: unsafe}
	}
}
