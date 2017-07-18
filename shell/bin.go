package shell

import (
	"github.com/lmorg/murex/lang/proc"
	"os"
)

func allExecutables() map[string]bool {
	exes := make(map[string]bool)
	envPath := proc.GlobalVars.GetString("PATH")
	if envPath == "" {
		envPath = os.Getenv("PATH")
	}

	dirs := splitPath(envPath)

	for i := range dirs {
		listExes(dirs[i], &exes)
	}

	for name := range proc.GoFunctions {
		exes[name] = true
	}

	return exes
}
