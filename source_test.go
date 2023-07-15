package main

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

func TestDiskSource(t *testing.T) {
	count.Tests(t, 1)

	file := "test/source.mx"
	test.Exists(t, file)

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
	test.Exists(t, file)

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

	source := `out: "testing"`
	srcRef := ref.History.AddSource("(builtin)", "source/builtin", []byte(source))
	execSource([]rune(source), srcRef, false)
}
