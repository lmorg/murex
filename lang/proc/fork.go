package proc

import (
	"errors"
	"sync"
)

/*type Fork {
	fid uint64
	state int
	Proc *Process
}*/

type funcID struct {
	procs  map[int]*Process
	mutex  sync.Mutex
	latest int
}

func newFuncID() funcID {
	var f funcID
	f.procs = make(map[int]*Process)
	return f
}

func (f *funcID) Register(p *Process) (fid int) {
	f.mutex.Lock()
	f.latest++
	f.procs[f.latest] = p
	fid = f.latest
	f.mutex.Unlock()
	p.Id = fid
	return
}

func (f *funcID) Deregister(fid int) {
	f.mutex.Lock()
	if f.procs[fid] != nil {
		delete(f.procs, fid)
	}
	f.mutex.Unlock()
	return
}

func (f *funcID) Proc(fid int) (*Process, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	if f.procs[fid] != nil {
		return f.procs[fid], nil
	}

	return nil, errors.New("Function ID does not exist.")
}

func (f *funcID) ListAll() (procs []*Process) {
	f.mutex.Lock()
	for id := range f.procs {
		procs = append(procs, f.procs[id])
	}
	f.mutex.Unlock()
	return
}
