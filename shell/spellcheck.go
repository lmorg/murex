package shell

import (
	"bytes"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/userdictionary"
	"github.com/lmorg/murex/utils/ansi"
)

func spellcheck(line []rune) []rune {
	r := line
	enabled, err := lang.ShellProcess.Config.Get("shell", "spellcheck-enabled", types.Boolean)
	if err != nil || !enabled.(bool) {
		return r
	}

	block, err := lang.ShellProcess.Config.Get("shell", "spellcheck-block", types.CodeBlock)
	if err != nil || len(block.(string)) == 0 {
		return r
	}

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_CREATE_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	fork.Name = "(spellcheck)"
	fork.Stdin.SetDataType(types.Generic)
	_, err = fork.Stdin.Writeln([]byte(string(r)))
	if err != nil && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		return r
	}

	_, err = fork.Execute([]rune(block.(string)))
	if err != nil && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		return r
	}

	b, err := fork.Stderr.ReadAll()
	if err != nil && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}
	if len(b) != 0 && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}

	err = fork.Stdout.ReadArray(func(bWord []byte) {
		if len(bWord) == 0 {
			return
		}

		sWord := string(bytes.TrimSpace(bWord))

		if autocomplete.GlobalExes[sWord] || lang.MxFunctions.Exists(sWord) || lang.GoFunctions[sWord] != nil || lang.GlobalAliases.Exists(sWord) {
			return
		}

		if lang.ShellProcess.Variables.GetValue(sWord) != nil {
			return
		}

		if userdictionary.IsInDictionary(sWord) {
			return
		}

		r = []rune(strings.ReplaceAll(string(r), sWord, ansi.ExpandConsts("{UNDERLINE}"+sWord+"{UNDEROFF}")))
	})
	if err != nil && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}

	return r
}
