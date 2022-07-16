package shell_test

import (
	"regexp"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/test/count"
)

func TestHintCodeBlock(t *testing.T) {
	tests := []struct {
		Block  string
		HintRx string
	}{
		{
			Block:  `{ out "bob" }`,
			HintRx: `^bob$`,
		},
		{
			Block:  `out "bob"`,
			HintRx: `^out "bob"$`,
		},
		{
			Block:  `{ $out "bob" }`,
			HintRx: `^Block returned false`,
		},
	}

	count.Tests(t, len(tests))

	lang.InitEnv()
	defaults.Defaults(lang.ShellProcess.Config, false)
	debug.Enabled = true

	for i, test := range tests {
		err := lang.ShellProcess.Config.Set("shell", "hint-text-func", test.Block, nil)
		if err != nil {
			panic(err.Error())
		}
		hint := string(shell.HintCodeBlock())
		rx := regexp.MustCompile(test.HintRx)
		if !rx.MatchString(hint) {
			t.Errorf("Hint doesn't match expected in test %d", i)
			t.Logf("  Block:    '%s'", test.Block)
			t.Logf("  Expected:  %s", test.HintRx)
			t.Logf("  Actual:   '%s'", hint)
		}
	}
}
