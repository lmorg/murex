//go:build openbsd
// +build openbsd

package integrations

import (
	"embed"
)

//go:embed *_openbsd.mx
var resOpenBSD embed.FS

func init() {
	resources["*_openbsd"] = &resOpenBSD
}
