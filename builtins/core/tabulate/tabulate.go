package tabulate

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineMethod("tabulate", cmdTabulate, types.Generic, types.Any)

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
	rxSplitComma = regexp.MustCompile(`[\s\t]*,[\s\t]*`)
	rxSplitSpace = regexp.MustCompile(`[\s\t]+-`)
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
	fSeparator:   "String, custom regex pattern for splitting fields (default: `" + constSeparator + "`)",
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
		//iKeyStart   int // where the key starts when column wraps and keyVal used
		processKey bool
		//iValStart   int // where the value starts when column wraps and keyVal used
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
			p.Name.String(), types.String, types.Generic, dt)
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
		split = rxTableSplit.Split(s, -1)
		if len(split) == 0 {
			continue
		}

		// table has indentation, lets remove that
		if len(split) > 1 && split[0] == "" {
			split = split[1:]
		}

		if keyVal && (len(split) < 2 || split[0] == "") {
			// is this a wrapped column?
			if columnWraps && last != "" {
				// is it a wrapped key? Check if indented flag
				for i := 0; i < len(s); i++ {
					if s[i] == '-' && i > 0 {
						// it's a key
						processKey = true
						if len(split) == 1 {
							split = []string{split[0], ""}
						}

						break
					}
					if s[i] != ' ' && s[i] != '\t' {
						break
					}
				}

				// look like it's just a wrapped column
				if !processKey {
					if len(colWrapsBuf) == 0 || colWrapsBuf[len(colWrapsBuf)-1] == ' ' {
						colWrapsBuf += strings.Join(split, joiner)
					} else {
						colWrapsBuf += joiner + strings.Join(split, joiner)
					}
				}
			}

			// else silently ignore heading
			if !processKey {
				continue
			}
		}

		if len(split) > 1 || processKey { // recheck because we've redefined the length of split
			processKey = false

			// looks like there's a new key, so lets write the colWrapsBuf
			if columnWraps && last != "" {
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

				/*for i, r := range s {
					if r != ' ' && r != '\t' {
						iKeyStart = i
						break
					}
				}*/
			}

			// split keys by comma
			if splitComma {
				keys = rxSplitComma.Split(split[0], -1)
			}

			// split keys by space
			if splitSpace {
				keys = rxSplitSpace.Split(split[0], 2)
				if len(keys) == 2 {
					keys[1] = "-" + keys[1]
				}
			}

			// remove the hint stuff
			if keyIncHint {
				var hint string
				if len(keys) != 0 {
					_, hint = stripKeyHint(keys)
				} else {
					split[0], hint = stripKeyHint([]string{split[0]})
				}
				if len(hint) != 0 {
					split[1] = "(args: " + hint + ") " + split[1]
				}
			}

		}

		if keyVal {
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

var rxSquareHints = regexp.MustCompile(`\[.*\]$`)

func stripKeyHint(keys []string) (string, string) {
	var (
		space, equ, square []string
		hint               string
	)

	for i := range keys {
		square = rxSquareHints.FindStringSubmatch(keys[i])
		if len(square) != 0 {
			keys[i] = strings.Replace(keys[i], square[0], "", 1)
		}

		space = strings.SplitN(keys[i], " ", 2)
		keys[i] = space[0]
		if strings.Contains(space[0], "=") {
			equ = strings.SplitN(keys[i], "=", 2)
			keys[i] = equ[0] + "="
		}
	}

	switch {
	case len(square) != 0:
		hint = square[0]

	case len(space) == 2:
		hint = space[1]

	case len(equ) == 2:
		hint = equ[1]
	}

	return keys[0], hint
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
