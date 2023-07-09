//go:build darwin
// +build darwin

package integrations

import (
	"embed"
)

//go:embed *_darwin.mx
var resDarwin embed.FS

func init() {
	resources["*_darwin"] = &resDarwin
}
