//go:build netbsd
// +build netbsd

package integrations

import (
	"embed"
)

//go:embed *_netbsd.mx
var resNetBSD embed.FS

func init() {
	resources["*_netbsd"] = &resNetBSD
}
