package examples

/*
	This file only exists to test that any parser changes do not introduce any
	regression errors. Using the suite of examples already provided we can
	expand the degree of parser tests for "free".
*/

import (
	"embed"
	"testing"

	"github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/test/count"
)

//go:embed murex_profile
//go:embed *.mx
var testcode embed.FS

func TestExampleCodeParses(t *testing.T) {
	dir, err := testcode.ReadDir(".")
	if err != nil {
		// not a bug in murex
		panic(err)
	}

	count.Tests(t, len(dir))

	for i := range dir {
		name := dir[i].Name()

		b, err := testcode.ReadFile(name)
		if err != nil {
			// not a bug in murex
			panic(err)
		}

		block := []rune(string(b))
		blk := expressions.NewBlock(block)
		err = blk.ParseBlock()
		if err != nil {
			// this _is_ a bug in murex!
			t.Errorf("example failed to parse: `%s`", name)
			t.Logf("  Error returned: %v", err)
		}
	}
}
