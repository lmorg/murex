package lang

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/envvars"
	"github.com/lmorg/murex/utils/path"
	"github.com/lmorg/readline/v4"
)

var envDataTypes = map[string][]string{
	types.Path:  {"HOME", "PWD", "OLDPWD", "SHELL", "TMPDIR", "HOMEBREW_CELLAR", "HOMEBREW_PREFIX", "HOMEBREW_REPOSITORY", "GOPATH", "GOROOT", "GOBIN"},
	types.Paths: {"PATH", "LD_LIBRARY_PATH", "MANPATH", "INFOPATH"},
}

func getVarSelf(p *Process) any {
	bg := p.Scope.Background.Get()
	return map[string]any{
		//"FID":         int(p.Parent.Id),
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

func getVarArgs(p *Process) any {
	return append([]string{p.Scope.Name.String()}, p.Scope.Parameters.StringArray()...)
}

func getVarMurexArgs() any {
	return os.Args
}

func getVarMurexExeValue() (any, error) {
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

func getPwdValue() (any, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return path.Unmarshal([]byte(pwd))
}

func getEnvVarValue(v *Variables) any {
	var err error
	evTable := make(map[string]any)
	envvars.All(evTable)
	for env, val := range evTable {
		val, err = v.getEnvValueValue(env, val.(string))
		if err == nil {
			evTable[env] = val
		}
	}
	return evTable
}

func getEnvVarString() any {
	evTable := make(map[string]any)
	envvars.All(evTable)
	return evTable
}

func getGlobalValues() any {
	m := make(map[string]any)

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

func getVarLinesValue() (int, error) {
	_, i, err := readline.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 25, fmt.Errorf("cannot assign value to $%s: %v", _VAR_LINES, err)
	}
	return i, nil
}

func getVarUserNameValue() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("cannot assign value to $%s nor $%s: %v", _VAR_USER, _VAR_LOGNAME, err)
	}

	return u.Username, nil
}

func getVarTmpDirValue() string {
	return os.TempDir()
}

func getVarRandomValue() int {
	return rand.Intn(32768)
}

func getVarOldPwdValue() (string, error) {
	v, err := ShellProcess.Variables.GetValue("PWDHIST")
	if err != nil {
		return "", fmt.Errorf("cannot assign value to $%s: %v", _VAR_OLDPWD, err)
	}

	slice, ok := v.([]any)
	if !ok {
		return "", fmt.Errorf("cannot assign value to $%s: $PWDHIST appears to be a %T, expecting []string", _VAR_OLDPWD, v)
	}

	switch len(slice) {
	case 0:
		return "", fmt.Errorf("cannot assign value to $%s: $PWDHIST appears to be a empty", _VAR_OLDPWD)

	case 1:
		return "", fmt.Errorf("cannot assign value to $%s: already at oldest entry in $PWDHIST", _VAR_OLDPWD)

	default:
		s, ok := slice[len(slice)-2].(string)
		if !ok {
			return "", fmt.Errorf("cannot assign value to $%s: $PWDHIST[-1] appears to be a %T, expecting string", _VAR_OLDPWD, slice[len(slice)-2])
		}
		return s, nil
	}
}
