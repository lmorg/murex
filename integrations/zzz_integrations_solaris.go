//go:build solaris
// +build solaris

package integrations

import (
	"embed"
)

//go:embed *_solaris.mx
var resSolaris embed.FS

func init() {
	resources["*_solaris"] = &resSolaris
}
