package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/debug"
)

//go:embed docs/apis/*.md
//go:embed docs/changelog/*.md
//go:embed docs/commands/*.md
//go:embed docs/events/*.md
//go:embed docs/mkarray/*.md
//go:embed docs/optional/*.md
//go:embed docs/parser/*.md
//go:embed docs/types/*.md
//go:embed docs/user-guide/*.md
//go:embed docs/variables/*.md
//go:embed docs/integrations/*.md
var docsEmbeded embed.FS

func init() {
	docs.Definition = docsImport
}

func docsImport(path string) []byte {
	if !strings.Contains(path, "/") {
		path = "commands/" + path
	}
	path = fmt.Sprintf("docs/%s.md", path)

	if debug.Enabled {
		// in debug mode lets output the actual error
		b, err := docsEmbeded.ReadFile(path)
		if err != nil {
			return append([]byte("error: "), []byte(err.Error())...)
		}
		return b
	}

	b, _ := docsEmbeded.ReadFile(path)
	return b
}
