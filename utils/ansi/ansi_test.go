package ansi

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

// TestAnsiColoured tests writing colours via the ansi package
func TestAnsiColoured(t *testing.T) {
	message := "This is only a test"
	stdtest := streams.NewStdtest()

	proc.ShellProcess.Config = config.NewConfiguration()
	proc.ShellProcess.Config.Define("shell", "add-colour", config.Properties{
		DataType:    types.Boolean,
		Default:     true,
		Description: "test data",
	})

	err := Stream(stdtest, FgRed, message)
	if err != nil {
		t.Error(err.Error())
	}

	//stdtest.Close()

	b, err := stdtest.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}

	if string(b) != FgRed+message+Reset {
		t.Error("Colourised config: Source string does not match output string: " + string(b))
	}
}

// TestAnsiNoColour tests the add-colour override disables the ansi package
func TestAnsiNoColour(t *testing.T) {
	message := "This is only a test"
	stdtest := streams.NewStdtest()

	proc.ShellProcess.Config = config.NewConfiguration()
	proc.ShellProcess.Config.Define("shell", "add-colour", config.Properties{
		DataType:    types.Boolean,
		Default:     false,
		Description: "test data",
	})

	err := Stream(stdtest, FgRed, message)
	if err != nil {
		t.Error(err.Error())
	}

	//stdtest.Close()

	b, err := stdtest.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}

	if string(b) != message {
		t.Error("No colour override: Source string does not match output string: " + string(b))
	}
}
