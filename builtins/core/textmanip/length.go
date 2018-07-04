package textmanip

import (
	"errors"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	proc.GoFunctions["left"] = cmdLeft
	proc.GoFunctions["right"] = cmdRight
	proc.GoFunctions["prefix"] = cmdPrefix
	proc.GoFunctions["suffix"] = cmdSuffix
}

func cmdLeft(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	left, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	if left == 0 {
		return errors.New("Cannot have zero characters.")
	}

	var output []string

	if left > 0 {
		p.Stdin.ReadArray(func(b []byte) {
			if len(b) < left {
				output = append(output, string(b))
			} else {
				output = append(output, string(b[:left]))
			}
		})
	} else {
		left = left * -1
		p.Stdin.ReadArray(func(b []byte) {
			if len(b) < left {
				output = append(output, string(b))
			} else {
				output = append(output, string(b[:len(b)-left]))
			}
		})
	}

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdRight(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	right, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	if right == 0 {
		return errors.New("Cannot have zero characters.")
	}

	var output []string

	if right > 0 {
		p.Stdin.ReadArray(func(b []byte) {
			if len(b) < right {
				output = append(output, string(b))
			} else {
				output = append(output, string(b[len(b)-right:]))
			}
		})
	} else {
		right = right * -1
		p.Stdin.ReadArray(func(b []byte) {
			if len(b) < right {
				output = append(output, string(b))
			} else {
				output = append(output, string(b[right:]))
			}
		})
	}

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdPrefix(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	prepend := p.Parameters.StringAll()

	var output []string
	p.Stdin.ReadArray(func(b []byte) {
		output = append(output, prepend+string(b))
	})

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdSuffix(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	s := p.Parameters.StringAll()

	var output []string
	p.Stdin.ReadArray(func(b []byte) {
		output = append(output, string(b)+s)
	})

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
