package ranges

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("@[", deprecatedRange, types.ReadArray, types.WriteArray)
}

const usage = "\nUsage: [start..end] / [start..end]se\n(start or end can be omitted)"

// if additional ranges are added here, they will also need to be added to
// /home/lau/dev/go/src/github.com/lmorg/murex/lang/parameters.go
var RxSplitRange = regexp.MustCompile(`^\s*(.*?)\s*\.\.\s*(.*?)\s*\]([bt8ernsiu]*)\s*$`)

func deprecatedRange(p *lang.Process) error {
	lang.FeatureDeprecatedBuiltin(p)
	return CmdRange(p)
}

func CmdRange(p *lang.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	s := p.Parameters.StringAll()

	split := RxSplitRange.FindStringSubmatch(s)
	if len(split) != 4 {
		err = indexAndExpand(p, dt)
		if err != nil {
			return fmt.Errorf("not a valid range: %v.%s\nnor a valid index: %v", split, usage, err)
		}
		return nil
	}

	r := &rangeParameters{
		Start: split[1],
		End:   split[2],
	}

	if strings.Contains(split[3], "e") {
		r.Exclude = true
		split[3] = strings.Replace(split[3], "e", "", -1)
	}

	if strings.Contains(split[3], "8") {
		r.RmBS = true
		split[3] = strings.Replace(split[3], "8", "", -1)
	}

	if strings.Contains(split[3], "b") {
		r.StripBlank = true
		split[3] = strings.Replace(split[3], "b", "", -1)
	}

	if strings.Contains(split[3], "t") {
		r.TrimSpace = true
		split[3] = strings.Replace(split[3], "t", "", -1)
	}

	if len(split[3]) > 1 {
		return fmt.Errorf("invalid syntax: you cannot combine the following flags: %s.%s", split[3], usage)
	}

	switch split[3] {
	case "r":
		err = newRegexp(r)

	case "s":
		err = newString(r)

	case "n":
		err = newNumber(r)

	case "i":
		err = newIndex(r)

	default:
		if p.Name.String() == "@[" {
			err = newNumber(r)
		} else {
			err = newIndex(r)
		}
	}

	if err != nil {
		return err
	}

	return readArray(p, r, dt)
}

func indexAndExpand(p *lang.Process, dt string) (err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic caught, please report this to https://github.com/lmorg/murex/issues : %s", r)
			}
		}()
	}

	// We will set data type from the index function but fallback to this just
	// in case it's forgotten about in the index function. This is safe because
	// SetDataType() cannot overwrite the data type once set.
	defer p.Stdout.SetDataType(dt)

	params := p.Parameters.StringArray()
	l := len(params) - 1
	if l < 0 {
		return errors.New("missing parameters. Please select 1 or more indexes")
	}

	switch {
	case params[l] == "]":
		params = params[:l]
	case strings.HasSuffix(params[l], "]"):
		params[l] = params[l][:len(params[l])-1]
	default:
		return errors.New("missing closing bracket, ` ]`")
	}

	f := lang.ReadIndexes[dt]
	if f == nil {
		return errors.New("i don't know how to get an index from this data type: `" + dt + "`")
	}

	silent, err := p.Config.Get("index", "silent", types.Boolean)
	if err != nil {
		silent = false
	}

	err = f(p, params)
	if silent.(bool) {
		return nil
	}

	return err
}
