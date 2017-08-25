package debug

// BadMutex is only used to test deadlocks and shouldn't be used in release code.
type BadMutex struct{}

// Lock is a fake mutex lock used to check deadlocks
func (bm *BadMutex) Lock() {}

// Lock is a fake mutex Unlock used to check deadlocks
func (bm *BadMutex) Unlock() {}
