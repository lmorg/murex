package ranges

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"

	"fmt"
	"regexp"
	"strings"
)

func init() {
	proc.GoFunctions["@["] = cmdRange
}

var rxSplitRange *regexp.Regexp = regexp.MustCompile(`^\s*(.*?)\s*\.\.\s*(.*?)\s*\]([erns]*)\s*$`)

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

func cmdRange(p *proc.Process) (err error) {
	const usage = "\nUsage: @[start..end] /  @[start..end]e\n(start or end can be omitted)"

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
		return fmt.Errorf("Invalid syntax: you cannot combile the following flags: %s.%s", split[3], usage)
	}

	debug.Json("split", split)

	var array []string

	switch split[3] {
	case "r":
		err = newRegexp(r)

	case "s":
		fallthrough

	default:
		err = newString(r)
	}

	if err != nil {
		return err
	}

	array, err = readArray(p, r)
	if err != nil {
		return err
	}

	b, err := define.MarshalData(p, dt, array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func readArray(p *proc.Process, r *rangeParameters) ([]string, error) {
	var (
		array   []string
		err     error
		started bool
		ended   bool
	)

	if r.Start == "" {
		started = true
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

		array = append(array, string(b))
	})

	return array, err
}
