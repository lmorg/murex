package test

import (
	"embed"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func ExistsFs(t *testing.T, name string, res embed.FS) {
	t.Helper()
	count.Tests(t, 1)

	b, err := res.ReadFile(name)
	if err != nil {
		t.Errorf("Cannot read file '%s': %s", name, err.Error())
		t.FailNow()
	}

	if len(b) == 0 {
		t.Errorf("File is empty: %s", name)
	}
}
