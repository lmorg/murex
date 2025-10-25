package structs

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/hintsummary"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("alias", cmdAlias, types.Null)
	lang.DefineFunction("!alias", cmdUnalias, types.Null)
}

const (
	usageAlias = `
usage: alias [ --copy ] alias_name = command parameter1 parameter2 ...
please note: command and parameters should not be quoted as one parameter to alias`
)

const (
	fAliasCopy = "--copy"
)

var argsAlias = &parameters.Arguments{
	Flags: map[string]string{
		fAliasCopy: types.Boolean,
		"-c":       fAliasCopy,
	},
	AllowAdditional:     true,
	IgnoreInvalidFlags:  false,
	StrictFlagPlacement: true,
}

func aliasErr(err error, name string) error {
	if err == nil {
		return nil
	}

	if name == "" {
		return fmt.Errorf("%v%s", err, usageAlias)
	}

	return fmt.Errorf("%v\nalias name: %s%s", err, name, usageAlias)
}

func cmdAlias(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)
		b, err := json.Marshal(lang.GlobalAliases.Dump(), p.Stdout.IsTTY())
		if err != nil {
			return err
		}
		_, err = p.Stdout.Writeln(b)
		return err

	}

	p.Stdout.SetDataType(types.Null)

	flags, aliasParams, err := p.Parameters.ParseFlags(argsAlias)
	if err != nil {
		return aliasErr(err, "")
	}

	name, params, err := aliasParseParams(aliasParams)
	if err != nil {
		return aliasErr(err, name)
	}

	if flags[fAliasCopy] == types.TrueString {
		cmdAliasCopy(p, name, params)
	}

	lang.GlobalAliases.Add(name, params, p.FileRef)
	return nil
}

var (
	rxAlias = regexp.MustCompile(`^([-_.a-zA-Z0-9]+)=(.*?)$`)

	errAliasMissingCommand = errors.New("missing command to alias")
	errInvalidSyntax       = errors.New("invalid syntax")
	errAliasUnknown        = errors.New("unknown error parsing alias")
)

func aliasParseParams(aliasParams []string) (string, []string, error) {
	if !rxAlias.MatchString(aliasParams[0]) && len(aliasParams) >= 2 && len(aliasParams[1]) >= 1 && aliasParams[1][0] != '=' {
		return "", nil, errInvalidSyntax
	}

	var (
		split  = rxAlias.FindStringSubmatch(aliasParams[0])
		name   string
		params []string
	)

	if len(split) == 0 {
		name = aliasParams[0]
		params = aliasParams[1:]
		switch {
		case len(params) == 0:
			return name, nil, errAliasMissingCommand
		case len(params[0]) == 1 && params[0] == "=":
			params = params[1:]
		case len(params[0]) > 0 && params[0][0] == '=':
			params[0] = params[0][1:]
		default:
			return name, nil, errAliasUnknown
		}

	} else {
		name = split[1]
		params = append([]string{split[2]}, aliasParams[1:]...)
	}

	if len(params) == 0 {
		return name, nil, errAliasMissingCommand
	}

	if params[0] == "" && len(params) > 0 {
		params = params[1:]
	}

	if len(params) == 0 || params[0] == "" {
		return name, nil, errAliasMissingCommand
	}

	return name, params, nil
}

func cmdAliasCopy(p *lang.Process, name string, params []string) {
	// summary
	summary := hintsummary.Summary.Get(params[0])
	if summary != "" {
		hintsummary.Summary.Set(name, summary)
	}

	// method
	dts := lang.MethodStdin.Types(params[0])
	for i := range dts {
		lang.MethodStdin.Define(name, dts[i])
	}

	// autocomplete
	if len(params) == 1 {
		flags, ok := autocomplete.ExesFlags[params[0]]
		if ok {
			autocomplete.ExesFlags[name] = flags
			autocomplete.ExesFlagsFileRef[name] = p.FileRef
		}
	}
}

func cmdUnalias(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	for _, name := range p.Parameters.StringArray() {
		err := lang.GlobalAliases.Delete(name)
		if err != nil {
			return err
		}
	}
	return nil
}
