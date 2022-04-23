package lang

import (
	"fmt"
	"regexp"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/home"
)

var rxTokenIndex = regexp.MustCompile(`(.*?)\[(.*?)\]`)

const errEmptyArray = "Array '@%s' is empty"

// ParseParameters is an internal function to parse parameters
func ParseParameters(prc *Process, p *parameters.Parameters) error {
	var namedPipeIsParam bool
	params := []string{}

	strictArrays, err := prc.Config.Get("proc", "strict-arrays", "bool")
	if err != nil {
		strictArrays = true
	}

	for i := range p.Tokens {
		params = append(params, "")

		var tCount bool
		for j := range p.Tokens[i] {
			switch p.Tokens[i][j].Type {
			case parameters.TokenTypeNil:
				// do nothing

			case parameters.TokenTypeNamedPipe:
				if !namedPipeIsParam {
					continue
				}
				p.Tokens[i][j].Type = parameters.TokenTypeValue
				fallthrough

			case parameters.TokenTypeValue:
				params[len(params)-1] += p.Tokens[i][j].Key
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeString:
				s, err := prc.Variables.GetString(p.Tokens[i][j].Key)
				if err != nil {
					return err
				}
				s = utils.CrLfTrimString(s)
				params[len(params)-1] += s
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeBlockString:
				fork := prc.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				exitNum, err := fork.Execute([]rune(p.Tokens[i][j].Key))
				if err != nil {
					return fmt.Errorf("subshell failed: %s", err.Error())
				}
				if exitNum > 0 &&
					(prc.RunMode == runmode.Try || prc.RunMode == runmode.TryPipe) {
					return fmt.Errorf("subshell exit status %d", exitNum)
				}
				b, err := fork.Stdout.ReadAll()
				if err != nil {
					return err
				}

				b = utils.CrLfTrim(b)

				params[len(params)-1] += string(b)
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeArray:
				data, err := prc.Variables.GetString(p.Tokens[i][j].Key)
				if err != nil {
					return err
				}

				if data == "" {
					if strictArrays.(bool) {
						return fmt.Errorf(errEmptyArray, p.Tokens[i][j].Key)
					} else {
						continue
					}
				}

				var array []string

				variable := streams.NewStdin()
				variable.SetDataType(prc.Variables.GetDataType(p.Tokens[i][j].Key))
				variable.Write([]byte(data))

				variable.ReadArray(func(b []byte) {
					array = append(array, string(b))
				})

				if len(array) == 0 && strictArrays.(bool) {
					return fmt.Errorf(errEmptyArray, p.Tokens[i][j].Key)
				}

				if !tCount {
					params = params[:len(params)-1]
				}

				params = append(params, array...)

				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeBlockArray:
				var array []string

				fork := prc.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				fork.Execute([]rune(p.Tokens[i][j].Key))
				fork.Stdout.ReadArray(func(b []byte) {
					array = append(array, string(b))
				})

				if len(array) == 0 && strictArrays.(bool) {
					return fmt.Errorf(errEmptyArray, "{"+p.Tokens[i][j].Key+"}")
				}

				if !tCount {
					params = params[:len(params)-1]
				}

				params = append(params, array...)

				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeIndex:
				//debug.Log("parameters.TokenTypeIndex:", p.Tokens[i][j].Key)
				match := rxTokenIndex.FindStringSubmatch(p.Tokens[i][j].Key)
				if len(match) != 3 {
					params[len(params)-1] = p.Tokens[i][j].Key
					tCount = true
					continue
				}

				block := []rune("$" + match[1] + "->[" + match[2] + "]")
				fork := prc.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				fork.Execute(block)
				b, err := fork.Stdout.ReadAll()
				if err != nil {
					return err
				}

				b = utils.CrLfTrim(b)

				params[len(params)-1] += string(b)
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeRange:
				// TODO: write me!
				debug.Log("parameters.TokenTypeRange:", p.Tokens[i][j].Key)
				//panic("TODO: write me!")

			case parameters.TokenTypeTilde:
				if len(p.Tokens[i][j].Key) == 0 {
					params[len(params)-1] += home.MyDir
				} else {
					params[len(params)-1] += home.UserDir(p.Tokens[i][j].Key)
				}
				tCount = true
				namedPipeIsParam = true

			default:
				err := fmt.Errorf(
					`unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"`,
					p.Tokens[i][j].Type, i, j, p.Tokens[i][j].Key,
				)
				return err
			}
		}

		if !tCount {
			params = params[:len(params)-1]
		}

	}

	p.DefineParsed(params)

	return nil
}
