package parameters

import (
	"encoding/json"
	"fmt"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"strings"
)

func (p *Parameters) Parse(vars *types.Vars) {
	l := len(p.Tokens)
	p.Params = make([]string, l)
	tCount := make([]int, l)

	for i := 0; i < l; i++ {
		for j := range p.Tokens[i] {
			switch p.Tokens[i][j].Type {
			case TokenTypeNil:
				// do nothing

			case TokenTypeValue:
				p.Params[i] += p.Tokens[i][j].Key
				tCount[i]++

			case TokenTypeString:
				p.Params[i] += vars.GetString(p.Tokens[i][j].Key)
				tCount[i]++

			case TokenTypeBlockString:
				stdout := streams.NewStdin()
				lang.ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, nil, types.Null)
				p.Params[i] += string(stdout.ReadAll())
				tCount[i]++

			case TokenTypeArray:
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
				p.Params = append(p.Params, array...)
				p.Params = append(p.Params, p.Params[i+1]...)
				tCount = append(tCount, make([]int, len(array))...)
				for k := range array {
					tCount[i+k] = 1
				}
				i += len(array)

			case TokenTypeBlockArray:
				var err error
				var array []string

				stdout := streams.NewStdin()
				lang.ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, nil, types.Null)

				b := []byte(stdout.ReadAll())
				if types.IsArray(b) {
					err = json.Unmarshal(b, &array)
				}
				if err != nil || !types.IsArray(b) {
					s := strings.Replace(p.Tokens[i][j].Key, "\r", "", -1)
					array = strings.Split(s, "\n")
				}
				// add to params
				p.Params = append(p.Params, array...)
				p.Params = append(p.Params, p.Params[i+1]...)
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
