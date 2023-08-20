//go:build !windows
// +build !windows

package man

import (
	"compress/gzip"
	"embed"
	"testing"

	"github.com/lmorg/murex/test/count"
)

//go:embed test_*.1.gz
var manPagesGz embed.FS

func TestManSummary(t *testing.T) {
	count.Tests(t, 1)

	files, err := manPagesGz.ReadDir(".")
	if err != nil {
		t.Error(err.Error())
	}

	count.Tests(t, len(files))

	for _, entry := range files {
		file, err := manPagesGz.Open(entry.Name())
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}

		gz, err := gzip.NewReader(file)
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}

		s := parseSummaryFile(gz)
		if s == "" {
			t.Errorf("No summary returned for '%s'", entry.Name())
		}

		gz.Close()
		file.Close()
	}
}
