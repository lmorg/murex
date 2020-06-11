package typemgmt

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils"
)

func init() {
	lang.GoFunctions["set"] = cmdSet
	lang.GoFunctions["!set"] = cmdUnset
	lang.GoFunctions["global"] = cmdGlobal
	lang.GoFunctions["!global"] = cmdUnglobal
	lang.GoFunctions["export"] = cmdExport
	lang.GoFunctions["!export"] = cmdUnexport
	lang.GoFunctions["unset"] = cmdUnexport
}

func cmdSet(p *lang.Process) error      { return set(p, p) }
func cmdUnset(p *lang.Process) error    { return unset(p, p) }
func cmdGlobal(p *lang.Process) error   { return set(p, lang.ShellProcess) }
func cmdUnglobal(p *lang.Process) error { return unset(p, lang.ShellProcess) }

func set(p *lang.Process, scope *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name; expected [data-type] name[=value]")
	}

	name, value, dataType, err := splitVarString(p.Parameters.Params)
	if err != nil {
		return err
	}

	// Set variable as method:
	if p.IsMethod {
		if value != "" {
			return errors.New("Unexpected parameters for calling `set` / `global` as method; value was set in parameters")
		}

		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		b = utils.CrLfTrim(b)

		if dataType == "" {
			dataType = p.Stdin.GetDataType()
		}
		return scope.Parent.Variables.Set(name, string(b), dataType)
	}

	// Set variable as parameters:
	if dataType == "" {
		dataType = types.String
	}

	return scope.Parent.Variables.Set(name, value, dataType)
}

func unset(p *lang.Process, scope *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name")
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = scope.Parent.Variables.Unset(varName)
	return err
}

var (
	rxSet     = regexp.MustCompile(`(?sm)^([_a-zA-Z0-9]+)=(.*$)`)
	rxVarName = regexp.MustCompile(`^([_a-zA-Z0-9]+)$`)
)

func cmdExport(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name")
	}

	params := p.Parameters.StringAll()

	// Set env as method:
	if p.IsMethod {
		if !rxVarName.MatchString(params) {
			return errors.New("Invalid variable name; unexpected parameters for calling `export` as method")
		}
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		b = utils.CrLfTrim(b)
		return os.Setenv(params, string(b))
	}

	// Set env as parameters:
	if rxVarName.MatchString(params) {
		return os.Setenv(params, "")
	}

	match := rxSet.FindAllStringSubmatch(params, -1)
	if len(match) == 0 || len(match[0]) < 3 {
		return errors.New("Error parsing export parameters. Expected: name[=value]")
	}
	err := os.Setenv(match[0][1], match[0][2])
	if err != nil {
		return err
	}

	if match[0][1] == "PATH" {
		autocomplete.UpdateGlobalExeList()
	}

	return nil
}

func cmdUnexport(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name")
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
		for _, r := range []rune(params[j]) {
			switch {
			case (r >= 'a' && 'z' >= r) || (r >= 'A' && 'Z' >= r) || (r >= '0' && '9' >= r) || r == '_':
				switch parserState {
				case parserStateExpValue:
					err = fmt.Errorf("Invalid space or tab in variable name")
					return
				default:
					runes[i] = r
					i++
				}

			case r == '=':
				switch parserState {
				case parserStateDataType:
					err = fmt.Errorf("Invalid character '=' in data-type name")
					return
				case parserStateName, parserStateExpValue:
					if dataType != "" && i > 0 {
						err = fmt.Errorf("Invalid space or tab in variable name / too many parameters")
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
					err = fmt.Errorf("Invalid space or tab in data type name")
					return
				case parserStateName:
					if i == 0 {
						err = fmt.Errorf("Invalid space or tab in variable name")
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
						err = fmt.Errorf("Invalid character '%s' in variable name", string([]rune{r}))
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
				err = fmt.Errorf("Invalid parameters; expecting: [data-type] name[=value]")
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
				err = fmt.Errorf("Invalid space or tab in variable name / too many parameters")
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
		err = fmt.Errorf("Invalid variable name. Names can only include alpha, numeric and underscore characters")
	}

	return
}
