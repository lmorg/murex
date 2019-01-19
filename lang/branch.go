package lang

// BranchFID is used to branch a process to create a sandboxed function ID.
// This function produces a `go vet` error due to copying mutexes. However this
// is expected behaviour in here and accounted for so the code should still be
// thread safe.
func (p *Process) BranchFID() *Branch {
	p.hasTerminatedM.Lock()

	// This copies a sync.Mutex value, but on this occasion it is perfectly safe.
	process := *p

	p.hasTerminatedM.Unlock()
	process.hasTerminatedM.Unlock()

	process.Name += " (branch)"
	process.Config = process.Config.Copy()
	process.Variables = ReferenceVariables(p.Variables)
	process.Stdout.Open()
	process.Stderr.Open()

	branch := new(Branch)
	branch.Process = &process

	GlobalFIDs.Register(branch.Process)

	return branch
}

// Branch is the structure returned from BranchFID. Use this to close a branch.
type Branch struct {
	*Process
}

// Close the branch
func (branch Branch) Close() {
	DeregisterProcess(branch.Process)
}
