package lang

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
)

func compile(tree *astNodes, parent *proc.Process) {
	for i := range *tree {
		(*tree)[i].Process.State = state.MemAllocated
		(*tree)[i].Process.Name = (*tree)[i].Name
		(*tree)[i].Process.Parameters.SetTokens((*tree)[i].ParamTokens)
		(*tree)[i].Process.IsMethod = (*tree)[i].Method
		(*tree)[i].Process.IsBackground = parent.IsBackground
		(*tree)[i].Process.Parent = parent
		(*tree)[i].Process.Scope = parent.Scope
		(*tree)[i].Process.WaitForTermination = make(chan bool)

		if (*tree)[i].LineNumber == 0 {
			(*tree)[i].Process.ColNumber = (*tree)[i].ColNumber + parent.ColNumber
		} else {
			(*tree)[i].Process.ColNumber = (*tree)[i].ColNumber
		}

		if parent.Id == 0 {
			(*tree)[i].Process.LineNumber = (*tree)[i].LineNumber + parent.LineNumber + 1
		} else {
			(*tree)[i].Process.LineNumber = (*tree)[i].LineNumber + parent.LineNumber
		}

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
		createProcess(&(*tree)[i].Process, !(*tree)[i].NewChain)
	}
}

// `evil` - Only use this if you are not concerned about STDERR nor exit number.
func runModeEvil(tree *astNodes) int {
	if len(*tree) == 0 {
		return 1
	}

	(*tree)[0].Process.Previous.SetTerminatedState(true)

	for i := range *tree {
		if i > 0 {
			if (*tree)[i].NewChain {
				waitProcess(&(*tree)[i-1].Process)
			} else {
				go waitProcess(&(*tree)[i-1].Process)
			}
		}

		(*tree)[i].Process.Stderr = new(streams.Null)
		go executeProcess(&(*tree)[i].Process)
	}

	waitProcess(&(*tree).Last().Process)
	return 0
}

func runModeNormal(tree *astNodes) (exitNum int) {
	if len(*tree) == 0 {
		return 1
	}

	(*tree)[0].Process.Previous.SetTerminatedState(true)

	for i := range *tree {
		if i > 0 {
			if (*tree)[i].NewChain {
				waitProcess(&(*tree)[i-1].Process)
			} else {
				go waitProcess(&(*tree)[i-1].Process)
			}
		}

		go executeProcess(&(*tree)[i].Process)
	}

	waitProcess(&(*tree).Last().Process)
	exitNum = (*tree).Last().Process.ExitNum
	return
}

// `try` - Last process in each pipe is checked.
func runModeTry(tree *astNodes) (exitNum int) {
	if len(*tree) == 0 {
		return 1
	}

	(*tree)[0].Process.Previous.SetTerminatedState(true)

	for i := range *tree {
		if i > 0 {
			if (*tree)[i].NewChain {
				waitProcess(&(*tree)[i-1].Process)
				exitNum = (*tree)[i-1].Process.ExitNum
				outSize, _ := (*tree)[i-1].Process.Stdout.Stats()
				errSize, _ := (*tree)[i-1].Process.Stderr.Stats()

				if exitNum == 0 && errSize > outSize {
					exitNum = 1
				}

				if exitNum != 0 {
					return
				}

			} else {
				go waitProcess(&(*tree)[i-1].Process)
			}
		}

		go executeProcess(&(*tree)[i].Process)
	}

	waitProcess(&(*tree).Last().Process)
	exitNum = (*tree).Last().Process.ExitNum
	outSize, _ := (*tree).Last().Process.Stdout.Stats()
	errSize, _ := (*tree).Last().Process.Stderr.Stats()

	if exitNum == 0 && errSize > outSize {
		exitNum = 1
	}

	return
}

// `trypipe` - Each process in the pipeline is tried sequentially. Breaks parallelisation.
func runModeTryPipe(tree *astNodes) (exitNum int) {
	debug.Log("Entering run mode `tryeach`")
	if len(*tree) == 0 {
		return 1
	}

	(*tree)[0].Process.Previous.SetTerminatedState(true)

	for i := range *tree {
		go executeProcess(&(*tree)[i].Process)
		waitProcess(&(*tree)[i].Process)

		exitNum = (*tree)[i].Process.ExitNum
		outSize, _ := (*tree)[i].Process.Stdout.Stats()
		errSize, _ := (*tree)[i].Process.Stderr.Stats()

		if exitNum == 0 && errSize > outSize {
			exitNum = 1
		}

		if exitNum != 0 {
			return
		}
	}

	return
}
