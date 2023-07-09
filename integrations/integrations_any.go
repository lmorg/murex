package integrations

import (
	"embed"
)

//go:embed *_any.mx
var resAny embed.FS

func init() {
	resources["*_any"] = &resAny
}
