package parameters

func (p *Parameters) SetPrepend(parameter string) {
	p.Params = append([]string{parameter}, p.Params...)
	p.Tokens = append([][]ParamToken{}, p.Tokens...)
}

func (p *Parameters) SetAll(parameters []string) {
	p.Params = parameters
}

func (p *Parameters) SetTokens(tokens [][]ParamToken) {
	p.Tokens = tokens
}
