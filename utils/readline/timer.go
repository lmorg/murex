package readline

import (
	"context"
	"sync/atomic"
)

func delayedSyntaxTimer(rl *Instance, i int64) {
	if rl.PasswordMask != 0 || rl.DelayedSyntaxWorker == nil {
		return
	}

	if len(rl.line)+rl.promptLen > rl.termWidth {
		// line wraps, which is hard to do with random ANSI escape sequences
		// so better we don't bother trying.
		return
	}

	newLine := rl.DelayedSyntaxWorker(rl.line)
	var sLine string

	count := atomic.LoadInt64(&rl.delayedSyntaxCount)
	if count != i {
		return
	}

	if rl.SyntaxHighlighter != nil {
		sLine = rl.SyntaxHighlighter(newLine)
	} else {
		sLine = string(newLine)
	}

	return
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
func (dtc DelayedTabContext) AppendSuggestions(suggestions []string) {
	dtc.rl.mutex.Lock()
	defer dtc.rl.mutex.Unlock()

	for i := range suggestions {
		select {
		case <-dtc.Context.Done():
			return

		default:
			dtc.rl.tcDescriptions[suggestions[i]] = dtc.rl.tcPrefix + suggestions[i]
			dtc.rl.tcSuggestions = append(dtc.rl.tcSuggestions, suggestions[i])
		}
	}

	dtc.rl.clearHelpers()
	dtc.rl.renderHelpers()
}

// AppendDescriptions updates the tab completions with additional suggestions + descriptions asynchronously
func (dtc DelayedTabContext) AppendDescriptions(suggestions map[string]string) {
	dtc.rl.mutex.Lock()
	defer dtc.rl.mutex.Unlock()

	for k := range suggestions {
		select {
		case <-dtc.Context.Done():
			return

		default:
			dtc.rl.tcDescriptions[k] = suggestions[k]
			dtc.rl.tcSuggestions = append(dtc.rl.tcSuggestions, k)
		}
	}

	dtc.rl.clearHelpers()
	dtc.rl.renderHelpers()
}
