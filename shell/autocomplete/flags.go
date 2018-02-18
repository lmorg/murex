package autocomplete

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/man"
)

// Flags is a struct to store auto-complete options
type Flags struct {
	IncFiles      bool               // `true` to disable file name completion
	IncDirs       bool               // `true` to disable directory navigation completion
	IncExePath    bool               // `true` to include binaries in $PATH
	Flags         []string           // known supported command line flags for executable
	Dynamic       string             // Use murex script to generate auto-complete options
	FlagValues    map[string][]Flags // Auto-complete possible values for known flags
	Optional      bool               // This nest of flags is optional
	AllowMultiple bool               // Allow multiple flags in this nest
	Alias         string             // Alias one []Flags to another
	NestedCommand bool               // Jump to another command's flag processing (derived from the previous parameter). eg `sudo command parameters...`
	AnyValue      bool               // Allow any value to be input (eg user input that cannot be pre-determined)
	AutoBranch    bool               // Autocomplete trees (eg directory structures) one branch at a time
	//NoFlags       bool               // `true` to disable Flags[] slice and man page parsing
}

// ExesFlags is map of executables and their supported auto-complete options.
var ExesFlags map[string][]Flags = make(map[string][]Flags)

// globalExes is a pre-populated list of all executables in $PATH.
// The point of this is to speed up exe auto-completion.
var globalExes map[string]bool = make(map[string]bool)

// UpdateGlobalExeList generates a list of executables in $PATH. This used to be called upon demand but it caused a
// slight but highly annoying pause if murex had been sat idle for a while. So now it's an exported function so it can
// be run as a background job or upon user request.
func UpdateGlobalExeList() {
	//envPath := proc.GlobalVars.GetString("PATH")
	//if envPath == "" {
	//	envPath = os.Getenv("PATH")
	//}
	envPath := proc.ShellProcess.VarGetString("PATH")

	dirs := splitPath(envPath)

	for i := range dirs {
		listExes(dirs[i], globalExes)
	}
}

// InitExeFlags initialises empty []Flags based on sane defaults and a quick scan of the man pages (OS dependant)
func InitExeFlags(exe string) {
	if len(ExesFlags[exe]) == 0 {
		ExesFlags[exe] = []Flags{{
			Flags:         man.ScanManPages(exe),
			IncFiles:      true,
			AllowMultiple: true,
			AnyValue:      true,
		}}
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

func match(f *Flags, partial string, args dynamicArgs) (items []string) {
	items = append(items, matchPartialFlags(f, partial)...)
	items = append(items, matchDynamic(f, partial, args)...)

	if f.IncExePath {
		pathexes := allExecutables(false)
		items = append(items, matchExes(partial, pathexes, false)...)
	}

	switch {
	case f.IncFiles:
		items = append(items, matchFilesAndDirs(partial)...)
	case f.IncDirs && !f.IncFiles:
		items = append(items, matchDirs(partial)...)
	}

	return
}

func matchFlags(flags []Flags, partial, exe string, params []string, pIndex *int, args dynamicArgs) (items []string) {
	var nest int

	defer func() {
		if debug.Enable {
			return
		}
		if r := recover(); r != nil {
			ansi.Stderrln(ansi.FgRed, fmt.Sprint("\nPanic caught:", r))
			ansi.Stderrln(ansi.FgRed, fmt.Sprint("Debug information (partial, exe, params, pIndex, nest): ", partial, exe, params, *pIndex, nest))
			b, _ := json.MarshalIndent(flags, "", "\t")
			ansi.Stderrln(ansi.FgRed, string(b))
		}
	}()

	if len(flags) > 0 {
		for ; *pIndex <= len(params); *pIndex++ {
		next:
			if *pIndex >= len(params) {
				break
			}

			if *pIndex > 0 && nest > 0 && flags[nest-1].NestedCommand {
				debug.Log("params:", params[*pIndex-1])
				InitExeFlags(params[*pIndex-1])
				if len(flags[nest-1].FlagValues) == 0 {
					flags[nest-1].FlagValues = make(map[string][]Flags)
				}

				// Only nest command if the command isn't present in Flags.Flags[]. Otherwise we then assume that flag
				// has already been defined by `autocomplete`.
				var doNotNest bool
				for i := range flags[nest-1].Flags {
					if flags[nest-1].Flags[i] == params[*pIndex-1] {
						doNotNest = true
					}
				}

				if !doNotNest {
					args.exe = params[*pIndex-1]
					args.params = params[*pIndex:]
					args.float = *pIndex
					flags[nest-1].FlagValues[args.exe] = ExesFlags[args.exe]
				}

			}

			if *pIndex > 0 && nest > 0 && len(flags[nest-1].FlagValues[params[*pIndex-1]]) > 0 {
				alias := flags[nest-1].FlagValues[params[*pIndex-1]][0].Alias
				if alias != "" {
					flags[nest-1].FlagValues[params[*pIndex-1]] = flags[nest-1].FlagValues[alias]
				}

				items = matchFlags(flags[nest-1].FlagValues[params[*pIndex-1]], partial, exe, params, pIndex, args)
				if len(items) > 0 {
					return
				}
			}

			if nest >= len(flags) {
				return
			}

			if flags[nest].AnyValue || len(match(&flags[nest], params[*pIndex], dynamicArgs{exe: args.exe, params: params[args.float:*pIndex]})) > 0 {
				if !flags[nest].AllowMultiple {
					nest++
				}
				continue
			}

			nest++
			goto next
		}
	}

	if nest > 0 {
		nest--
	}
	for ; nest <= len(flags); nest++ {
		debug.Log("nest", nest, "partial", partial, "exe", exe, "params", params)
		debug.Json("&flags", &flags)
		items = append(items, match(&flags[nest], partial, args)...)
		if !flags[nest].Optional {
			break
		}
	}

	return
}

func matchPartialFlags(f *Flags, partial string) (items []string) {
	for i := range f.Flags {
		flag := f.Flags[i]
		if flag == "" {
			continue
		}
		if strings.HasPrefix(flag, partial) {
			//if flag[len(flag)-1] != '=' {
			//	flag += " "
			//}
			items = append(items, flag[len(partial):])
		}
	}
	sort.Strings(items)
	return
}
