package readline

func (rl *Instance) AddEvent(keyPress string, callback func(string, []rune, int) (bool, bool)) {
	rl.evtKeyPress[keyPress] = callback
}

func (rl *Instance) DelEvent(keyPress string) {
	delete(rl.evtKeyPress, keyPress)
}
