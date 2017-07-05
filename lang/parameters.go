package lang

import (
	"fmt"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

func parseParameters(p *parameters.Parameters, vars *types.Vars) {
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
				ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, nil, types.Null)
				stdout.Close()
				b := stdout.ReadAll()

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

				variable := new(streams.Stdin)
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
				ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, nil, types.Null)
				stdout.Close()

				stdout.ReadArray(func(b []byte) {
					array = append(array, string(b))
				})

				if !tCount {
					p.Params = p.Params[:len(p.Params)-1]
				}

				p.Params = append(p.Params, array...)

				tCount = true

			default:
				panic(fmt.Sprintf(
					`Unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"`,
					p.Tokens[i][j].Type, i, j, p.Tokens[i][j].Key,
				))
			}
		}

		if !tCount {
			p.Params = p.Params[:len(p.Params)-1]
		}

	}
}
