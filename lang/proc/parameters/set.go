package parameters

// Internal function: Prepend the pre-parsed parameter token tree with new tokens
func (p *Parameters) SetPrepend(value string) {
	pt := make([][]ParamToken, 1)
	pt[0] = make([]ParamToken, 1)
	pt[0][0].Key = value
	pt[0][0].Type = TokenTypeValue
	p.Tokens = append(pt, p.Tokens...)
}

// Internal function: define a parameter token when parsing
func (p *Parameters) SetTokens(tokens [][]ParamToken) {
	p.Tokens = tokens
}
