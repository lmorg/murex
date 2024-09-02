//go:build js || windows || plan9
// +build js windows plan9

package session

func UnixSetSid() {
	// not supported on this platform
}

func UnixIsSession() bool { return false }
