package ranges

import (
	"bytes"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/rmbs"
)

type rangeParameters struct {
	Exclude    bool
	RmBS       bool
	StripBlank bool
	TrimSpace  bool
	Buffer     bool
	Start      string
	End        string
	Match      rangeFuncs
}

type rangeFuncs interface {
	Start([]byte) bool
	End([]byte) bool
	SetLength(int)
}

func readArray(p *lang.Process, r *rangeParameters, dt string) error {
	var (
		nestedErr      error
		started, ended bool
		stdin          = p.Stdin
		length         int
	)

	if r.Start == "" {
		started = true
	}

	array, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	if r.Buffer {
		stdin, length, err = buffer(p, dt)
		if err != nil {
			return err
		}
		r.Match.SetLength(length)
	}

	//not := !p.IsNot

	err = stdin.ReadArray(p.Context, func(b []byte) {
		if ended {
			return
		}

		if r.RmBS {
			b = []byte(rmbs.Remove(string(b)))
		}

		if r.TrimSpace {
			b = bytes.TrimSpace(b)
		}

		if r.StripBlank && len(b) == 0 {
			return
		}

		if !started {
			if r.Match.Start(b) {
				started = true
				if r.Exclude {
					return
				}

			} else {
				return
			}
		}

		if r.End != "" && r.Match.End(b) {
			ended = true
			if r.Exclude {
				return
			}
		}

		nestedErr = array.Write(b)
		if nestedErr != nil {
			p.Done()
			return
		}
	})

	if nestedErr != nil {
		return nestedErr
	}

	if err != nil {
		return err
	}

	return array.Close()
}
