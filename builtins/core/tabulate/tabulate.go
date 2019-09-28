package tabulate

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["tabulate"] = cmdTabulate

	defaults.AppendProfile(`
		autocomplete set tabulate { [{
			"DynamicDesc": ({ tabulate --help }),
			"AllowMultiple": true
		}] }
	`)
}

const (
	constSeparator = `(\t|\s[\s]+)+`
)

var (
	rxWhitespaceLeft  = regexp.MustCompile(`^[\t\s]+`)
	rxWhitespaceRight = regexp.MustCompile(`[\t\s]+$`)
)

// flags

const (
	//fNoTrim     = "--no-trim"
	fSeparator  = "--separator"
	fSplitComma = "--split-comma"
	fMap        = "--map"
	fJoiner     = "--joiner"
	fHelp       = "--help"
)

var flags = map[string]string{
	//fNoTrim:     types.Boolean,
	fSeparator:  types.String,
	fSplitComma: types.Boolean,
	fMap:        types.Boolean,
	fJoiner:     types.String,
	fHelp:       types.Boolean,
}

var desc = map[string]string{
	//fNoTrim:     "Disable ",
	fSeparator:  "String, custom regex pattern for spliting fields (default: `" + constSeparator + "`)",
	fSplitComma: "Boolean, split first field and duplicate the line if comma found in first field (eg parsing flags in help pages)",
	fMap:        "Boolean, return JSON map instead of table",
	fJoiner:     "String, used with --map to concatenate any trailing records in a given field",
	fHelp:       "Boolean, displays this help message",
}

func cmdTabulate(p *lang.Process) error {
	f, _, err := p.Parameters.ParseFlags(
		&parameters.Arguments{
			Flags:           flags,
			AllowAdditional: false,
		},
	)

	if err != nil {
		return err
	}

	var (
		//trim       = true
		separator  = constSeparator
		splitComma = false
		joiner     = " "
		w          writer

		firstRec bool
		offByOne bool
	)

	for flag, value := range f {
		switch flag {
		//case fNoTrim:
		//	trim = false
		case fSeparator:
			separator = value
		case fSplitComma:
			splitComma = true
		case fJoiner:
			joiner = value
		case fMap:
			// check this afterwards just in case fJoiner
			// hasn't yet been processed (which can change
			// the bahavior of fMap)
		case fHelp:
			return help(p)
		}
	}

	if err := p.ErrIfNotAMethod(); err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	dt := p.Stdin.GetDataType()
	if dt != types.Generic && dt != types.String {
		p.Stdout.SetDataType(types.Null)
		return fmt.Errorf("`%s` is designed to only take string (%s) or generic (%s) data-types from STDIN. Instead it received '%s'",
			p.Name, types.String, types.Generic, dt)
	}

	if f[fMap] == "" {
		p.Stdout.SetDataType("csv")
		w = csv.NewWriter(p.Stdout)
	} else {
		p.Stdout.SetDataType(types.Json)
		w = newMapWriter(p.Stdout, joiner)
	}

	rxTableSplit, err := regexp.Compile(separator)
	if err != nil {
		return err
	}

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		/*if trim {
			s = strings.TrimSpace(s)
			//s = rxWhitespaceLeft.ReplaceAllString(s, "")
			//s = rxWhitespaceRight.ReplaceAllString(s, "")
		}*/

		if !rxTableSplit.MatchString(s) {
			continue
		}

		split := rxTableSplit.Split(scanner.Text(), -1)
		if len(split) == 0 {
			continue
		}

		firstRec = true
		if firstRec && split[0] == "" {
			offByOne = true
		}

		if offByOne && len(split) > 1 && split[0] == "" {
			split = split[1:]
		}

		if splitComma && len(split) > 1 {
			comma := strings.Split(split[0], ",")
			if len(comma) == 1 {
				goto noSplit
			}

			for i := range comma {
				flag := strings.TrimSpace(comma[i])
				new := append([]string{flag}, split[1:]...)
				err := w.Write(new)
				if err != nil {
					return err
				}
			}
			continue
		}

	noSplit:
		err := w.Write(split)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	w.Flush()
	return w.Error()
}

func help(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)
	b, err := json.Marshal(desc, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
