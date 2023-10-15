package streams

import (
	"context"
)

type FakeTTY struct {
	Stdin
}

func NewFakeTTY() *FakeTTY {
	ft := new(FakeTTY)
	ft.max = DefaultMaxBufferSize
	//stdin.buffer = make([]byte, 0, 1024*1024)
	ft.ctx, ft.forceClose = context.WithCancel(context.Background())
	return ft
}

func (ft *FakeTTY) IsTTY() bool {
	return true
}
