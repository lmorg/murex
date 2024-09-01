//go:build !js && !windows && !plan9
// +build !js,!windows,!plan9

package session

import "syscall"

//var UnixSessionID int

func UnixSetSid() {
	/*var err error
	UnixSessionID, err = syscall.Setsid()
	if err != nil {
		panic(err)
		}*/

	// Create a new session
	_, _ = syscall.Setsid()
}
