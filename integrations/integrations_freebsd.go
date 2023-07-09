//go:build freebsd
// +build freebsd

package integrations

import (
	"embed"
)

//go:embed *_freebsd.mx
var resFreeBSD embed.FS

func init() {
	resources["*_freebsd"] = &resFreeBSD
}
