package lang

import (
	"sync"
	"time"
)

type ForkManagement struct {
	forks map[int64]*[]Process
	mutex sync.Mutex
}

func NewForkManagement() *ForkManagement {
	fm := new(ForkManagement)
	fm.forks = make(map[int64]*[]Process)

	return fm
}

func (fm *ForkManagement) add(procs *[]Process) (id int64) {
	fm.mutex.Lock()

getId:
	id = time.Now().UnixMicro()
	if fm.forks[id] != nil {
		goto getId
	}

	fm.forks[id] = procs

	fm.mutex.Unlock()
	return
}

func (fm *ForkManagement) delete(id int64) {
	fm.mutex.Lock()
	delete(fm.forks, id)
	fm.mutex.Unlock()
}

func (fm *ForkManagement) GetForks() []*[]Process {
	fm.mutex.Lock()

	forks := make([]*[]Process, len(fm.forks))
	var i int
	for _, procs := range fm.forks {
		forks[i] = procs
		i++
	}
	fm.mutex.Unlock()

	return forks
}
