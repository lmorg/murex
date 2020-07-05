package builtins

import (
	"testing"

	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/gopath"
)

// sourceFile is the names of the auto-generated files produced from docgen. This is used for the testing.
// This will be also auto-populated by docgen
var sourceFile map[string]string

// TestCoreDocs tests documentation has been written for core builtins
func TestCoreDocs(t *testing.T) {
	count.Tests(t, 1)
	test.Exists(t, gopath.Source([]string{"builtins"})+"docgen_test.go")

	path := gopath.Source([]string{"builtins", "docs"})

	count.Tests(t, len(lang.GoFunctions)*2)
	for name := range lang.GoFunctions {
		syn := docs.Synonym[name]
		if syn == "" {
			t.Errorf("Synonym for `%s` does not exist", name)
			continue
		}

		src := sourceFile[syn]
		if src == "" {
			t.Errorf("docgen failed to write a source path for `%s`", syn)
		}

		test.Exists(t, path+src)
	}
}
