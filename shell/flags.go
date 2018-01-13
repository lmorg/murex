package shell

import (
	"bytes"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
	"os"
	"sort"
	"strings"
)

// Flags is a struct to store auto-complete options
type Flags struct {
	NoFiles    bool               // `true` to disable file name completion
	NoDirs     bool               // `true` to disable directory navigation completion
	NoFlags    bool               // `true` to disable Flags[] slice and man page parsing
	IncExePath bool               // `true` to include binaries in $PATH
	Flags      []string           // known supported command line flags for executable
	Dynamic    string             // Use murex script to generate auto-complete options
	FlagValues map[string][]Flags // Auto-complete possible values for known flags
	Optional   bool
}

// ExesFlags is map of executables and their supported auto-complete options.
// We might as well pre-populate the structure with a few base commands we might expect.
var ExesFlags map[string][]Flags = map[string][]Flags{
	"cd":      {{Flags: []string{}, NoFiles: true}},
	"mkdir":   {{Flags: []string{}, NoFiles: true}},
	"rmdir":   {{Flags: []string{}, NoFiles: true}},
	"man":     {{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true}},
	"which":   {{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true}},
	"whereis": {{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true}},
	"sudo":    {{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true}},
	"exec":    {{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true}},
	"pty":     {{Flags: []string{}, NoFiles: true, NoDirs: true, IncExePath: true}},
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

func match(f *Flags, partial, exe string, params []string) (items []string) {
	items = append(items, matchPartialFlags(f, partial, exe)...)
	items = append(items, matchDynamic(f, partial, exe, params)...)

	if f.IncExePath {
		pathexes := allExecutables(false)
		items = append(items, matchExes(partial, pathexes, false)...)
	}

	switch {
	case !f.NoFiles:
		items = append(items, matchFilesAndDirs(partial)...)
	case !f.NoDirs:
		items = append(items, matchDirs(partial)...)
	}

	return
}

func matchFlags(flags []Flags, partial, exe string, params []string, pIndex *int) (items []string) {
	var nest int

	if len(flags) > 0 {
		for ; *pIndex <= len(params); *pIndex++ {
		next:
			if *pIndex >= len(params) {
				break
			}

			//fmt.Println(flags, nest, *pIndex, params)

			if *pIndex > 0 && len(flags[nest].FlagValues[params[*pIndex-1]]) > 0 {
				items = matchFlags(flags[nest].FlagValues[params[*pIndex-1]], partial, exe, params, pIndex)
				if len(items) > 0 {
					return
				}
				continue
			}

			//fmt.Println(*pIndex, params)

			if len(match(&flags[nest],
				params[*pIndex],
				exe,
				params[:*pIndex])) > 0 {
				continue
			}

			nest++
			goto next
		}
	}

	for ; nest <= len(flags); nest++ {
		items = append(items, match(&flags[nest], partial, exe, params)...)
		if !flags[nest].Optional {
			break
		}
	}

	return
}

func matchPartialFlags(f *Flags, partial, exe string) (items []string) {
	//if len(ExesFlags[exe][nest].Flags) == 0 && !ExesFlags[exe][nest].NoFlags {
	//	f := ExesFlags[exe][nest]
	//	f.Flags = man.ScanManPages(exe)
	//	ExesFlags[exe] = append([]Flags{},f)
	//}

	for i := range f.Flags {
		flag := f.Flags[i]
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

func matchDynamic(f *Flags, partial, exe string, params []string) (items []string) {
	//if len(f.Dynamic) == 0 {
	if f.Dynamic == "" {
		return
	}

	p := &proc.Process{
		Name:       exe,
		Parameters: parameters.Parameters{Params: params},
		Parent:     proc.ShellProcess,
	}
	p.Scope = p

	if !types.IsBlock([]byte(f.Dynamic)) {
		ansi.Stderrln(ansi.FgRed, "Dynamic autocompleter is not a code block!")
		return
	}
	block := []rune(f.Dynamic[1 : len(f.Dynamic)-1])

	stdout := streams.NewStdin()
	stderr := streams.NewStdin()
	exitNum, err := lang.ProcessNewBlock(block, nil, stdout, stderr, p)
	stdout.Close()
	stderr.Close()

	b, _ := stderr.ReadAll()
	ansi.Stderrln(ansi.FgRed, string(b))

	if err != nil {
		ansi.Stderrln(ansi.FgRed, "Error in dynamic autocomplete code: "+err.Error())
	}
	if exitNum != 0 {
		ansi.Stderrln(ansi.FgRed, "None zero exit number!")
	}

	stdout.ReadArray(func(b []byte) {
		s := string(bytes.TrimSpace(b))
		if len(s) == 0 {
			return
		}
		if strings.HasPrefix(s, partial) {
			items = append(items, s[len(partial):]+" ")
		}
	})

	return
}
