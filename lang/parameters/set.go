package parameters

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
	for i := range params {
		p.PreParsed = append([][]rune{[]rune(params[i])}, p.PreParsed...)
	}
	p.mutex.Unlock()
}

func (p *Parameters) CopyFrom(src *Parameters) {
	p.mutex.Lock()
	src.mutex.Lock()

	p.params = make([]string, len(src.params))
	copy(p.params, src.params)

	p.mutex.Unlock()
	src.mutex.Unlock()
}
