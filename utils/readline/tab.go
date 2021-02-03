package readline

import "context"

// TabDisplayType defines how the autocomplete suggestions display
type TabDisplayType int

const (
	// TabDisplayGrid is the default. It's where the screen below the prompt is
	// divided into a grid with each suggestion occupying an individual cell.
	TabDisplayGrid = iota

	// TabDisplayList is where suggestions are displayed as a list with a
	// description. The suggestion gets highlighted but both are searchable (ctrl+f)
	TabDisplayList

	// TabDisplayMap is where suggestions are displayed as a list with a
	// description however the description is what gets highlighted and only
	// that is searchable (ctrl+f). The benefit of TabDisplayMap is when your
	// autocomplete suggestions are IDs rather than human terms.
	TabDisplayMap
)

func (rl *Instance) getTabCompletion() {
	rl.tcOffset = 0
	if rl.TabCompleter == nil {
		return
	}

	if rl.delayedTabContext.cancel != nil {
		rl.delayedTabContext.cancel()
	}

	rl.delayedTabContext = DelayedTabContext{rl: rl}
	rl.delayedTabContext.Context, rl.delayedTabContext.cancel = context.WithCancel(context.Background())

	rl.tcPrefix, rl.tcSuggestions, rl.tcDescriptions, rl.tcDisplayType = rl.TabCompleter(rl.line, rl.pos, rl.delayedTabContext)
	/*if len(rl.tcSuggestions) == 0 && delayed {
		return
	}*/
	//panic(rl.tcDisplayType)

	if len(rl.tcDescriptions) == 0 {
		// probably not needed, but just in case someone doesn't initialise the
		// map in their API call.
		rl.tcDescriptions = make(map[string]string)
	}

	/*if len(rl.tcSuggestions) == 1 && !rl.modeTabCompletion {
		if len(rl.tcSuggestions[0]) == 0 || rl.tcSuggestions[0] == " " || rl.tcSuggestions[0] == "\t" {
			return
		}
		rl.insert([]byte(rl.tcSuggestions[0]))
		return
	}*/

	rl.initTabCompletion()
}

func (rl *Instance) initTabCompletion() {
	if rl.tcDisplayType == TabDisplayGrid {
		rl.initTabGrid()
	} else {
		rl.initTabMap()
	}
}

func (rl *Instance) moveTabCompletionHighlight(x, y int) {
	if rl.tcDisplayType == TabDisplayGrid {
		rl.moveTabGridHighlight(x, y)
	} else {
		rl.moveTabMapHighlight(x, y)
	}
}

func (rl *Instance) writeTabCompletion() {
	if !rl.modeTabCompletion {
		return
	}

	switch rl.tcDisplayType {
	case TabDisplayGrid:
		rl.writeTabGrid()

	case TabDisplayMap:
		rl.writeTabMap()

	case TabDisplayList:
		rl.writeTabMap()

	default:
		rl.writeTabGrid()
	}
}

func (rl *Instance) resetTabCompletion() {
	rl.modeTabCompletion = false
	rl.tcOffset = 0
	rl.tcUsedY = 0
	rl.modeTabFind = false
	rl.tfLine = []rune{}
}
