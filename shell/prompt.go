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
		//out := streams.NewStdin()
		//branch := lang.ShellProcess.BranchFID()
		//defer branch.Close()
		//branch.Variables.Set("linenum", 1, types.Integer)
		//exitNum, err = lang.RunBlockExistingConfigSpace([]rune(prompt.(string)), nil, out, nil, branch.Process)

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
		fork.Variables.Set("linenum", 1, types.Integer)
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
		//out := streams.NewStdin()
		//branch := lang.ShellProcess.BranchFID()
		//defer branch.Close()
		//branch.Variables.Set("linenum", nLines, types.Integer)
		//exitNum, err = lang.RunBlockExistingConfigSpace([]rune(prompt.(string)), nil, out, nil, branch.Process)

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
		fork.Variables.Set("linenum", nLines, types.Integer)
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
