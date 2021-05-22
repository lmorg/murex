package lang

import (
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/ref"
)

func compile(tree *AstNodes, parent *Process) (procs []Process, errNo int) {
	if parent == nil {
		panic("nil parent")
	}

	if tree == nil {
		panic("nil tree")
	}

	procs = make([]Process, len(*tree))

	for i := range *tree {
		procs[i].State.Set(state.MemAllocated)
		procs[i].Name = (*tree)[i].Name
		procs[i].IsMethod = (*tree)[i].Method
		procs[i].IsBackground = parent.IsBackground
		procs[i].Parent = parent
		procs[i].Scope = parent.Scope
		procs[i].WaitForTermination = make(chan bool)
		procs[i].RunMode = parent.RunMode
		procs[i].Config = parent.Config
		procs[i].Tests = parent.Tests
		procs[i].Variables = parent.Variables
		procs[i].Parameters.SetTokens((*tree)[i].ParamTokens)
		procs[i].Done = func() {}
		procs[i].Kill = func() {}
		//procs[i].hasTerminated = make(chan bool, 1)
		procs[i].PromptId = parent.PromptId

		procs[i].FileRef = &ref.File{Source: parent.FileRef.Source}

		if (*tree)[i].LineNumber == 0 {
			procs[i].FileRef.Column = (*tree)[i].ColNumber + parent.FileRef.Column
		} else {
			procs[i].FileRef.Column = (*tree)[i].ColNumber
		}

		if parent.Id == 0 {
			procs[i].FileRef.Line = (*tree)[i].LineNumber + parent.FileRef.Line + 1
		} else {
			procs[i].FileRef.Line = (*tree)[i].LineNumber + parent.FileRef.Line
		}

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

		case (*tree)[i].NewChain:
			// new chain
			procs[i].Stdin = streams.NewStdin()
		}

		// Define stdout / stderr interfaces:
		switch {
		case (*tree)[i].PipeOut:
			if i+1 == len(procs) {
				errNo = ErrPipingToNothing
				return
			}
			procs[i+1].Stdin = streams.NewStdin()
			procs[i].Stdout = procs[i].Next.Stdin
			procs[i].Stderr = procs[i].Parent.Stderr

		case (*tree)[i].PipeErr:
			if i+1 == len(procs) {
				errNo = ErrPipingToNothing
				return
			}
			procs[i+1].Stdin = streams.NewStdin()
			procs[i].Stdout = procs[i].Parent.Stderr //Stdout
			procs[i].Stderr = procs[i].Next.Stdin

		default:
			procs[i].Stdout = procs[i].Parent.Stdout
			procs[i].Stderr = procs[i].Parent.Stderr
		}

	}

	for i := range *tree {
		createProcess(&procs[i], !(*tree)[i].NewChain)
	}

	return
}
