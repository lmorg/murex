package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
)

func (tree *expTreeT) parseSubShell(exec bool, strOrVal bool) ([]rune, interface{}, string, error) {
	var (
		brackets = 1
		escape   bool
	)

	start := tree.charPos

	tree.charPos += 2

	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case escape:
			escape = false

		case r == '\\':
			escape = true

		case r == '{':
			return nil, "", "", raiseError(
				tree.expression, nil, tree.charPos, fmt.Sprintf(
					"too many nested brackets '{' %d",
					tree.charPos))

		case r == '}':
			brackets--
			if brackets == 0 {
				goto endSubShell
			}
		}
	}

	return nil, nil, "", raiseError(
		tree.expression, nil, tree.charPos, fmt.Sprintf(
			"missing closing bracket '}' at %d\nExpression: %s",
			tree.charPos, `...`+string(tree.expression[start:])))

endSubShell:
	value := tree.expression[start : tree.charPos+1]
	key := tree.expression[start+2 : tree.charPos]

	if !exec {
		return value, nil, "", nil
	}

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_PARENT_VARTABLE)
	exitNum, err := fork.Execute(key)
	if err != nil {
		return nil, nil, "", fmt.Errorf("subshell failed: %s", err.Error())
	}
	if exitNum > 0 && tree.p.RunMode.IsStrict() {
		return nil, nil, "", fmt.Errorf("subshell exit status %d", exitNum)
	}
	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return nil, nil, "", err
	}

	b = utils.CrLfTrim(b)
	dataType := fork.Stdout.GetDataType()

	v, err := formatBytes(b, dataType, strOrVal)
	return value, v, dataType, err
}
