package lang

type self struct {
	Parent     uint32
	Scope      uint32
	TTY        bool
	Method     bool
	Not        bool
	Background bool
	Module     string
}

func getVarSelf(p *Process) interface{} {
	return self{
		Parent:     p.Scope.Parent.Id,
		Scope:      p.Scope.Id,
		TTY:        p.Scope.Stdout.IsTTY(),
		Method:     p.Scope.IsMethod,
		Not:        p.Scope.IsNot,
		Background: p.Scope.IsBackground,
		Module:     p.Scope.FileRef.Source.Module,
	}
}

func getVarParams(p *Process) interface{} {
	return append([]string{p.Scope.Name}, p.Scope.Parameters.Params...)
}
