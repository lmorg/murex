package readline

// EventState presents a simplified view of the current readline state
type EventState struct {
	Line        string
	CursorPos   int
	KeyPress    string
	IsMasked    bool
	InputMode   string
	PreviewMode string
}

// EventReturn is a structure returned by the callback event function.
// This is used by readline to determine what state the API should
// return to after the readline event.
type EventReturn struct {
	Actions    []func(rl *Instance)
	HintText   []rune
	SetLine    []rune
	SetPos     int
	Continue   bool
	MoreEvents bool
}

// keyPressEventCallbackT: keyPress, eventId, line, pos
type keyPressEventCallbackT func(int, *EventState) *EventReturn

// AddEvent registers a new keypress handler
func (rl *Instance) AddEvent(keyPress string, callback keyPressEventCallbackT) {
	rl.evtKeyPress[keyPress] = callback
}

// DelEvent deregisters an existing keypress handler
func (rl *Instance) DelEvent(keyPress string) {
	delete(rl.evtKeyPress, keyPress)
}

func (rl *Instance) newEventState(keyPress string) *EventState {
	return &EventState{
		Line:        rl.line.String(),
		CursorPos:   rl.line.RunePos(),
		KeyPress:    keyPress,
		IsMasked:    rl.PasswordMask > 0,
		InputMode:   rl._getInputMode(),
		PreviewMode: rl._getPreviewMode(),
	}
}

const (
	EventModeInputDefault        = "Normal"
	EventModeInputVimKeys        = "VimKeys"
	EventModeInputVimReplaceOnce = "VimReplaceOnce"
	EventModeInputVimReplaceMany = "VimReplaceMany"
	EventModeInputVimDelete      = "VimDelete"
	EventModeInputVimCommand     = "VimCommand"
	EventModeInputAutocomplete   = "Autocomplete"
	EventModeInputFuzzyFind      = "FuzzyFind"
)

// _getInputMode is used purely for event reporting
func (rl *Instance) _getInputMode() string {
	switch {
	case rl.modeViMode == vimKeys:
		return EventModeInputVimKeys
	case rl.modeViMode == vimReplaceOnce:
		return EventModeInputVimReplaceOnce
	case rl.modeViMode == vimReplaceMany:
		return EventModeInputVimReplaceMany
	case rl.modeViMode == vimDelete:
		return EventModeInputVimDelete
	case rl.modeViMode == vimCommand:
		return EventModeInputVimCommand
	case rl.modeTabFind:
		return EventModeInputFuzzyFind
	case rl.modeTabCompletion:
		return EventModeInputAutocomplete
	default:
		return EventModeInputDefault
	}
}

const (
	EventModePreviewOff     = "Disabled"
	EventModePreviewItem    = "Autocomplete"
	EventModePreviewLine    = "CmdLine"
	EventModePreviewUnknown = "Unknown"
)

// _getPreviewMode is used purely for event reporting
func (rl *Instance) _getPreviewMode() string {
	switch {
	case rl.previewMode == previewModeClosed:
		return EventModePreviewOff
	case rl.previewRef == previewRefLine:
		return EventModePreviewLine
	case rl.previewRef == previewRefDefault:
		return EventModePreviewItem
	default:
		return EventModePreviewUnknown
	}
}
