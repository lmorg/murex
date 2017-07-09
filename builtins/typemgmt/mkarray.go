package typemgmt

import (
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"strconv"
	"strings"
)

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
				//groups[len(groups)-1] = append(groups[len(groups)-1], ast{})
				groups = append(groups, []ast{})
			}

		default:
			groups[len(groups)-1] = append(groups[len(groups)-1], nodes[i])
		}
	}

	// Now do your magic
	var array []string

	for g := range groups {
		var (
			template string
			//static   []string
			variable map[int][]string = make(map[int][]string)
			l        int              = len(variable) - 1
		)

		for n := range groups[g] {
			switch groups[g][n].Type {
			case astTypeString:
				if open {
					variable[len(static)-1] = append(variable[len(static)-1], string(groups[g][n].Data))
					continue
				}
				static = append(static, string(groups[g][n].Data))

			case astTypeRange:
				a, err := rangeToArray(groups[g][n].Data)
				if err != nil {
					return err
				}
				variable[len(static)-1] = append(variable[len(static)-1], a...)
				continue

			case astTypeOpen:
				variable[len(static)-1] = make([]string, 0)
				open = true

			case astTypeClose:
				open = false

			}
		}

		var (
			s       string
			i       int
			counter map[int]int = make(map[int]int)
		)

		for {
			s += static[i] + variable[i][counter[i]]
			i++
			if i == len(static)-1 {
				if s != "" {
					array = append(array, s)
				}
				i = 0
				j := len(static) - 1
				for {
					counter[j]++
					if counter[j] >= len(variable[j]) {
						if j == 0 {
							goto nextGroup
						}
						counter[j] = 0
						if counter[j-1] < len(variable[j-1])-1 {
							counter[j-1]++
							s = ""
							break
						}
						j--
					}
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

func rangeToArray(b []byte) ([]string, error) {
	split := strings.Split(string(b), "..")
	if len(split) > 2 {
		return nil, errors.New("Invalid syntax. Too many periods in range, `..`. Please escape periods, `\\.`, if you wish to include period in your range.")
	}

	if len(split) < 2 {
		return nil, errors.New("Invalid syntax. Range periods, `..`, found but cannot determine start and end range.")
	}

	i1, e1 := strconv.Atoi(split[0])
	i2, e2 := strconv.Atoi(split[1])

	if e1 == nil && e2 == nil {
		switch {
		case i1 < i2:
			a := make([]string, i2-i1+1)
			for i := range a {
				a[i] = strconv.Itoa(i + i1)
			}
			return a, nil
		case i1 > i2:
			a := make([]string, i1-i2+1)
			for i := range a {
				a[i] = strconv.Itoa(i1 - i)
			}
			return a, nil
		default:
			return nil, errors.New("Invalid range. Start and end of range are the same.")
		}
	}

	return nil, errors.New("TODO: write code pleases")
}
