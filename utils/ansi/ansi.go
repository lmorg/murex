package ansi

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

var rxAnsiConsts = regexp.MustCompile(`\{([-\^A-Z0-9]+)\}`)

// IsAllowed returns a boolean value depending on whether the shell is configured to allow ANSI colours
func IsAllowed() bool {
	v, err := lang.ShellProcess.Config.Get("shell", "color", types.Boolean)
	if err != nil {
		return false
	}
	return v.(bool)
}

// ExpandConsts writes a new string with the {CONST} values replaced
func ExpandConsts(s string) string {
	return expandConsts(s, !IsAllowed())
}

// ForceExpandConsts expands consts irrespective of user preferences. It is not
// recommended that you use this apart from in other testing functions.
func ForceExpandConsts(s string, noColour bool) string {
	return expandConsts(s, noColour)
}

func expandConsts(s string, noColour bool) string {
	match := rxAnsiConsts.FindAllStringSubmatch(s, -1)
	for i := range match {

		// misc escape sequences
		b := constants[match[i][1]]
		if len(b) != 0 {
			s = strings.Replace(s, match[i][0], string(b), -1)
			continue
		}

		// SGR (Select Graphic Rendition) parameters
		b = sgr[match[i][1]]
		if len(b) != 0 {
			if noColour {
				b = []byte{}
			}
			s = strings.Replace(s, match[i][0], string(b), -1)
			continue
		}

	}

	return s
}
