package expressions

import "regexp"

var (
	ExpFullName = regexp.MustCompile(`^[\p{L}\p{M}\p{Z}\p{Cf}\p{Braille}\p{Pd}\p{Pc}\p{Ps}\p{Pf}'・]+$`)
)
