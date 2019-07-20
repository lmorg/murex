package ranges

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
)

func init() {
	lang.GoFunctions["@["] = cmdRange
}

const usage = "\nUsage: @[start..end] / @[start..end]se\n(start or end can be omitted)"

var rxSplitRange = regexp.MustCompile(`^\s*(.*?)\s*\.\.\s*(.*?)\s*\]([erns]*)\s*$`)

type rangeParameters struct {
	Exclude bool
	Start   string
	End     string
	Match   rangeFuncs
}

type rangeFuncs interface {
	Start([]byte) bool
	End([]byte) bool
}

func cmdRange(p *lang.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	s := p.Parameters.StringAll()

	split := rxSplitRange.FindStringSubmatch(s)
	if len(split) != 4 {
		return fmt.Errorf("Invalid syntax: could not separate component values: %v.%s", split, usage)
	}

	r := &rangeParameters{
		Start: split[1],
		End:   split[2],
	}

	if strings.Contains(split[3], "e") {
		r.Exclude = true
		split[3] = strings.Replace(split[3], "e", "", -1)
	}

	if len(split[3]) > 1 {
		return fmt.Errorf("Invalid syntax: you cannot combine the following flags: %s.%s", split[3], usage)
	}

	switch split[3] {
	case "r":
		err = newRegexp(r)

	case "s":
		err = newString(r)

	case "n":
		fallthrough

	default:
		err = newNumber(r)
	}

	if err != nil {
		return err
	}

	return readArray(p, r, dt)
}

func readArray(p *lang.Process, r *rangeParameters, dt string) error {
	var (
		nestedErr      error
		started, ended bool
	)

	if r.Start == "" {
		started = true
	}

	array, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	err = p.Stdin.ReadArray(func(b []byte) {
		if ended {
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
			return
		}
	})

	if nestedErr != nil {
		return nestedErr
	}

	return array.Close()
}
