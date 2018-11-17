package proc

import (
	"errors"
	"sort"
	"sync"

	"github.com/lmorg/murex/debug"
)

// FID (Function ID) table: ie table of murex `proc.Process` processes
type funcID struct {
	procs  map[int]*Process
	mutex  sync.Mutex
	latest int
}

// fidList is the list of exported FIDs
type fidList []*Process

// Len returns the length of fidList - used purely for sorting FIDs
func (f fidList) Len() int { return len(f) }

// Less checks if one FID comes before another FID - used purely for sorting FIDs
func (f fidList) Less(i, j int) bool { return f[i].Id < f[j].Id }

// Swap alters the order of the exported FIDs - used purely for sorting FIDs
func (f fidList) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

// newFuncID creates new FID (Function ID) table of `proc.Process`es
func newFuncID() *funcID {
	f := new(funcID)
	f.procs = make(map[int]*Process)
	f.procs[0] = ShellProcess
	f.latest++
	return f
}

// Register process to assign it a FID (Function ID)
func (f *funcID) Register(p *Process) (fid int) {
	f.mutex.Lock()
	f.latest++
	f.procs[f.latest] = p
	fid = f.latest
	f.mutex.Unlock()
	p.Id = fid
	p.FidTree = append(p.FidTree, fid)
	p.Variables.process = p
	return
}

// Deregister removes function from the FID table
func (f *funcID) Deregister(fid int) {
	if debug.Enable {
		return
	}

	f.mutex.Lock()
	if f.procs[fid] != nil {
		delete(f.procs, fid)
	}
	f.mutex.Unlock()
	return
}

// Proc gets process by FID
func (f *funcID) Proc(fid int) (*Process, error) {
	if fid == 0 {
		return nil, errors.New("FID 0 is reserved for the shell.")
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.procs[fid] != nil {
		return f.procs[fid], nil
	}

	return nil, errors.New("Function ID does not exist.")
}

// ListAll processes registered in the FID (Function ID) table - return as a ordered list
func (f *funcID) ListAll() (procs fidList) {
	f.mutex.Lock()
	for id := range f.procs {
		procs = append(procs, f.procs[id])
	}
	f.mutex.Unlock()

	sort.Sort(procs)
	return
}

// Dump lists all processes registered in the FID (Function ID) table - return as an unsorted list (faster but less useful)
func (f *funcID) Dump() map[int]*Process {
	f.mutex.Lock()
	f.mutex.Unlock()
	return f.procs
}
