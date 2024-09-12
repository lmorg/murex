package expressions

import (
	"errors"
	"fmt"
	"os/user"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/home"
)

func (tree *ParserT) parseVarScalarExpr(exec, execScalars bool) ([]rune, interface{}, string, primitives.FunctionT, error) {
	runes, v, mxDt, err := tree.parseVarScalar(execScalars, varAsValue)

	if exec {
		fn := func() (*primitives.Value, error) {
			switch t := v.(type) {
			case *getVarIndexOrElementT:
				val, mxDt, err := tree.getVarIndexOrElement(t)
				return &primitives.Value{Value: val, DataType: mxDt}, err

			case parseLambdaExecTrueT:
				_, val, mxDt, err := t()
				return &primitives.Value{Value: val, DataType: mxDt}, err

			default:
				val, mxDt, err := tree.getVar(scalarNameDetokenised(runes), varAsValue)
				return &primitives.Value{Value: val, DataType: mxDt}, err
			}
		}
		return runes, v, mxDt, fn, err
	}

	return runes, v, mxDt, nil, err
}

func (tree *ParserT) parseVarScalar(execScalars bool, strOrVal varFormatting) ([]rune, interface{}, string, error) {
	if tree.nextChar() == '(' {
		tree.charPos++
		return tree.parseVarParenthesis(execScalars, strOrVal)
	}

	if !isBareChar(tree.nextChar()) {
		// always print $
		return []rune{'$'}, "$", types.String, nil
	}

	tree.charPos++
	value := tree.parseBareword()

	if tree.charPos < len(tree.expression) && tree.expression[tree.charPos] == '[' {
		return tree.parseVarIndexElement(execScalars, '$', value, strOrVal)
	}

	tree.charPos--

	r := append([]rune{'$'}, value...)

	// don't getVar() until we come to execute the expression, skip when only
	// parsing syntax
	if execScalars {
		v, dataType, err := tree.getVar(value, strOrVal)
		return r, v, dataType, err
	}

	return r, nil, "", nil
}

func (tree *ParserT) parseVarParenthesis(exec bool, strOrVal varFormatting) ([]rune, interface{}, string, error) {
	start := tree.charPos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case r == ')':
			goto endParenthesis
		}
	}

	return nil, nil, "", raiseError(
		tree.expression, nil, tree.charPos, "expecting closing parenthesis, ')', after variable reference")

endParenthesis:
	path := tree.expression[start+1 : tree.charPos]
	value := tree.expression[start-1 : tree.charPos+1]

	if !exec {
		return value, nil, "", nil
	}

	v, dt, err := tree.getVar(path, strOrVal)
	if err != nil {
		return value, nil, "", err
	}
	return value, v, dt, nil
}

func (tree *ParserT) parseVarIndexElement(exec bool, sigil rune, varName []rune, strOrVal varFormatting) ([]rune, interface{}, string, error) {
	if tree.nextChar() == '{' {
		//return tree.parseLambdaScala(exec, '$', varName, strOrVal)
		//return tree.parseLambdaStatement(exec, '$')
		if exec {
			return tree.parseLambdaExecTrue(varName, sigil, "")
		}
		r, err := tree.parseLambdaExecFalse(sigil)
		value := append([]rune{sigil}, varName...)
		value = append(value, r...)
		value = append(value, ']')
		var fn parseLambdaExecTrueT = func() ([]rune, any, string, error) {
			return _parseLambdaExecTrue(tree.p, varName, sigil, "", value, r[1:], tree.StrictArrays())
		}
		return value, fn, "", err
	}

	var (
		brackets = 1
		escape   bool
		getIorE  = new(getVarIndexOrElementT)
	)

	start := tree.charPos

	if tree.nextChar() == '[' {
		brackets++
		tree.charPos++
	}

	tree.charPos++

	getIorE.isIorE = brackets

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

	return nil, nil, "", raiseError(
		tree.expression, nil, tree.charPos, "missing closing bracket ']'")

endIndexElement:
	value := tree.expression[start-len(varName)-1 : tree.charPos+1]
	getIorE.varName = varName
	getIorE.key = tree.expression[start+getIorE.isIorE : tree.charPos-getIorE.isIorE+1]
	getIorE.strOrVal = strOrVal

	if !exec {
		return value, getIorE, "", nil
	}

	v, dt, err := tree.getVarIndexOrElement(getIorE)
	if err != nil {
		return nil, nil, "", err
	}
	return nil, v, dt, nil
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
			r, v, _, err := tree.parseLambdaExecTrue(varName, '@', "") // TODO: test me
			return r, v, err
		} else {
			// just parsing source
			//r, _, err := tree.parseSubShell(false, '@', varAsValue)
			r, err := tree.parseLambdaExecFalse('@')
			return r, nil, err

			//r, v, _, err := tree.parseLambdaScala(false, '@', varName, varAsValue) // just parsing source
			//return r, v, err
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

func isUserNameChar(r rune) bool {
	return isBareChar(r) || r == '.' || r == '-'
}

func (tree *ParserT) parseVarTilde(exec bool) (string, error) {
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
	username := string(tree.expression[start:tree.charPos])
	tree.charPos--

	if !exec {
		return "~" + username, nil
	}

	if len(username) == 0 {
		return home.MyDir, nil
	}

	usr, err := user.Lookup(username)
	if err != nil {
		return "", fmt.Errorf("cannot expand variable `~%s`: %s", username, err.Error())
	}

	return usr.HomeDir, nil
}
