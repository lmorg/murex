package shell

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/readline"
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
		fork.Name = "(prompt)"
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
		fork.Name = "(prompt-multiline)"
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

func leftMost() []byte {
	fd := int(os.Stdout.Fd())
	w, _, err := readline.GetSize(fd)
	if err != nil {
		return utils.NewLineByte
	}

	b := make([]byte, w+1)
	for i := 0; i < w; i++ {
		b[i] = ' '
	}
	b[w] = '\r'

	return b
}
