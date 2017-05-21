package lang

import (
	"encoding/json"
	"fmt"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"regexp"
)

var rxNewLine *regexp.Regexp = regexp.MustCompile(`[\r\n]+`)

func parseParameters(p *parameters.Parameters, vars *types.Vars) {
	//p.Params = make([]string, 0)
	//tCount := make([]bool, len(p.Tokens))
	//var tCount []bool
	debug.Json("####################################", p.Tokens)
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
				var err error
				var array []string

				b := []byte(vars.GetString(p.Tokens[i][j].Key))

				if types.IsArray(b) {
					err = json.Unmarshal(b, &array)
				}
				if err != nil || !types.IsArray(b) {
					array = rxNewLine.Split(string(b), -1)
				}
				if array[0] == "" {
					array = array[1:]
				}

				if array[len(array)-1] == "" {
					array = array[:len(array)-1]
				}

				if !tCount {
					p.Params = p.Params[:len(p.Params)-1]
				}

				p.Params = append(p.Params, array...)
				t := make([]bool, len(array))
				for ti := range t {
					t[ti] = true
				}

				tCount = true

			case parameters.TokenTypeBlockArray:
				var err error
				var array []string

				stdout := streams.NewStdin()
				ProcessNewBlock([]rune(p.Tokens[i][j].Key), nil, stdout, nil, types.Null)
				stdout.Close()

				b := []byte(stdout.ReadAll())

				if types.IsArray(b) {
					err = json.Unmarshal(b, &array)
				}
				if err != nil || !types.IsArray(b) {
					array = rxNewLine.Split(string(b), -1)
				}
				if array[0] == "" {
					array = array[1:]
				}

				if array[len(array)-1] == "" {
					array = array[:len(array)-1]
				}

				if !tCount {
					p.Params = p.Params[:len(p.Params)-1]
				}

				p.Params = append(p.Params, array...)
				t := make([]bool, len(array))
				for ti := range t {
					t[ti] = true
				}

				tCount = true

			default:
				panic(fmt.Sprintf(
					`Unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"`,
					p.Tokens[i][j].Type, i, j, p.Tokens[i][j].Key,
				))
			}
		}
		debug.Log("#######################################", tCount)
		debug.Json("############### before ", p.Params)
		if !tCount {
			p.Params = p.Params[:len(p.Params)-1]
		}
		debug.Json("############### after ", p.Params)
	}

	/*if len(p.Tokens) != 0 && tCount[len(tCount)-1] == false {
		if len(p.Tokens) == 1 {
			p.Params = make([]string, 0)
		} else {
			p.Params = p.Params[:len(p.Tokens)-1]
		}
	}*/
	/*for i := 0; i < len(tCount); i++ {
		if !tCount[i] {
			switch {
			case i == 0:
				p.Params = p.Params[1:]
				tCount = tCount[1:]
			case i == len(tCount):
				p.Params = p.Params[:i]
				tCount = tCount[:i]
			default:
				p.Params = append(p.Params[:i], p.Params[i+1:]...)
				tCount = append(tCount[:i], tCount[i+1:]...)
			}
			i--
		}
	}*/
}
