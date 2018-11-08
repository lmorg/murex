package ansi

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

var rxAnsiConsts *regexp.Regexp = regexp.MustCompile(`\{([-\^A-Z0-9]+)\}`)

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

// Stderr writes colourised output to p.Stderr
func Stderr(p *proc.Process, ansiCode, message string) (err error) {
	if IsAllowed() {
		_, err = p.Stderr.Write([]byte(ansiCode + message + Reset))
		return
	}
	_, err = p.Stderr.Write([]byte(message + utils.NewLineString))
	return
}

// Stderrln writes colourised output to p.Stderr with an OS specific carriage return
func Stderrln(p *proc.Process, ansiCode, message string) (err error) {
	if IsAllowed() {
		_, err = p.Stderr.Write([]byte(ansiCode + message + utils.NewLineString + Reset))
		return
	}
	_, err = p.Stderr.Write([]byte(message + utils.NewLineString))
	return
}

// ExpandConsts writes a new string with the {CONST} values replaced
func ExpandConsts(s string) string {
	match := rxAnsiConsts.FindAllStringSubmatch(s, -1)
	for i := range match {
		b := constants[match[i][1]]
		if len(b) == 0 {
			continue
		}

		s = strings.Replace(s, match[i][0], string(b), -1)
	}

	return s
}
