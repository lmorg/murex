//go:build !windows
// +build !windows

package man

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
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

func parseDescriptionsLines(r io.Reader, fMap *map[string]string) {
	var pl *parsedLineT
	var section string

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		b := scanner.Bytes()
		//b := append(scanner.Bytes(), utils.NewLineByte...)

		b = []byte(rmbs.Remove(string(b)))
		b = utils.CrLfTrim(b)

		heading := rxHeading.Find(b)
		if len(heading) > 0 {
			section = string(heading)
		}

		if !lists.Match(validSections, section) {
			continue
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
			continue

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
	}
	if err := scanner.Err(); err != nil && debug.Enabled {
		panic(err)
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
		fmt.Println(json.LazyLoggingPretty(*pl), "-->"+string(b)+"<--")
		if pl.Position == len(b) {
			return pl
		}

		switch b[pl.Position] {
		case ',':
			pl.Position += countWhiteSpace(b[pl.Position+1:]) + 1
			//fallthrough

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

		case '=', '[', '<':
			start := pl.Position
			group := true
			grpC := b[pl.Position]
			for ; pl.Position < len(b); pl.Position++ {
				switch b[pl.Position] {
				case '[', '<':
					switch grpC {
					case '=':
						grpC = b[pl.Position]
					}
				case ']':
					switch grpC {
					case '[':
						group = false
					}
				case '>':
					switch grpC {
					case '<':
						group = false
					}
				case ' ', '\t', ',':
					if grpC == '[' || grpC == '<' {
						continue
					}
					group = false
				}
				if !group {
					break
				}
			}
			pl.Example = string(b[start:pl.Position])
			goto start

		case ' ', '\t':
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
	fmt.Println("parseFlag", string(b))
	var (
		split   bool
		bracket byte = 0
	)

	for i, c := range b {
		fmt.Printf("i==%d c=='%s' bracket=%d\n", i, string(c), bracket)
		switch {
		case isValidFlagChar(c):
			continue

		case c == '[', c == '<':
			switch {
			case bracket == c:
				return 0
			case i+1 == len(b):
				return 0
			case b[i+1] == '=':
				splitFlags(b[:i], split, pl)
				return i
			//case isValidFlagChar(b[i+1]), b[i+1] == '=', b[i+1] == '<':
			//	bracket = true
			case bracket != 0:
				continue
			default:
				//	return 0
				bracket = c
			}

		case c == ']':
			switch bracket {
			case 0:
				return 0
			case '[':
				bracket = 0
				split = true
			default:
				continue
			}

		case c == '>':
			switch bracket {
			case 0:
				return 0
			case '<':
				bracket = 0
				split = true
			default:
				continue
			}

		default:
			if bracket != 0 {
				continue
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
