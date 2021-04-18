package net

import (
	"context"
	"net"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	stdio.RegisterPipe("udp-dial", newDialerUDP)
	stdio.RegisterPipe("tcp-dial", newDialerTCP)
	stdio.RegisterPipe("udp-listen", newListenerUDP)
	stdio.RegisterPipe("tcp-listen", newListenerTCP)
}

func newDialerUDP(arguments string) (stdio.Io, error) {
	return NewDialer("udp", arguments)
}

func newDialerTCP(arguments string) (stdio.Io, error) {
	return NewDialer("tcp", arguments)
}

func newListenerUDP(arguments string) (stdio.Io, error) {
	return NewListener("udp", arguments)
}

func newListenerTCP(arguments string) (stdio.Io, error) {
	return NewListener("tcp", arguments)
}

// NewDialer creates a new net.Dial-based stream.Io pipe
func NewDialer(protocol, address string) (n *Net, err error) {
	n = new(Net)
	n.protocol = protocol

	if protocol == "udp" || protocol == "tcp" {
		n.dataType = types.Generic
	} else {
		protocol = "tcp"
	}

	n.conn, err = net.Dial(protocol, address)
	if err != nil {
		return nil, err
	}

	n.ctx, n.forceClose = context.WithCancel(context.Background())

	return
}

// NewListener creates a new net.Listen-based stream.Io pipe
func NewListener(protocol, address string) (n *Net, err error) {
	n = new(Net)
	n.protocol = protocol

	if protocol == "udp" || protocol == "tcp" {
		n.dataType = types.Generic
	} else {
		protocol = "tcp"
	}

	listen, err := net.Listen(protocol, address)
	if err != nil {
		return nil, err
	}
	defer listen.Close()

	n.conn, err = listen.Accept()
	if err != nil {
		return nil, err
	}

	n.ctx, n.forceClose = context.WithCancel(context.Background())

	return
}
