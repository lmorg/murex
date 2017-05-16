package parameters

func (p *Parameters) SetPrepend(value string) {
	pt := make([][]ParamToken, 1)
	pt[0] = make([]ParamToken, 1)
	pt[0][0].Key = value
	pt[0][0].Type = TokenTypeValue
	p.tokens = append(pt, p.tokens...)
}

func (p *Parameters) SetTokens(tokens [][]ParamToken) {
	p.tokens = tokens
}
