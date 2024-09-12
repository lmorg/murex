//go:build js || windows || plan9
// +build js windows plan9

package session

import (
	"runtime"

	"github.com/lmorg/murex/debug"
)

func UnixOpenTTY() {
	// not supported on this platform
}

func UnixIsSession() bool { return false }

func UnixCreateSession() {
	debug.Logf("!!! UnixCreateSession is not supported on %s", runtime.GOOS)
}
