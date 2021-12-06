package parameters

// SetPrepend - internal function to prepend the pre-parsed parameter token tree with new tokens
func (p *Parameters) SetPrepend(value string) {
	pt := make([][]ParamToken, 1)
	pt[0] = make([]ParamToken, 1)
	pt[0][0].Key = value
	pt[0][0].Type = TokenTypeValue
	p.Tokens = append(pt, p.Tokens...)
}

// SetTokens - internal function to define a parameter token when parsing
func (p *Parameters) SetTokens(tokens [][]ParamToken) {
	p.Tokens = make([][]ParamToken, len(tokens))

	for i := range tokens {
		p.Tokens[i] = make([]ParamToken, len(tokens[i]))
		copy(p.Tokens[i], tokens[i])
	}
}

// DefineParsed overrides all of the parsed parameters for a given process
func (p *Parameters) DefineParsed(params []string) {
	p.mutex.Lock()
	p.params = params
	p.mutex.Unlock()
}

// Prepend inserts one or multiple parameters at the top of stack
func (p *Parameters) Prepend(params []string) {
	p.mutex.Lock()
	p.params = append(params, p.params...)
	p.mutex.Unlock()
}

func (p *Parameters) PointTo(cf *Parameters) {
	p.mutex.Lock()
	cf.mutex.Lock()
	p.params = cf.params
	p.Tokens = cf.Tokens
	p.mutex.Unlock()
	cf.mutex.Unlock()
}
