package process

import (
	"os"
	"sync"
)

type systemProcessInheritance interface {
	Signal(sig os.Signal) error
	Kill() error
	Pid() int
	ExitNum() int
	State() *os.ProcessState
	ForcedTTY() bool
}

type SystemProcess struct {
	mutex       sync.Mutex
	inheritance systemProcessInheritance
}

func NewSystemProcessStruct() *SystemProcess {
	return &SystemProcess{inheritance: nil}
}

func (sp *SystemProcess) getInheritance() systemProcessInheritance {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	return sp.inheritance
}

func (sp *SystemProcess) Set(i systemProcessInheritance) {
	sp.mutex.Lock()
	sp.inheritance = i
	sp.mutex.Unlock()
}

func (sp *SystemProcess) Defined() bool {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	return sp.inheritance != nil
}

func (sp *SystemProcess) Signal(sig os.Signal) error { return sp.getInheritance().Signal(sig) }
func (sp *SystemProcess) Kill() error                { return sp.getInheritance().Kill() }
func (sp *SystemProcess) Pid() int                   { return sp.getInheritance().Pid() }
func (sp *SystemProcess) ExitNum() int               { return sp.getInheritance().ExitNum() }
func (sp *SystemProcess) State() *os.ProcessState    { return sp.getInheritance().State() }
func (sp *SystemProcess) ForcedTTY() bool            { return sp.getInheritance().ForcedTTY() }
