package parameters

import (
	"github.com/lmorg/murex/lang/types"
)

type Parameters struct {
	params []string
	tokens [][]ParamToken
}

func (p *Parameters) SetPrepend(parameter string) {
	p.params = append([]string{parameter}, p.params...)
	p.tokens = append([][]ParamToken{}, p.tokens...)
}

func (p *Parameters) SetAll(parameters []string) {
	p.params = parameters
}

func (p *Parameters) SetTokens(tokens [][]ParamToken) {
	p.tokens = tokens
}

func (p *Parameters) Parse(vars *types.Vars) {
	for i := range p.params {
		if len(p.params[i]) > 1 && (p.params[i][0] != '{' || p.params[i][len(p.params[i])-1] != '}') {
			// TODO: I shouldn't need to check for code blocks when I tokenise the vars
			vars.KeyValueReplace(&p.params[i])
		}
	}
}
