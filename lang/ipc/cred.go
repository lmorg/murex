//go:build ignore
// +build ignore

// Package peercred is an easy to use wrapper around the Linux
// SO_PEERCRED socket option on Unix domain sockets.  It degrades
// gracefully with an unsupported error on other operating systems.
package ipc

// Cred is the structure containing peer credentials.  Linux passes the
// process ID (pid), user ID (uid), and group ID (gid) from the other
// side of the connection.
//
// This structure mimics the ucred structure from the <linux/socket.h>
// file.
//
// These values are populated at socket creation time, and will not
// reflect privileges dropped after the creation of the socket.
type Cred struct {
	PID int32
	UID uint32
	GID uint32
}
