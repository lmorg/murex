package readline

import (
	"context"
	"sync/atomic"
)

func delayedSyntaxTimer(rl *Instance, i int32) {
	if rl.PasswordMask != 0 || rl.DelayedSyntaxWorker == nil {
		return
	}

	if rl.cacheSyntax.Get(rl.line) != "" {
		return
	}

	if len(rl.line)+rl.promptLen > rl.termWidth {
		// line wraps, which is hard to do with random ANSI escape sequences
		// so better we don't bother trying.
		return
	}

	newLine := rl.DelayedSyntaxWorker(rl.line)
	var sLine string

	if rl.SyntaxHighlighter != nil {
		sLine = rl.SyntaxHighlighter(newLine)
	} else {
		sLine = string(newLine)
	}
	rl.cacheSyntax.Append(rl.line, sLine)

	if atomic.LoadInt32(&rl.delayedSyntaxCount) != i {
		return
	}

	rl.moveCursorToStart()
	print(sLine)
	rl.moveCursorFromEndToLinePos()
}

// DelayedTabContext is a custom context interface for async updates to the tab completions
type DelayedTabContext struct {
	rl      *Instance
	Context context.Context
	cancel  context.CancelFunc
}

// AppendSuggestions updates the tab completions with additional suggestions asynchronously
func (dtc *DelayedTabContext) AppendSuggestions(suggestions []string) {
	if dtc == nil || dtc.rl == nil {
		return
	}

	if !dtc.rl.modeTabCompletion {
		return
	}

	dtc.rl.tabMutex.Lock()

	for i := range suggestions {
		select {
		case <-dtc.Context.Done():
			dtc.rl.tabMutex.Unlock()
			return

		default:
			if dtc.rl.tcDescriptions == nil {
				dtc.rl.tcDescriptions = make(map[string]string)
			}
			dtc.rl.tcDescriptions[suggestions[i]] = dtc.rl.tcPrefix + suggestions[i]
			dtc.rl.tcSuggestions = append(dtc.rl.tcSuggestions, suggestions[i])
		}
	}

	dtc.rl.tabMutex.Unlock()
	dtc.rl.ForceHintTextUpdate(" ")
	dtc.rl.clearHelpers()
	dtc.rl.renderHelpers()
}

// AppendDescriptions updates the tab completions with additional suggestions + descriptions asynchronously
func (dtc *DelayedTabContext) AppendDescriptions(suggestions map[string]string) {
	if dtc.rl == nil {
		// This might legitimately happen with some tests
		return
	}

	if !dtc.rl.modeTabCompletion {
		return
	}

	dtc.rl.tabMutex.Lock()

	for k := range suggestions {
		select {
		case <-dtc.Context.Done():
			dtc.rl.tabMutex.Unlock()
			return

		default:
			dtc.rl.tcDescriptions[k] = suggestions[k]
			dtc.rl.tcSuggestions = append(dtc.rl.tcSuggestions, k)
		}
	}

	dtc.rl.tabMutex.Unlock()
	dtc.rl.ForceHintTextUpdate(" ")
	dtc.rl.clearHelpers()
	dtc.rl.renderHelpers()
}
