package ansi

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

// TestAnsiColoured tests writing colours via the ansi package
func TestAnsiColoured(t *testing.T) {
	proc.ShellProcess.Config = config.NewConfiguration()
	proc.ShellProcess.Config.Define("shell", "add-colour", config.Properties{
		DataType:    types.Boolean,
		Default:     true,
		Description: "test data",
	})

	text := "This is only a test"
	message := "{RED}" + text + "{RESET}"
	output := ExpandConsts(message)

	if output != FgRed+text+Reset {
		t.Error("Colourised config: Source string does not match expected output string: " + output)
	}
	if output == message {
		t.Error("Colourised config: Source string has had no variables substituted")
	}
	if output == text {
		t.Error("Colourised config: Source string variables substituted with zero length string")
	}

}

// TestAnsiNoColour tests the add-colour override disables the ansi package
func TestAnsiNoColour(t *testing.T) {
	proc.ShellProcess.Config = config.NewConfiguration()
	proc.ShellProcess.Config.Define("shell", "add-colour", config.Properties{
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
