package lang

import "os"

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
		Background: p.Scope.Background.Get(),
		Module:     p.Scope.FileRef.Source.Module,
	}
}

func getVarParams(p *Process) interface{} {
	return append([]string{p.Scope.Name.String()}, p.Scope.Parameters.Params...)
}

func getVarMurexExe() interface{} {
	path, err := os.Executable()
	if err != nil {
		return err.Error()
	}

	return path
}
