package vars

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/home"
	"regexp"
	"strings"
)

var (
	rxVars *regexp.Regexp = regexp.MustCompile(`(\$[_a-zA-Z0-9]+)`)
	rxHome *regexp.Regexp = regexp.MustCompile(`(~[_\-.a-zA-Z0-9]+)`)
)

// ExpandVariablesString finds variables in a string and replaces it with the value of the variable
func ExpandVariablesString(line string) string {
	match := rxVars.FindAllString(line, -1)
	for i := range match {
		line = strings.Replace(line, match[i], proc.ShellProcess.VarGetString(match[i][1:]), -1)
	}

	match = rxHome.FindAllString(line, -1)
	for i := range match {
		line = rxHome.ReplaceAllString(line, home.UserDir(match[i][1:]))
	}

	line = strings.Replace(line, "~", home.MyDir, -1)
	return line
}

// ExpandVariables finds variables in a line and replaces it with the value of the variable
func ExpandVariables(line []rune) []rune {
	return []rune(ExpandVariablesString(string(line)))
}
