package shell

import (
	"encoding/json"
	"github.com/lmorg/murex/lang/proc"
	"os"
	"sort"
	"strings"
)

var ExesFlags map[string]string = make(map[string]string)

type Flags []string

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

func matchFlags(partial, exe string) (items []string) {
	if ExesFlags[exe] == "" {
		return
	}

	//flags := make(Flags)
	//err = json.Unmarshal([]byte(ExesFlags[exe]), &flags)
	//if err != nil {
	//	return
	//}

	var flags []string
	json.Unmarshal([]byte(ExesFlags[exe]), &flags)

	for i := range flags {
		flag := strings.TrimSpace(flags[i])
		if flag == "" {
			continue
		}
		if strings.HasPrefix(flag, partial) {
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
