package docs

import (
	"os"
	"strings"
	"testing"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/consts"
)

// TestDocs tests the docs directory has been populated
func TestDocs(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		t.Error("GOPATH not set. GOPATH needs to be set for docs package to be tested")
	}

	path := []string{
		"src",
		"github.com",
		"lmorg",
		"murex",
		"builtins",
		"core",
		"docs",
	}

	docs := GOPATH + consts.PathSlash + strings.Join(path, consts.PathSlash) + consts.PathSlash

	exists(t, docs+"autogen-digests.go")
	exists(t, docs+"autogen-synonyms.go")

	for name := range proc.GoFunctions {
		exists(t, docs+"autogen-func-"+name+".go")
	}
}

func exists(t *testing.T, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("Missing file", path)
	}
}
