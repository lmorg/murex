package lang

import (
	"fmt"
	"regexp"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/home"
)

var rxTokenIndex = regexp.MustCompile(`(.*?)\[(.*?)\]`)

// ParseParameters is an internal function to parse parameters
func ParseParameters(prc *proc.Process, p *parameters.Parameters) {
	for i := range p.Tokens {
		p.Params = append(p.Params, "")

		var tCount bool
		for j := range p.Tokens[i] {
			switch p.Tokens[i][j].Type {
			case parameters.TokenTypeNil:
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
				stdout := streams.NewStdin()
				RunBlockExistingConfigSpace([]rune(p.Tokens[i][j].Key), nil, stdout, prc.Stderr, prc)
				b, err := stdout.ReadAll()
				if err != nil {
					ansi.Stderrln(prc, ansi.FgRed, err.Error())
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

				stdout := streams.NewStdin()
				RunBlockExistingConfigSpace([]rune(p.Tokens[i][j].Key), nil, stdout, prc.Stderr, prc)

				stdout.ReadArray(func(b []byte) {
					array = append(array, string(b))
				})

				if !tCount {
					p.Params = p.Params[:len(p.Params)-1]
				}

				p.Params = append(p.Params, array...)

				tCount = true

			case parameters.TokenTypeIndex:
				debug.Log("parameters.TokenTypeIndex:", p.Tokens[i][j].Key)
				match := rxTokenIndex.FindStringSubmatch(p.Tokens[i][j].Key)
				if len(match) != 3 {
					p.Params[len(p.Params)-1] = p.Tokens[i][j].Key
					tCount = true
					continue
				}

				block := []rune("$" + match[1] + "->[" + match[2] + "]")
				stdout := streams.NewStdin()
				RunBlockExistingConfigSpace(block, nil, stdout, prc.Stderr, prc)

				b, err := stdout.ReadAll()
				if err != nil {
					ansi.Stderrln(prc, ansi.FgRed, err.Error())
				}

				b = utils.CrLfTrim(b)

				p.Params[len(p.Params)-1] += string(b)
				tCount = true

			case parameters.TokenTypeRange:
				debug.Log("parameters.TokenTypeRange:", p.Tokens[i][j].Key)

			case parameters.TokenTypeTilde:
				if len(p.Tokens[i][j].Key) == 0 {
					p.Params[len(p.Params)-1] += home.MyDir
				} else {
					p.Params[len(p.Params)-1] += home.UserDir(p.Tokens[i][j].Key)
				}
				tCount = true

			default:
				//os.Stderr.WriteString(fmt.Sprintf(
				ansi.Stderrln(prc, ansi.FgRed, fmt.Sprintf(
					`Unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"%s`,
					p.Tokens[i][j].Type, i, j, p.Tokens[i][j].Key,
					utils.NewLineString,
				))
			}
		}

		if !tCount {
			p.Params = p.Params[:len(p.Params)-1]
		}

	}
}
