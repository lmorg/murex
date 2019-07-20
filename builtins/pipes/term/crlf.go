package term

import (
	"os"
	"sync/atomic"

	"github.com/lmorg/murex/utils"
)

// This function is just a way for readline to guarantee that it will always start on a new line.
// (it actually only works for murex builtins so i'll need to get cleverer with readline at some point)

type appendCrLf struct {
	r int32
}

func (lf *appendCrLf) set(b byte) {
	// Wrapping this up inside a goroutine may case instances where lf.r gets set out of sequence.
	// However the bug presented will - at worst - just mean the "\n" character might not get
	// written to the terminal when it should have. The risk of this is low and the but it presents
	// is pretty mild however the performance improvement is in the range of 2 to 5 seconds on jobs
	// which use the terminal output heavily for around 20 seconds. So the risk is, in my opinion,
	// worth the reward.
	//
	// Update: maybe not. Seeing far more crlf glitches than previous testing had suggested.
	//go func() {
	atomic.StoreInt32(&lf.r, int32(b))
	//}()
}

func (lf *appendCrLf) Write() {
	r := atomic.SwapInt32(&lf.r, int32('\n'))
	if r != '\n' {
		os.Stderr.Write(utils.NewLineByte)
	}
}

// CrLf function to append a line feed character at the end of text piped to the terminal to aid readability.
var CrLf = appendCrLf{r: '\n'}
