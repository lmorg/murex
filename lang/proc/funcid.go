package proc

import (
	"errors"
	"sort"
	"sync"
)

type funcID struct {
	procs  map[int]*Process
	mutex  sync.Mutex
	latest int
}

type fidList []*Process

func (f fidList) Len() int           { return len(f) }
func (f fidList) Less(i, j int) bool { return f[i].Id < f[j].Id }
func (f fidList) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

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

func (f *funcID) ListAll() (procs fidList) {
	f.mutex.Lock()
	for id := range f.procs {
		procs = append(procs, f.procs[id])
	}
	f.mutex.Unlock()

	sort.Sort(procs)
	return
}
