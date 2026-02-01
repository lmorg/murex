//go:build js || windows || plan9 || tinygo
// +build js windows plan9 tinygo

package session

import (
	"os"
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

func UnixTTY() *os.File {
	return os.Stdin
}
