package process

import (
	"os/exec"
	"sync"
)

type Exec struct {
	mutex sync.RWMutex
	pid   int
	cmd   *exec.Cmd
	Env   []string
}

func (exec *Exec) Set(pid int, cmd *exec.Cmd) {
	exec.mutex.Lock()
	exec.pid = pid
	exec.cmd = cmd
	exec.mutex.Unlock()
}

func (exec *Exec) Get() (int, *exec.Cmd) {
	exec.mutex.RLock()
	pid := exec.pid
	cmd := exec.cmd
	exec.mutex.RUnlock()

	return pid, cmd
}

func (exec *Exec) Pid() int {
	exec.mutex.RLock()
	pid := exec.pid
	exec.mutex.RUnlock()

	return pid
}
