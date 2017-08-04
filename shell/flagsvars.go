package shell

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/man"
	"os"
	"sort"
	"strings"
)

// Struct to store auto-complete options
type Flags struct {
	NoFiles bool     // `true` to disable file name completion
	NoDirs  bool     // `true` to disable directory navigation completion
	Flags   []string // known supported command line flags for executable
}

// Map of executables and their supported auto-complete options
var ExesFlags map[string]Flags = make(map[string]Flags)

func allExecutables(includeBuiltins bool) map[string]bool {
	exes := make(map[string]bool)
	envPath := proc.GlobalVars.GetString("PATH")
	if envPath == "" {
		envPath = os.Getenv("PATH")
	}

	dirs := splitPath(envPath)

	for i := range dirs {
		listExes(dirs[i], &exes)
	}

	if !includeBuiltins {
		return exes
	}

	for name := range proc.GoFunctions {
		exes[name] = true
	}

	return exes
}

func matchFlags(partial, exe string) (items []string) {
	if len(ExesFlags[exe].Flags) == 0 {
		ExesFlags[exe] = Flags{Flags: man.ScanManPages(exe)}
	}

	for i := range ExesFlags[exe].Flags {
		//flag := strings.TrimSpace(ExesFlags[exe].Flags[i])
		flag := ExesFlags[exe].Flags[i]
		if flag == "" {
			continue
		}
		if strings.HasPrefix(flag, partial) {
			if flag[len(flag)-1] != '=' {
				flag += " "
			}
			items = append(items, flag[len(partial):])
		}
	}
	sort.Strings(items)
	return
}

func matchVars(partial string) (items []string) {
	vars := proc.GlobalVars.DumpMap()

	envVars := os.Environ()
	for i := range envVars {
		v := strings.Split(envVars[i], "=")
		vars[v[0]] = true
	}

	for name := range vars {
		if strings.HasPrefix(name, partial[1:]) {
			items = append(items, name[len(partial)-1:])
		}
	}
	sort.Strings(items)
	return
}
