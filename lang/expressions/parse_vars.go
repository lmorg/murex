package expressions

import (
	"errors"

	"github.com/lmorg/murex/utils/home"
)

func (tree *expTreeT) parseVarScalar(exec bool, strOrVal varFormatting) ([]rune, interface{}, string, error) {
	if !isBareChar(tree.nextChar()) {
		return nil, nil, "", errors.New("'$' symbol found but no variable name followed")
	}

	tree.charPos++
	value := tree.parseBareword()

	if tree.charPos < len(tree.expression) && tree.expression[tree.charPos] == '[' {
		return tree.parseVarIndexElement(exec, value, strOrVal)
	}

	tree.charPos--

	if !exec {
		// don't getVar() until we come to execute the expression, skip when only
		// parsing syntax
		return append([]rune{'$'}, value...), nil, "", nil
	}

	v, dataType, err := tree.getVar(value, strOrVal)
	return value, v, dataType, err
}

func (tree *expTreeT) parseVarIndexElement(exec bool, varName []rune, strOrVal varFormatting) ([]rune, interface{}, string, error) {
	var (
		brackets = 1
		escape   bool
	)

	start := tree.charPos

	if tree.nextChar() == '[' {
		brackets++
		tree.charPos++
	}

	tree.charPos++

	isIorE := brackets

	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case escape:
			escape = false

		case r == '\\':
			escape = true

		case r == '[':
			return nil, "", "", raiseError(
				tree.expression, nil, tree.charPos, "too many nested square '[' brackets")

		case r == ']':
			brackets--
			if brackets == 0 {
				goto endIndexElement
			}
		}
	}

	return nil, "", "", raiseError(
		tree.expression, nil, tree.charPos, "missing closing bracket ']'")

endIndexElement:
	value := tree.expression[start-len(varName)-1 : tree.charPos+1]
	key := tree.expression[start+isIorE : tree.charPos-isIorE+1]

	if !exec {
		return value, "", "", nil
	}

	v, dt, err := tree.getVarIndexOrElement(varName, key, isIorE, strOrVal)
	if err != nil {
		return nil, "", "", err
	}
	return nil, v, dt, nil
}

func (tree *expTreeT) parseVarArray(exec bool) ([]rune, interface{}, error) {
	if !isBareChar(tree.nextChar()) {
		return nil, nil, errors.New("'@' symbol found but no variable name followed")
	}

	tree.charPos++
	value := tree.parseBareword()

	tree.charPos--

	if !exec {
		// don't getArray() until we come to execute the expression, skip when only
		// parsing syntax
		return append([]rune{'@'}, value...), nil, nil
	}

	v, err := tree.getArray(value)
	return value, v, err
}

func isUserNameChar(r rune) bool {
	return isBareChar(r) || r == '.' || r == '-'
}

func (tree *expTreeT) parseVarTilde(exec bool) string {
	tree.charPos++
	start := tree.charPos

	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		switch {
		case isUserNameChar(tree.expression[tree.charPos]):
			// valid user name

		default:
			// not a valid username character
			goto endTilde
		}
	}

endTilde:
	user := string(tree.expression[start:tree.charPos])
	tree.charPos--

	if !exec {
		return "~" + user
	}

	if len(user) == 0 {
		return home.MyDir
	}

	return home.UserDir(user)
}
