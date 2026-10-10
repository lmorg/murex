package autocomplete

import "sync"

type globalExesT struct {
	ptr   *map[string]bool
	mutex sync.Mutex
}

func NewGlobalExes() *globalExesT {
	ge := new(globalExesT)
	ptr := make(map[string]bool)
	ge.ptr = &ptr
	return ge
}

func (ge *globalExesT) Get() *map[string]bool {
	ge.mutex.Lock()
	ptr := ge.ptr
	ge.mutex.Unlock()
	return ptr
}

func (ge *globalExesT) Set(ptr *map[string]bool) {
	ge.mutex.Lock()
	ge.ptr = ptr
	ge.mutex.Unlock()
}

func (ge *globalExesT) List() []string {
	ge.mutex.Lock()
	defer ge.mutex.Unlock()

	slice := make([]string, len(*ge.ptr))
	var i int
	for exe := range *ge.ptr {
		slice[i] = exe
		i++
	}

	return slice
}
