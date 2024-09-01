//go:build js || windows || plan9
// +build js windows plan9

package lang

func UnixCreateSession() {
	// not supported on this platform
}
