package docs

import (
	"testing"

	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/utils/gopath"
)

// TestSummaries tests the docs directory has been populated with summaries.
// All other files are tested from the builtins/docs_test.go to avoid
// cyclic package import paths
func TestSummaries(t *testing.T) {
	path := []string{
		"builtins",
		"docs",
	}

	docs := gopath.Source(path)

	test.Exists(t, docs+"000_summaries_commands_docgen.go")
}
