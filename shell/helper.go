package shell

import (
	"strings"

	"github.com/lmorg/murex/builtins/core/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/man"
)

var manDescript map[string]string = make(map[string]string)

func tabCompletion(line []rune, pos int) (prefix string, items []string) {
	if len(line) > pos-1 {
		line = line[:pos]
	}

	pt, _ := parse(line)

	switch {
	case pt.Variable != "":
		var s string
		if pt.VarLoc < len(line) {
			s = strings.TrimSpace(string(line[pt.VarLoc:]))
		}
		s = pt.Variable + s
		//retPos = len(s)
		prefix = s

		items = autocomplete.MatchVars(s)

	case pt.ExpectFunc:
		var s string
		if pt.Loc < len(line) {
			s = strings.TrimSpace(string(line[pt.Loc:]))
		}
		//retPos = len(s)
		prefix = s
		items = autocomplete.MatchFunction(s)

	default:
		var s string
		if len(pt.Parameters) > 0 {
			s = pt.Parameters[len(pt.Parameters)-1]
		}
		//retPos = len(s)
		prefix = s

		autocomplete.InitExeFlags(pt.FuncName)

		pIndex := 0
		items = autocomplete.MatchFlags(autocomplete.ExesFlags[pt.FuncName], s, pt.FuncName, pt.Parameters, &pIndex)
	}

	v, err := proc.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = 4
	}

	limitSuggestions := v.(int)
	if len(items) < limitSuggestions || limitSuggestions < 0 {
		limitSuggestions = len(items)
	}
	//Instance.Config.MaxCompleteLines = limitSuggestions
	Prompt.MaxTabCompleterRows = limitSuggestions

	/*suggest = make([][]rune, len(items))
	for i := range items {
		if len(items[i]) == 0 {
			continue
		}

		if !pt.QuoteSingle && !pt.QuoteDouble && len(items[i]) > 1 && strings.Contains(items[i][:len(items[i])-1], " ") {
			items[i] = strings.Replace(items[i], " ", `\ `, -1)
		}

		if items[i][len(items[i])-1] == '/' || items[i][len(items[i])-1] == '=' {
			suggest[i] = []rune(items[i])
		} else {
			suggest[i] = []rune(items[i] + " ")
		}
	}*/

	for i := range items {
		if len(items[i]) == 0 {
			items[i] = " "
		}
		if items[i][len(items[i])-1] != ' ' && items[i][len(items[i])-1] != '=' && items[i][len(items[i])-1] != '/' {
			items[i] += " "
		}
	}

	return
}

func syntaxCompletion(line []rune, pos int) ([]rune, int) {
	pt, _ := parse(line)
	switch {
	case pt.QuoteSingle:
		if pos < len(line)-1 || line[pos] != '\'' {
			return append(line, '\''), pos
		}

	case pt.QuoteDouble:
		if pos < len(line)-1 || line[pos] != '"' {
			return append(line, '"'), pos
		}

	case pt.QuoteBrace > 0:
		if pos < len(line)-1 || line[pos] != '(' {
			return append(line, ')'), pos
		}

	case pt.QuoteBrace < 0:
		if line[pos] == ')' && line[len(line)-1] == ')' && pos != len(line)-1 {
			return line[:len(line)-1], pos
		}

	case pt.NestedBlock > 0:
		if pos < len(line)-1 || line[pos] != '{' {
			return append(line, '}'), pos
		}

	case pt.NestedBlock < 0:
		if line[pos] == '}' && line[len(line)-1] == '}' && pos != len(line)-1 {
			return line[:len(line)-1], pos
		}

	case pos > 0 && line[pos-1] == '[':
		if pos < len(line)-1 {
			r := append(line[:pos+1], ']')
			return append(r, line[pos+2:]...), pos
		}
		return append(line, ']'), pos

	}
	return line, pos
}

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

	ht, err := proc.ShellProcess.Config.Get("shell", "hint-text-func", types.CodeBlock)
	if err != nil {
		return []rune{}
	}

	out := streams.NewStdin()
	exitNum, err := lang.RunBlockShellNamespace([]rune(ht.(string)), nil, out, nil)
	out.Close()

	b, err2 := out.ReadAll()
	if len(b) > 1 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	if len(b) > 1 && b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}

	if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Block returned false.")
	}

	return []rune(string(b))
}
