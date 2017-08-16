package shell

import (
	"bytes"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
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
	Dynamic string   // Use murex script to generate auto-complete options
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
		f := ExesFlags[exe]
		f.Flags = man.ScanManPages(exe)
		ExesFlags[exe] = f
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

func matchDynamic(partial, exe string, params []string) (items []string) {
	if len(ExesFlags[exe].Dynamic) == 0 {
		return
	}

	p := &proc.Process{
		Name:       exe,
		Parameters: parameters.Parameters{Params: params},
		Parent:     proc.ShellProcess,
	}
	p.Scope = p

	if !types.IsBlock([]byte(ExesFlags[exe].Dynamic)) {
		os.Stdout.WriteString("Dynamic autocompleter is not a code block!" + utils.NewLineString)
		return
	}
	block := []rune(ExesFlags[exe].Dynamic[1 : len(ExesFlags[exe].Dynamic)-1])

	stdout := streams.NewStdin()
	stderr := streams.NewStdin()
	exitNum, err := lang.ProcessNewBlock(block, nil, stdout, stderr, p)
	stdout.Close()
	stderr.Close()

	b, _ := stderr.ReadAll()
	os.Stderr.Write(b)

	if err != nil {
		os.Stdout.WriteString(err.Error() + utils.NewLineString)
		return
	}
	if exitNum != 0 {
		os.Stdout.WriteString("None zero exit number!" + utils.NewLineString)
		//return
	}

	stdout.ReadArray(func(b []byte) {
		s := string(bytes.TrimSpace(b))
		if len(s) == 0 {
			return
		}
		if strings.HasPrefix(s, partial) {
			items = append(items, s[len(partial):])
		}
	})

	return
}
