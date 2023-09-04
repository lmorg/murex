package management

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/hintsummary"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/murex/utils/which"
)

func init() {
	lang.DefineFunction("which", cmdWhich, types.String)
}

func cmdWhich(p *lang.Process) error {
	if p.Stdout.IsTTY() {
		return cmdWhichTty(p)
	}

	return cmdWhichFunction(p)
}

func cmdWhichTty(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)
	cmds := p.Parameters.StringArray()
	if len(cmds) == 0 {
		return whichUsage(p)
	}

	for i := range cmds {
		s := whichTtyString(cmds[i])
		if len(s) == 0 {
			p.Stdout.Writeln([]byte(fmt.Sprintf("%s => unknown", cmds[i])))
			p.ExitNum = 1
			continue
		}

		p.Stdout.Writeln([]byte(fmt.Sprintf("%s %s", cmds[i], s)))
	}

	return nil
}

const whichUsageMessage = "no commands specified"

func whichUsage(p *lang.Process) error {
	_, err := p.Stderr.Writeln([]byte(whichUsageMessage))
	p.ExitNum = 1
	return err
}

/////

func whichTtyString(cmd string) (s string) {
	summary := hintsummary.Summary.Get(cmd)
	var (
		aliasPrefix string
		exists      bool
	)

	if lang.GlobalAliases.Exists(cmd) {
		a := lang.GlobalAliases.Get(cmd)
		alias := make([]string, len(a))
		copy(alias, a)
		escape.CommandLine(alias)
		args := strings.Join(alias, " ")
		aliasPrefix = "=> (alias) " + args + " "
		cmd = alias[0]
		exists = true
	}

	if lang.MxFunctions.Exists(cmd) {
		if summary == "" {
			summary, _ = lang.MxFunctions.Summary(cmd)
		}
		if summary == "" {
			summary = "no summary written"
		}
		return aliasPrefix + "=> (murex function) " + summary
	}

	if lang.GoFunctions[cmd] != nil {
		if summary == "" {
			synonym := docs.Synonym[cmd]
			summary = docs.Summary[synonym]
		}
		if summary == "" {
			summary = "no doc written"
		}
		return aliasPrefix + "=> (builtin) " + summary
	}

	path := which.Which(cmd)
	if path != "" {
		if summary == "" {
			summary = man.ParseSummary(man.GetManPages(cmd))
		}
		if summary == "" {
			summary = "no man page found"
		}
		path = filepath.Clean(path)
		if resolved, err := os.Readlink(path); err == nil {
			path = path + " -> " + filepath.Clean(resolved)
		}
		return aliasPrefix + "=> (" + path + ") " + summary
	}

	if exists {
		return aliasPrefix
	}
	return ""
}

/////

func cmdWhichFunction(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	cmds := p.Parameters.StringArray()
	if len(cmds) == 0 {
		return whichUsage(p)
	}

	var success bool
	for i := range cmds {
		s := whichFunctionString(cmds[i])
		p.Stdout.Writeln([]byte(s))
		if s != "unknown" {
			success = true
		}
	}

	if !success {
		p.ExitNum = 1
	}
	return nil
}

func whichFunctionString(cmd string) (s string) {
	switch {
	case lang.GlobalAliases.Exists(cmd):
		return "alias"

	case lang.MxFunctions.Exists(cmd):
		return "function"

	case lang.GoFunctions[cmd] != nil:
		return "builtin"

	default:
		path := which.Which(cmd)
		if path != "" {
			return path
		}
		return "unknown"
	}
}
