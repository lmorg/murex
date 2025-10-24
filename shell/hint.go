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
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/murex/utils/parser"
)

func hintText(line []rune, pos int) []rune {
	r := hintExpandVariables(line)
	if len(r) > 0 {
		return r
	}

	pt, _ := parser.Parse(line, pos)

	if pt.FuncName == "" {
		return HintCodeBlockCached()
	}

	cmd := pt.FuncName

	if cmd == "cd" && len(pt.Parameters) > 0 && len(pt.Parameters[0]) > 0 {
		path := variables.ExpandString(pt.Parameters[0])
		if path == "-" {
			hint, ok := hintCdPrevious()
			if ok {
				return []rune("Change directory: " + string(hint))
			}
			return hint
		}
		path = utils.NormalisePath(path)
		return []rune("Change directory: " + path)
	}

	// check if a custom summary has been set
	globalExes := autocomplete.GlobalExes.Get()
	r = hintsummary.Get(cmd, (*globalExes)[cmd])
	if len(r) > 0 {
		return r
	}

	return HintCodeBlockCached()
}

func hintCdPreviousPwdHistErr(err error) []rune {
	return []rune(fmt.Sprintf("unable to decode $PWDHIST: %s", err.Error()))
}

func hintCdPrevious() ([]rune, bool) {
	pwdHist, err := lang.ShellProcess.Variables.GetValue("PWDHIST")
	if err != nil {
		return hintCdPreviousPwdHistErr(err), false
	}

	pwdStrings, err := lists.GenericToString(pwdHist)
	if err != nil {
		return hintCdPreviousPwdHistErr(err), false
	}

	if len(pwdStrings) < 2 {
		return []rune("already at first directory in $PWDHIST"), false
	}

	return []rune(pwdStrings[len(pwdStrings)-2]), true
}

func hintExpandVariables(line []rune) []rune {
	r, err := history.ExpandVariables(line, Prompt)
	if err != nil {
		return []rune(ansi.ExpandConsts("{RED}") + err.Error())
	}

	// don't update if no changes
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

var _cachedHintText []rune

func HintCodeBlockCached() []rune {
	if len(_cachedHintText) > 0 {
		return _cachedHintText
	}

	return HintCodeBlock()
}

func HintCodeBlock() []rune {
	ht, fileRef, err := lang.ShellProcess.Config.GetFileRef("shell", "hint-text-func", types.CodeBlock)
	if err != nil || len(ht.(string)) == 0 || ht.(string) == "{}" {
		return []rune{}
	}

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.Name.Set("(hint-text-func)")
	fork.FileRef = fileRef
	exitNum, err := fork.Execute([]rune(ht.(string)))

	b, err2 := fork.Stdout.ReadAll()
	if len(b) > 1 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	if len(b) > 1 && b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}

	if debug.Enabled && (exitNum != 0 || err != nil || err2 != nil) {
		return ([]rune(fmt.Sprintf(
			"Block returned false: Exit Num: %d, Stdout length: %d, Stdout read error: %s, Stderr: %s",
			exitNum, len(b), err2, err)))
	}

	_cachedHintText = []rune(string(b))
	return _cachedHintText
}
