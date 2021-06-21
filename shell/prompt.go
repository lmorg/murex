package shell

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func getPrompt() {
	var (
		err, err2 error
		exitNum   int
		b         []byte
	)

	prompt, err := lang.ShellProcess.Config.Get("shell", "prompt", types.CodeBlock)
	if err == nil {
		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
		fork.Variables.Set(fork.Process, "linenum", 1, types.Integer)
		fork.Name.Set("(prompt)")
		fork.Execute([]rune(prompt.(string)))

		b, err2 = fork.Stdout.ReadAll()
		b = utils.CrLfTrim(b)
	}

	if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Invalid prompt. Block returned false."))
		b = []byte("murex » ")
	}

	Prompt.SetPrompt(string(b))
}

func getMultilinePrompt(nLines int) {
	var (
		err, err2 error
		exitNum   int
		b         []byte
	)

	prompt, err := lang.ShellProcess.Config.Get("shell", "prompt-multiline", types.CodeBlock)
	if err == nil {
		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
		fork.Variables.Set(fork.Process, "linenum", nLines, types.Integer)
		fork.Name.Set("(prompt-multiline)")
		fork.Execute([]rune(prompt.(string)))

		b, err2 = fork.Stdout.ReadAll()
		b = utils.CrLfTrim(b)
	}

	if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Invalid prompt. Block returned false."))
		b = []byte(fmt.Sprintf("%5d » ", nLines))
	}

	Prompt.SetPrompt(string(b))
}

// ConfigReadGetCursorPos is a dynamic config wrapper function for Prompt.EnableGetCursorPos
func ConfigReadGetCursorPos() (interface{}, error) {
	return Prompt.EnableGetCursorPos, nil
}

// ConfigWriteGetCursorPos is a dynamic config wrapper function for Prompt.EnableGetCursorPos
func ConfigWriteGetCursorPos(v interface{}) error {
	switch v.(type) {
	case bool:
		Prompt.EnableGetCursorPos = v.(bool)

	case string:
		switch v.(string) {
		case types.TrueString:
			Prompt.EnableGetCursorPos = true

		case types.FalseString:
			Prompt.EnableGetCursorPos = false

		default:
			return fmt.Errorf("Expecting 'true' or 'false'. Instead received '%s'", v.(string))
		}

	default:
		return fmt.Errorf("Expecting boolean value. Instead received %T", v)
	}

	return nil
}
