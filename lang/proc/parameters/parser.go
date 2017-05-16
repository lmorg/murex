package parameters

import (
	"fmt"
	"github.com/lmorg/murex/lang/types"
)

func (p *Parameters) Parse(vars *types.Vars) {
	p.Params = make([]string, len(p.Tokens))
	tCount := make([]int, len(p.Tokens))

	for i := range p.Tokens {
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
				tCount[i]++

			case TokenTypeArray:
				tCount[i]++

			case TokenTypeBlockArray:
				tCount[i]++

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

	/*if len(p.Tokens) != 0 && len(p.Tokens[len(p.Tokens)-1]) != 0 && p.Tokens[len(p.Tokens)-1][0].Type == TokenTypeNil {
		if len(p.Tokens) == 1 {
			p.Params = make([]string, 0)
		} else {
			p.Params = p.Params[:len(p.Tokens)-1]
		}
	}*/
}
