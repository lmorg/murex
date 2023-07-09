//go:build !windows
// +build !windows

package man

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/rmbs"
)

func parseDescriptions(command string, fMap *map[string]string) {
	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.Name.Set("(man)")
	err := fork.Variables.Set(fork.Process, "command", command, types.String)
	if err != nil {
		if debug.Enabled {
			panic(err)
		}
		return
	}
	_, err = fork.Execute(manBlock)
	if err != nil {
		if debug.Enabled {
			panic(err)
		}
		return
	}

	parseDescriptionsLines(fork.Stdout, fMap)
}

var rxHeading = regexp.MustCompile(`^[A-Z]+$`)

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

		if section != "DESCRIPTION" && section != "OPTIONS" {
			return
		}

		ws := countWhiteSpace(b)
		switch {
		case ws == 0:
			updateFlagMap(pl, fMap)
			pl = nil

		case ws == len(b)-1:
			updateFlagMap(pl, fMap)
			pl = nil

		case b[ws] == '-':
			updateFlagMap(pl, fMap)
			pl = parseLineFlags(b[ws:])

		case pl == nil:
			return

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
	for i := range pl.Flags {
		if pl.Example == "" {
			(*fMap)[pl.Flags[i]] = strings.TrimSpace(pl.Description)
		} else {
			(*fMap)[pl.Flags[i]] = fmt.Sprintf(
				"(%s) %s",
				pl.Example, strings.TrimSpace(pl.Description))
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
			pl.Flags = append(pl.Flags, string(match))
			pl.Position += len(match)

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
