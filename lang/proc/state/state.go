package state

import "sync/atomic"

// State is a thread safe FunctionState struct
type State struct {
	fs int32
}

// Set the state in a thread safe way
func (s *State) Set(fs FunctionState) {
	atomic.StoreInt32(&s.fs, int32(fs))
}

// Get the state in a thread safe way
func (s *State) Get() FunctionState {
	fs := atomic.LoadInt32(&s.fs)
	return FunctionState(fs)
}

// String is a stringer function for Get()
func (s *State) String() string {
	fs := atomic.LoadInt32(&s.fs)
	return FunctionState(fs).String()
}
