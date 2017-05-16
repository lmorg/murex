package parameters

import (
	"fmt"
	"github.com/lmorg/murex/lang/types"
)

func (p *Parameters) Parse(vars *types.Vars) {
	p.params = make([]string, len(p.Tokens))
	for i := range p.Tokens {
		for j := range p.Tokens[i] {
			switch p.Tokens[i][j].Type {
			case TokenTypeNil:
				// do nothing

			case TokenTypeValue:
				p.params[i] += p.Tokens[i][j].Key

			case TokenTypeString:
				p.params[i] += vars.GetString(p.Tokens[i][j].Key)

			case TokenTypeBlockString:

			case TokenTypeArray:

			case TokenTypeBlockArray:

			default:
				panic(fmt.Sprintf(
					`Unexpected token type (%d) in parsed parameters. Param[%d][%d] == "%s"`,
					p.Tokens[i][j].Type, i, j, p.Tokens[i][j].Key,
				))
			}
		}
	}

	if len(p.Tokens) != 0 && len(p.Tokens[len(p.Tokens)-1]) != 0 && p.Tokens[len(p.Tokens)-1][0].Type == TokenTypeNil {
		if len(p.Tokens) == 1 {
			p.params = make([]string, 0)
		} else {
			p.params = p.params[:len(p.Tokens)-1]
		}
	}
}
