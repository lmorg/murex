//go:build !windows
// +build !windows

package man

import (
	"compress/gzip"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestManSummary(t *testing.T) {
	count.Tests(t, 1)

	files, err := manPages.ReadDir(".")
	if err != nil {
		t.Error(err.Error())
	}

	for _, entry := range files {
		file, err := manPages.Open(entry.Name())
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
