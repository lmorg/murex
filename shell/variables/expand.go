package variables

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/home"
)

var (
	rxVars = regexp.MustCompile(`(\$[_a-zA-Z0-9]+)`)
	rxHome = regexp.MustCompile(`(~[_\-.a-zA-Z0-9]+)`)
)

// ExpandString finds variables in a string and replaces it with the value of the variable
func ExpandString(line string) string {
	match := rxVars.FindAllString(line, -1)
	for i := range match {
		line = strings.Replace(line, match[i], proc.ShellProcess.Variables.GetString(match[i][1:]), -1)
	}

	match = rxHome.FindAllString(line, -1)
	for i := range match {
		line = rxHome.ReplaceAllString(line, home.UserDir(match[i][1:]))
	}

	line = strings.Replace(line, "~", home.MyDir, -1)
	return line
}

// Expand finds variables in a line and replaces it with the value of the variable
func Expand(line []rune) []rune {
	return []rune(ExpandString(string(line)))
}
