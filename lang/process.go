package lang

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"os"
	"regexp"
	"strings"
)

var rxNamedPipeStdinOnly *regexp.Regexp = regexp.MustCompile(`^<[a-zA-Z0-9]+>$`)

func createProcess(p *proc.Process, f proc.Flow) {
	if rxNamedPipeStdinOnly.MatchString(p.Name) {
		p.Parameters.SetPrepend(p.Name[1 : len(p.Name)-1])
		p.Name = "<read-pipe>"
	}

	if p.Name[0] == '!' {
		p.IsNot = true
	}

	if (!proc.GlobalAliases.Exists(p.Name) || p.Parent.Name == "alias") &&
		p.Name[0] != '$' && proc.GoFunctions[p.Name].Func == nil {
		p.Parameters.SetPrepend(p.Name)
		// Make a special case of excluding `printf` from running inside a PTY as it hangs murex.
		// Obviously this shouldn't happen and in an ideal world I would fix murex instead of implementing this
		// horrible kludge. But I can live without `printf` being inside a PTY so I will class this bug as a low
		// priority.
		if f.NewChain && !f.PipeOut && !f.PipeErr && p.Name != "printf" {
			p.Name = "pty"
		} else {
			p.Name = "exec"
		}
	}

	p.IsMethod = !f.NewChain

	return
}

func executeProcess(p *proc.Process) {
	debug.Json("Executing:", p)

	parseRedirection(p)
	parseParameters(&p.Parameters, &proc.GlobalVars)

	// Echo
	echo, err := proc.GlobalConf.Get("shell", "Echo", types.Boolean)
	if err != nil {
		panic(err.Error())
	}
	if echo.(bool) {
		params := strings.Replace(strings.Join(p.Parameters.Params, `", "`), "\n", "\n# ", -1)
		os.Stdout.WriteString("# " + p.Name + `("` + params + `");` + utils.NewLineString)
	}

	p.Stderr.SetDataType(types.String)

	// Execute function.
	switch {
	case proc.GlobalAliases.Exists(p.Name) && p.Parent.Name != "alias":
		r := append(proc.GlobalAliases.Get(p.Name), []rune(" "+p.Parameters.StringAll())...)
		p.ExitNum, err = ProcessNewBlock(r, p.Stdin, p.Stdout, p.Stderr, "alias")

	case p.Name[0] == '$' && len(p.Name) > 1:
		s := proc.GlobalVars.GetString(p.Name[1:])
		p.Stdout.SetDataType(proc.GlobalVars.GetType(p.Name[1:]))
		_, err = p.Stdout.Write([]byte(s))

	//case proc.GoFunctions[p.Name].Func == nil:
	//	err = proc.GoFunctions[p.Name].Func(p)

	default:
		err = proc.GoFunctions[p.Name].Func(p)
	}

	p.Stdout.DefaultDataType(err != nil)

	if err != nil {
		p.Stderr.Writeln([]byte("Error in `" + p.Name + "`: " + err.Error()))
		if p.ExitNum == 0 {
			p.ExitNum = 1
		}
	}

	for !p.Previous.HasTerminated {
		// Code shouldn't really get stuck here.
		// This would only happen if someone abuses pipes on a function that has no stdin.
	}

	destroyProcess(p)
}

func waitProcess(p *proc.Process) {
	debug.Log("Waiting for", p.Name)
	p.HasTerminated = <-p.WaitForTermination
}

func destroyProcess(p *proc.Process) {
	debug.Json("Destroying:", p)
	p.Stdout.Close()
	p.Stderr.Close()
	p.WaitForTermination <- true
	debug.Log("Destroyed " + p.Name)
}
