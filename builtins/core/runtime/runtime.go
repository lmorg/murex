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
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/json"
)

const (
	fVars          = "--variables"
	fAliases       = "--aliases"
	fConfig        = "--config"
	fNamedPipes    = "--named-pipes"
	fPipes         = "--pipes"
	fFunctions     = "--functions"
	fPrivates      = "--privates"
	fOpenAgents    = "=--open-agents"
	fFids          = "--fids"
	fReadArrays    = "--readarray"
	fReadMaps      = "--readmap"
	fWriteArrays   = "--writearray"
	fIndexes       = "--indexes"
	fMarshallers   = "--marshallers"
	fUnmarshallers = "--unmarshallers"
	fEvents        = "--events"
	fAutocomplete  = "--autocomplete"
	fMemstats      = "--memstats"
	fAstCache      = "--astcache"
	fTests         = "--tests"
	fTestResults   = "--test-results"
	fModules       = "--modules"
	fDebug         = "--debug"
	fSources       = "--sources"
	fHelp          = "--help"

	// inspect
	inspVariables = "--inspect-variables"
)

var flags = map[string]string{
	fVars:          types.Boolean,
	fAliases:       types.Boolean,
	fConfig:        types.Boolean,
	fPipes:         types.Boolean,
	fNamedPipes:    types.Boolean,
	fFunctions:     types.Boolean,
	fPrivates:      types.Boolean,
	fOpenAgents:    types.Boolean,
	fFids:          types.Boolean,
	fReadArrays:    types.Boolean,
	fReadMaps:      types.Boolean,
	fWriteArrays:   types.Boolean,
	fIndexes:       types.Boolean,
	fMarshallers:   types.Boolean,
	fUnmarshallers: types.Boolean,
	fEvents:        types.Boolean,
	fAutocomplete:  types.Boolean,
	fMemstats:      types.Boolean,
	fAstCache:      types.Boolean,
	fTests:         types.Boolean,
	fTestResults:   types.Boolean,
	fModules:       types.Boolean,
	fDebug:         types.Boolean,
	fSources:       types.Boolean,
	fHelp:          types.Boolean,

	// inspect flags are defined in cmdRuntime() below
}

func init() {
	lang.GoFunctions["runtime"] = cmdRuntime

	defaults.AppendProfile(`
        autocomplete set runtime { [{
            "Dynamic": ({ runtime --help }),
            "AllowMultiple": true
        }] }
    `)
}

func help() (s []string) {
	for f := range flags {
		s = append(s, f)
	}

	sort.Strings(s)
	return
}

func cmdRuntime(p *lang.Process) error {
	// inspect
	if debug.Inspect {
		flags[inspVariables] = types.Boolean
	}

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
		return errors.New("Please include one or more parameters")
	}

	ret := make(map[string]interface{})
	for flag := range f {
		switch flag {
		case fVars:
			ret[fVars[2:]] = p.Variables.Dump()
		case fAliases:
			ret[fAliases[2:]] = lang.GlobalAliases.Dump()
		case fConfig:
			ret[fConfig[2:]] = p.Config.DumpRuntime()
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
		case fReadArrays:
			ret[fReadArrays[2:]] = stdio.DumpArray()
		case fReadMaps:
			ret[fReadMaps[2:]] = stdio.DumpMap()
		case fWriteArrays:
			ret[fWriteArrays[2:]] = stdio.DumpArray()
		case fIndexes:
			ret[fIndexes[2:]] = define.DumpIndex()
		case fMarshallers:
			ret[fMarshallers[2:]] = define.DumpMarshaller()
		case fUnmarshallers:
			ret[fUnmarshallers[2:]] = define.DumpUnmarshaller()
		case fEvents:
			ret[fEvents[2:]] = events.DumpEvents()
		case fAutocomplete:
			ret[fAutocomplete[2:]] = autocomplete.ExesFlags
		case fMemstats:
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			ret[fMemstats[2:]] = mem
		case fAstCache:
			ret[fAstCache[2:]] = lang.AstCache.Dump()
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
		case fHelp:
			ret[fHelp[2:]] = help()
		default:
			if !debug.Inspect {
				return errors.New("Unrecognised parameter: " + flag)
			}

			// inspect
			switch flag {
			case inspVariables:
				ret[inspVariables[2:]] = p.Variables.Inspect()
			default:
				return errors.New("Unrecognised parameter: " + flag)
			}

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
