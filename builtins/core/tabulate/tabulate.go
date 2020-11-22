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
	rxSplitComma      = regexp.MustCompile(`[\s\t]*,[\s\t]*`)
	rxSplitSpace      = regexp.MustCompile(`[\s\t]+-`)
)

// flags

const (
	fSeparator   = "--separator"
	fSplitComma  = "--split-comma"
	fSplitSpace  = "--split-space"
	fKeyIncHint  = "--key-inc-hint"
	fKeyVal      = "--key-value"
	fMap         = "--map"
	fJoiner      = "--joiner"
	fColumnWraps = "--column-wraps"
	fHelp        = "--help"
)

var flags = map[string]string{
	fSeparator:   types.String,
	fSplitComma:  types.Boolean,
	fSplitSpace:  types.Boolean,
	fKeyIncHint:  types.Boolean,
	fKeyVal:      types.Boolean,
	fMap:         types.Boolean,
	fJoiner:      types.String,
	fColumnWraps: types.Boolean,
	fHelp:        types.Boolean,
}

var desc = map[string]string{
	fSeparator:   "String, custom regex pattern for spliting fields (default: `" + constSeparator + "`)",
	fSplitComma:  "Boolean, split first field and duplicate the line if comma found in first field (eg parsing flags in help pages)",
	fSplitSpace:  "Boolean, split first field and duplicate the line if white space found in first field (eg parsing flags in help pages)",
	fKeyIncHint:  "Boolean, used with " + fMap + " to split any space or equal delimited hints/examples (eg parsing flags)",
	fKeyVal:      "Boolean, discard any records that don't appear key value pairs (auto-enabled when " + fMap + " used)",
	fMap:         "Boolean, return JSON map instead of table",
	fJoiner:      "String, used with " + fMap + " to concatenate any trailing records in a given field",
	fColumnWraps: "Boolean, used with " + fMap + " or " + fKeyVal + " to merge trailing lines if the text wraps within the same column",
	fHelp:        "Boolean, displays this help message",
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
		separator   = constSeparator
		splitComma  bool
		splitSpace  bool
		keyIncHint  bool
		keyVal      = false
		joiner      = " "
		columnWraps = false
		colWrapsBuf string // buffer for wrapped columns
		keys        []string
		w           writer
		last        string
		split       []string
	)

	for flag, value := range f {
		switch flag {
		case fSeparator:
			separator = value
		case fSplitComma:
			splitComma = true
		case fSplitSpace:
			splitSpace = true
		case fKeyIncHint:
			keyIncHint = true
		case fKeyVal:
			keyVal = true
		case fJoiner:
			joiner = value
		case fMap:
			keyVal = true
		case fColumnWraps:
			columnWraps = true
		case fHelp:
			return help(p)
		}
	}

	if splitSpace && splitComma {
		return fmt.Errorf("Cannot have %s and %s both enabled. Please pick one or the other", fSplitComma, fSplitSpace)
	}

	if !keyVal && keyIncHint {
		return fmt.Errorf("Cannot use %s without %s or %s being set", fKeyIncHint, fKeyVal, fMap)
	}

	if !keyVal && columnWraps {
		return fmt.Errorf("Cannot use %s without %s or %s being set", fColumnWraps, fKeyVal, fMap)
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

		// not a table row
		if !rxTableSplit.MatchString(s) {
			continue
		}

		// still not a table row
		split = rxTableSplit.Split(scanner.Text(), -1)
		if len(split) == 0 {
			continue
		}

		// table has indentation, lets remove that
		if len(split) > 1 && split[0] == "" {
			split = split[1:]
		}

		// looks like there's a new key, so lets write the colWrapsBuf
		if columnWraps && len(split) > 1 && last != "" {
			if len(keys) == 0 {

				err = w.Write([]string{last, colWrapsBuf})
				if err != nil {
					return err
				}

			} else {

				for i := range keys {
					err = w.Write([]string{keys[i], colWrapsBuf})
					if err != nil {
						return err
					}
				}
			}

			colWrapsBuf = ""
		}

		// split keys by comma
		if splitComma && len(split) > 1 {
			keys = rxSplitComma.Split(split[0], -1)
		}

		// split keys by space
		if splitSpace && len(split) > 1 {
			keys = rxSplitSpace.Split(split[0], -1)
			for i := range keys {
				keys[i] = "-" + keys[i]
			}
		}

		// remove the hint stuff
		if keyIncHint {
			if len(keys) != 0 {
				for i := range keys {
					keys[i] = strings.SplitN(keys[i], " ", 2)[0]
					if strings.Contains(keys[i], "=") {
						keys[i] = strings.SplitN(keys[i], "=", 2)[0] + "="
					}
				}
			} else {
				split[0] = strings.SplitN(split[0], " ", 2)[0]
				if strings.Contains(split[0], "=") {
					split[0] = strings.SplitN(split[0], "=", 2)[0] + "="
				}
			}
		}

		if keyVal {
			if len(split) < 2 || split[0] == "" {
				// is this a wrapped column? If not, it's clearly a field
				// heading so we should skip over it
				if columnWraps && last != "" {
					colWrapsBuf += joiner + strings.Join(split, joiner)
				}

				// else silently ignore heading
				continue
			}
			last = split[0]
		}

		// only write if columns not wrapped, otherwise loop round to check for
		// any wrapped columns
		if !columnWraps {
			if len(keys) == 0 {

				err = w.Write(split)
				if err != nil {
					return err
				}

			} else {

				for i := range keys {
					split[0] = keys[i]
					err = w.Write(split)
					if err != nil {
						return err
					}
				}
			}

			keys = nil

		} else {
			colWrapsBuf = strings.Join(split[1:], joiner)
		}
	}

	// clean up any trailing wrapped columns
	if columnWraps && len(colWrapsBuf) != 0 {
		if len(keys) == 0 {

			err = w.Write([]string{last, colWrapsBuf})
			if err != nil {
				return err
			}

		} else {

			for i := range keys {
				err = w.Write([]string{keys[i], colWrapsBuf})
				if err != nil {
					return err
				}
			}
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
