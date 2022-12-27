package mkarray

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

// This code is ugly. Read at your own risk.

func init() {
	lang.DefineFunction("a", cmdA, types.String)
	lang.DefineFunction("ja", cmdJa, types.Json)
	lang.DefineFunction("ta", cmdTa, types.WriteArray)

	defaults.AppendProfile(`alias ja=a`)
}

const (
	astTypeString = iota
	astTypeOpen
	astTypeClose
	astTypeSeparator
	astTypeRange
)

type ast struct {
	Data []byte
	Type int
}

func cmdA(p *lang.Process) error {
	return mkArray(p, types.String)
}

func cmdJa(p *lang.Process) error {
	return mkArray(p, types.Json)
}

func cmdTa(p *lang.Process) error {
	dataType, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	params := p.Parameters.StringArray()[1:]
	p.Parameters.DefineParsed(params)

	return mkArray(p, dataType)
}

// echo %[1..3]
// a: [1..10] -> ...
func mkArray(p *lang.Process, dataType string) error {
	p.Stdout.SetDataType(dataType)

	expression := p.Parameters.ByteAll()

	ok, err := isNumberArray(p, expression, dataType)
	if ok || err != nil {
		return err
	}

	var (
		escaped, open, dots bool
		nodes               = make([]ast, 1)
		node                = &nodes[0]
	)

	writer, err := p.Stdout.WriteArray(dataType)
	if err != nil {
		return err
	}

	// Parse the parameters
	for i, b := range expression {
		switch b {
		case '\\':
			dots = false
			if escaped {
				node.Data = append(node.Data, b)
			}
			escaped = !escaped

		case ',':
			dots = false
			if escaped {
				node.Data = append(node.Data, b)
				escaped = !escaped
				continue
			}
			nodes = append(nodes,
				ast{
					Data: []byte{','},
					Type: astTypeSeparator,
				},
				ast{},
			)
			node = &nodes[len(nodes)-1]

		case '[':
			dots = false
			if escaped {
				node.Data = append(node.Data, b)
				escaped = !escaped
				continue
			}
			if open {
				return fmt.Errorf("cannot open bracket (char %d) inside of open bracket.\nIf you wanted to print the bracket then please escape it: `\\[``", i)
			}
			open = true
			nodes = append(nodes,
				ast{
					Data: []byte{'['},
					Type: astTypeOpen,
				},
				ast{},
			)
			node = &nodes[len(nodes)-1]

		case ']':
			dots = false
			if escaped {
				node.Data = append(node.Data, b)
				escaped = !escaped
				continue
			}
			if !open {
				return fmt.Errorf("cannot close bracket (char %d) with an open bracket.\nIf you wanted to print the bracket then please escape it: `\\]``", i)
			}
			open = false
			nodes = append(nodes,
				ast{
					Data: []byte{']'},
					Type: astTypeClose,
				},
				ast{},
			)
			node = &nodes[len(nodes)-1]

		case '.':
			if open {
				if dots {
					node.Type = astTypeRange
				}
				dots = !dots
			}
			node.Data = append(node.Data, b)

		default:
			dots = false
			escaped = false
			node.Data = append(node.Data, b)
		}
	}

	if open {
		return fmt.Errorf("missing closing square bracket, ']', in: %s", string(expression))
	}

	// Group the parameters to handle recursive matching
	groups := make([][]ast, 1)
	for i := range nodes {
		switch nodes[i].Type {
		case astTypeOpen:
			open = true
			groups[len(groups)-1] = append(groups[len(groups)-1], nodes[i])

		case astTypeClose:
			open = false
			groups[len(groups)-1] = append(groups[len(groups)-1], nodes[i])

		case astTypeSeparator:
			if open {
				groups[len(groups)-1] = append(groups[len(groups)-1], nodes[i])
			} else {
				groups = append(groups, []ast{})
			}

		default:
			groups[len(groups)-1] = append(groups[len(groups)-1], nodes[i])
		}
	}

	// Now do your magic
	var (
		marker = string([]byte{0})
	)

	for g := range groups {
		var (
			template string
			variable = make(map[int][]string)
			l        = -1
		)

		for n := range groups[g] {
			if p.HasCancelled() {
				goto cancelled
			}

			switch groups[g][n].Type {
			case astTypeString:
				if open {
					variable[l] = append(variable[l], string(groups[g][n].Data))
					continue
				}
				template += string(groups[g][n].Data)

			case astTypeRange:
				a, err := rangeToArray(groups[g][n].Data)
				if err != nil {
					return err
				}
				variable[l] = append(variable[l], a...)
				continue

			case astTypeOpen:
				template += marker
				l++
				variable[l] = make([]string, 0)
				open = true

			case astTypeClose:
				open = false
			}
		}

		counter := make([]int, len(variable))

		for {
		nextIndex:
			if p.HasCancelled() {
				goto cancelled
			}

			s := template
			for t := 0; t < len(counter); t++ {
				c := counter[t]
				s = strings.Replace(s, marker, variable[t][c], 1)
			}
			writer.WriteString(s)

			i := len(counter) - 1
			if i < 0 {
				goto nextGroup
			}

			counter[i]++
			if counter[i] == len(variable[i]) {
			nextCounter:
				counter[i] = 0
				i--
				if i < 0 {
					goto nextGroup
				}
				counter[i]++
				if counter[i] < len(variable[i]) {
					goto nextIndex
				} else {
					goto nextCounter
				}
			} else {
				goto nextIndex
			}

		}
	nextGroup:
	}

cancelled:
	return writer.Close()
}

var rxIsNumberArray = regexp.MustCompile(`^\[([0-9]+)..([0-9]+)\]$`)

func isNumberArray(p *lang.Process, expression []byte, dataType string) (bool, error) {
	// these data types are all strings anyway. So no point making them numeric
	if dataType == types.String || dataType == types.Generic {
		return false, nil
	}

	match := rxIsNumberArray.FindAllSubmatch(expression, -1)
	if len(match) != 1 {
		return false, nil
	}

	if len(match[0]) != 3 {
		return false, nil
	}

	// [0n..] should be strings
	if len(match[0][1]) > 1 && match[0][1][0] == '0' {
		return false, nil
	}

	// [..0n] should be strings
	if len(match[0][2]) > 1 && match[0][2][0] == '0' {
		return false, nil
	}

	left, err := strconv.Atoi(string(match[0][1]))
	if err != nil {
		return true, err
	}

	right, err := strconv.Atoi(string(match[0][2]))
	if err != nil {
		return true, err
	}

	var (
		slice []int
		i     int
	)

	switch {
	case left < right:
		slice = make([]int, right-left+1)
		for v := left; v != right+1; v++ {
			slice[i] = v
			i++
		}

	case left > right:
		slice = make([]int, left-right+1)
		for v := left; v != right-1; v-- {
			slice[i] = v
			i++
		}

	default:
		return false, nil
	}

	b, err := lang.MarshalData(p, dataType, slice)
	if err != nil {
		return true, err
	}

	_, err = p.Stdout.Write(b)
	return true, err
}
