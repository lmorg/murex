package process

import "sync"

type Name struct {
	mutex sync.RWMutex
	name  string
}

func (n *Name) Set(s string) {
	n.mutex.Lock()
	n.name = s
	n.mutex.Unlock()
}

func (n *Name) SetRune(r []rune) {
	n.mutex.Lock()
	n.name = string(r)
	n.mutex.Unlock()
}

func (n *Name) String() string {
	n.mutex.RLock()
	s := n.name
	n.mutex.RUnlock()

	return s
}

func (n *Name) Append(s string) {
	n.mutex.Lock()
	n.name += s
	n.mutex.Unlock()
}
