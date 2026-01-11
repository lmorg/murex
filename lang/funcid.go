package lang

import (
	"errors"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/state"
)

// FID (Function ID) table: ie table of murex `Process` processes
type funcID struct {
	list   map[uint32]*Process
	mutex  sync.Mutex
	latest uint32
}

// newFuncID creates new FID (Function ID) table of `Process`es
func newFuncID() *funcID {
	f := new(funcID)
	f.list = make(map[uint32]*Process)
	return f
}

// Register process to assign it a FID (Function ID)
func (f *funcID) Register(p *Process) (fid uint32) {
	fid = atomic.AddUint32(&f.latest, 1)

	f.mutex.Lock()

	f.list[fid] = p
	p.Id = fid
	p.Variables.process = p

	f.mutex.Unlock()

	return
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
	f.mutex.Unlock()

	if p != nil {
		return p, nil
	}

	return nil, errors.New("function ID does not exist")
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

// List processes registered in the FID (Function ID) table - return as a ordered list
func (f *funcID) List(options ...OptFidList) fidList {
	fids := f.ListAll()

	var procs fidList

	for _, p := range fids {
		//fmt.Printf("start: %d %s\n", p.Id, p.Name.String())
		inc := true
		for opt := range options {
			//fmt.Printf("opt %d: %d %s\n", opt, p.Id, p.Name.String())
			inc = inc && options[opt](p)
		}
		if inc {
			procs = append(procs, p)
		}
		//fmt.Printf("end: %d %s\n", p.Id, p.Name.String())
	}

	return procs
}

// Dump returns a list of FIDs in a format for `runtime` builtin
func (f *funcID) Dump() any {
	fidList := f.ListAll()

	dump := make([]any, len(fidList))

	for i, p := range fidList {
		dump[i] = p.Dump()
	}

	return dump
}

func (f *funcID) WaitOnChildState(parent *Process, state state.State) {
	for {
		time.Sleep(100 * time.Millisecond)
		fids := f.ListAll()
		for _, f := range fids {
			if f.Parent.Id == parent.Id && f.State.Get() >= state.Get() {
				return
			}
		}
	}
}
