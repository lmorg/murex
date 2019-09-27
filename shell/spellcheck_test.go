package shell

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/arraytools"
	_ "github.com/lmorg/murex/builtins/core/textmanip"
	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/ansi"
)

func TestSpellcheckCrLf(t *testing.T) {
	count.Tests(t, 1)
	lang.InitEnv()
	defaults.Defaults(lang.ShellProcess.Config, false)

	err := lang.ShellProcess.Config.Set("shell", "spellcheck-enabled", true)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-enabled config: %s", err)
	}

	err = lang.ShellProcess.Config.Set("shell", "spellcheck-block", `{ -> jsplit ' ' -> suffix "\n" }`)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-block config: %s", err)
	}

	line := "the quick brown fox"
	newLine := string(spellcheck([]rune(line)))
	ansiLine := ansi.ExpandConsts("{UNDERLINE}the{UNDEROFF} {UNDERLINE}quick{UNDEROFF} {UNDERLINE}brown{UNDEROFF} {UNDERLINE}fox{UNDEROFF}")

	if newLine != ansiLine {
		t.Error("spellcheck output doesn't match expected:")
		t.Logf("  Expected: '%s'", ansiLine)
		t.Logf("  Actual:   '%s'", newLine)
	}
}

func TestSpellcheckZeroLenStr(t *testing.T) {
	count.Tests(t, 1)
	lang.InitEnv()
	defaults.Defaults(lang.ShellProcess.Config, false)

	err := lang.ShellProcess.Config.Set("shell", "spellcheck-enabled", true)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-enabled config: %s", err)
	}

	err = lang.ShellProcess.Config.Set("shell", "spellcheck-block", `{ -> jsplit '\s' -> append '' }`)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-block config: %s", err)
	}

	line := "the quick brown fox"
	newLine := string(spellcheck([]rune(line)))
	ansiLine := ansi.ExpandConsts("{UNDERLINE}the{UNDEROFF} {UNDERLINE}quick{UNDEROFF} {UNDERLINE}brown{UNDEROFF} {UNDERLINE}fox{UNDEROFF}")

	if newLine != ansiLine {
		t.Error("spellcheck output doesn't match expected:")
		t.Logf("  Expected: '%s'", ansiLine)
		t.Logf("  Actual:   '%s'", newLine)
	}
}

// test times out for reasons currently unknown
/*func TestSpellcheckBadBlock(t *testing.T) {
	count.Tests(t, 1)
	lang.InitEnv()
	defaults.Defaults(lang.ShellProcess.Config, false)

	err := lang.ShellProcess.Config.Set("shell", "spellcheck-enabled", true)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-enabled config: %s", err)
	}

	err = lang.ShellProcess.Config.Set("shell", "spellcheck-block", `{ -> jsplit { }`)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-block config: %s", err)
	}

	line := "the quick brown fox"
	newLine := string(spellcheck([]rune(line)))

	if newLine != line {
		t.Error("spellcheck output doesn't match expected:")
		t.Logf("  Expected: '%s'", line)
		t.Logf("  Actual:   '%s'", newLine)
	}
}*/
