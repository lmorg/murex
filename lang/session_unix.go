//go:build !js && !windows && !plan9
// +build !js,!windows,!plan9

package lang

import "syscall"

//var UnixSessionID int

func UnixCreateSession() {
	/*var err error
	UnixSessionID, err = syscall.Setsid()
	if err != nil {
		panic(err)
		}*/

	// Create a new session
	_, _ = syscall.Setsid()
}
