package docs_test

import (
	"embed"
	"testing"

	"github.com/lmorg/murex/test"
)

//go:embed *.go
var docs embed.FS

// TestSummaries tests the docs directory has been populated with summaries.
// All other files are tested from the builtins/docs_test.go to avoid
// cyclic package import paths
func TestSummaries(t *testing.T) {
	test.ExistsFs(t, "000_summaries_commands_docgen.go", docs)
}
