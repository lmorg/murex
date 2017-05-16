package parameters

import (
	"fmt"
	"github.com/lmorg/murex/lang/types"
)

func (p *Parameters) Parse(vars *types.Vars) {
	p.params = make([]string, len(p.tokens))
	for i := range p.tokens {
		for j := range p.tokens[i] {
			switch p.tokens[i][j].Type {
			case TokenTypeNil:
				// do nothing

			case TokenTypeValue:
				p.params[i] += p.tokens[i][j].Key

			case TokenTypeString:
				p.params[i] += vars.GetString(p.tokens[i][j].Key)

			case TokenTypeBlockString:

			case TokenTypeArray:

			case TokenTypeBlockArray:

			default:
				panic(fmt.Sprintf(
					`Unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"`,
					p.tokens[i][j].Type, i, j, p.tokens[i][j].Key,
				))
			}
		}
	}

	if len(p.tokens) != 0 && len(p.tokens[len(p.tokens)-1]) != 0 && p.tokens[len(p.tokens)-1][0].Type == TokenTypeNil {
		if len(p.tokens) == 1 {
			p.params = make([]string, 0)
		} else {
			p.params = p.params[:len(p.tokens)-1]
		}
	}
}
