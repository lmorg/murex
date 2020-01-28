package lang

import (
	"errors"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/lmorg/murex/debug"
)

// FID (Function ID) table: ie table of murex `Process` processes
type funcID struct {
	list   map[uint32]*Process
	init   map[uint32]*Process
	mutex  sync.Mutex
	latest uint32
}

// newFuncID creates new FID (Function ID) table of `Process`es
func newFuncID() *funcID {
	f := new(funcID)
	f.init = make(map[uint32]*Process)
	f.list = make(map[uint32]*Process)
	return f
}

// Register process to assign it a FID (Function ID)
func (f *funcID) Register(p *Process) (fid uint32) {
	fid = atomic.AddUint32(&f.latest, 1)

	f.mutex.Lock()
	f.init[fid] = p
	f.mutex.Unlock()

	p.Id = fid
	p.FidTree = append(p.FidTree, fid)
	p.Variables.process = p
	return
}

// Executing moves the function from init to list
func (f *funcID) Executing(fid uint32) error {

	f.mutex.Lock()
	p := f.init[fid]

	if p == nil {
		f.mutex.Unlock()
		return errors.New("Function ID not in init map")
	}

	delete(f.init, fid)
	f.list[fid] = p
	f.mutex.Unlock()

	return nil
}

// Deregister removes function from the FID table
func (f *funcID) Deregister(fid uint32) {
	if debug.Enabled {
		return
	}

	f.mutex.Lock()
	delete(f.list, fid)
	f.mutex.Unlock()
}

// Proc gets process by FID
func (f *funcID) Proc(fid uint32) (*Process, error) {
	if fid == 0 {
		return nil, errors.New("FID 0 is reserved for the shell")
	}

	f.mutex.Lock()
	p := f.list[fid]

	if p != nil {
		f.mutex.Unlock()
		return p, nil
	}

	p = f.init[fid]
	f.mutex.Unlock()

	if p != nil {
		return nil, errors.New("Function hasn't started yet")
	}

	return nil, errors.New("Function ID does not exist")
}

// fidList is the list of exported FIDs
type fidList []*Process

// Len returns the length of fidList - used purely for sorting FIDs
func (f fidList) Len() int { return len(f) }

// Less checks if one FID comes before another FID - used purely for sorting FIDs
func (f fidList) Less(i, j int) bool { return f[i].Id < f[j].Id }

// Swap alters the order of the exported FIDs - used purely for sorting FIDs
func (f fidList) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

// ListAll processes registered in the FID (Function ID) table - return as a ordered list
func (f *funcID) ListAll() fidList {
	f.mutex.Lock()
	procs := make(fidList, len(f.list))
	var i int
	for _, p := range f.list {
		procs[i] = p
		i++
	}
	f.mutex.Unlock()

	sort.Sort(procs)
	return procs
}

/*// Dump lists all processes registered in the FID (Function ID) table - return as an unsorted list (faster but less useful)
func (f *funcID) Dump() dump map[int]*Process {
	f.mutex.Lock()
	r := f.procs
	f.mutex.Unlock()
	return r
}*/
