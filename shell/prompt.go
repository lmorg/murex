package shell

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
)

func getPrompt() {
	var (
		err, err2 error
		exitNum   int
		b         []byte
	)

	proc.ShellProcess.Variables.Set("linenum", 1, types.Number)
	prompt, err := proc.ShellProcess.Config.Get("shell", "prompt", types.CodeBlock)
	if err == nil {
		out := streams.NewStdin()
		exitNum, err = lang.RunBlockShellNamespace([]rune(prompt.(string)), nil, out, nil)
		//out.Close()

		b, err2 = out.ReadAll()
		if len(b) > 1 && b[len(b)-1] == '\n' {
			b = b[:len(b)-1]
		}

		if len(b) > 1 && b[len(b)-1] == '\r' {
			b = b[:len(b)-1]
		}

	}

	if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Invalid prompt. Block returned false.")
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

	proc.ShellProcess.Variables.Set("linenum", nLines, types.Number)
	prompt, err := proc.ShellProcess.Config.Get("shell", "prompt-multiline", types.CodeBlock)
	if err == nil {
		out := streams.NewStdin()
		exitNum, err = lang.RunBlockShellNamespace([]rune(prompt.(string)), nil, out, nil)
		//out.Close()

		b, err2 = out.ReadAll()
		if len(b) > 1 && b[len(b)-1] == '\n' {
			b = b[:len(b)-1]
		}

		if len(b) > 1 && b[len(b)-1] == '\r' {
			b = b[:len(b)-1]
		}
	}

	if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Invalid prompt. Block returned false.")
		b = []byte(fmt.Sprintf("%5d » ", nLines))
	}

	//Instance.SetPrompt(string(b))
	Prompt.SetPrompt(string(b))
}
