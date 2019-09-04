package lang

import (
	"fmt"
	"regexp"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/home"
)

var rxTokenIndex = regexp.MustCompile(`(.*?)\[(.*?)\]`)

// ParseParameters is an internal function to parse parameters
func ParseParameters(prc *Process, p *parameters.Parameters) {
	for i := range p.Tokens {
		p.Params = append(p.Params, "")

		var tCount bool
		for j := range p.Tokens[i] {
			switch p.Tokens[i][j].Type {
			case parameters.TokenTypeNil:
				// do nothing

			case parameters.TokenTypeNamedPipe:
				// do nothing

			case parameters.TokenTypeValue:
				p.Params[len(p.Params)-1] += p.Tokens[i][j].Key
				tCount = true

			case parameters.TokenTypeString:
				s := prc.Variables.GetString(p.Tokens[i][j].Key)
				s = utils.CrLfTrimString(s)
				p.Params[len(p.Params)-1] += s
				tCount = true

			case parameters.TokenTypeBlockString:
				fork := prc.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				fork.Execute([]rune(p.Tokens[i][j].Key))
				b, err := fork.Stdout.ReadAll()
				if err != nil {
					prc.Stderr.Writeln([]byte(err.Error()))
				}

				b = utils.CrLfTrim(b)

				p.Params[len(p.Params)-1] += string(b)
				tCount = true

			case parameters.TokenTypeArray:
				var array []string

				variable := streams.NewStdin()
				variable.SetDataType(prc.Variables.GetDataType(p.Tokens[i][j].Key))
				variable.Write([]byte(prc.Variables.GetString(p.Tokens[i][j].Key)))

				variable.ReadArray(func(b []byte) {
					array = append(array, string(b))
				})

				if !tCount {
					p.Params = p.Params[:len(p.Params)-1]
				}

				p.Params = append(p.Params, array...)

				tCount = true

			case parameters.TokenTypeBlockArray:
				var array []string

				fork := prc.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				fork.Execute([]rune(p.Tokens[i][j].Key))
				fork.Stdout.ReadArray(func(b []byte) {
					array = append(array, string(b))
				})

				if !tCount {
					p.Params = p.Params[:len(p.Params)-1]
				}

				p.Params = append(p.Params, array...)

				tCount = true

			case parameters.TokenTypeIndex:
				//debug.Log("parameters.TokenTypeIndex:", p.Tokens[i][j].Key)
				match := rxTokenIndex.FindStringSubmatch(p.Tokens[i][j].Key)
				if len(match) != 3 {
					p.Params[len(p.Params)-1] = p.Tokens[i][j].Key
					tCount = true
					continue
				}

				block := []rune("$" + match[1] + "->[" + match[2] + "]")
				fork := prc.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				fork.Execute(block)
				b, err := fork.Stdout.ReadAll()
				if err != nil {
					prc.Stderr.Writeln([]byte(err.Error()))
				}

				b = utils.CrLfTrim(b)

				p.Params[len(p.Params)-1] += string(b)
				tCount = true

			case parameters.TokenTypeRange:
				// TODO: write me!
				//debug.Log("parameters.TokenTypeRange:", p.Tokens[i][j].Key)

			case parameters.TokenTypeTilde:
				if len(p.Tokens[i][j].Key) == 0 {
					p.Params[len(p.Params)-1] += home.MyDir
				} else {
					p.Params[len(p.Params)-1] += home.UserDir(p.Tokens[i][j].Key)
				}
				tCount = true

			default:
				prc.Stderr.Writeln([]byte(fmt.Sprintf(
					`Unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"%s`,
					p.Tokens[i][j].Type, i, j, p.Tokens[i][j].Key,
					utils.NewLineString,
				)))
			}
		}

		if !tCount {
			p.Params = p.Params[:len(p.Params)-1]
		}

	}
}
