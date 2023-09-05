package lang

import (
	"os"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/envvars"
	"github.com/lmorg/murex/utils/path"
	"github.com/lmorg/murex/utils/readline"
)

var envDataTypes = map[string][]string{
	types.Path:  {"HOME", "PWD", "OLDPWD", "SHELL", "HOMEBREW_CELLAR", "HOMEBREW_PREFIX", "HOMEBREW_REPOSITORY"},
	types.Paths: {"PATH", "LD_LIBRARY_PATH", "MANPATH", "INFOPATH"},
}

func getVarSelf(p *Process) interface{} {
	bg := p.Scope.Background.Get()
	return map[string]interface{}{
		"Parent":      int(p.Scope.Parent.Id),
		"Scope":       int(p.Scope.Id),
		"TTY":         p.Scope.Stdout.IsTTY(),
		"Method":      p.Scope.IsMethod,
		"Interactive": Interactive && !bg,
		"Not":         p.Scope.IsNot,
		"Background":  bg,
		"Module":      p.Scope.FileRef.Source.Module,
	}
}

func getVarArgs(p *Process) interface{} {
	return append([]string{p.Scope.Name.String()}, p.Scope.Parameters.StringArray()...)
}

func getVarMurexArgs() interface{} {
	return os.Args
}

func getVarMurexExeValue() (interface{}, error) {
	pwd, err := os.Executable()
	if err != nil {
		return nil, err
	}

	return path.Unmarshal([]byte(pwd))
}

func getHostname() string {
	name, _ := os.Hostname()
	return name
}

func getPwdValue() (interface{}, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return path.Unmarshal([]byte(pwd))
}

func getEnvVarValue() interface{} {
	v := make(map[string]interface{})
	envvars.All(v)
	return v
}

func getGlobalValues() interface{} {
	m := make(map[string]interface{})

	GlobalVariables.mutex.Lock()
	for name, v := range GlobalVariables.vars {
		m[name] = v.Value
	}
	GlobalVariables.mutex.Unlock()

	return m
}

func getVarColumnsValue() int {
	return readline.GetTermWidth()
}
