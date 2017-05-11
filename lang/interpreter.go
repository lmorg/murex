package lang

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
)

func compile(tree *Nodes, parent *proc.Process) {
	for i := range *tree {
		(*tree)[i].Process.Name = (*tree)[i].Name
		(*tree)[i].Process.Parameters.SetAll((*tree)[i].Parameters)
		(*tree)[i].Process.Method = (*tree)[i].Method
		(*tree)[i].Process.Parent = parent

		// Define previous and next processes:
		switch {
		case i == 0:
			// first
			(*tree)[0].Process.Previous = parent
			if i == len(*tree)-1 {
				(*tree)[0].Process.Next = parent
			} else {
				(*tree)[0].Process.Next = &(*tree)[1].Process
			}

		case i == len(*tree)-1:
			// last
			(*tree)[i].Process.Previous = &(*tree)[i-1].Process
			(*tree)[i].Process.Next = parent

		case i > 0:
			// everything in the middle
			(*tree)[i].Process.Previous = &(*tree)[i-1].Process
			(*tree)[i].Process.Next = &(*tree)[i+1].Process

		default:
			// This condition should never happen,
			// but lets but a default catch and stack trace in just in case.
			panic("Failed in an unexpected way: Compile()->switch{default}")
		}

		// Define stdin interface:
		switch {
		case i == 0:
			// first
			(*tree)[0].Process.Stdin = parent.Stdin

		case (*tree)[i].NewChain:
			// new chain
			(*tree)[i].Process.Stdin = streams.NewStdin()
			(*tree)[i].Process.Stdin.Close()
		}

		// Define stdout / stderr interfaces:
		switch {
		case (*tree)[i].PipeOut:
			(*tree)[i+1].Process.Stdin = streams.NewStdin()
			(*tree)[i].Process.Stdout = (*tree)[i].Process.Next.Stdin
			(*tree)[i].Process.Stderr = (*tree)[i].Process.Parent.Stderr

		case (*tree)[i].PipeErr:
			(*tree)[i+1].Process.Stdin = streams.NewStdin()
			(*tree)[i].Process.Stdout = (*tree)[i].Process.Parent.Stdout
			(*tree)[i].Process.Stderr = (*tree)[i].Process.Next.Stdin

		default:
			(*tree)[i].Process.Stdout = (*tree)[i].Process.Parent.Stdout
			(*tree)[i].Process.Stderr = (*tree)[i].Process.Parent.Stderr
		}

		// Not required for a single pass interpreter,
		// but I keep this code hanging about just in case I decide to expand the parser.
		//if len((*tree)[i].Children) > 0 {
		//	compile(&(*tree)[i].Children, &(*tree)[i].Process)
		//}
	}

	for i := range *tree {
		proc.CreateProcess(&(*tree)[i].Process, proc.Flow{
			NewChain: (*tree)[i].NewChain,
			PipeOut:  (*tree)[i].PipeOut,
			PipeErr:  (*tree)[i].PipeErr,
			Last:     i == len(*tree)-1,
		})
	}
}

func runNormal(tree *Nodes) (exitNum int) {
	if len(*tree) == 0 {
		return 1
	}

	(*tree)[0].Process.Previous.Terminated = true

	for i := range *tree {
		if (*tree)[i].NewChain && i > 0 {
			(*tree)[i-1].Process.Wait()
		}

		go (*tree)[i].Process.Execute()
	}

	(*tree).Last().Process.Wait()
	exitNum = (*tree).Last().Process.ExitNum
	return
}

func runHyperSensitive(tree *Nodes) (exitNum int) {
	debug.Log("Entering Hyper Sensitive mode!!!")
	if len(*tree) == 0 {
		return 1
	}

	(*tree)[0].Process.Previous.Terminated = true

	for i := range *tree {
		(*tree)[i].Process.Execute()
		exitNum = (*tree)[i].Process.ExitNum
		if exitNum != 0 {
			return
		}
	}

	return
}
