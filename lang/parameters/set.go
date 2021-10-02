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

// ShiftLeft returns the top most parameter and removes it from the stack
/*func (p *Parameters) ShiftLeft() (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	switch len(p.params) {
	case 0:
		return "", fmt.Errorf("cannot ShiftLeft() when p.Len() == 0")
	case 1:
		left := p.params[0]
		p.params = []string{}
		return left, nil
	default:
		left := p.params[0]
		p.params = p.params[1:]
		return left, nil
	}
}*/

// DefineParsed overrides all of the parsed parameters for a given process
func (p *Parameters) DefineParsed(params []string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.params = params
}

// Prepend inserts one or multiple parameters at the top of stack
func (p *Parameters) Prepend(params []string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.params = append(params, p.params...)
}
