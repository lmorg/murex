package proc

import "sync"

// Setting a fixed PID pool below without any recycling of PIDs does severely hamper the usefulness of this shell.
// However since this is still very much an alpha project, I'm using this hard limit as a catch for runaway processes.
const pidPool int = 1024

type pid struct {
	sync.Mutex
	Process [pidPool]*Process
	count   int
}

func (pid *pid) Add(process *Process) {
	pid.Lock()
	pid.count++
	pid.Process[pid.count] = process
	pid.Unlock()
}

func (pid *pid) CountRunning() (i int) {
	pid.Lock()
	for j := range pid.Process {
		if !pid.Process[j].HasTerminated {
			i++
		}
	}
	pid.Unlock()
	return
}
