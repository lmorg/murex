package shell

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/hintsummary"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/parser"
)

var (
	cachedHintText []rune
)

func hintText(line []rune, pos int) []rune {
	r := hintExpandVariables(line)
	if len(r) > 0 {
		return r
	}

	pt, _ := parser.Parse(line, 0)
	cmd := pt.FuncName
	//var summary string

	/*if pt.ExpectFunc == false {
		// autocomplete hint
	}*/

	if cmd == "cd" && len(pt.Parameters) > 0 && len(pt.Parameters[0]) > 0 {
		path := utils.NormalisePath(pt.Parameters[0])
		return []rune("Change directory: " + path)
	}

	// check if a custom summary has been set
	r = hintsummary.Get(cmd, autocomplete.GlobalExes[cmd])

	if pt.ExpectParam && len(r) == 0 {
		r = append(r, []rune(
			fmt.Sprintf("%starget `"+cmd+"` was not found%s", ansi.FgRed, ansi.Reset),
		)...)
	}

	if len(r) > 0 {
		return r
	}

	return hintCodeBlock()
}

func hintExpandVariables(line []rune) []rune {
	r, err := history.ExpandVariables(line, Prompt)
	if err != nil {
		return []rune(ansi.ExpandConsts("{RED}") + err.Error())
	}

	vars := variables.Expand(r)
	disclaimer := []rune{}
	if string(r) != string(vars) {
		disclaimer = []rune("(example only) ")
	}
	r = append(disclaimer, vars...)
	if string(line) == string(r) {
		r = []rune{}
	}

	if len(r) > 0 {
		s := strings.Replace(string(r), "\r", `\r`, -1)
		s = strings.Replace(s, "\n", `\n`, -1)
		s = strings.Replace(s, "\t", `\t`, -1)
		return []rune(s)
	}

	return []rune{}
}

func hintCodeBlock() []rune {
	if len(cachedHintText) > 0 {
		return cachedHintText
	}
	ht, err := lang.ShellProcess.Config.Get("shell", "hint-text-func", types.CodeBlock)
	if err != nil || len(ht.(string)) == 0 || ht.(string) == "{}" {
		return []rune{}
	}

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.Name.Set("shell (hint-text-func)")
	exitNum, err := fork.Execute([]rune(ht.(string)))

	b, err2 := fork.Stdout.ReadAll()
	if len(b) > 1 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	if len(b) > 1 && b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}

	if debug.Enabled && ( /*exitNum != 0 ||*/ err != nil || err2 != nil) {
		lang.ShellProcess.Stderr.Write([]byte(fmt.Sprintf(
			"Block returned false:\nExit Num: %d\nStdout length: %d\nStdout read error: %s\nStderr: %s\n",
			exitNum,
			len(b),
			err2,
			err,
		)))
	}

	cachedHintText = []rune(string(b))

	return cachedHintText
}
