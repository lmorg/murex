//go:build ignore
// +build ignore

////go:build !windows && !plan9 && !js
//// +build !windows,!plan9,!js

package ipc

import (
	"fmt"
	"net"

	"golang.org/x/sys/unix"
)

// Read returns peer credentials using the SO_PEERCRED socket option on
// Unix domain sockets.
//
// On unsupported operating systems it returns an error.
func Read(conn *net.UnixConn) (*Cred, error) {
	raw, err := conn.SyscallConn()
	if err != nil {
		return nil, fmt.Errorf("unable to get raw socket connection: %w", err)
	}

	var (
		ucred      *unix.Ucred
		sockoptErr error
	)

	err = raw.Control(func(fd uintptr) {
		ucred, sockoptErr = unix.GetsockoptUcred(
			int(fd),
			unix.SOL_SOCKET,
			unix.SO_PEERCRED,
		)
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get raw socket: %w", err)
	}

	if sockoptErr != nil {
		return nil, fmt.Errorf("unable to get peer credential: %w", sockoptErr)
	}

	return &Cred{
		PID: ucred.Pid,
		UID: ucred.Uid,
		GID: ucred.Gid,
	}, nil
}
