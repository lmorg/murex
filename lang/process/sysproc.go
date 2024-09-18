package process

import (
	"fmt"
	"os"
	"sync"
)

type systemProcessInheritance interface {
	Signal(sig os.Signal) error
	Kill() error
	Pid() int
	ExitNum() int
	State() *os.ProcessState
}

type SystemProcess struct {
	mutex       sync.Mutex
	inheritance systemProcessInheritance
	pid         chan int
}

func NewSystemProcessStruct() *SystemProcess {
	sp := new(SystemProcess)
	sp.pid = make(chan int, 2)
	return sp
}

func (sp *SystemProcess) Set(i systemProcessInheritance) {
	sp.mutex.Lock()
	sp.inheritance = i
	sp.pid <- i.Pid()
	sp.mutex.Unlock()
}

func (sp *SystemProcess) External() bool {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	return sp.inheritance != nil
}

var errNotDefined = fmt.Errorf("no system process defined")

func (sp *SystemProcess) Signal(sig os.Signal) error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	if sp.inheritance != nil {
		return sp.inheritance.Signal(sig)
	}
	return errNotDefined
}

func (sp *SystemProcess) Kill() error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	if sp.inheritance != nil {
		return sp.inheritance.Kill()
	}
	return errNotDefined
}

const NOT_A_SYSTEM_PROCESS = -1

func (sp *SystemProcess) Pid() int {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	if sp.inheritance != nil {
		return sp.inheritance.Pid()
	}
	return NOT_A_SYSTEM_PROCESS
}

func (sp *SystemProcess) ExitNum() int {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	if sp.inheritance != nil {
		return sp.inheritance.ExitNum()
	}
	return 1
}

func (sp *SystemProcess) State() *os.ProcessState {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	if sp.inheritance != nil {
		return sp.inheritance.State()
	}
	return nil
}

// WaitForPid should only be used in redirection.go
func (sp *SystemProcess) WaitForPid() int {
	return <-sp.pid
}
