package readline

import "sync/atomic"

func timer(rl *Instance, i int64) {
	if rl.PasswordMask != 0 || rl.DelayedWorker == nil {
		return
	}

	newLine := rl.DelayedWorker(rl.line)
	var sLine string

	count := atomic.LoadInt64(&rl.delayedCount)
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
