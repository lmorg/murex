package readline

// AddEvent registers a new keypress handler
func (rl *Instance) AddEvent(keyPress string, callback func(string, []rune, int) (bool, bool, bool, []rune)) {
	rl.evtKeyPress[keyPress] = callback
}

// DelEvent deregisters an existing keypress handler
func (rl *Instance) DelEvent(keyPress string) {
	delete(rl.evtKeyPress, keyPress)
}
