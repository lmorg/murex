package ansi

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
)

var rxAnsiConsts = regexp.MustCompile(`\{([-\^A-Z0-9]+)\}`)

// IsAllowed returns a boolean value depending on whether the shell is configured to allow ANSI colours
func IsAllowed() bool {
	v, err := proc.ShellProcess.Config.Get("shell", "add-colour", types.Boolean)
	if err != nil {
		return false
	}
	return v.(bool)
}

// Stream writes colourised output to a stdio.Io interface
func Stream(std stdio.Io, ansiCode, message string) (err error) {
	if std.IsTTY() && IsAllowed() {
		_, err = std.Write([]byte(ansiCode + message + Reset))
		return
	}

	_, err = std.Write([]byte(message))
	return
}

// Streamln writes colourised output to a stdio.Io interface with an OS specific carriage return
func Streamln(std stdio.Io, ansiCode, message string) (err error) {
	if std.IsTTY() && IsAllowed() {
		_, err = std.Writeln([]byte(ansiCode + message + Reset))
		return
	}

	_, err = std.Writeln([]byte(message))
	return
}

// ExpandConsts writes a new string with the {CONST} values replaced
func ExpandConsts(s string) string {
	noColour := !IsAllowed()

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
