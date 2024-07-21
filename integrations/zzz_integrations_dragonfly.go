//go:build dragonfly
// +build dragonfly

package integrations

import (
	"embed"
)

//go:embed *_dragonfly.mx
var resDragonfly embed.FS

func init() {
	resources["*_dragonfly"] = &resDragonfly
}
