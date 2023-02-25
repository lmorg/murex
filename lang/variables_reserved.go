package lang

import "os"

func getVarSelf(p *Process) interface{} {
	return map[string]interface{}{
		"Parent":     int(p.Scope.Parent.Id),
		"Scope":      int(p.Scope.Id),
		"TTY":        p.Scope.Stdout.IsTTY(),
		"Method":     p.Scope.IsMethod,
		"Not":        p.Scope.IsNot,
		"Background": p.Scope.Background.Get(),
		"Module":     p.Scope.FileRef.Source.Module,
	}
}

func getVarArgs(p *Process) interface{} {
	return append([]string{p.Scope.Name.String()}, p.Scope.Parameters.StringArray()...)
}

func getVarMurexExe() interface{} {
	path, err := os.Executable()
	if err != nil {
		return err.Error()
	}

	return path
}

func getHostname() string {
	name, _ := os.Hostname()
	return name
}

func getPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}
