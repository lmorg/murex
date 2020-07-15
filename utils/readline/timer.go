package readline

import (
	"context"
	"sync/atomic"
)

func delayedSyntaxTimer(rl *Instance, i int64) {
	if rl.PasswordMask != 0 || rl.DelayedSyntaxWorker == nil {
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

	moveCursorBackwards(rl.pos)
	print(sLine)
	moveCursorBackwards(len(rl.line) - rl.pos)
}

type DelayedTabContext struct {
	rl      *Instance
	Context context.Context
	cancel  context.CancelFunc
}

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

	//dtc.rl.updateTabFind([]rune{})
	//dtc.rl.writeTabCompletion()
	dtc.rl.clearHelpers()
	dtc.rl.renderHelpers()
}
