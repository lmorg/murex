package lang

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/consts"
)

var (
	rxNamedPipeStdinOnly *regexp.Regexp = regexp.MustCompile(`^<[a-zA-Z0-9]+>$`)
	rxVariables          *regexp.Regexp = regexp.MustCompile(`^\$([_a-zA-Z0-9]+)(\[(.*?)\]|)$`)
)

func createProcess(p *proc.Process, isMethod bool) {
	// Create empty function so we don't have to check nil state when invoking kill, ie you try to kill a process
	// before it's fully started
	p.Kill = func() {}

	proc.GlobalFIDs.Register(p) // This also registers the variables process
	p.CreationTime = time.Now()

	parseRedirection(p)

	if rxNamedPipeStdinOnly.MatchString(p.Name) {
		p.Parameters.SetPrepend(p.Name[1 : len(p.Name)-1])
		p.Name = consts.NamedPipeProcName
	}

	if p.Name[0] == '!' {
		p.IsNot = true
	}

	p.IsMethod = isMethod

	// We do stderr first so we can log errors in the stdout pipe to stderr
	switch p.NamedPipeErr {
	case "":
		p.NamedPipeErr = "err"
	case "err":
		//p.Stderr.Writeln([]byte("Invalid usage of named pipes: stderr defaults to <err>."))
	case "out":
		p.Stderr.SetDataType(types.Generic)
		p.Stderr = p.Next.Stdout
	default:
		p.Stderr.SetDataType(types.Generic)
		pipe, err := proc.GlobalPipes.Get(p.NamedPipeErr)
		if err == nil {
			p.Stderr = pipe
		} else {
			p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
		}
	}

	// We do stdout last so we can log errors in the stdout pipe to stderr
	switch p.NamedPipeOut {
	case "":
		p.NamedPipeOut = "out"
	case "err":
		p.Stdout.SetDataType(types.Null)
		p.Stdout = p.Next.Stderr
	case "out":
		//p.Stderr.Writeln([]byte("Invalid usage of named pipes: stdout defaults to <out>."))
	default:
		p.Stdout.SetDataType(types.Null)
		pipe, err := proc.GlobalPipes.Get(p.NamedPipeOut)
		if err == nil {
			p.Stdout = pipe
		} else {
			p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
		}
	}

	// Test cases
	if p.NamedPipeTest != "" {
		var stdout2, stderr2 *streams.Stdin
		p.Stdout, stdout2 = streams.NewTee(p.Stdout)
		p.Stderr, stderr2 = streams.NewTee(p.Stderr)
		err := p.Tests.SetStreams(p.NamedPipeTest, stdout2, stderr2, &p.ExitNum)
		if err != nil {
			p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
		}
	}

	p.Stdout.Open()
	p.Stderr.Open()

	p.Stderr.SetDataType(types.Generic)

	p.State = state.Assigned

	// Lets run `pipe` and `test` ahead of time to fudge the use of named pipes
	if p.Name == "pipe" || p.Name == "test" {
		ParseParameters(p, &p.Parameters)
		err := proc.GoFunctions[p.Name](p)
		if err != nil {
			ansi.Streamln(proc.ShellProcess.Stderr, ansi.FgRed, fmt.Sprintf("Error in `%s` (%d,%d): %s", p.Name, p.LineNumber, p.ColNumber, err.Error()))
			if p.ExitNum == 0 {
				p.ExitNum = 1
			}
		}
		p.SetTerminatedState(true)
		p.State = state.Executed
	}

	return
}

func executeProcess(p *proc.Process) {
	if p.HasTerminated() {
		destroyProcess(p)
		return
	}

	var err error

	p.State = state.Starting

	echo, err := p.Config.Get("shell", "echo", types.Boolean)
	if err != nil {
		echo = false
	}

	//debug.Json("Executing:", p)

	// Create a kill switch
	if p.Name != "exec" {
		p.Kill = func() { destroyProcess(p) }
	}

	//if !p.IsBackground {
	//	proc.KillForeground = p.Kill
	//	proc.ForegroundProc = p
	//}

	ParseParameters(p, &p.Parameters)

	// Execute function.
	var parsedAlias bool
	p.State = state.Executing
	p.StartTime = time.Now()

executeProcess:
	if echo.(bool) {
		params := strings.Replace(strings.Join(p.Parameters.Params, `", "`), "\n", "\n# ", -1)
		os.Stdout.WriteString("# " + p.Name + `("` + params + `");` + utils.NewLineString)
	}

	// execution mode:
	switch {
	case proc.GlobalAliases.Exists(p.Name) && p.Parent.Name != "alias" && parsedAlias == false:
		// murex aliases
		alias := proc.GlobalAliases.Get(p.Name)
		p.Name = alias[0]
		p.Parameters.Params = append(alias[1:], p.Parameters.Params...)
		parsedAlias = true
		goto executeProcess

	case proc.MxFunctions.Exists(p.Name):
		// murex functions
		var r []rune
		p.Scope = p
		r, err = proc.MxFunctions.Block(p.Name)
		if err == nil {
			p.ExitNum, err = RunBlockNewConfigSpace(r, p.Stdin, p.Stdout, p.Stderr, p)
		}

	case p.Name[0] == '$':
		// variables as functions
		match := rxVariables.FindAllStringSubmatch(p.Name+p.Parameters.StringAll(), -1)
		switch {
		case len(p.Name) == 1:
			err = errors.New("Variable token, `$`, used without specifying variable name.")
		case len(match) == 0 || len(match[0]) == 0:
			b, _ := json.MarshalIndent(match, "", "\t")
			fmt.Println(p.Name, p.Parameters.StringArray(), string(b))
			err = errors.New("`" + p.Name[1:] + "` is not a valid variable name.")
		case match[0][2] == "":
			s := p.Variables.GetString(match[0][1])
			p.Stdout.SetDataType(p.Variables.GetDataType(match[0][1]))
			_, err = p.Stdout.Write([]byte(s))
		default:
			block := []rune("$" + match[0][1] + "->[" + match[0][3] + "]")
			RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stderr, p)
		}

	case proc.GoFunctions[p.Name] != nil:
		// murex builtins
		err = proc.GoFunctions[p.Name](p)

	default:
		// shell execute
		p.Parameters.Params = append([]string{p.Name}, p.Parameters.Params...)
		p.Name = "exec"
		err = proc.GoFunctions[p.Name](p)
	}

	p.Stdout.DefaultDataType(err != nil)

	if err != nil {
		//ansi.Streamln(p.Stderr, ansi.FgRed, "Error in `"+p.Name+"`: "+err.Error())
		ansi.Streamln(p.Stderr, ansi.FgRed, fmt.Sprintf("Error in `%s` (%d,%d): %s", p.Name, p.LineNumber, p.ColNumber, err.Error()))
		if p.ExitNum == 0 {
			p.ExitNum = 1
		}
	}

	p.State = state.Executed

	if p.NamedPipeTest != "" {
		testEnabled, err := p.Config.Get("test", "enabled", types.Boolean)
		if err == nil && testEnabled.(bool) {
			p.Tests.Compare(p.NamedPipeTest, p)
		}
	}

	for !p.Previous.HasTerminated() {
		// Code shouldn't really get stuck here.
		// This would only happen if someone abuses pipes on a function that has no stdin.
	}

	destroyProcess(p)
}

func waitProcess(p *proc.Process) {
	//debug.Log("Waiting for", p.Name)
	<-p.WaitForTermination
}

func destroyProcess(p *proc.Process) {
	// Make special case for `bg` because that doesn't wait. Also make a special
	// case for `pipe` and `test` because they run out-of-band
	if p.Name != "bg" {
		p.WaitForTermination <- false
	}

	proc.DeregisterProcess(p)
}
