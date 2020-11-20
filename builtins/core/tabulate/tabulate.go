package tabulate

import (
	"bufio"
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
	fSeparator   = "--separator"
	fSplitComma  = "--split-comma"
	fKeyVal      = "--key-value"
	fMap         = "--map"
	fJoiner      = "--joiner"
	fColumnWraps = "--column-wraps"
	fHelp        = "--help"
)

var flags = map[string]string{
	//fNoTrim:     types.Boolean,
	fSeparator:   types.String,
	fSplitComma:  types.Boolean,
	fKeyVal:      types.Boolean,
	fMap:         types.Boolean,
	fJoiner:      types.String,
	fColumnWraps: types.Boolean,
	fHelp:        types.Boolean,
}

var desc = map[string]string{
	fSeparator:   "String, custom regex pattern for spliting fields (default: `" + constSeparator + "`)",
	fSplitComma:  "Boolean, split first field and duplicate the line if comma found in first field (eg parsing flags in help pages)",
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
		splitComma  = false
		keyVal      = false
		joiner      = " "
		columnWraps = false
		w           writer

		offByOne bool
		last     string
	)

	for flag, value := range f {
		switch flag {
		case fSeparator:
			separator = value
		case fSplitComma:
			splitComma = true
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
		w = newCsvWriter(p.Stdout, joiner, columnWraps)
	} else {
		p.Stdout.SetDataType(types.Json)
		w = newMapWriter(p.Stdout, joiner, columnWraps)
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

		if !rxTableSplit.MatchString(s) {
			continue
		}

		split := rxTableSplit.Split(scanner.Text(), -1)
		if len(split) == 0 {
			continue
		}

		if split[0] == "" {
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
		if keyVal {
			if len(split) < 2 || split[0] == "" {
				if columnWraps && last != "" {
					err := w.Merge(last, strings.Join(split, ""))
					if err != nil {
						return err
					}
				}
				// else silently ignore
				continue
			} else {
				last = split[0]
			}
		}
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
