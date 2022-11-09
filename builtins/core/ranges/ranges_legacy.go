package ranges

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("@[", CmdRangeLegacy, types.ReadArray, types.WriteArray)
}

const usageLegacy = "\nUsage: [start..end] / [start..end]se\n(start or end can be omitted)"

// if additional ranges are added here, they will also need to be added to
// /home/lau/dev/go/src/github.com/lmorg/murex/lang/parameters.go
var RxSplitRangeLegacy = regexp.MustCompile(`^\s*(.*?)\s*\.\.\s*(.*?)\s*\]([bt8ernsi]*)\s*$`)

func CmdRangeLegacy(p *lang.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	s := p.Parameters.StringAll()

	split := RxSplitRangeLegacy.FindStringSubmatch(s)
	if len(split) != 4 {
		return fmt.Errorf("invalid syntax: could not separate component values: %v.%s", split, usageLegacy)
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
		return fmt.Errorf("invalid syntax: you cannot combine the following flags: %s.%s", split[3], usageLegacy)
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
