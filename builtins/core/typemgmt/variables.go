package typemgmt

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func init() {
	lang.DefineMethod("set", cmdSet, types.Any, types.Null)
	lang.DefineFunction("!set", cmdUnset, types.Null)

	lang.DefineMethod("global", cmdGlobal, types.Any, types.Null)
	lang.DefineFunction("!global", cmdUnglobal, types.Null)

	lang.DefineMethod("export", cmdExport, types.Any, types.Null)
	lang.DefineFunction("!export", cmdUnexport, types.Null)
	lang.DefineFunction("unset", cmdUnexport, types.Null)
}

func cmdSet(p *lang.Process) error      { return set(p, p.Variables) }
func cmdUnset(p *lang.Process) error    { return unset(p, p.Variables) }
func cmdGlobal(p *lang.Process) error   { return set(p, lang.GlobalVariables) }
func cmdUnglobal(p *lang.Process) error { return unset(p, lang.GlobalVariables) }

const errMissingName = "missing variable name"

func set(p *lang.Process, v *lang.Variables) error {
	if p.Parameters.Len() == 0 {
		return fmt.Errorf("%s, expected: set|global [data-type] name[=value]", errMissingName)
	}

	name, value, dataType, err := splitVarString(p.Parameters.StringArray())
	if err != nil {
		return err
	}

	// Set variable as method:
	if p.IsMethod {
		if value != "" {
			return errors.New("unexpected parameters for calling `set` / `global` as method; value was set in parameters")
		}

		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		b = utils.CrLfTrim(b)

		if dataType == "" {
			dataType = p.Stdin.GetDataType()
		}

		if dataType == types.String {
			return v.Set(p, name, string(b), dataType)
		}

		iface, err := types.ConvertGoType(string(b), dataType)
		if err != nil {
			return fmt.Errorf("unable to convert parameters into data type: %s", err.Error())
		}
		return v.Set(p, name, iface, dataType)
	}

	// Set variable as parameters:
	if dataType == "" {
		dataType = types.String
	}

	if dataType == types.String {
		return v.Set(p, name, value, dataType)
	}

	iface, err := types.ConvertGoType(value, dataType)
	if err != nil {
		return fmt.Errorf("unable to convert parameters into data type: %s", err.Error())
	}
	return v.Set(p, name, iface, dataType)
}

func unset(p *lang.Process, v *lang.Variables) error {
	if p.Parameters.Len() == 0 {
		return errors.New(errMissingName)
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return v.Unset(varName)
}

var rxVarName = regexp.MustCompile(`^([_a-zA-Z0-9]+)$`)

func cmdExport(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		return errors.New(errMissingName)
	}

	params := p.Parameters.StringAll()

	// Set env as method:
	if p.IsMethod {
		if !rxVarName.MatchString(params) {
			return errors.New("invalid variable name; unexpected parameters for calling `export` as method")
		}
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		b = utils.CrLfTrim(b)
		return os.Setenv(params, string(b))
	}

	name, value, err := parseExportParameters(params)
	if err != nil {
		return fmt.Errorf("error parsing export parameters: %s\nexpected: name[=value]", err.Error())
	}

	if value == nil {
		*value, _ = p.Variables.GetString(name)

	}

	err = os.Setenv(name, *value)
	if err != nil {
		return err
	}

	//if name == "PATH" {
	//	autocomplete.UpdateGlobalExeList()
	//}

	return nil
}

func parseExportParameters(s string) (string, *string, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "", nil, errors.New(errMissingName)
	}

	split := strings.SplitN(s, "=", 2)
	switch len(split) {
	case 0:
		return "", nil, errors.New(errMissingName)

	case 1:
		return strings.TrimSpace(split[0]), nil, nil

	case 2:
		s := strings.TrimSpace(split[1])
		return strings.TrimSpace(split[0]), &s, nil

	default:
		s := strings.TrimSpace(strings.Join(split[1:], "="))
		return strings.TrimSpace(split[0]), &s, nil
	}
}

func cmdUnexport(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		return errors.New(errMissingName)
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = os.Unsetenv(varName)
	return err
}

const (
	parserStateName = iota
	parserStateExpValue
	parserStateValue
	parserStateDataType
)

func splitVarString(params []string) (name, value, dataType string, err error) {
	var (
		parserState int
		max         int
	)

	for i := range params {
		max += len(params[i]) + 1
	}

	runes := make([]rune, max)
	i := 0
	for j := range params {
		for _, r := range params[j] {
			switch {
			case (r >= 'a' && 'z' >= r) || (r >= 'A' && 'Z' >= r) || (r >= '0' && '9' >= r) || r == '_' || r == '.':
				switch parserState {
				case parserStateExpValue:
					err = fmt.Errorf("invalid space or tab in variable name")
					return
				default:
					runes[i] = r
					i++
				}

			case r == '=':
				switch parserState {
				case parserStateDataType:
					err = fmt.Errorf("invalid character '=' in data-type name")
					return
				case parserStateName, parserStateExpValue:
					if dataType != "" && i > 0 {
						err = fmt.Errorf("invalid space or tab in variable name / too many parameters")
						return
					}
					if name != "" {
						if i == 0 {
							parserState = parserStateValue
							continue
						}
						dataType = name
					}
					name = string(runes[:i])
					i = 0
					runes = make([]rune, max)
					parserState = parserStateValue
				case parserStateValue:
					runes[i] = r
					i++
				}

			case r == ' ' || r == '\t':
				switch parserState {
				case parserStateDataType:
					err = fmt.Errorf("invalid space or tab in data type name")
					return
				case parserStateName:
					if i == 0 {
						err = fmt.Errorf("invalid space or tab in variable name")
						return
					}
					parserState = parserStateExpValue
					continue
				case parserStateValue:
					runes[i] = r
					i++
				case parserStateExpValue:
					// do nothing
				}

			default:
				switch parserState {
				case parserStateName, parserStateExpValue:
					if len(params) > 1 && dataType != "" {
						err = fmt.Errorf("invalid character '%s' in variable name", string([]rune{r}))
						return
					}
					parserState = parserStateDataType
					fallthrough
				case parserStateDataType:
					runes[i] = r
					i++
				case parserStateValue:
					runes[i] = r
					i++
				}

			}
		}

		switch parserState {
		case parserStateDataType:
			if len(params) == 0 {
				err = fmt.Errorf("invalid parameters; expecting: [data-type] name[=value]")
				return
			}
			dataType = string(runes[:i])
			i = 0
			runes = make([]rune, max)
			parserState = parserStateName

		case parserStateName:
			switch {
			case name == "":
				name = string(runes[:i])
			case dataType == "":
				dataType = name
				name = string(runes[:i])
			default:
				err = fmt.Errorf("invalid space or tab in variable name / too many parameters")
				return
			}
			i = 0
			runes = make([]rune, max)

		case parserStateValue:
			if value == "" {
				value = string(runes[:i])
			} else {
				value += " " + string(runes[:i])
			}
			i = 0
			runes = make([]rune, max)

		}
	}

	if name == "" {
		err = fmt.Errorf("invalid variable name. Names can only include alpha, numeric and underscore characters")
	}

	return
}
