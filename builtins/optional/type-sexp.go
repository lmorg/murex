//go:build opt_type_sexp
// +build opt_type_sexp

package optional

// This is an optional builtin for S-Expressions support.
// Uses a 3rd party library: github.com/abesto/sexp
// (included in vendor directory)
import _ "github.com/lmorg/murex/builtins/types/sexp" // compile data type
