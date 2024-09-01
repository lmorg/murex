//go:build !js && !windows && !plan9
// +build !js,!windows,!plan9

package session

import (
	"fmt"
	"syscall"

	"github.com/lmorg/murex/debug"
)

func UnixSetSid() {
	var err error

	// Create a new session
	_, err = syscall.Setsid()
	if err != nil {
		debug.Log(fmt.Sprintf("!!! syscall.Setsid() failed: %s", err.Error()))
	}
}
