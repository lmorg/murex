//go:build !windows && !plan9
// +build !windows,!plan9

package integrations

import (
	"embed"
)

//go:embed *_posix.mx
var resPOSIX embed.FS

func init() {
	resources["*_posix"] = &resPOSIX
}
