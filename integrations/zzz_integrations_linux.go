//go:build linux
// +build linux

package integrations

import (
	"embed"
)

//go:embed *_linux.mx
var resLinux embed.FS

func init() {
	resources["*_linux"] = &resLinux
}
