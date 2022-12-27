package expressions

import (
	"errors"

	"github.com/lmorg/murex/lang/expressions/noglob"
	"github.com/lmorg/murex/utils/lists"
)

type StatementT struct {
	command    []rune
	parameters [][]rune
	paramTemp  []rune
	namedPipes []string

	// parser states
	canHaveZeroLenStr bool // to get around $VARS being empty or unset
	possibleGlob      bool // to signal to NextParameter of a possible glob
	asStatement       bool // force murex to parse expression as statement
	ignoreCrLf        bool // allow '\' to escape a new line
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

func (tree *ParserT) nextParameter() error {
	st := tree.statement

	switch {

	case len(st.command) == 0:
		// no command yet so this must be a command
		st.command = st.paramTemp
		st.possibleGlob = false

	case st.possibleGlob:
		// glob
		st.possibleGlob = false
		st.canHaveZeroLenStr = false
		if !tree.ExpandGlob() || lists.Match(noglob.GetNoGlobCmds(), st.String()) {
			st.parameters = append(st.parameters, st.paramTemp)
			break
		}
		v, err := tree.parseGlob(st.paramTemp)
		if err != nil {
			return err
		}
		if v == nil {
			st.parameters = append(st.parameters, st.paramTemp)
			break
		}
		for i := range v {
			st.parameters = append(st.parameters, []rune(v[i]))
		}

	case st.canHaveZeroLenStr:
		// variable, possibly zero length
		st.parameters = append(st.parameters, st.paramTemp)
		st.canHaveZeroLenStr = false

	case len(st.paramTemp) == 0:
		// just empty space. Nothing to do
		return nil

	default:
		// just a regular old parameter
		st.parameters = append(st.parameters, st.paramTemp)
	}

	st.paramTemp = []rune{}
	return nil
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
