package shell

import (
	"encoding/json"
	"github.com/lmorg/murex/lang/proc"
	"os"
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

func getExeFlags(exe string) (flags []string) {
	if ExesFlags[exe] == "" {
		return
	}

	//flags := make(Flags)
	//err = json.Unmarshal([]byte(ExesFlags[exe]), &flags)
	//if err != nil {
	//	return
	//}

	json.Unmarshal([]byte(ExesFlags[exe]), &flags)
	return
}
