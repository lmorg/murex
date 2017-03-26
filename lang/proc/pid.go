package proc

import "sync"

// Setting a fixed PID pool below without any recycling of PIDs does severely hamper the usefulness of this shell.
// However since this is still very much an alpha project, I'm using this hard limit as a catch for runaway processes.
const pidPool int = 1024

type Pid struct {
	sync.Mutex
	Process [pidPool]*Process
	count   int
}

func (pid *Pid) Add(process *Process) {
	pid.Lock()
	pid.count++
	pid.Process[pid.count] = process
	pid.Unlock()
}

func (pid *Pid) CountRunning() (i int) {
	pid.Lock()
	for j := range pid.Process {
		if !pid.Process[j].Terminated {
			i++
		}
	}
	pid.Unlock()
	return
}
