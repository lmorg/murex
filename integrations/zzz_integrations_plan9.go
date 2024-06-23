//go:build plan9
// +build plan9

package integrations

import (
	"embed"
)

//go:embed *_plan9.mx
var resPlan9 embed.FS

func init() {
	resources["*_plan9"] = &resPlan9
}
