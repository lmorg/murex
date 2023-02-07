package readline

import (
	"os"
	"sync"
)

var (
	primary *os.File = os.Stdout
	replica *os.File = os.Stdin
)

func SetTTY(primaryTTY, replicaTTY *os.File) {
	primary = primaryTTY
	replica = replicaTTY
}

var ForceCrLf = true

// Instance is used to encapsulate the parameter group and run time of any given
// readline instance so that you can reuse the readline API for multiple entry
// captures without having to repeatedly unload configuration.
type Instance struct {
	fdMutex sync.Mutex

	Active        bool
	closeSigwinch func()

	// PasswordMask is what character to hide password entry behind.
	// Once enabled, set to 0 (zero) to disable the mask again.
	PasswordMask rune

	// SyntaxHighlight is a helper function to provide syntax highlighting.
	// Once enabled, set to nil to disable again.
	SyntaxHighlighter func([]rune) string

	// History is an interface for querying the readline history.
	// This is exposed as an interface to allow you the flexibility to define how
	// you want your history managed (eg file on disk, database, cloud, or even
	// no history at all). By default it uses a dummy interface that only stores
	// historic items in memory.
	History History

	// HistoryAutoWrite defines whether items automatically get written to
	// history.
	// Enabled by default. Set to false to disable.
	HistoryAutoWrite bool // = true

	// TabCompleter is a simple function that offers completion suggestions.
	// It takes the readline line ([]rune) and cursor pos. Returns a prefix
	// string, an array of suggestions and a map of definitions (optional).
	TabCompleter      func([]rune, int, DelayedTabContext) (string, []string, map[string]string, TabDisplayType)
	delayedTabContext DelayedTabContext

	MinTabItemLength int
	MaxTabItemLength int

	// MaxTabCompletionRows is the maximum number of rows to display in the tab
	// completion grid.
	MaxTabCompleterRows int // = 4

	// SyntaxCompletion is used to autocomplete code syntax (like braces and
	// quotation marks). If you want to complete words or phrases then you might
	// be better off using the TabCompletion function.
	// SyntaxCompletion takes the line ([]rune), change (string) and cursor
	// position, and returns the new line and cursor position.
	SyntaxCompleter func([]rune, string, int) ([]rune, int)

	// DelayedSyntaxWorker allows for syntax highlighting happen to the line
	// after the line has been drawn.
	DelayedSyntaxWorker func([]rune) []rune
	delayedSyntaxCount  int32

	// HintText is a helper function which displays hint text the prompt.
	// HintText takes the line input from the prompt and the cursor position.
	// It returns the hint text to display.
	HintText func([]rune, int) []rune

	// HintColor any ANSI escape codes you wish to use for hint formatting. By
	// default this will just be blue.
	HintFormatting string

	// AutocompleteHistory is another customization allowing for alternative
	// results when [ctrl]+[r]
	AutocompleteHistory func(string) ([]string, map[string]string)

	// TempDirectory is the path to write temporary files when editing a line in
	// $EDITOR. This will default to os.TempDir()
	TempDirectory string

	// GetMultiLine is a callback to your host program. Since multiline support
	// is handled by the application rather than readline itself, this callback
	// is required when calling $EDITOR. However if this function is not set
	// then readline will just use the current line.
	GetMultiLine func([]rune) []rune

	MaxCacheSize int
	cacheHint    cacheSliceRune
	cacheSyntax  cacheString
	//cacheSyntaxHighlight cacheString
	//cacheSyntaxDelayed   cacheSliceRune

	// readline operating parameters
	prompt        string //  = ">>> "
	promptLen     int    //= 4
	line          []rune
	lineChange    string // cache what had changed from previous line
	pos           int
	termWidth     int
	multiline     []byte
	multiSplit    []string
	skipStdinRead bool

	// history
	lineBuf string
	histPos int

	// hint text
	hintY    int //= 0
	hintText []rune

	ShowPreviews  bool
	previewCache  *previewCacheT
	PreviewImages bool

	// tab completion
	modeTabCompletion bool
	tabMutex          sync.Mutex
	tcPrefix          string
	tcSuggestions     []string
	tcDescriptions    map[string]string
	tcDisplayType     TabDisplayType
	tcOffset          int
	tcPosX            int
	tcPosY            int
	tcMaxX            int
	tcMaxY            int
	tcUsedY           int
	tcMaxLength       int

	// tab find
	modeTabFind   bool
	tfLine        []rune
	tfSuggestions []string
	modeAutoFind  bool // for when invoked via ^R or ^F outside of [tab]

	// vim
	modeViMode       viMode //= vimInsert
	viIteration      string
	viUndoHistory    []undoItem
	viUndoSkipAppend bool
	viYankBuffer     string

	// event
	evtKeyPress map[string]func(string, []rune, int) *EventReturn

	//ForceCrLf          bool
	EnableGetCursorPos bool
}

// NewInstance is used to create a readline instance and initialise it with sane
// defaults.
func NewInstance() *Instance {
	rl := new(Instance)

	rl.History = new(ExampleHistory)
	rl.HistoryAutoWrite = true
	rl.MaxTabCompleterRows = 4
	rl.prompt = ">>> "
	rl.promptLen = 4
	rl.HintFormatting = seqFgBlue
	rl.evtKeyPress = make(map[string]func(string, []rune, int) *EventReturn)

	rl.TempDirectory = os.TempDir()

	rl.MaxCacheSize = 256
	rl.cacheHint.Init(rl)
	rl.cacheSyntax.Init(rl)

	//rl.ForceCrLf = true

	return rl
}
