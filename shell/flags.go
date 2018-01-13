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
	FlagValues map[string]Flags // Auto-complete possible values for known flags
	Optional bool
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

func matchFlags( partial, exe string, params []string) (items []string) {
	/*if len(ExesFlags[exe])==0 {
		items = matchPartialFlags(partial, exe)
		items = append(items, matchDynamic(partial, exe, params)...)

		if ExesFlags[exe][i].IncExePath {
			pathexes := allExecutables(false)
			items = append(items, matchExes(partial, pathexes, false)...)
		}

		switch {
		case !ExesFlags[exe][0].NoFiles:
			items = append(items, matchFilesAndDirs(partial)...)
		case !ExesFlags[exe][0].NoDirs:
			items = append(items, matchDirs(partial)...)
		}

		return
	}*/

	match := func(nest int, partial, exe string, params []string) (items []string){
		//fmt.Println(nest, partial,exe,params, ExesFlags[exe])

		//if nest >= len(ExesFlags[exe]) {
		//	return
		//}



		items = append(items, matchPartialFlags(nest, partial, exe)...)
		items = append(items, matchDynamic(nest, partial, exe, params)...)

		if ExesFlags[exe][nest].IncExePath {
			pathexes := allExecutables(false)
			items = append(items, matchExes(partial, pathexes, false)...)
		}

		switch {
		case !ExesFlags[exe][nest].NoFiles:
			items = append(items, matchFilesAndDirs(partial)...)
		case !ExesFlags[exe][nest].NoDirs:
			items = append(items, matchDirs(partial)...)
		}

		return
	}

	var nest int

	if  len(ExesFlags[exe])==0 {

		ExesFlags[exe]= []Flags{{
				Flags: man.ScanManPages(exe),
			}}

	} else {

		for param := range params {
			//for ; nest<len(ExesFlags[exe] );nest++ { // we deliberately don't want `nest` to reset within each iteration of `range params`
		next:
			if len(match(nest, params[param], exe, params[:param])) > 0 {
				continue
			}

			nest++
			goto next

		}
	}

		//items = match(nest,partial,exe,params)
		for ;nest<=len(ExesFlags[exe] );nest++ {
			items = append(items,match(nest,partial,exe,params)...)
			if !ExesFlags[exe][nest].Optional {
				break
			}
	}

		//ExesFlags[exe][nest].Optional

	/*
	for i:=range ExesFlags[exe] {

	}

		for _,param:= range params {
			var match bool
			for j:= range ExesFlags[exe][i].Flags {
				if param == ExesFlags[exe][i].Flags[j] {

				}
			}
		}
		items = ExesFlags[exe][i].Flags
		//items = matchPartialFlags(partial, exe)
	}*/
	return
}

func matchPartialFlags(nest int, partial, exe string) (items []string) {
	//if len(ExesFlags[exe][nest].Flags) == 0 && !ExesFlags[exe][nest].NoFlags {
	//	f := ExesFlags[exe][nest]
	//	f.Flags = man.ScanManPages(exe)
	//	ExesFlags[exe] = append([]Flags{},f)
	//}

	for i := range ExesFlags[exe][nest].Flags {
		flag := ExesFlags[exe][nest].Flags[i]
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

func matchDynamic(nest int,partial, exe string, params []string) (items []string) {
	if len(ExesFlags[exe][nest].Dynamic) == 0 {
		return
	}

	p := &proc.Process{
		Name:       exe,
		Parameters: parameters.Parameters{Params: params},
		Parent:     proc.ShellProcess,
	}
	p.Scope = p

	if !types.IsBlock([]byte(ExesFlags[exe][nest].Dynamic)) {
		ansi.Stderrln(ansi.FgRed, "Dynamic autocompleter is not a code block!")
		return
	}
	block := []rune(ExesFlags[exe][nest].Dynamic[1 : len(ExesFlags[exe][nest].Dynamic)-1])

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
