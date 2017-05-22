package proc

import "sync"

// TODO: need to write some method of recycling terminated PIDs.

type pid struct {
	sync.Mutex
	Process []*Process
	count   int
}

func (pid *pid) Add(process *Process) {
	pid.Lock()
	pid.count++
	pid.Process = append(pid.Process, process)
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
