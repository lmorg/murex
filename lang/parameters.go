package lang

import (
	"fmt"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"regexp"
)

var rxNewLine *regexp.Regexp = regexp.MustCompile(`[\r\n]+`)

func parseParameters(p *parameters.Parameters, vars *types.Vars) {
	//debug.Json("####################################", p.Tokens)
	for i := range p.Tokens {
		p.Params = append(p.Params, "")
		//tCount = append(tCount, false)
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

				if b[len(b)-1] == '\n' {
					b = b[:len(b)-1]
				}

				if b[len(b)-1] == '\r' {
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

				/*if types.IsArray(b) {
					err = json.Unmarshal(b, &array)
					debug.Log("json.Unmarshal(b, &array) return:", err)
				}
				if err != nil || !types.IsArray(b) {
					array = rxNewLine.Split(string(b), -1)
				}
				if array[0] == "" {
					array = array[1:]
				}

				if array[len(array)-1] == "" {
					array = array[:len(array)-1]
				}*/

				if !tCount {
					p.Params = p.Params[:len(p.Params)-1]
				}

				p.Params = append(p.Params, array...)
				// i can't remember what this does...
				//t := make([]bool, len(array))
				//for ti := range t {
				//	t[ti] = true
				//}

				tCount = true

			case parameters.TokenTypeBlockArray:
				var array []string

				stdout := streams.NewStdin()
				ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, nil, types.Null)
				stdout.Close()

				/*b := []byte(stdout.ReadAll())

				if types.IsArray(b) {
					err = json.Unmarshal(b, &array)
					debug.Log("json.Unmarshal(b, &array) return:", err)
				}
				if err != nil || !types.IsArray(b) {
					array = rxNewLine.Split(string(b), -1)
				}*/

				stdout.ReadArray(func(b []byte) {
					array = append(array, string(b))
				})

				//if array[0] == "" {
				//	array = array[1:]
				//}

				//if array[len(array)-1] == "" {
				//	array = array[:len(array)-1]
				//}

				if !tCount {
					p.Params = p.Params[:len(p.Params)-1]
				}

				p.Params = append(p.Params, array...)
				// i can't remember what this does...
				//t := make([]bool, len(array))
				//for ti := range t {
				//	t[ti] = true
				//}

				tCount = true

			default:
				panic(fmt.Sprintf(
					`Unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"`,
					p.Tokens[i][j].Type, i, j, p.Tokens[i][j].Key,
				))
			}
		}
		//debug.Log("#######################################", tCount)
		//debug.Json("############### before ", p.Params)
		if !tCount {
			p.Params = p.Params[:len(p.Params)-1]
		}
		//debug.Json("############### after ", p.Params)
	}
}
