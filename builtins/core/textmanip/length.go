package textmanip

import (
	"github.com/lmorg/murex/lang"
)

func init() {
	lang.GoFunctions["left"] = cmdLeft
	lang.GoFunctions["right"] = cmdRight
	lang.GoFunctions["prefix"] = cmdPrefix
	lang.GoFunctions["suffix"] = cmdSuffix
}

func cmdLeft(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	left, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	switch {
	case left > 0:
		p.Stdin.ReadArray(func(b []byte) {
			if len(b) < left {
				err = aw.Write(b)
			} else {
				err = aw.Write(b[:left])
			}

			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}
		})

	case left < 0:
		left = left * -1
		p.Stdin.ReadArray(func(b []byte) {
			if len(b) < left {
				err = aw.WriteString("")
			} else {
				err = aw.Write(b[:len(b)-left])
			}

			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}
		})

	default:
		p.Stdin.ReadArray(func([]byte) {
			err = aw.WriteString("")
			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}
		})
	}

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}

func cmdRight(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	right, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	switch {
	case right > 0:
		p.Stdin.ReadArray(func(b []byte) {
			if len(b) < right {
				err = aw.Write(b)
			} else {
				err = aw.Write(b[len(b)-right:])
			}

			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}
		})

	case right < 0:
		right = right * -1
		p.Stdin.ReadArray(func(b []byte) {
			if len(b) < right {
				err = aw.WriteString("")
			} else {
				err = aw.Write(b[right:])
			}

			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}
		})

	default:
		p.Stdin.ReadArray(func([]byte) {
			err = aw.WriteString("")

			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}
		})
	}

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}

func cmdPrefix(p *lang.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	prepend := p.Parameters.ByteAll()

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(func(b []byte) {
		err = aw.Write(append(prepend, b...))
		if err != nil {
			p.Stdin.ForceClose()
			p.Done()
		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}

func cmdSuffix(p *lang.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	suffix := p.Parameters.ByteAll()

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(func(b []byte) {
		err = aw.Write(append(b, suffix...))
		if err != nil {
			p.Stdin.ForceClose()
			p.Done()
		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}
