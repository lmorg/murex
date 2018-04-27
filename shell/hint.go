package shell

import (
	"strings"

	"github.com/lmorg/murex/builtins/core/docs"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/man"
)

var (
	manDescript    map[string]string = make(map[string]string)
	cachedHintText []rune
)

func hintText(line []rune, pos int) []rune {
	//deleteme1, _ := parse(line)
	//deleteme2, _ := utils.JsonMarshal(deleteme1, false)
	//return []rune(string(deleteme2))

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

	var cmd string

	pt, _ := parse(line)
	cmd = pt.FuncName

	if proc.GlobalAliases.Exists(cmd) {
		s := proc.GlobalAliases.Get(cmd)
		r = []rune("(alias: '" + strings.Join(s, "' '") + "') => ")
		cmd = s[0]
	}

	if proc.MxFunctions.Exists(cmd) {
		dig, _ := proc.MxFunctions.Digest(cmd)
		//return append(r, []rune("(murex function - preview not developed yet)")...)
		return append(r, []rune("(murex function) "+dig)...)
	}

	if proc.GoFunctions[cmd] != nil {
		syn := docs.Synonym[cmd]
		r = append(r, []rune(docs.Digest[syn])...)
		return r
	}

	var manPage []rune

	s := manDescript[cmd]
	if s != "" && s != "!" {
		manPage = []rune(manDescript[cmd])
	}

	if s != "!" && len(manPage) == 0 {
		f := man.GetManPages(cmd)
		manPage = []rune(man.ParseDescription(f))
		if len(manPage) == 0 {
			manDescript[cmd] = "!"
		} else {
			manDescript[cmd] = string(manPage)
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
	stderr := streams.NewStdin()
	/*exitNum, err := */ lang.RunBlockShellNamespace([]rune(ht.(string)), nil, stdout, stderr)

	b, _ /*err2*/ := stdout.ReadAll()
	if len(b) > 1 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	if len(b) > 1 && b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}

	if debug.Enable {
		b, _ := stderr.ReadAll()
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, string(b))
	}

	/*if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Block returned false.")
	}*/

	cachedHintText = []rune(string(b))

	return cachedHintText
}
