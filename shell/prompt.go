package shell

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func getPrompt() {
	var (
		err, err2 error
		exitNum   int
		b         []byte
	)

	prompt, err := proc.ShellProcess.Config.Get("shell", "prompt", types.CodeBlock)
	if err == nil {
		out := streams.NewStdin()
		branch := proc.ShellProcess.BranchFID()
		defer branch.Close()
		branch.Variables.Set("linenum", 1, types.Integer)
		exitNum, err = lang.RunBlockExistingConfigSpace([]rune(prompt.(string)), nil, out, nil, branch.Process)

		b, err2 = out.ReadAll()
		b = utils.CrLfTrim(b)
	}

	if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
		proc.ShellProcess.Stderr.Writeln([]byte("Invalid prompt. Block returned false."))
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

	prompt, err := proc.ShellProcess.Config.Get("shell", "prompt-multiline", types.CodeBlock)
	if err == nil {
		out := streams.NewStdin()
		branch := proc.ShellProcess.BranchFID()
		defer branch.Close()
		branch.Variables.Set("linenum", nLines, types.Integer)
		exitNum, err = lang.RunBlockExistingConfigSpace([]rune(prompt.(string)), nil, out, nil, branch.Process)

		b, err2 = out.ReadAll()
		b = utils.CrLfTrim(b)
	}

	if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
		proc.ShellProcess.Stderr.Writeln([]byte("Invalid prompt. Block returned false."))
		b = []byte(fmt.Sprintf("%5d » ", nLines))
	}

	Prompt.SetPrompt(string(b))
}
