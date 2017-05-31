package debug

// This is only used to test blocking conditions and shouldn't be used in release code.
type BadMutex struct{}

func (bm *BadMutex) Lock()   {}
func (bm *BadMutex) Unlock() {}
