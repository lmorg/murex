package parser_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/parser"
)

func TestIsSafeCmdsBuiltins(t *testing.T) {
	safeCmdsLocal := parser.GetSafeCmds()

	count.Tests(t, len(safeCmdsLocal))

	for _, cmd := range safeCmdsLocal {
		if lang.GoFunctions[cmd] == nil {
			t.Errorf("Command hardcoded in safe whitelist but is not a builtin: %s", cmd)
		}
	}
}
