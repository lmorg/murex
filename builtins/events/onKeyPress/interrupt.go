package onkeypress

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Line         string
	Raw          string
	Pos          int
	KeySequence  string
	Invoker      string
	IsMasked     bool
	PreviewMode  string
	ReadlineMode string
}

const (
	modeDefault      = "Normal"
	modeVimKeys      = "VimKeys"
	modeAutocomplete = "Autocomplete"
	modeFuzzyFind    = "FuzzyFind"
)

const (
	modePreviewOff  = "false"
	modePreviewItem = "Item"
	modePreviewLine = "Line"
)
