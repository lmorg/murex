package builtins

import (
	"embed"
	"testing"

	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

//go:embed *_test.go
var docgen embed.FS

// TestSummaries tests the docs directory has been populated with summaries.
// All other files are tested from the builtins/docs_test.go to avoid
// cyclic package import paths
func TestDocgenTest(t *testing.T) {
	test.ExistsFs(t, "docgen_test.go", docgen)
}

// sourceFile is the names of the auto-generated files produced from docgen. This is used for the testing.
// This will be also auto-populated by docgen
var sourceFile map[string]string

//go:embed docs/*.go
var coreDocs embed.FS

// TestCoreDocs tests documentation has been written for core builtins
func TestCoreDocs(t *testing.T) {
	count.Tests(t, len(lang.GoFunctions)*2)

	for name := range lang.GoFunctions {

		syn := docs.Synonym[name]
		if syn == "" {
			t.Errorf("Synonym for `%s` does not exist", name)
			continue
		}

		src := sourceFile[syn]
		if src == "" {
			src = syn + "_commands_docgen.go"
			t.Logf("docgen failed to write a source path for `%s`. Guessing it at '%s'", syn, src)
		}

		test.ExistsFs(t, "docs/"+src, coreDocs)
	}
}
