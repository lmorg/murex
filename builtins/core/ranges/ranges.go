package ranges

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("range", cmdRange, types.ReadArray, types.WriteArray)

	defaults.AppendProfile(`
		autocomplete set range { [{
			"Dynamic": ({ range -h }),
			"AllowMultiple": true,
			"AllowAny": true
		}] }
	`)
}

var rxSplitRange = regexp.MustCompile(`^(.*?)\.\.(.*?)$`)

const (
	f_e = "--exclude"
	f_8 = "--crop-backspace"
	f_b = "--no-empty-lines"
	f_t = "--trim-space"
	f_r = "--regexp"
	f_s = "--string"
	f_n = "--zero-index"
	f_i = "--index"
	f_h = "--help"
)

var args = &parameters.Arguments{
	Flags: map[string]string{
		f_h:  types.Boolean,
		"-h": f_h,
		f_e:  types.Boolean,
		"-e": f_e,
		f_8:  types.Boolean,
		"-8": f_8,
		f_b:  types.Boolean,
		"-b": f_b,
		f_t:  types.Boolean,
		"-t": f_t,
		f_r:  types.Boolean,
		"-r": f_r,
		f_s:  types.Boolean,
		"-s": f_s,
		f_n:  types.Boolean,
		"-n": f_n,
		f_i:  types.Boolean,
		"-i": f_i,
	},
	AllowAdditional: true,
}

func cmdRange(p *lang.Process) (err error) {
	flags, additional, err := p.Parameters.ParseFlags(args)
	if err != nil {
		return err
	}

	if flags[f_h] == types.TrueString {
		p.Stdout.SetDataType(types.Json)
		aw, err := p.Stdout.WriteArray(types.Json)
		if err != nil {
			return err
		}
		for f := range args.Flags {
			err = aw.WriteString(f)
			if err != nil {
				return err
			}
		}
		return aw.Close()
	}

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if len(additional) != 1 {
		return fmt.Errorf("missing parameter. Expecting something like `start..end`")
	}

	split := rxSplitRange.FindStringSubmatch(additional[0])
	if len(split) != 3 {
		return fmt.Errorf("invalid syntax: could not determine range values: %v", split)
	}

	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	r := &rangeParameters{
		Start:      split[1],
		End:        split[2],
		Exclude:    flags[f_e] == types.TrueString,
		RmBS:       flags[f_8] == types.TrueString,
		StripBlank: flags[f_b] == types.TrueString,
		TrimSpace:  flags[f_t] == types.TrueString,
	}

	var n int

	if flags[f_r] == types.TrueString {
		err = newRegexp(r)
		n++
	}
	if flags[f_s] == types.TrueString {
		err = newString(r)
		n++
	}
	if flags[f_n] == types.TrueString {
		err = newNumber(r)
		n++
	}
	if flags[f_i] == types.TrueString || n == 0 {
		err = newIndex(r)
		n++
	}

	if n > 1 {
		return errors.New("multiple modes selected. Please chose either regexp, string or index")
	}

	if err != nil {
		return err
	}

	return readArray(p, r, dt)
}
