package parameters

import (
	"fmt"
	"github.com/lmorg/murex/lang/types"
)

/*func subStrReplace(old, new string, start, end int) (s string) {
	s = old[:start]
	s += new
	s += old[end:]
	return
}*/

func newString(old string, start int, length int, new string)

func (p *Parameters) Parse(vars *types.Vars) {
	/*for i := range p.Params {
		if len(p.Params[i]) > 1 && (p.Params[i][0] != '{' || p.Params[i][len(p.Params[i])-1] != '}') {
			// TODO: I shouldn't need to check for code blocks when I tokenise the vars
			vars.KeyValueReplace(&p.Params[i])
		}
	}*/

	for i := range p.Tokens {
		var updated []string
		for j := range p.Tokens[i] {
			switch p.Tokens[i][j].Type {
			case 0: // do nothing

			case TokenTypeString:
				old := p.Params[i]

				/*old := p.Params[i]
				new := vars.GetString(p.Tokens[i][j].Key)
				start := p.Tokens[i][j].StrLoc
				end := start + len(p.Tokens[i][j].Key) + 1
				adjust += len(new) - len(p.Tokens[i][j].Key) - 1
				p.Params[i] = subStrReplace(old, new, start+adjust, end+adjust)

				s := p.Params[i][:p.Tokens[i][j].StrLoc] + vars.GetString(p.Tokens[i][j].Key)
				s += p.Params[i][p.Tokens[i][j].StrLoc+len(p.Tokens[i][j].Key)+adjust]
				adjust += p.P*/

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
}
