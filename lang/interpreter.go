package lang

import (
	"context"
	"strings"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/expressions/functions"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/lang/state"
)

var ParseBlock func(block []rune) (*[]functions.FunctionT, error)
var ParseExpression func([]rune, int, bool) (int, error)
var ParseStatementParameters func([]rune, *Process) (string, []string, error)

const ExpressionFunctionName = "expr"

func compile(tree *[]functions.FunctionT, parent *Process) (*[]Process, int) {
	if parent == nil {
		panic("nil parent")
	}

	if tree == nil {
		panic("nil tree")
	}

	rm := parent.RunMode
	if len(*tree) > 0 && (string((*tree)[0].Command) == "runmode" ||
		string((*tree)[0].Command) == "runmode:") {
		_, params, err := ParseStatementParameters((*tree)[0].Raw, parent)
		if err != nil {
			return nil, ErrUnableToParseParametersInRunmode
		}

		switch strings.Join(params, " ") {

		// function wide scopes

		case "unsafe function":
			rm = runmode.FunctionUnsafe

		case "try function":
			rm = runmode.FunctionTry

		case "trypipe function":
			rm = runmode.FunctionTryPipe

		case "tryerr function":
			rm = runmode.FunctionTryErr

		case "trypipeerr function":
			rm = runmode.FunctionTryPipeErr

		// module wide scopes

		case "unsafe module":
			rm = runmode.ModuleUnsafe
			ModuleRunModes[parent.FileRef.Source.Module] = rm

		case "try module":
			rm = runmode.ModuleTry
			ModuleRunModes[parent.FileRef.Source.Module] = rm

		case "trypipe module":
			rm = runmode.ModuleTryPipe
			ModuleRunModes[parent.FileRef.Source.Module] = rm

		case "tryerr module":
			rm = runmode.ModuleTryErr
			ModuleRunModes[parent.FileRef.Source.Module] = rm

		case "trypipeerr module":
			rm = runmode.ModuleTryPipeErr
			ModuleRunModes[parent.FileRef.Source.Module] = rm

		default:
			return nil, ErrInvalidParametersInRunmode
		}

		parent.Scope.RunMode = rm
		*tree = (*tree)[1:]
	}

	procs := make([]Process, len(*tree))

	for i := range *tree {
		procs[i].State.Set(state.MemAllocated)
		procs[i].raw = (*tree)[i].Raw
		procs[i].Name.SetRune((*tree)[i].CommandName())
		procs[i].Parameters.PreParsed = (*tree)[i].Parameters
		procs[i].namedPipes = (*tree)[i].NamedPipes
		procs[i].IsMethod = (*tree)[i].Properties.Method()
		procs[i].OperatorLogicAnd = (*tree)[i].Properties.LogicAnd()
		procs[i].OperatorLogicOr = (*tree)[i].Properties.LogicOr()
		procs[i].Background.Set(parent.Background.Get())
		procs[i].Parent = parent
		procs[i].Scope = parent.Scope
		procs[i].WaitForTermination = make(chan bool)
		procs[i].WaitForStopped = make(chan bool)
		procs[i].HasStopped = make(chan bool)
		procs[i].RunMode = rm //parent.RunMode
		procs[i].Config = parent.Config
		procs[i].Tests = parent.Tests
		procs[i].Variables = parent.Variables
		procs[i].CCEvent = parent.CCEvent
		procs[i].CCExists = parent.CCExists
		procs[i].FileRef = &ref.File{Source: parent.FileRef.Source}
		procs[i].Forks = NewForkManagement()

		procs[i].FileRef.Column = parent.FileRef.Column + (*tree)[i].ColumnN
		procs[i].FileRef.Line = (*tree)[i].LineN

		trace(&procs[i])

		// Define previous and next processes:
		switch {
		case i == 0:
			// first
			procs[0].Previous = parent
			if i == len(*tree)-1 {
				procs[0].Next = parent
			} else {
				procs[0].Next = &procs[1]
			}

		case i == len(*tree)-1:
			// last
			procs[i].Previous = &procs[i-1]
			procs[i].Next = parent

		case i > 0:
			// everything in the middle
			procs[i].Previous = &procs[i-1]
			procs[i].Next = &procs[i+1]

		default:
			// This condition should never happen,
			// but lets but a default catch and stack trace in just in case.
			panic("Failed in an unexpected way: Compile()->switch{default}")
		}

		// Define stdin interface:
		switch {
		case i == 0:
			// first
			procs[0].Stdin = parent.Stdin

		case (*tree)[i].Properties.NewChain():
			// new chain
			procs[i].Stdin = streams.NewStdin()
		}

		// Define stdout / stderr interfaces:
		switch {
		case (*tree)[i].Properties.PipeOut():
			if i+1 == len(procs) {
				return nil, ErrPipingToNothing
			}
			procs[i+1].Stdin = streams.NewStdin()
			procs[i].Stdout = procs[i].Next.Stdin
			procs[i].Stderr = procs[i].Parent.Stderr

		case (*tree)[i].Properties.PipeErr():
			if i+1 == len(procs) {
				return nil, ErrPipingToNothing
			}
			procs[i+1].Stdin = streams.NewStdin()
			procs[i].Stdout = procs[i].Parent.Stderr //Stdout
			procs[i].Stderr = procs[i].Next.Stdin

		default:
			procs[i].Stdout = procs[i].Parent.Stdout
			procs[i].Stderr = procs[i].Parent.Stderr
		}

		procs[i].Context, procs[i].Done = context.WithCancel(context.Background())
		procs[i].Kill = func() {
			procs[i].Stdin.ForceClose()
			procs[i].Stdout.ForceClose()
			procs[i].Stderr.ForceClose()
			procs[i].Done()
		}

		if len((*tree)[i].Cast) != 0 {
			procs[i].Stdin.SetDataType(string((*tree)[i].Cast))
		}
	}

	for i := range *tree {
		createProcess(&procs[i], !(*tree)[i].Properties.NewChain())
	}

	return &procs, 0
}

func checkTryErr(p *Process, exitNum *int) {
	outSize, _ := p.Stdout.Stats()
	errSize, _ := p.Stderr.Stats()

	if *exitNum < 1 && errSize > outSize {
		*exitNum = 1
	}
}
