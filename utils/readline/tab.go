package readline

import (
	"context"
	"os"
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

	rl.delayedTabContext = DelayedTabContext{rl: rl}
	rl.delayedTabContext.Context, rl.delayedTabContext.cancel = context.WithCancel(context.Background())

	rl.tcr = rl.TabCompleter(rl.line.Runes(), rl.line.RunePos(), rl.delayedTabContext)
	if rl.tcr == nil {
		return
	}

	rl.tabMutex.Lock()
	rl.tcPrefix, rl.tcSuggestions, rl.tcDescriptions, rl.tcDisplayType = rl.tcr.Prefix, rl.tcr.Suggestions, rl.tcr.Descriptions, rl.tcr.DisplayType
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
	rl.autocompleteHeightAdjust()

	if rl.tcDisplayType == TabDisplayGrid {
		rl.initTabGrid()
	} else {
		rl.initTabMap()
	}
}

func (rl *Instance) autocompleteHeightAdjust() {
	_, height, err := GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 25
	}

	switch {
	case height <= 4:
		rl.MaxTabCompleterRows = 1
	case height-4 <= rl.MaxTabCompleterRows:
		rl.MaxTabCompleterRows = height - 4
	}

}

func (rl *Instance) moveTabCompletionHighlight(x, y int) {
	if rl.tcDisplayType == TabDisplayGrid {
		rl.moveTabGridHighlight(x, y)
	} else {
		rl.moveTabMapHighlight(x, y)
	}
}

func (rl *Instance) writeTabCompletionStr() string {
	if !rl.modeTabCompletion {
		return ""
	}

	posX, posY := rl.lineWrapCellPos()
	_, lineY := rl.lineWrapCellLen()
	output := moveCursorDownStr(rl.hintY + lineY - posY)
	output += "\r\n" + seqClearScreenBelow

	switch rl.tcDisplayType {
	case TabDisplayGrid:
		output += rl.writeTabGridStr()

	case TabDisplayMap:
		output += rl.writeTabMapStr()

	case TabDisplayList:
		output += rl.writeTabMapStr()

	default:
		output += rl.writeTabGridStr()
	}

	output += moveCursorUpStr(rl.hintY + rl.tcUsedY + lineY - posY)
	output += "\r" + moveCursorForwardsStr(posX)

	return output
}

func (rl *Instance) resetTabCompletion() {
	rl.modeTabCompletion = false
	rl.tcOffset = 0
	rl.tcUsedY = 0
	rl.modeTabFind = false
	rl.modeAutoFind = false
	rl.tfLine = []rune{}
}
