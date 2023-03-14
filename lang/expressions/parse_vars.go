package expressions

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/home"
)

var errUnexpectedClosingParenthesis = fmt.Errorf("expecting closing parenthesis, ')', after variable reference")

func (tree *ParserT) parseVarScalar(exec bool, strOrVal varFormatting) ([]rune, interface{}, string, error) {
	var paren bool

	if tree.nextChar() == '(' {
		tree.charPos++
		paren = true
	}

	if !isBareChar(tree.nextChar()) {
		// always print $
		return []rune{'$'}, "$", types.String, nil
	}

	tree.charPos++
	value := tree.parseBareword()

	if tree.charPos < len(tree.expression) && tree.expression[tree.charPos] == '[' {
		if paren {
			r, v, s, err := tree.parseVarIndexElement(exec, value, strOrVal)
			if tree.nextChar() == ')' {
				tree.charPos++
				return r, v, s, err
			}
			return nil, nil, "", errUnexpectedClosingParenthesis

		} else {
			return tree.parseVarIndexElement(exec, value, strOrVal)
		}
	}

	if paren {
		if tree.charPos < len(tree.expression) && tree.expression[tree.charPos] == ')' {
			tree.charPos++
		} else {
			return nil, nil, "", errUnexpectedClosingParenthesis
		}
	}

	tree.charPos--

	var r []rune
	if paren {
		r = make([]rune, len(value)+3)
		copy(r, []rune{'$', '('})
		copy(r[2:], value)
		r[len(r)-1] = ')'
	} else {
		r = append([]rune{'$'}, value...)
	}

	// don't getVar() until we come to execute the expression, skip when only
	// parsing syntax
	if exec {
		v, dataType, err := tree.getVar(value, strOrVal)
		return r, v, dataType, err
	}

	return r, nil, "", nil
}

func (tree *ParserT) parseVarIndexElement(exec bool, varName []rune, strOrVal varFormatting) ([]rune, interface{}, string, error) {
	if tree.nextChar() == '{' {
		return tree.parseLambda(exec, '$', varName, strOrVal)
	}

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

func treePlusPlus(tree *ParserT) { tree.charPos++ }
func (tree *ParserT) parseLambda(exec bool, prefix rune, varName []rune, strOrVal varFormatting) ([]rune, interface{}, string, error) {
	defer treePlusPlus(tree)
	if exec {
		if tree.p == nil {
			panic("`tree.p` is undefined")
		}

		path := string(varName)

		value, err := tree.p.Variables.GetValue(path)
		if err != nil {
			return nil, nil, "", err
		}

		dataType := tree.p.Variables.GetDataType(path)

		err = tree.p.Variables.Set(tree.p, "", value, dataType)
		if err != nil {
			return nil, nil, "", fmt.Errorf("unable to set `$.`: %s", err.Error())
		}
	}

	return tree.parseSubShell(exec, prefix, strOrVal)
}

func (tree *ParserT) parseVarArray(exec bool) ([]rune, interface{}, error) {
	if !isBareChar(tree.nextChar()) {
		return nil, nil, errors.New("'@' symbol found but no variable name followed")
	}

	tree.charPos++
	value := tree.parseBareword()

	if tree.charPos < len(tree.expression) && tree.expression[tree.charPos] == '[' {
		return tree.parseVarRange(exec, value)
	}

	tree.charPos--

	if !exec {
		// don't getArray() until we come to execute the expression, skip when only
		// parsing syntax
		return append([]rune{'@'}, value...), nil, nil
	}

	v, err := tree.getArray(value)
	return value, v, err
}

func (tree *ParserT) parseVarRange(exec bool, varName []rune) ([]rune, interface{}, error) {
	if tree.nextChar() == '{' {
		if exec {
			return tree.parseLambdaArray(varName)
		} else {
			r, v, _, err := tree.parseLambda(exec, '@', varName, varAsValue)
			return r, v, err
		}
	}

	var escape bool

	start := tree.charPos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case escape:
			escape = false

		case r == '\\':
			escape = true

		case r == '[':
			return nil, "", raiseError(
				tree.expression, nil, tree.charPos, "too many nested square '[' brackets")

		case r == ']':
			goto endRange
		}
	}

	return nil, "", raiseError(
		tree.expression, nil, tree.charPos, "missing closing bracket ']'")

endRange:
	key := tree.expression[start+1 : tree.charPos]
	flags := []rune{}
	if isBareChar(tree.nextChar()) {
		tree.charPos++
		flags = tree.parseBareword()
		tree.charPos--
	}
	value := tree.expression[start-len(varName)-1 : tree.charPos]

	if !exec {
		return value, "", nil
	}

	v, err := tree.getVarRange(varName, key, flags)
	if err != nil {
		return nil, "", err
	}
	return nil, v, nil
}

func (tree *ParserT) parseLambdaArray(varName []rune) ([]rune, interface{}, error) {
	// no `exec` boolean here because this method should only be invoked when `exec == true`
	defer treePlusPlus(tree)
	if tree.p == nil {
		panic("`tree.p` is undefined")
	}

	path := string(varName)
	v, err := tree.p.Variables.GetValue(path)
	if err != nil {
		return nil, nil, err
	}

	switch t := v.(type) {
	case []interface{}:
		pos := tree.charPos
		array := make([]interface{}, len(t))

		var (
			item interface{}
			r    []rune
			j    int
		)

		for i := range t {
			tree.charPos = pos
			element := fmt.Sprintf("%s.%d", path, i)

			value, err := tree.p.Variables.GetValue(element)
			if err != nil {
				return nil, nil, err
			}

			dataType := tree.p.Variables.GetDataType(element)

			err = tree.p.Variables.Set(tree.p, "", value, dataType)
			if err != nil {
				return nil, nil, fmt.Errorf("unable to set `$.`: %s", err.Error())
			}

			r, item, _, err = tree.parseSubShell(true, '$', varAsValue)
			if err != nil {
				return nil, nil, err
			}
			switch item.(type) {
			case string:
				if len(item.(string)) > 0 {
					array[j] = item
					j++
				}
			case bool:
				if item.(bool) {
					array[j] = value
					j++
				}
			default:
				array[j] = item
				j++
			}
		}

		return r, array[:j], nil

	default:
		return nil, nil, fmt.Errorf("cannot run lambda. Expecting an array, instead got '%T' in '%s'", t, path)
	}
}

func isUserNameChar(r rune) bool {
	return isBareChar(r) || r == '.' || r == '-'
}

func (tree *ParserT) parseVarTilde(exec bool) string {
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
