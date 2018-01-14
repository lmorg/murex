package lang

import (
	"fmt"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/home"
	"regexp"
)

var rxTokenIndex = regexp.MustCompile(`(.*?)\[(.*?)\]`)

// ParseParameters is an internal function to parse parameters
func ParseParameters(prc *proc.Process, p *parameters.Parameters, vars *types.Vars) {
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
				p.Params[len(p.Params)-1] += vars.GetString(p.Tokens[i][j].Key)
				tCount = true

			case parameters.TokenTypeBlockString:
				stdout := streams.NewStdin()
				ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, prc.Stderr, prc)
				stdout.Close()
				b, err := stdout.ReadAll()
				if err != nil {
					//os.Stderr.WriteString(err.Error() + utils.NewLineString)
					//ansi.Stderrln(ansi.FgRed, err.Error())
					prc.Stderr.Writeln([]byte(err.Error()))
				}

				if len(b) > 0 && b[len(b)-1] == '\n' {
					b = b[:len(b)-1]
				}

				if len(b) > 0 && b[len(b)-1] == '\r' {
					b = b[:len(b)-1]
				}

				p.Params[len(p.Params)-1] += string(b)
				tCount = true

			case parameters.TokenTypeArray:
				var array []string

				variable := streams.NewStdin()
				variable.SetDataType(vars.GetType(p.Tokens[i][j].Key))
				variable.Write([]byte(vars.GetString(p.Tokens[i][j].Key)))
				variable.Close()

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
				ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, prc.Stderr, prc)
				stdout.Close()

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
				ProcessNewBlock(block, nil, stdout, nil, prc)
				stdout.Close()
				b, err := stdout.ReadAll()
				if err != nil {
					ansi.Stderrln(ansi.FgRed, err.Error())
				}

				if len(b) > 0 && b[len(b)-1] == '\n' {
					b = b[:len(b)-1]
				}

				if len(b) > 0 && b[len(b)-1] == '\r' {
					b = b[:len(b)-1]
				}

				p.Params[len(p.Params)-1] += string(b)
				tCount = true

			case parameters.TokenTypeTilde:
				if len(p.Tokens[i][j].Key) == 0 {
					p.Params[len(p.Params)-1] += home.MyDir
				} else {
					p.Params[len(p.Params)-1] += home.UserDir(p.Tokens[i][j].Key)
				}
				tCount = true

			default:
				//os.Stderr.WriteString(fmt.Sprintf(
				ansi.Stderrln(ansi.FgRed, fmt.Sprintf(
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
