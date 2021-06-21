package mkarray

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

// This code is ugly. Read at your own risk.

func init() {
	lang.GoFunctions["a"] = cmdA
	lang.GoFunctions["ja"] = cmdJa
	lang.GoFunctions["ta"] = cmdTa
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

// echo @{a: abc[1,2,3],[1..3]}
// a: [1..10] -> ...
func mkArray(p *lang.Process, dataType string) error {
	p.Stdout.SetDataType(dataType)

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
	for i, b := range p.Parameters.ByteAll() {
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
				return fmt.Errorf("Cannot open bracket (char %d) inside of open bracket.\nIf you wanted to print the bracket then please escape it: `\\[``", i)
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
				return fmt.Errorf("Cannot close bracket (char %d) with an open bracket.\nIf you wanted to print the bracket then please escape it: `\\]``", i)
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
