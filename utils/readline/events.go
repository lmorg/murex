package readline

// EventReturn is a structure returned by the callback event function.
// This is used by readline to determine what state the API should
// return to after the readline event.
type EventReturn struct {
	Actions    []func(rl *Instance)
	HintText   []rune
	SetLine    []rune
	SetPos     int
	NextEvent  bool
	MoreEvents bool
}

// keyPressEventCallbackT: keyPress, eventId, line, pos
type keyPressEventCallbackT func(string, int, []rune, int) *EventReturn

// AddEvent registers a new keypress handler
func (rl *Instance) AddEvent(keyPress string, callback keyPressEventCallbackT) {
	rl.evtKeyPress[keyPress] = callback
}

// DelEvent deregisters an existing keypress handler
func (rl *Instance) DelEvent(keyPress string) {
	delete(rl.evtKeyPress, keyPress)
}
