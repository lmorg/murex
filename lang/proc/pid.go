package proc

import "sync"

// TODO: need to completely rethink this as it's causing locking; so just disabled at this point in time.

type pid struct {
	sync.Mutex
	Process []*Process
	count   int
}

func (pid *pid) Add(process *Process) {
	return
	pid.Lock()
	pid.count++
	pid.Process = append(pid.Process, process)
	pid.Unlock()
}

func (pid *pid) CountRunning() (i int) {
	return
	pid.Lock()
	for j := range pid.Process {
		if !pid.Process[j].HasTerminated {
			i++
		}
	}
	pid.Unlock()
	return
}
