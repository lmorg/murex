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
	//procs map[int]*Process
	//mutex sync.Mutex
	//mutex  debug.BadMutex
	//latest int

	procs  sync.Map
	latest uint32
}

// newFuncID creates new FID (Function ID) table of `Process`es
func newFuncID() *funcID {
	f := new(funcID)
	/*f.procs = make(map[int]*Process)
	f.procs[0] = ShellProcess
	f.latest++*/
	f.procs.Store(uint32(0), ShellProcess)
	return f
}

// Register process to assign it a FID (Function ID)
func (f *funcID) Register(p *Process) (fid uint32) {
	/*f.mutex.Lock()
	f.latest++
	f.procs[f.latest] = p
	fid = f.latest
	f.mutex.Unlock()

	p.Id = fid
	p.FidTree = append(p.FidTree, fid)
	p.Variables.process = p
	return*/

	i := atomic.AddUint32(&f.latest, 1)
	f.procs.Store(i, p)
	p.Id = i
	p.FidTree = append(p.FidTree, i)
	p.Variables.process = p
	return i
}

// Deregister removes function from the FID table
func (f *funcID) Deregister(fid uint32) {
	if debug.Enabled {
		return
	}

	/*f.mutex.Lock()
	if f.procs[fid] != nil {
		delete(f.procs, fid)
	}
	f.mutex.Unlock()
	return*/
	f.procs.Delete(fid)
}

// Proc gets process by FID
func (f *funcID) Proc(fid uint32) (*Process, error) {
	if fid == 0 {
		return nil, errors.New("FID 0 is reserved for the shell")
	}

	/*f.mutex.Lock()
	p := f.procs[fid]
	f.mutex.Unlock()

	if p != nil {
		return p, nil
	}

	return nil, errors.New("Function ID does not exist")*/

	p, ok := f.procs.Load(fid)
	if ok {
		return p.(*Process), nil
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
func (f *funcID) ListAll() (procs fidList) {
	/*f.mutex.Lock()
	for id := range f.procs {
		procs = append(procs, f.procs[id])
	}
	f.mutex.Unlock()*/

	f.procs.Range(func(key interface{}, val interface{}) bool {
		p, ok := f.procs.Load(key.(uint32))
		if ok {
			procs = append(procs, p.(*Process))
		}
		return true
	})

	sort.Sort(procs)
	return
}

/*// Dump lists all processes registered in the FID (Function ID) table - return as an unsorted list (faster but less useful)
func (f *funcID) Dump() dump map[int]*Process {
	f.mutex.Lock()
	r := f.procs
	f.mutex.Unlock()
	return r
}*/
