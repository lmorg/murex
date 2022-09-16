package readline

import (
	"context"
)

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

	rl.tabMutex.Lock()

	rl.delayedTabContext = DelayedTabContext{rl: rl}
	rl.delayedTabContext.Context, rl.delayedTabContext.cancel = context.WithCancel(context.Background())

	rl.tcPrefix, rl.tcSuggestions, rl.tcDescriptions, rl.tcDisplayType = rl.TabCompleter(rl.line, rl.pos, rl.delayedTabContext)

	if len(rl.tcDescriptions) == 0 {
		// probably not needed, but just in case someone doesn't initialize the
		// map in their API call.
		rl.tcDescriptions = make(map[string]string)
	}

	rl.tabMutex.Unlock()

	rl.initTabCompletion()
}

func (rl *Instance) initTabCompletion() {
	rl.modeTabCompletion = true
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

func (rl *Instance) writeTabCompletion(resetCursorPos bool) {
	if !rl.modeTabCompletion {
		return
	}

	_, posY := lineWrapPos(rl.promptLen, rl.pos, rl.termWidth)
	_, lineY := lineWrapPos(rl.promptLen, len(rl.line), rl.termWidth)
	moveCursorDown(rl.hintY + lineY - posY)
	print("\r\n" + seqClearScreenBelow)

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

	if resetCursorPos {
		moveCursorUp(rl.hintY + rl.tcUsedY)
		print("\r")
		rl.moveCursorFromStartToLinePos()
	}
}

func (rl *Instance) resetTabCompletion() {
	rl.modeTabCompletion = false
	rl.tcOffset = 0
	rl.tcUsedY = 0
	rl.modeTabFind = false
	rl.tfLine = []rune{}
}
