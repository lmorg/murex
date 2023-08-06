package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func (tree *ParserT) parseFunction(exec bool, cmd []rune) ([]rune, any, string, error) {
	r, err := tree.parseParen(false)
	if err != nil {
		return r, nil, types.Null, fmt.Errorf("cannot parse `function(parameters...)`: %s", err.Error())
	}

	if !exec {
		return r, nil, types.Null, nil
	}

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
	params := append([]rune{' '}, r[1:len(r)-1]...)
	block := append(cmd, params...)
	exitNum, err := fork.Execute(block)
	if err != nil {
		return r, nil, types.Null,
			fmt.Errorf("function `%s` compilation error: %s",
				string(cmd), err.Error())
	}

	if exitNum != 0 {
		return r, nil, types.Null,
			fmt.Errorf("function `%s` returned non-zero exit number (%d)",
				string(cmd), exitNum)
	}

	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return r, nil, types.Null, fmt.Errorf("function `%s` STDOUT read error: %s",
			string(cmd), err.Error())
	}
	dt := fork.Stdout.GetDataType()
	v, err := formatBytes(b, dt, varAsValue)
	if err != nil {
		return r, nil, types.Null, fmt.Errorf("function `%s` STDOUT conversion error: %s",
			string(cmd), err.Error())
	}
	return r, v, dt, nil
}
