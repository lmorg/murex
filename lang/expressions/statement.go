package expressions

import (
	"errors"
)

type StatementT struct {
	command    []rune
	parameters [][]rune
	paramTemp  []rune
	namedPipes []string

	canHaveZeroLenStr bool // to get around $VARS being empty or unset
}

func (st *StatementT) String() string {
	return string(st.command)
}

func (st *StatementT) Parameters() []string {
	params := make([]string, len(st.parameters))

	for i := range st.parameters {
		params[i] = string(st.parameters[i])
	}

	return params
}

func (st *StatementT) NextParameter() {
	switch {

	case len(st.command) == 0:
		// no command yet so this must be a command
		st.command = st.paramTemp

	case st.canHaveZeroLenStr:
		st.parameters = append(st.parameters, st.paramTemp)
		st.canHaveZeroLenStr = false

	case len(st.paramTemp) == 0:
		// just empty space. Nothing to do
		return

	default:
		// just a regular old parameter
		st.parameters = append(st.parameters, st.paramTemp)
	}

	st.paramTemp = []rune{}
}

func (st *StatementT) validate() error {
	switch {
	case len(st.command) == 0:
		return errors.New("no command specified (empty command property)")

	case st.command[0] == '$':
		return errors.New("commands cannot begin with '$'. Please quote or escape this character")

	case st.command[0] == '@' && len(st.command) > 1 && st.command[1] != '[':
		return errors.New("commands cannot begin with '@'. Please quote or escape this character")

	case st.command[0] == '%':
		return errors.New("commands cannot begin with '%'. Please quote or escape this character")

	default:
		return nil
	}
}
