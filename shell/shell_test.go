package shell_test

import (
	"testing"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/test/count"
)

func TestSpellcheckBadBlock(t *testing.T) {
	count.Tests(t, 1)
	lang.InitEnv()
	defaults.Config(lang.ShellProcess.Config, false)

	err := lang.ShellProcess.Config.Set("shell", "spellcheck-enabled", true, lang.NewTestProcess().FileRef)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-enabled config: %s", err)
	}

	err = lang.ShellProcess.Config.Set("shell", "spellcheck-func", `{ -> jsplit { }`, lang.NewTestProcess().FileRef)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-func config: %s", err)
	}

	line := "the quick brown fox"
	newLine := string(shell.Spellchecker([]rune(line)))

	if newLine != line {
		t.Error("spellcheck output doesn't match expected:")
		t.Logf("  Expected: '%s'", line)
		t.Logf("  Actual:   '%s'", newLine)
	}
}
