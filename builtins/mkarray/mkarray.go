package mkarray

import (
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"strings"
)

// This code is ugly. Read at your own risk.

func init() {
	proc.GoFunctions["a"] = proc.GoFunction{Func: mkArray, TypeIn: types.Generic, TypeOut: types.Csv}
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

// echo @{a: abc[1,2,3],[1..3]}
//a: [1..10] -> ...
func mkArray(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	var (
		escaped, open, dots bool
		nodes               []ast = make([]ast, 1)
		node                *ast  = &nodes[0]
	)

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
				return errors.New(fmt.Sprintf("Cannot open bracket (char %d) inside of open bracket.\nIf you wanted to print the bracket then please escape it: `\\[``", i))
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
				return errors.New(fmt.Sprintf("Cannot close bracket (char %d) with an open bracket.\nIf you wanted to print the bracket then please escape it: `\\]``", i))
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
			if dots {
				node.Type = astTypeRange
			}
			dots = !dots
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
		array  []string
		marker string = string([]byte{0})
	)

	for g := range groups {
		var (
			template string
			variable map[int][]string = make(map[int][]string)
			l        int              = -1
		)

		for n := range groups[g] {
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
			s := template
			for t := 0; t < len(counter); t++ {
				c := counter[t]
				s = strings.Replace(s, marker, variable[t][c], 1)
			}
			array = append(array, s)

			i := len(counter) - 1
			for {
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
		}
	nextGroup:
	}

	b, err := utils.JsonMarshal(array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}
