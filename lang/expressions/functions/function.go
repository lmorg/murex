// go:generate stringify

package functions

import "strings"

type FunctionT struct {
	Command    []rune
	Parameters [][]rune
	NamedPipes []string
	Cast       []rune
	Properties Property
	LineN      int
	ColumnN    int
	Raw        []rune
}

func (fn FunctionT) CommandName() []rune {
	name := fn.Command
	if len(name) > 0 && name[len(name)-1] == ':' {
		name = name[:len(name)-1]
	}
	return name
}

type Property int

const (
	P_NEW_CHAIN Property = 1 << iota
	P_METHOD
	P_FOLLOW_ON
	P_PIPE_OUT
	P_PIPE_ERR
	P_LOGIC_AND
	P_LOGIC_OR
)

func (prop Property) NewChain() bool   { return prop&P_NEW_CHAIN != 0 }
func (prop Property) Method() bool     { return prop&P_METHOD != 0 }
func (prop Property) FollowOnFn() bool { return prop&P_FOLLOW_ON != 0 }
func (prop Property) PipeOut() bool    { return prop&P_PIPE_OUT != 0 }
func (prop Property) PipeErr() bool    { return prop&P_PIPE_ERR != 0 }
func (prop Property) LogicAnd() bool   { return prop&P_LOGIC_AND != 0 }
func (prop Property) LogicOr() bool    { return prop&P_LOGIC_OR != 0 }

func (prop Property) Decompose() string {
	var a []string

	if prop.NewChain() {
		a = append(a, "new pipeline (`\\n`, `;`)")
	}

	if prop.PipeOut() || prop.Method() {
		a = append(a, "pipe out (`|`, `->`, `=>`, `|>`, `>>`)")
	}

	if prop.PipeErr() || prop.Method() {
		a = append(a, "pipe err (`?`)")
	}

	if prop.LogicAnd() {
		a = append(a, "logic AND (`&&`)")
	}

	if prop.LogicOr() {
		a = append(a, "logic OR (`||`)")
	}

	s := strings.Join(a, ", ")

	return s
}
