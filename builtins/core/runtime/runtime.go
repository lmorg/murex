package cmdruntime

import (
	"errors"
	"runtime"
	"sort"

	"github.com/lmorg/murex/builtins/core/open"
	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/hintsummary"
	"github.com/lmorg/murex/utils/cache"
	cdcache "github.com/lmorg/murex/utils/cd/cache"
	"github.com/lmorg/murex/utils/envvars"
	"github.com/lmorg/murex/utils/json"
)

const (
	fVars               = "--variables"
	fGlobals            = "--globals"
	fExports            = "--exports"
	fAliases            = "--aliases"
	fBuiltins           = "--builtins"
	fMethods            = "--methods"
	fConfig             = "--config"
	fNamedPipes         = "--named-pipes"
	fPipes              = "--pipes"
	fFunctions          = "--functions"
	fPrivates           = "--privates"
	fOpenAgents         = "--open-agents"
	fFids               = "--fids"
	fShellProc          = "--shell-proc"
	fReadArrays         = "--readarray"
	fReadArrayWithTypes = "--readarraywithtype"
	fWriteArrays        = "--writearray"
	fReadMaps           = "--readmap"
	fIndexes            = "--indexes"
	fNotIndexes         = "--not-indexes"
	fMarshallers        = "--marshallers"
	fUnmarshallers      = "--unmarshallers"
	fEvents             = "--events"
	fAutocomplete       = "--autocomplete"
	fMemstats           = "--memstats"
	fTests              = "--tests"
	fTestResults        = "--test-results"
	fModules            = "--modules"
	fDebug              = "--debug"
	fSources            = "--sources"
	fSummaries          = "--summaries"
	fCachedFilePaths    = "--cached-file-paths"
	fCacheDump          = "--cache"
	fCacheTrim          = "--trim-cache"
	fHelp               = "--help"
)

var flags = map[string]string{
	fVars:               types.Boolean,
	fGlobals:            types.Boolean,
	fExports:            types.Boolean,
	fAliases:            types.Boolean,
	fBuiltins:           types.Boolean,
	fMethods:            types.Boolean,
	fConfig:             types.Boolean,
	fPipes:              types.Boolean,
	fNamedPipes:         types.Boolean,
	fFunctions:          types.Boolean,
	fPrivates:           types.Boolean,
	fOpenAgents:         types.Boolean,
	fFids:               types.Boolean,
	fShellProc:          types.Boolean,
	fReadArrays:         types.Boolean,
	fReadArrayWithTypes: types.Boolean,
	fReadMaps:           types.Boolean,
	fWriteArrays:        types.Boolean,
	fIndexes:            types.Boolean,
	fNotIndexes:         types.Boolean,
	fMarshallers:        types.Boolean,
	fUnmarshallers:      types.Boolean,
	fEvents:             types.Boolean,
	fAutocomplete:       types.Boolean,
	fMemstats:           types.Boolean,
	fTests:              types.Boolean,
	fTestResults:        types.Boolean,
	fModules:            types.Boolean,
	fDebug:              types.Boolean,
	fSources:            types.Boolean,
	fSummaries:          types.Boolean,
	fCachedFilePaths:    types.Boolean,
	fCacheDump:          types.Boolean,
	fCacheTrim:          types.Boolean,
	fHelp:               types.Boolean,
}

func init() {
	lang.DefineFunction("runtime", cmdRuntime, types.Json)

	defaults.AppendProfile(`
        autocomplete set runtime { [{
            "Dynamic": ({ runtime --help }),
            "AllowMultiple": true
        }] }
    `)
}

// Help returns an array of flags supported by `runtime`
func Help() (s []string) {
	for f := range flags {
		s = append(s, f)
	}

	sort.Strings(s)
	return
}

func cmdRuntime(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	f, _, err := p.Parameters.ParseFlags(
		&parameters.Arguments{
			Flags:           flags,
			AllowAdditional: false,
		},
	)

	if err != nil {
		return err
	}

	if len(f) == 0 {
		return errors.New("please include one or more parameters")
	}

	ret := make(map[string]interface{})
	for flag := range f {
		switch flag {
		case fVars:
			ret[fVars[2:]] = p.Scope.Variables.Dump()
		case fGlobals:
			ret[fGlobals[2:]] = lang.GlobalVariables.Dump()
		case fExports:
			m := make(map[string]interface{})
			envvars.All(m)
			ret[fExports[2:]] = m
		case fAliases:
			ret[fAliases[2:]] = lang.GlobalAliases.Dump()
		case fBuiltins:
			var s []string
			for name := range lang.GoFunctions {
				s = append(s, name)
			}
			sort.Strings(s)
			ret[fBuiltins[2:]] = s
		case fMethods:
			ret[fMethods[2:]] = map[string]map[string][]string{
				"in":  lang.MethodStdin.Dump(),
				"out": lang.MethodStdout.Dump(),
			}
		case fConfig:
			ret[fConfig[2:]] = lang.ShellProcess.Config.DumpRuntime()
		case fNamedPipes:
			ret[fNamedPipes[2:]] = lang.GlobalPipes.Dump()
		case fPipes:
			ret[fPipes[2:]] = stdio.DumpPipes()
		case fFunctions:
			ret[fFunctions[2:]] = lang.MxFunctions.Dump()
		case fPrivates:
			ret[fPrivates[2:]] = lang.PrivateFunctions.Dump()
		case fOpenAgents:
			ret[fOpenAgents[2:]] = open.OpenAgents.Dump()
		case fFids:
			ret[fFids[2:]] = lang.GlobalFIDs.ListAll()
		case fShellProc:
			ret[fShellProc[2:]] = lang.ShellProcess.Dump()
		case fReadArrays:
			ret[fReadArrays[2:]] = stdio.DumpReadArray()
		case fReadArrayWithTypes:
			ret[fReadArrays[2:]] = stdio.DumpReadArrayWithType()
		case fReadMaps:
			ret[fReadMaps[2:]] = stdio.DumpMap()
		case fWriteArrays:
			ret[fWriteArrays[2:]] = stdio.DumpWriteArray()
		case fIndexes:
			ret[fIndexes[2:]] = lang.DumpIndex()
		case fNotIndexes:
			ret[fIndexes[2:]] = lang.DumpNotIndex()
		case fMarshallers:
			ret[fMarshallers[2:]] = lang.DumpMarshaller()
		case fUnmarshallers:
			ret[fUnmarshallers[2:]] = lang.DumpUnmarshaller()
		case fEvents:
			ret[fEvents[2:]] = events.DumpEvents()
		case fAutocomplete:
			ret[fAutocomplete[2:]] = autocomplete.RuntimeDump()
		case fMemstats:
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			ret[fMemstats[2:]] = mem
		case fTests:
			ret[fTests[2:]] = p.Tests.Dump()
		case fTestResults:
			ret[fTestResults[2:]] = dumpTestResults(p)
		case fModules:
			ret[fModules[2:]] = profile.Packages
		case fDebug:
			ret[fDebug[2:]] = debug.Dump()
		case fSources:
			ret[fSources[2:]] = ref.History.Dump()
		case fSummaries:
			ret[fSummaries[2:]] = hintsummary.Summary.Dump()
		case fCachedFilePaths:
			ret[fCachedFilePaths[2:]] = cdcache.DumpCompletions()
		case fCacheDump:
			v, err := cache.Dump(p.Context)
			if err != nil {
				return err
			}
			ret[fCacheDump[2:]] = v
		case fCacheTrim:
			v, err := cache.Trim(p.Context)
			if err != nil {
				return err
			}
			ret[fCacheTrim[2:]] = v
		case fHelp:
			ret[fHelp[2:]] = Help()
		default:
			return errors.New("unrecognized parameter: " + flag)
		}
	}

	var b []byte
	if len(ret) == 1 {
		var obj interface{}
		for _, obj = range ret {
		}
		b, err = json.Marshal(obj, p.Stdout.IsTTY())
		if err != nil {
			return err
		}

	} else {
		b, err = json.Marshal(ret, p.Stdout.IsTTY())
		if err != nil {
			return err
		}
	}

	_, err = p.Stdout.Write(b)
	return err
}

func dumpTestResults(p *lang.Process) interface{} {
	return map[string]interface{}{
		"shell":   lang.ShellProcess.Tests.Results.Dump(),
		"process": p.Tests.Results.Dump(),
	}
}
