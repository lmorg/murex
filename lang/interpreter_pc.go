//go:build !js
// +build !js

package lang

import (
	"github.com/lmorg/murex/lang/state"
)

//////////////////
//  Schedulers  //
//////////////////

// `evil` - Only use this if you are not concerned about STDERR nor exit number.
/*func runModeEvil(procs []Process) int {
	if len(procs) == 0 {
		return 1
	}

	procs[0].Previous.SetTerminatedState(true)

	for i := range procs {

		if i > 0 {
			if !procs[i].IsMethod {
				waitProcess(&procs[i-1])
			} else {
				go waitProcess(&procs[i-1])
			}
		}

		//if procs[i].Name == "break" {
		//	exitNum, _ := procs[i].Parameters.Int(0)
		//	return exitNum
		//}
		procs[i].Stderr = new(null.Null)
		go executeProcess(&procs[i])
	}

	waitProcess(&procs[len(procs)-1])
	return 0
}*/

func runModeNormal(procs []Process) (exitNum int) {
	//debug.Json("runModeNormal ()", procs)
	var (
		//skipToNextPipeline bool
		prev int
	)

	if len(procs) == 0 {
		return 1
	}

	procs[0].Previous.SetTerminatedState(true)

	for i := range procs {
		if i > 0 {
			prev = i - 1

			if procs[i].IsMethod {
				go waitProcess(&procs[prev])
			} else {
				waitProcess(&procs[prev])
			}

			if (procs[i].OperatorLogicAnd && procs[prev].ExitNum > 0) ||
				(procs[i].OperatorLogicOr && procs[prev].ExitNum < 1) {

				procs[i].hasTerminatedM.Lock()
				procs[i].hasTerminatedV = true
				procs[i].hasTerminatedM.Unlock()
				procs[i].ExitNum = procs[prev].ExitNum
			}
		}

		/*if procs[i].Name == "break" {
			exitNum, _ = procs[i].Parameters.Int(0)
			return
		}*/
		go executeProcess(&procs[i])
	}

	//debug.Json("runModeNormal (final waitProcess)", procs)
	waitProcess(&procs[len(procs)-1])
	exitNum = procs[len(procs)-1].ExitNum
	//debug.Json("runModeNormal (end)", procs)
	return
}

// `try` - Last process in each pipe is checked.
func runModeTry(procs []Process) (exitNum int) {
	//defer panic("fin")
	if len(procs) == 0 {
		return 1
	}

	procs[0].Previous.SetTerminatedState(true)

	for i := 0; i < len(procs); i++ {
		//fmt.Println(i, json.LazyLogging(procs))
		go executeProcess(&procs[i])
		next := i + 1

		if next == len(procs) || !procs[next].IsMethod {
			//fmt.Println("not", i)
			waitProcess(&procs[i])
			exitNum = procs[i].ExitNum
			outSize, _ := procs[i].Stdout.Stats()
			errSize, _ := procs[i].Stderr.Stats()

			if exitNum < 1 && errSize > outSize {
				exitNum = 1
			}

			if next < len(procs) {
				if exitNum < 1 && procs[next].OperatorLogicOr {
					//panic(json.LazyLogging(procs))
					i++
					procs[i].hasTerminatedM.Lock()
					procs[i].hasTerminatedV = true
					procs[i].hasTerminatedM.Unlock()
					procs[i].Stdout.Close()
					procs[i].Stderr.Close()
					GlobalFIDs.Deregister(procs[i].Id)
					procs[i].State.Set(state.AwaitingGC)
					//panic(i)
					continue
				}

				if exitNum > 0 && !procs[next].OperatorLogicOr {
					for i++; i < len(procs); i++ {
						procs[i].Stdout.Close()
						procs[i].Stderr.Close()
						GlobalFIDs.Deregister(procs[i].Id)
						procs[i].State.Set(state.AwaitingGC)
					}
					return
				}
			}

		} else {
			//fmt.Println("isa", i)
			go waitProcess(&procs[i])
		}
	}

	return
}

// `trypipe` - Each process in the pipeline is tried sequentially. Breaks parallelization.
func runModeTryPipe(procs []Process) (exitNum int) {
	//debug.Log("Entering run mode `trypipe`")
	if len(procs) == 0 {
		return 1
	}

	procs[0].Previous.SetTerminatedState(true)

	for i := 0; i < len(procs); i++ {
		/*if procs[i].Name == "break" {
			exitNum, _ = procs[i].Parameters.Int(0)
			return
		}*/
		go executeProcess(&procs[i])
		waitProcess(&procs[i])

		exitNum = procs[i].ExitNum
		outSize, _ := procs[i].Stdout.Stats()
		errSize, _ := procs[i].Stderr.Stats()

		if exitNum == 0 && errSize > outSize {
			exitNum = 1
		}

		next := i + 1
		if next < len(procs) {
			if exitNum < 1 && procs[next].OperatorLogicOr {
				//panic(json.LazyLogging(procs))
				i++
				procs[i].hasTerminatedM.Lock()
				procs[i].hasTerminatedV = true
				procs[i].hasTerminatedM.Unlock()
				procs[i].Stdout.Close()
				procs[i].Stderr.Close()
				GlobalFIDs.Deregister(procs[i].Id)
				procs[i].State.Set(state.AwaitingGC)
				//panic(i)
				continue
			}

			if exitNum > 0 && !procs[next].OperatorLogicOr {
				for i++; i < len(procs); i++ {
					procs[i].Stdout.Close()
					procs[i].Stderr.Close()
					GlobalFIDs.Deregister(procs[i].Id)
					procs[i].State.Set(state.AwaitingGC)
				}
				return
			}
		}
	}

	return
}
