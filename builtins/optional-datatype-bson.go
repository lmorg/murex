// +build ignore
// This is disabled by default because it calls an external command, `bzr`

package builtins

// Uses a 3rd party library: labix.org/v2/mgo/bson
// (included in vendor directory)
import _ "github.com/lmorg/murex/builtins/types/bson" // compile data type
