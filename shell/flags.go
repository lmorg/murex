package shell

import (
	"bytes"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/man"
	"os"
	"sort"
	"strings"
)

// Flags is a struct to store auto-complete options
type Flags struct {
	NoFiles    bool     // `true` to disable file name completion
	NoDirs     bool     // `true` to disable directory navigation completion
	NoFlags    bool     // `true` to disable Flags[] slice and man page parsing
	IncExePath bool     // `true` to include binaries in $PATH
	Flags      []string // known supported command line flags for executable
	Dynamic    string   // Use murex script to generate auto-complete options
}

// ExesFlags is map of executables and their supported auto-complete options.
// We might as well pre-populate the structure with a few base commands we might expect.
var ExesFlags map[string]Flags = map[string]Flags{
	"cd":      Flags{Flags: []string{}, NoFiles: true},
	"mkdir":   Flags{Flags: []string{}, NoFiles: true},
	"rmdir":   Flags{Flags: []string{}, NoFiles: true},
	"man":     Flags{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true},
	"which":   Flags{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true},
	"whereis": Flags{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true},
	"sudo":    Flags{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true},
	"exec":    Flags{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true},
	"pty":     Flags{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true},
}

// globalExes is a pre-populated list of all executables in $PATH.
// The point of this is to speed up exe auto-completion.
var globalExes map[string]bool = make(map[string]bool)

// UpdateGlobalExeList generates a list of executables in $PATH. This used to be called upon demand but it caused a
// slight but highly annoying pause if murex had been sat idle for a while. So now it's an exported function so it can
// be run as a background job or upon user request.
func UpdateGlobalExeList() {
	envPath := proc.GlobalVars.GetString("PATH")
	if envPath == "" {
		envPath = os.Getenv("PATH")
	}

	dirs := splitPath(envPath)

	for i := range dirs {
		listExes(dirs[i], globalExes)
	}
}

func allExecutables(includeBuiltins bool) map[string]bool {
	exes := make(map[string]bool)
	for k, v := range globalExes {
		exes[k] = v
	}

	if !includeBuiltins {
		return exes
	}

	for name := range proc.GoFunctions {
		exes[name] = true
	}

	proc.MxFunctions.UpdateMap(exes)
	proc.GlobalAliases.UpdateMap(exes)

	return exes
}

func matchFlags(partial, exe string) (items []string) {
	if len(ExesFlags[exe].Flags) == 0 && !ExesFlags[exe].NoFlags {
		f := ExesFlags[exe]
		f.Flags = man.ScanManPages(exe)
		ExesFlags[exe] = f
	}

	for i := range ExesFlags[exe].Flags {
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
		ansi.Stderrln(ansi.FgRed, "Dynamic autocompleter is not a code block!")
		//os.Stdout.WriteString("Dynamic autocompleter is not a code block!" + utils.NewLineString)
		return
	}
	block := []rune(ExesFlags[exe].Dynamic[1 : len(ExesFlags[exe].Dynamic)-1])

	stdout := streams.NewStdin()
	stderr := streams.NewStdin()
	exitNum, err := lang.ProcessNewBlock(block, nil, stdout, stderr, p)
	stdout.Close()
	stderr.Close()

	b, _ := stderr.ReadAll()
	ansi.Stderrln(ansi.FgRed, string(b))
	//os.Stderr.Write(b)

	if err != nil {
		ansi.Stderrln(ansi.FgRed, "Error in dynamic autocomplete code: "+err.Error())
		//os.Stdout.WriteString(err.Error() + utils.NewLineString)
		//return
	}
	if exitNum != 0 {
		ansi.Stderrln(ansi.FgRed, "None zero exit number!")
		//os.Stdout.WriteString("None zero exit number!" + utils.NewLineString)
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
