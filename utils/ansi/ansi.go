package ansi

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"os"
)

func allowAnsi() bool {
	v, err := proc.GlobalConf.Get("shell", "add-colour", types.Boolean)
	if err != nil {
		return false
	}
	return v.(bool)
}

// Stream writes colourised output to a streams.Io interface
func Stream(std streams.Io, ansiCode, message string) (err error) {
	if std.IsTTY() && allowAnsi() {
		_, err = std.Write([]byte(ansiCode + message + Reset))
		return
	}

	_, err = std.Write([]byte(message))
	return
}

// Streamln writes colourised output to a streams.Io interface with an OS specific carriage return
func Streamln(std streams.Io, ansiCode, message string) (err error) {
	if std.IsTTY() && allowAnsi() {
		_, err = std.Writeln([]byte(ansiCode + message + Reset))
		return
	}

	_, err = std.Writeln([]byte(message))
	return
}

// Stderr writes colourised output to os.Stderr
func Stderr(ansiCode, message string) (err error) {
	if allowAnsi() {
		_, err = os.Stderr.WriteString(ansiCode + message + Reset)
		return
	}
	_, err = os.Stderr.WriteString(message + utils.NewLineString)
	return
}

// Stderrln writes colourised output to os.Stderr with an OS specific carriage return
func Stderrln(ansiCode, message string) (err error) {
	if allowAnsi() {
		_, err = os.Stderr.WriteString(ansiCode + message + utils.NewLineString + Reset)
		return
	}
	_, err = os.Stderr.WriteString(message + utils.NewLineString)
	return
}
