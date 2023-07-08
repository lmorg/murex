//go:build windows
// +build windows

package integrations

import (
	"embed"
)

//go:embed *_windows.mx
var resWindows embed.FS

func init() {
	resources["*_windows"] = &resWindows
}
