//go:build windows && plan9 && js
// +build windows,plan9,js

package ipc

import (
	"errors"
	"net"
)

var errUnsupported = errors.New("peercred is unsupported on this operating system")

// Read returns peer credentials using the SO_PEERCRED socket option on
// Unix domain sockets.
//
// On unsupported operating systems it returns an error.
func Read(conn *net.UnixConn) (*Cred, error) {
	return nil, errUnsupported
}
