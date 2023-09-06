package management

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/murex/utils/which"
)

const (
	typeAlias      = "alias"
	typeBuiltin    = "builtin"
	typePrivate    = "private"
	typeFunction   = "function"
	typeExecutable = "executable"
	typeUnknown    = "unknown"
)

func init() {
	lang.DefineFunction("type", cmdType, types.String)
}

func cmdType(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		return typeUsage(p)
	}

	if p.Stdout.IsTTY() {
		return cmtTypeTty(p)
	}

	return cmtTypeFunction(p)
}

func cmtTypeTty(p *lang.Process) error {
	p.Stdout.SetDataType(types.Generic)

	cmd, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	var (
		b   []byte
		msg string
		t   = typeOf(p, cmd)
	)

	switch t {
	case typeAlias:
		s := lang.GlobalAliases.Get(cmd)
		escape.CommandLine(s)
		b = []byte(strings.Join(s, " "))
		msg = fmt.Sprintf("`%s` is a shell %s:\n", cmd, t)

	case typePrivate:

	case typeFunction:
		r, err := lang.MxFunctions.Block(cmd)
		if err != nil {
			return err
		}
		b = []byte(string(r))
		msg = fmt.Sprintf("`%s` is a shell %s:\n", cmd, t)

	case typeBuiltin:
		s := fmt.Sprintf("run `murex-docs %s` for help", cmd)
		b = []byte(s)
		msg = fmt.Sprintf("`%s` is a Murex %s:\n", cmd, t)

	case typeExecutable:
		b = []byte(which.Which(cmd))
		msg = fmt.Sprintf("`%s` is an external %s:\n", cmd, t)

	case typeUnknown:
		p.ExitNum = 1
		msg = fmt.Sprintf("`%s` is %s :(\n", cmd, t)
	}

	_, _ = p.Stdout.Write([]byte(msg))
	_, err = p.Stdout.Writeln(b)
	return err
}

func cmtTypeFunction(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	cmd, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(typeOf(p, cmd)))
	return err
}

const typeUsageMessage = "no commands specified"

func typeUsage(p *lang.Process) error {
	_, err := p.Stderr.Writeln([]byte(typeUsageMessage))
	p.ExitNum = 1
	return err
}

func typeOf(p *lang.Process, cmd string) (s string) {
	switch {
	case lang.GlobalAliases.Exists(cmd):
		return typeAlias

	case lang.MxFunctions.Exists(cmd):
		return typeFunction

	case lang.PrivateFunctions.Exists(cmd, p.FileRef):
		return typePrivate

	case lang.GoFunctions[cmd] != nil:
		return typeBuiltin

	default:
		path := which.Which(cmd)
		if path != "" {
			return typeExecutable
		}

		return typeUnknown
	}
}
