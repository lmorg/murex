package ansi

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/ansi/codes"
)

// TestAnsiColoured tests writing colours via the ansi package
func TestAnsiColoured(t *testing.T) {
	count.Tests(t, 1)

	lang.ShellProcess.Config = config.InitConf
	lang.ShellProcess.Config.Define("shell", "color", config.Properties{
		DataType:    types.Boolean,
		Default:     true,
		Description: "test data",
	})

	text := "This is only a test"
	message := "{RED}" + text + "{RESET}"
	output := ExpandConsts(message)

	if output != codes.FgRed+text+codes.Reset {
		t.Error("Colourised config: Source string does not match expected output string: " + output)
	}
	if output == message {
		t.Error("Colourised config: Source string has had no variables substituted")
	}
	if output == text {
		t.Error("Colourised config: Source string variables substituted with zero length string")
	}

}

// TestAnsiNoColour tests the color override disables the ansi package
func TestAnsiNoColour(t *testing.T) {
	count.Tests(t, 1)

	lang.ShellProcess.Config = config.InitConf
	lang.ShellProcess.Config.Define("shell", "color", config.Properties{
		DataType:    types.Boolean,
		Default:     false,
		Description: "test data",
	})

	text := "This is only a test"
	message := "{RED}" + text + "{RESET}"
	output := ExpandConsts(message)

	if output != text {
		t.Error("No colour override: Source string does not match expected output string: " + output)
	}
}

func TestChar20Leak(t *testing.T) {
	tests := []string{
		`{ESC}123`,
		` {ESC}123`,
		`  {ESC}123`,
		"\t{ESC}123",
		`123{ESC}123`,
		`{ESC}123`,
		`123123`,
		`{ESC}`,
	}

	lang.ShellProcess.Config = config.InitConf
	lang.ShellProcess.Config.Define("shell", "color", config.Properties{
		DataType:    types.Boolean,
		Default:     true,
		Description: "test data",
	})

	for i, test := range tests {
		expected := strings.ReplaceAll(test, "{ESC}", "\x1b")
		actual := ExpandConsts(test)
		if expected != actual {
			t.Errorf("mismatch in test %d", i)
			t.Logf("  test:    '%s'", test)
			t.Logf("  expected: %v", []byte(expected))
			t.Logf("  actual:   %v", []byte(actual))
		}
	}

	count.Tests(t, len(tests))
}
