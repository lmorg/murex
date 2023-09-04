// +build opt_select

package optional

// This optinal builtin provides SQL inlining via the `select` command.
// It uses sqlite3 as a backend, which will be compiled into the murex executable.

import _ "github.com/lmorg/murex/builtins/optional/select" // compile optional builtin
