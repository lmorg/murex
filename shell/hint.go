package shell

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/builtins/core/docs"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/man"
)

var (
	manDesc        = make(map[string]string)
	cachedHintText []rune
)

func hintText(line []rune, pos int) []rune {
	r, err := history.ExpandVariables(line, Prompt)
	if err != nil {
		return []rune("Error: " + err.Error())
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

	pt, _ := parse(line)
	cmd := pt.FuncName

	if cmd == "cd" && len(pt.Parameters) > 0 && len(pt.Parameters[0]) > 0 {
		path := utils.NormalisePath(pt.Parameters[0])
		return []rune("Change directory: " + path)
	}

	if proc.GlobalAliases.Exists(cmd) {
		s := proc.GlobalAliases.Get(cmd)
		r = []rune("(alias: '" + strings.Join(s, "' '") + "') => ")
		cmd = s[0]
	}

	if proc.MxFunctions.Exists(cmd) {
		summary, _ := proc.MxFunctions.Summary(cmd)
		return append(r, []rune("(murex function) "+summary)...)
	}

	if proc.GoFunctions[cmd] != nil {
		syn := docs.Synonym[cmd]
		r = append(r, []rune(docs.Summary[syn])...)
		return r
	}

	var manPage []rune

	s := manDesc[cmd]
	if s != "" && s != "!" {
		manPage = []rune(manDesc[cmd])
	}

	if s != "!" && len(manPage) == 0 {
		f := man.GetManPages(cmd)
		manPage = []rune(man.ParseDescription(f))
		if len(manPage) == 0 {
			manDesc[cmd] = "!"
		} else {
			manDesc[cmd] = string(manPage)
		}

	}

	if len(manPage) == 0 && len(r) > 0 {
		manPage = []rune("(no man page for `" + cmd + "` installed)")
	}

	r = append(r, manPage...)
	if len(r) > 0 {
		return r
	}

	if len(cachedHintText) > 0 {
		return cachedHintText
	}
	ht, err := proc.ShellProcess.Config.Get("shell", "hint-text-func", types.CodeBlock)
	if err != nil || len(ht.(string)) == 0 || ht.(string) == "{}" {
		return []rune{}
	}

	stdout := streams.NewStdin()
	branch := proc.ShellProcess.BranchFID()
	branch.IsBackground = true
	defer branch.Close()
	exitNum, err := lang.RunBlockExistingConfigSpace([]rune(ht.(string)), nil, stdout, nil, branch.Process)

	b, err2 := stdout.ReadAll()
	if len(b) > 1 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	if len(b) > 1 && b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}

	if debug.Enable && (exitNum != 0 || err != nil || len(b) == 0 || err2 != nil) {
		proc.ShellProcess.Stderr.Write([]byte(fmt.Sprintf(
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
