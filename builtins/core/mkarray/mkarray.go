package mkarray

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("a", cmdA, types.String)
	lang.DefineFunction("ja", cmdJa, types.Json)
	lang.DefineFunction("ta", cmdTa, types.WriteArray)
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

type arrayT struct {
	p          *lang.Process
	expression []byte
	dataType   string
	writer     stdio.ArrayWriter
	groups     [][]ast
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
// ja: [1..10] -> ...
// ta: dataType [1..10] -> ...
func mkArray(p *lang.Process, dataType string) error {
	p.Stdout.SetDataType(dataType)

	a := new(arrayT)
	a.p = p
	a.expression = p.Parameters.ByteAll()
	a.dataType = dataType

	ok, err := a.isNumberArray()
	if ok || err != nil {
		return err
	}

	return a.isStringArray()
}

func (a *arrayT) parseExpression() error {
	var (
		escaped, open, dots bool
		nodes               = make([]ast, 1)
		node                = &nodes[0]
	)

	for i, b := range a.expression {
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
		return fmt.Errorf("missing closing square bracket, ']', in: %s", string(a.expression))
	}

	// Group the parameters to handle recursive matching
	a.groups = make([][]ast, 1)
	for i := range nodes {
		switch nodes[i].Type {
		case astTypeOpen:
			open = true
			a.groups[len(a.groups)-1] = append(a.groups[len(a.groups)-1], nodes[i])

		case astTypeClose:
			open = false
			a.groups[len(a.groups)-1] = append(a.groups[len(a.groups)-1], nodes[i])

		case astTypeSeparator:
			if open {
				a.groups[len(a.groups)-1] = append(a.groups[len(a.groups)-1], nodes[i])
			} else {
				a.groups = append(a.groups, []ast{})
			}

		default:
			a.groups[len(a.groups)-1] = append(a.groups[len(a.groups)-1], nodes[i])
		}
	}

	return nil
}
