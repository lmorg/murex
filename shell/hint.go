package shell

import (
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/builtins/core/docs"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/man"
)

var (
	manDesc        = make(map[string]string)
	cachedHintText []rune
)

func hintText(line []rune, pos int) []rune {
	r := hintExpandVariables(line)
	if len(r) > 0 {
		return r
	}

	pt, _ := parse(line)
	cmd := pt.FuncName

	if cmd == "cd" && len(pt.Parameters) > 0 && len(pt.Parameters[0]) > 0 {
		path := utils.NormalisePath(pt.Parameters[0])
		return []rune("Change directory: " + path)
	}

	if proc.GlobalAliases.Exists(cmd) {
		s := proc.GlobalAliases.Get(cmd)
		r = []rune("(alias) '" + strings.Join(s, "' '") + "' => ")
		cmd = s[0]
	}

	if proc.MxFunctions.Exists(cmd) {
		summary, _ := proc.MxFunctions.Summary(cmd)
		if summary == "" {
			summary = "no summary written"
		}
		return append(r, []rune("(murex function) "+summary)...)
	}

	if proc.GoFunctions[cmd] != nil {
		synonym := docs.Synonym[cmd]
		summary := docs.Summary[synonym]
		if summary == "" {
			summary = "no doc written"
		}
		r = append(r, []rune("(builtin) "+summary)...)
		return r
	}

	if autocomplete.GlobalExes[cmd] {
		manPage := manDesc[cmd]

		if manPage == "" {
			manPage = man.ParseDescription(man.GetManPages(cmd))
		}

		if manPage == "" {
			manPage = "no man page found"
		}

		manDesc[cmd] = manPage
		which := readlink(which(cmd))

		r = append(r, []rune("("+which+") "+manPage)...)
		if len(r) > 0 {
			return r
		}
	}

	return hintCodeBlock()
}

func hintExpandVariables(line []rune) []rune {
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

	return []rune{}
}

func which(cmd string) string {
	envPath := proc.ShellProcess.Variables.GetString("PATH")

	for _, path := range autocomplete.SplitPath(envPath) {
		filepath := path + consts.PathSlash + cmd
		_, err := os.Stat(filepath)
		if !os.IsNotExist(err) {
			return filepath
		}
	}

	return ""
}

func readlink(path string) string {
	/*f, err := os.Stat(path)
	if err != nil {
		return err.Error()
	}

	if f.Mode()&os.ModeSymlink != 0 {
		return path
	}*/

	ln, err := os.Readlink(path)
	if err != nil {
		return path
	}

	//if ln[0] != consts.PathSlash[0] {
	return ln //path + " => " + ln
	//}
}

func hintCodeBlock() []rune {
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
