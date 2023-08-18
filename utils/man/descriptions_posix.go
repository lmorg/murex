//go:build !windows
// +build !windows

package man

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/murex/utils/rmbs"
)

func GetManPage(command string, width int) stdio.Io {
	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.Name.Set("(man)")
	err := fork.Variables.Set(fork.Process, "command", command, types.String)
	if err != nil {
		if debug.Enabled {
			panic(err)
		}
		return nil
	}

	_, err = fork.Execute(ManPageExecBlock(width))
	if err != nil {
		if debug.Enabled {
			panic(err)
		}
		return nil
	}

	return fork.Stdout
}

func parseDescriptions(command string, fMap *map[string]string) {
	stdout := GetManPage(command, 1000)
	parseDescriptionsLines(stdout, fMap)
}

var rxHeading = regexp.MustCompile(`^[A-Z]+$`)

var validSections = []string{
	"DESCRIPTION",
	"OPTIONS",
	"PRIMARIES",  // required for `find` on macOS
	"EXPRESSION", // required for `find` on GNU
}

func parseDescriptionsLines(io stdio.Io, fMap *map[string]string) {
	var pl *parsedLineT
	var section string
	err := io.ReadLine(func(b []byte) {
		b = []byte(rmbs.Remove(string(b)))
		b = utils.CrLfTrim(b)

		heading := rxHeading.Find(b)
		if len(heading) > 0 {
			section = string(heading)
		}

		if !lists.Match(validSections, section) {
			return
		}

		ws := countWhiteSpace(b)
		switch {
		case ws == 0:
			fallthrough

		case ws == len(b)-1:
			updateFlagMap(pl, fMap)
			pl = nil

		case b[ws] == '-':
			updateFlagMap(pl, fMap)
			pl = parseLineFlags(b[ws:])

		case pl == nil:
			return

		case pl.Description != "" && len(pl.Description) < 30 && ws >= 8: // kludge for `find` style flags
			pl.Example += " " + pl.Description
			pl.Description = ""
			fallthrough

		case pl.Description == "":
			pl.Position = ws
			fallthrough

		case ws == pl.Position:
			pl.Description += " " + string(b[ws:])

		default:
			updateFlagMap(pl, fMap)
			pl = nil
		}
	})
	if err != nil {
		if debug.Enabled {
			panic(err)
		}
		return
	}
}

func updateFlagMap(pl *parsedLineT, fMap *map[string]string) {
	if pl == nil {
		return
	}

	pl.Description = strings.TrimSpace(pl.Description)
	pl.Description = strings.ReplaceAll(pl.Description, "  ", " ")

	for i := range pl.Flags {
		if pl.Example == "" {
			(*fMap)[pl.Flags[i]] = pl.Description
		} else {
			(*fMap)[pl.Flags[i]] = fmt.Sprintf(
				"eg: %s -- %s",
				strings.TrimSpace(pl.Example), pl.Description)
		}
	}
}

func countWhiteSpace(b []byte) int {
	for i := range b {
		if b[i] == ' ' || b[i] == '\t' {
			continue
		}
		return i
	}
	return 0
}

var (
	//rxLineMatchFlag = regexp.MustCompile(`^-[-_a-zA-Z0-9]+`)
	rxLineMatchFlag = regexp.MustCompile(`^-[-_a-zA-Z0-9]+`)
	rxExampleCaps   = regexp.MustCompile(`^[A-Z]+([\t, ]|$)`)
)

type parsedLineT struct {
	Position    int
	Description string
	Example     string
	Flags       []string
}

func parseLineFlags(b []byte) *parsedLineT {
	//defer recover()

	pl := new(parsedLineT)

	for {
	start:
		if pl.Position == len(b) {
			return pl
		}

		switch b[pl.Position] {
		case ',':
			pl.Position += countWhiteSpace(b[pl.Position+1:]) + 1
			fallthrough

		case '-':
			match := rxLineMatchFlag.Find(b[pl.Position:])
			if len(match) == 0 {
				pl.Description = string(b[pl.Position:])
				return pl
			}
			//pl.Flags = append(pl.Flags, string(match))
			//pl.Position += len(match)
			i := parseFlag(b[pl.Position:], pl)
			pl.Position += i

		case '=', '[':
			start := pl.Position
			var group bool
			for ; pl.Position < len(b); pl.Position++ {
				switch b[pl.Position] {
				case '[':
					group = true
				case ']':
					group = false
				case ' ', '\t', ',':
					if group {
						continue
					}
					pl.Example = string(b[start:pl.Position])
					goto start
				}
			}

		case ' ':
			example := rxExampleCaps.Find(b[pl.Position+1:])
			switch {
			case len(example) == 0:
				// start of description
				pl.Description = string(b[pl.Position+1:])
				return pl
			case pl.Position+len(example) == len(b)-1:
				// end of line
				pl.Example = string(b[pl.Position:])
				pl.Position += len(example) + 1
				return pl
			default:
				pl.Example = string(b[pl.Position : pl.Position+len(example)])
				pl.Position += len(example)
			}

		default:
			pl.Description = string(b[pl.Position:])
			return pl
		}
	}
}

func parseFlag(b []byte, pl *parsedLineT) int {
	var bracket, split bool
	for i, c := range b {
		switch {
		case isValidFlagChar(c):
			continue

		case c == '[':
			switch {
			case bracket:
				return 0
			case i+1 == len(b):
				return 0
			case b[i+1] == '=':
				splitFlags(b[:i], split, pl)
				return i
			case isValidFlagChar(b[i+1]):
				bracket = true
			default:
				return 0
			}

		case c == ']':
			if !bracket {
				return 0
			}
			bracket = false
			split = true

		default:
			if bracket {
				return 0
			}
			splitFlags(b[:i], split, pl)
			return i
		}
	}

	splitFlags(b, split, pl)
	return len(b)
}

var (
	empty      = []byte{}
	braceOpen  = []byte{'['}
	braceClose = []byte{']'}
	rxNoBrace  = regexp.MustCompile(`\[.*?\]`)
)

func splitFlags(b []byte, split bool, pl *parsedLineT) {
	if !split {
		pl.Flags = append(pl.Flags, string(b))
		return
	}

	full := bytes.ReplaceAll(b, braceOpen, empty)
	full = bytes.ReplaceAll(full, braceClose, empty)

	removed := rxNoBrace.ReplaceAll(b, empty)

	pl.Flags = append(pl.Flags, string(full), string(removed))
}

func isValidFlagChar(c byte) bool {
	return c == '-' ||
		(c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}
