package autocomplete

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/murex/utils/readline"
)

// Flags is a struct to store auto-complete options
type Flags struct {
	IncFiles      bool               // `true` to include file name completion
	FileRegexp    string             // Regexp match for files if IncFiles set
	IncDirs       bool               // `true` to include directory navigation completion
	IncExePath    bool               // `true` to include binaries in $PATH
	IncExeAll     bool               // `true` to include all executable names
	IncManPage    bool               // `true` to include man page lookup
	Flags         []string           // known supported command line flags for executable
	FlagsDesc     map[string]string  // known supported command line flags for executable with descriptions
	Dynamic       string             // Use murex script to generate auto-complete suggestions
	DynamicDesc   string             // Use murex script to generate auto-complete suggestions with descriptions
	ListView      bool               // Display the helps as a "popup menu-like" list rather than grid
	MapView       bool               // Like ListView but the description is highlighted instead
	FlagValues    map[string][]Flags // Auto-complete possible values for known flags
	Optional      bool               // This nest of flags is optional
	AllowMultiple bool               // Allow multiple flags in this nest
	Alias         string             // Alias one []Flags to another
	NestedCommand bool               // Jump to another command's flag processing (derived from the previous parameter). eg `sudo command parameters...`
	AnyValue      bool               // Allow any value to be input (eg user input that cannot be pre-determined)
	AutoBranch    bool               // Autocomplete trees (eg directory structures) one branch at a time
	ExecCmdline   bool               // Execute the commandline and pass it to STDIN when Dynamic/DynamicDesc used (potentially dangerous)
	//NoFlags       bool             // `true` to disable Flags[] slice and man page parsing
}

var (
	// ExesFlags is map of executables and their supported auto-complete options.
	ExesFlags = make(map[string][]Flags)

	// ExesFlagsFileRef is a map of which module defined ExesFlags
	ExesFlagsFileRef = make(map[string]*ref.File)

	// GlobalExes is a pre-populated list of all executables in $PATH.
	// The point of this is to speed up exe auto-completion.
	GlobalExes = make(map[string]bool)
)

// UpdateGlobalExeList generates a list of executables in $PATH. This used to be called upon demand but it caused a
// slight but highly annoying pause if murex had been sat idle for a while. So now it's an exported function so it can
// be run as a background job or upon user request.
func UpdateGlobalExeList() {
	envPath, _ := lang.ShellProcess.Variables.GetString("PATH")

	dirs := SplitPath(envPath)

	for i := range dirs {
		listExes(dirs[i], GlobalExes)
	}
}

// InitExeFlags initializes empty []Flags based on sane defaults and a quick scan of the man pages (OS dependant)
func InitExeFlags(exe string) {
	if len(ExesFlags[exe]) == 0 {
		ExesFlags[exe] = []Flags{{
			Flags:         scanManPages(exe),
			IncFiles:      true,
			AllowMultiple: true,
			AnyValue:      true,
		}}
	}
}

type runtimeDump struct {
	FlagValues []Flags
	FileRef    *ref.File
}

// RuntimeDump exports the autocomplete flags and FileRef metadata in a JSON
// compatible struct for `runtime` to consume
func RuntimeDump() interface{} {
	dump := make(map[string]runtimeDump)

	for exe := range ExesFlags {
		dump[exe] = runtimeDump{
			FlagValues: ExesFlags[exe],
			FileRef:    ExesFlagsFileRef[exe],
		}
	}

	return dump
}

func scanManPages(exe string) []string {
	f := man.GetManPages(exe)
	return man.ParseFlags(f)
}

func allExecutables(includeBuiltins bool) map[string]bool {
	exes := make(map[string]bool)
	for k, v := range GlobalExes {
		exes[k] = v
	}

	if !includeBuiltins {
		return exes
	}

	for name := range lang.GoFunctions {
		exes[name] = true
	}

	lang.MxFunctions.UpdateMap(exes)
	lang.GlobalAliases.UpdateMap(exes)

	return exes
}

func match(f *Flags, partial string, args dynamicArgs, act *AutoCompleteT) int {
	matchPartialFlags(f, partial, act)
	matchDynamic(f, partial, args, act)

	if f.IncExeAll {
		pathall := allExecutables(true)
		act.append(matchExes(partial, pathall, false)...)

	} else if f.IncExePath {
		pathexes := allExecutables(false)
		act.append(matchExes(partial, pathexes, false)...)
	}

	if f.IncManPage {
		flags := scanManPages(args.exe)
		act.append(lists.CropPartial(flags, partial)...)
	}

	switch {
	case act.CacheDynamic:
		// do nothing
	case f.IncFiles:
		act.append(matchFilesAndDirsWithRegexp(partial, f.FileRegexp, act)...)
	case f.IncDirs && !f.IncFiles:
		act.append(matchDirs(partial, act)...)
	}

	if /*len(f.FlagsDesc) > 0 &&*/ f.ListView {
		act.TabDisplayType = readline.TabDisplayList
	} else if f.MapView {
		act.TabDisplayType = readline.TabDisplayMap
	}

	return len(act.Items)
}

func matchFlags(flags []Flags, partial, exe string, params []string, pIndex *int, args dynamicArgs, act *AutoCompleteT) int {
	var nest int

	defer func() {
		if debug.Enabled {
			return
		}
		if r := recover(); r != nil {
			lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprint("\nPanic caught:", r)))
			lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf("Debug information:\n- partial: '%s'\n- exe: '%s'\n- params: %s\n- pIndex: %d\n- nest: %d\nAutocompletion syntax:", partial, exe, params, *pIndex, nest)))
			b, _ := json.MarshalIndent(flags, "", "\t")
			lang.ShellProcess.Stderr.Writeln([]byte(string(b)))

		}
	}()

	if len(flags) > 0 {
		for ; *pIndex <= len(params); *pIndex++ {
		next:
			if *pIndex >= len(params) {
				break
			}

			if *pIndex > 0 && nest > 0 && flags[nest-1].NestedCommand {
				//debug.Log("params:", params[*pIndex-1])
				InitExeFlags(params[*pIndex-1])
				if len(flags[nest-1].FlagValues) == 0 {
					flags[nest-1].FlagValues = make(map[string][]Flags)
				}

				// Only nest command if the command isn't present in Flags.Flags[]. Otherwise we then assume that flag
				// has already been defined by `autocomplete`.
				// NOTE TO SELF: I can't remember what this does? And is it required for FlagsDesc?
				var doNotNest bool

				if flags[nest-1].FlagsDesc[params[*pIndex-1]] != "" {
					doNotNest = true
				}
				for i := range flags[nest-1].Flags {
					if flags[nest-1].Flags[i] == params[*pIndex-1] {
						doNotNest = true
						break
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

				length := matchFlags(flags[nest-1].FlagValues[params[*pIndex-1]], partial, exe, params, pIndex, args, act)
				if length > 0 {
					return len(act.Items)
				}
			}

			if nest >= len(flags) {
				return len(act.Items)
			}

			length := match(&flags[nest], params[*pIndex], dynamicArgs{exe: args.exe, params: params[args.float:*pIndex]}, act.disposable())
			if flags[nest].AnyValue || length > 0 {
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
		if nest >= len(flags) {
			/* I don't know why this is needed but it catches a segfault with the following code:

			autocomplete set docgen { [
				{
					"AllowMultiple": true,
					"Optional": true,
					"FlagsDesc": {
						"-panic": "Write a stack trace on error",
						"-readonly": "Don't write output to disk. Use this to test the config",
						"-verbose": "Verbose output (all log messages inc warnings)",
						"-version": "Output docgen version number and exit",
						"-warning": "Display warning messages (will also return a non-zero exit status if warnings found)",
						"-config": "Location of the base docgen config file"
					},
					"FlagValues": {
						"-config": [{
							"IncFiles": true
						}]
					}
				}
			] } */
			break
		}
		match(&flags[nest], partial, args, act)
		if !flags[nest].Optional {
			break
		}
	}

	return len(act.Items)
}

func matchPartialFlags(f *Flags, partial string, act *AutoCompleteT) {
	var flag string

	for i := range f.Flags {
		flag = f.Flags[i]
		if flag == "" {
			continue
		}
		if strings.HasPrefix(flag, partial) {
			act.append(flag[len(partial):])
		}
	}

	for flag := range f.FlagsDesc {
		if !strings.HasPrefix(flag, partial) {
			continue
		}

		act.appendDef(flag[len(partial):], f.FlagsDesc[flag])
	}
}
