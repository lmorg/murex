package psuedotty

import (
	"context"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
)

func (p *PTY) Read(b []byte) (int, error) {
	return p.out.Read(b)
}

func (p *PTY) ReadLine(callback func([]byte)) error {
	return p.out.ReadLine(callback)
}

func (p *PTY) ReadArray(ctx context.Context, callback func([]byte)) error {
	return p.out.ReadArray(ctx, callback)
}

func (p *PTY) ReadArrayWithType(ctx context.Context, callback func(interface{}, string)) error {
	return p.out.ReadArrayWithType(ctx, callback)
}

func (p *PTY) ReadMap(conf *config.Config, callback func(*stdio.Map)) error {
	return p.out.ReadMap(conf, callback)
}

func (p *PTY) ReadAll() ([]byte, error) {
	b, err := p.out.ReadAll()
	_ = p.in.Close()
	p.out.Close()
	return b, err
}
