package variables

import (
	"os/user"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/home"
)

var (
	rxVars = regexp.MustCompile(`(\$[_.a-zA-Z0-9]+)`)
	rxHome = regexp.MustCompile(`(~[_\-.a-zA-Z0-9]+)`)
)

// ExpandString finds variables in a string and replaces it with the value of the variable
func ExpandString(line string) string {
	escape := string([]byte{0, 1, 2, 0})
	hat := string([]byte{0, 1, 2, 1})

	line = strings.ReplaceAll(line, `\$`, escape)
	line = strings.ReplaceAll(line, `^$`, hat)

	match := rxVars.FindAllString(line, -1)
	for i := range match {
		s, _ := lang.ShellProcess.Variables.GetString(match[i][1:])
		line = strings.Replace(line, match[i], s, -1)
	}

	line = strings.ReplaceAll(line, escape, `\$`)
	line = strings.ReplaceAll(line, hat, `^$`)

	match = rxHome.FindAllString(line, -1)
	for i := range match {
		var home string
		usr, err := user.Lookup((match[i][1:]))
		if err == nil {
			home = usr.HomeDir
		}
		line = rxHome.ReplaceAllString(line, home)
	}

	line = strings.Replace(line, "~", home.MyDir, -1)
	return line
}

// Expand finds variables in a line and replaces it with the value of the variable
func Expand(line []rune) []rune {
	return []rune(ExpandString(string(line)))
}
