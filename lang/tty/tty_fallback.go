//go:build windows || plan9 || js
// +build windows plan9 js

package tty

import (
	"errors"
	"runtime"
)

const errMessage = "This isn't supported on " + runtime.GOOS

func Enabled() bool {
	return false
}

func CreatePTY() error {
	return errors.New(errMessage)
}

func DestroyPty() {
	// not supported on this platform
}

func BufferRecall(_ []byte, _ string) {
	// not supported on this platform
}

func ConfigRead() (interface{}, error) {
	return false, nil
}

func ConfigWrite(_ interface{}) error {
	return errors.New(errMessage)
}

func MissingCrLf() bool {
	return false
}

func WriteCrLf() {
	// empty function for cross compiling compatibility
}
