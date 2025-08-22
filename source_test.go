package main

import (
	"embed"
	"testing"

	"github.com/lmorg/murex/config/profile/source"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

//go:embed test/source*
var testSource embed.FS

func TestDiskSource(t *testing.T) {
	count.Tests(t, 1)

	file := "test/source.mx"
	test.ExistsFs(t, file, testSource)

	disk, err := diskSource(file)
	if err != nil {
		t.Error(err)
	}

	if len(disk) == 0 {
		t.Error(err)
	}
}

func TestDiskSourceGz(t *testing.T) {
	count.Tests(t, 1)

	file := "test/source.mx.gz"
	test.ExistsFs(t, file, testSource)

	disk, err := diskSource(file)
	if err != nil {
		t.Error(err)
	}

	if len(disk) == 0 {
		t.Error(err)
	}
}

func TestExecSource(t *testing.T) {
	count.Tests(t, 1)

	lang.InitEnv()

	block := `out: "testing"`
	srcRef := ref.History.AddSource("(builtin)", "source/builtin", []byte(block))
	source.Exec([]rune(block), srcRef, false)
}
