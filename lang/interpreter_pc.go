//go:build !js
// +build !js

package lang

import (
	"github.com/lmorg/murex/lang/state"
)

//////////////////
//  Schedulers  //
//////////////////

func runModeNormal(procs *[]Process) (exitNum int) {
	var (
		prev         int
		skipPipeline bool
	)

	if len(*procs) == 0 {
		return 1
	}

	for i := range *procs {
		if i > 0 {

			prev = i - 1

			if (*procs)[i].IsMethod {
				go waitProcess(&(*procs)[prev])
			} else {
				waitProcess(&(*procs)[prev])
			}

			if ((*procs)[i].OperatorLogicAnd && (*procs)[prev].ExitNum != 0) ||
				((*procs)[i].OperatorLogicOr && (*procs)[prev].ExitNum == 0) ||
				(skipPipeline && ((*procs)[i].OperatorLogicAnd || (*procs)[i].OperatorLogicOr)) {

				(*procs)[i].SetTerminatedState(true)
				(*procs)[i].ExitNum = (*procs)[prev].ExitNum
				skipPipeline = true
			} else {
				skipPipeline = false
			}
		}

		go executeProcess(&(*procs)[i])
	}

	waitProcess(&(*procs)[len(*procs)-1])
	exitNum = (*procs)[len(*procs)-1].ExitNum

	return
}

// `try` - Last process in each pipe is checked.
func runModeTry(procs *[]Process, tryErr bool) (exitNum int) {
	if len((*procs)) == 0 {
		return 1
	}

	for i := 0; i < len(*procs); i++ {
		go executeProcess(&(*procs)[i])
		next := i + 1

		if next == len((*procs)) || !(*procs)[next].IsMethod {
			waitProcess(&(*procs)[i])
			exitNum = (*procs)[i].ExitNum

			if tryErr {
				checkTryErr(&(*procs)[i], &exitNum)
			}

			if next < len(*procs) {
				if exitNum < 1 && (*procs)[next].OperatorLogicOr {
					i++
					(*procs)[i].SetTerminatedState(true)
					(*procs)[i].Stdout.Close()
					(*procs)[i].Stderr.Close()
					GlobalFIDs.Deregister((*procs)[i].Id)
					(*procs)[i].State.Set(state.AwaitingGC)
					continue
				}

				if exitNum > 0 && !(*procs)[next].OperatorLogicOr {
					for i++; i < len(*procs); i++ {
						(*procs)[i].Stdout.Close()
						(*procs)[i].Stderr.Close()
						GlobalFIDs.Deregister((*procs)[i].Id)
						(*procs)[i].State.Set(state.AwaitingGC)
					}
					return
				}
			}

		} else {
			go waitProcess(&(*procs)[i])
		}
	}

	return
}

// `trypipe` - Each process in the pipeline is tried sequentially. Breaks parallelization.
func runModeTryPipe(procs *[]Process, tryPipeErr bool) (exitNum int) {
	if len(*procs) == 0 {
		return 1
	}

	for i := 0; i < len(*procs); i++ {
		go executeProcess(&(*procs)[i])
		waitProcess(&(*procs)[i])

		exitNum = (*procs)[i].ExitNum

		if tryPipeErr {
			checkTryErr(&(*procs)[i], &exitNum)
		}

		next := i + 1
		if next < len(*procs) {
			if exitNum < 1 && (*procs)[next].OperatorLogicOr {
				i++
				(*procs)[i].SetTerminatedState(true)
				(*procs)[i].Stdout.Close()
				(*procs)[i].Stderr.Close()
				GlobalFIDs.Deregister((*procs)[i].Id)
				(*procs)[i].State.Set(state.AwaitingGC)
				continue
			}

			if exitNum > 0 && !(*procs)[next].OperatorLogicOr {
				for i++; i < len(*procs); i++ {
					(*procs)[i].Stdout.Close()
					(*procs)[i].Stderr.Close()
					GlobalFIDs.Deregister((*procs)[i].Id)
					(*procs)[i].State.Set(state.AwaitingGC)
				}
				return
			}
		}
	}

	return
}
