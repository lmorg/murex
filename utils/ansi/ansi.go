package ansi

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"os"
)

func AllowAnsi() bool {
	v, err := proc.GlobalConf.Get("shell", "add-colour", types.Boolean)
	if err != nil {
		return false
	}
	return v.(bool)
}

func Stream(std streams.Io, ansiCode, message string) (err error) {
	if std.IsTTY() && AllowAnsi() {
		_, err = std.Write([]byte(ansiCode + message + Reset))
		return
	}

	_, err = std.Write([]byte(message))
	return
}

func Streamln(std streams.Io, ansiCode, message string) (err error) {
	if std.IsTTY() && AllowAnsi() {
		_, err = std.Writeln([]byte(ansiCode + message + Reset))
		return
	}

	_, err = std.Writeln([]byte(message))
	return
}

func Stderr(ansiCode, message string) (err error) {
	if AllowAnsi() {
		_, err = os.Stderr.WriteString(ansiCode + message + Reset)
		return
	}
	_, err = os.Stderr.WriteString(message + utils.NewLineString)
	return
}

func Stderrln(ansiCode, message string) (err error) {
	if AllowAnsi() {
		_, err = os.Stderr.WriteString(ansiCode + message + utils.NewLineString + Reset)
		return
	}
	_, err = os.Stderr.WriteString(message + utils.NewLineString)
	return
}
