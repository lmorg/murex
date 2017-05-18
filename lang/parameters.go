package lang

import (
	"encoding/json"
	"fmt"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"strings"
)

func parseParameters(p *parameters.Parameters, vars *types.Vars) {
	l := len(p.Tokens)
	p.Params = make([]string, l)
	tCount := make([]int, l)

	for i := 0; i < l; i++ {
		for j := range p.Tokens[i] {
			switch p.Tokens[i][j].Type {
			case parameters.TokenTypeNil:
				// do nothing

			case parameters.TokenTypeValue:
				p.Params[i] += p.Tokens[i][j].Key
				tCount[i]++

			case parameters.TokenTypeString:
				p.Params[i] += vars.GetString(p.Tokens[i][j].Key)
				tCount[i]++

			case parameters.TokenTypeBlockString:
				stdout := streams.NewStdin()
				ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, nil, types.Null)
				p.Params[i] += string(stdout.ReadAll())
				tCount[i]++

			case parameters.TokenTypeArray:
				var err error
				var array []string
				b := []byte(vars.GetString(p.Tokens[i][j].Key))
				if types.IsArray(b) {
					err = json.Unmarshal(b, &array)
				}
				if err != nil || !types.IsArray(b) {
					s := strings.Replace(p.Tokens[i][j].Key, "\r", "", -1)
					array = strings.Split(s, "\n")
				}
				// add to params
				params := append(p.Params[i:], array...)
				p.Params = append(params, p.Params[i+1:]...)
				tCount = append(tCount, make([]int, len(array))...)
				for k := range array {
					tCount[i+k] = 1
				}
				i += len(array)

			case parameters.TokenTypeBlockArray:
				var err error
				var array []string

				stdout := streams.NewStdin()
				ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, nil, types.Null)

				b := []byte(stdout.ReadAll())
				if types.IsArray(b) {
					err = json.Unmarshal(b, &array)
				}
				if err != nil || !types.IsArray(b) {
					s := strings.Replace(p.Tokens[i][j].Key, "\r", "", -1)
					array = strings.Split(s, "\n")
				}
				// add to params
				params := append(p.Params[i:], array...)
				p.Params = append(params, p.Params[i+1:]...)
				tCount = append(tCount, make([]int, len(array))...)
				for k := range array {
					tCount[i+k] = 1
				}
				i += len(array)

			default:
				panic(fmt.Sprintf(
					`Unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"`,
					p.Tokens[i][j].Type, i, j, p.Tokens[i][j].Key,
				))
			}
		}
	}

	if len(p.Tokens) != 0 && tCount[len(p.Tokens)-1] == 0 {
		if len(p.Tokens) == 1 {
			p.Params = make([]string, 0)
		} else {
			p.Params = p.Params[:len(p.Tokens)-1]
		}
	}
}
