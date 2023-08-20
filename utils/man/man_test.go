//go:build !windows
// +build !windows

package man

import (
	"embed"
	"testing"

	"github.com/lmorg/murex/test/count"
)

//go:embed test_*.txt
var manPagesTxt embed.FS

func TestMan(t *testing.T) {
	files, err := manPagesTxt.ReadDir(".")
	if err != nil {
		t.Error(err.Error())
	}

	count.Tests(t, len(files)*2)

	for _, entry := range files {
		file, err := manPagesTxt.Open(entry.Name())
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}

		flags, descs := ParseByStdio(file)
		if len(flags) == 0 {
			t.Errorf("%d flags returned for '%s'", len(flags), entry.Name())
		}
		if len(descs) == 0 {
			t.Errorf("%d descriptions returned for '%s'", len(descs), entry.Name())
		}

		err = file.Close()
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}
	}
}
