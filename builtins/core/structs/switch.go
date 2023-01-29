package structs

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/humannumbers"
)

func init() {
	lang.DefineFunction("switch", cmdSwitch, types.Any)
}

const (
	errSwitchParameters = "%s parameters supplied. Please read the `switch` docs for how to use. eg `murex-docs switch`"
	errReferToDocs      = "Please read the `switch` docs for how to use. eg `murex-docs switch`"
)

func cmdSwitch(p *lang.Process) error {
	switch p.Parameters.Len() {
	case 0:
		return fmt.Errorf(errSwitchParameters, "no")
	case 1:
		return switchLogic(p, false, "")
	case 2:
		param, _ := p.Parameters.String(0)
		return switchLogic(p, true, param)
	default:
		return fmt.Errorf(errSwitchParameters, "too many")
	}
}

func switchLogic(p *lang.Process, byVal bool, val string) error {
	var loc int
	if byVal {
		loc = 1
	}

	block, err := p.Parameters.Block(loc)
	if err != nil {
		return err
	}

	swt, err := expressions.ParseSwitch(p, block)
	if err != nil {
		return err
	}

	var prevIfPassed bool

	for i, token := range swt {
		switch token.Condition {
		case "if", "case":
			caseIf, thenBlock, err := validateStatementParameters(token, i, byVal)
			if err != nil {
				return err
			}

			var pass bool
			if byVal {
				pass, err = compareConditional(p, val, caseIf)
				if err != nil {
					return fmt.Errorf("error comparing %s statement, %s conditional:\n%s",
						humannumbers.Ordinal(i+1), token.Condition, err.Error())
				}
			} else {
				pass, err = executeConditional(p, caseIf)
				if err != nil {
					return fmt.Errorf("error executing %s statement, %s conditional:\n%s",
						humannumbers.Ordinal(i+1), token.Condition, err.Error())
				}
			}

			if pass {
				err = executeThen(p, thenBlock)
				if err != nil {
					return fmt.Errorf("error executing %s statement, then block:\n%s",
						humannumbers.Ordinal(i+1), err.Error())
				}

				switch swt[i].Condition {
				case "if":
					prevIfPassed = true
					continue
				case "case":
					return nil
				}
			}

		case "default", "catch", "else":
			if prevIfPassed {
				return nil
			}

			_, thenBlock, err := validateStatementParameters(token, i, byVal)
			if err != nil {
				return err
			}

			err = executeThen(p, thenBlock)
			if err != nil {
				return fmt.Errorf("error executing %s statement, catch block:\n%s",
					humannumbers.Ordinal(i+1), err.Error())
			}

			return nil

		default:
			return fmt.Errorf("error executing %s statement, `%s` is not a valid statement.\nExpecting `case`, `if`, `catch`",
				humannumbers.Ordinal(i+1), token.Condition)
		}
	}

	if !prevIfPassed {
		p.ExitNum = 1
	}

	return nil
}

func validateStatementParameters(token *expressions.SwitchT, i int, byVal bool) ([]rune, []rune, error) {
	var adjust int

	switch token.Condition {
	case "if", "case":
		switch token.ParametersLen() {
		case 0:
			return nil, nil,
				fmt.Errorf("missing parameters for %s statement (%s)\n%s",
					humannumbers.Ordinal(i+1), token.Condition, errReferToDocs)
		case 1:
			return nil, nil,
				fmt.Errorf("too few parameters for %s statement (%s)\nExpected: conditional then { code block }\nFound: %s\n%s",
					humannumbers.Ordinal(i+1), token.Condition, token.ParametersStringAll(), errReferToDocs)

		case 3:
			if token.ParametersString(1) != "then" {
				return nil, nil,
					fmt.Errorf("too many parameters for %s statement (%s) or typo in statements.\nExpecting 'then' statement\nFound: '%s'\n%s",
						humannumbers.Ordinal(i+1), token.Condition, token.ParametersStringAll(), errReferToDocs)
			}
			adjust = 1
			fallthrough

		case 2:
			thenBlock, err := token.Block(1 + adjust)
			if err != nil {
				return nil, nil,
					fmt.Errorf("cannot compile %s statement (%s): %s\nExpecting code block, found: '%s'",
						humannumbers.Ordinal(i+1), token.Condition, err.Error(), token.ParametersString(1+adjust))
			}

			if byVal {
				return token.Parameters[0], thenBlock, nil
			}

			caseIf, err := token.Block(0)
			if err != nil {
				return nil, nil, fmt.Errorf("cannot compile %s statement (%s): %s\nExpecting %s conditional block, found: '%s'",
					humannumbers.Ordinal(i+1), token.Condition, err.Error(), token.Condition, token.ParametersString(0))
			}
			return caseIf, thenBlock, nil

		default:
			return nil, nil,
				fmt.Errorf("too many parameters for %s statement (%s)\nFound: '%s'\n%s",
					humannumbers.Ordinal(i+1), token.Condition, token.ParametersStringAll(), errReferToDocs)
		}

	case "catch", "default":
		switch token.ParametersLen() {
		case 0:
			return nil, nil, fmt.Errorf("missing parameters for %s statement (%s)\n%s",
				humannumbers.Ordinal(i+1), token.Condition, errReferToDocs)

		case 1:
			thenBlock, err := token.Block(0)
			if err != nil {
				return nil, nil,
					fmt.Errorf("cannot compile %s statement (%s): %s\nExpecting code block, found: '%s'",
						humannumbers.Ordinal(i+1), token.Condition, err.Error(), token.ParametersString(0))
			}
			return nil, thenBlock, nil

		default:
			return nil, nil,
				fmt.Errorf("too many parameters for %s statement (%s)\nFound: '%s'\n%s",
					humannumbers.Ordinal(i+1), token.Condition, token.ParametersStringAll(), errReferToDocs)
		}

	default:
		return nil, nil,
			fmt.Errorf("invalid %s statement '%s'", humannumbers.Ordinal(i+1), token.Condition)
	}
}

func compareConditional(p *lang.Process, val string, caseIf []rune) (bool, error) {
	if !types.IsBlockRune(caseIf) {
		return val == string(caseIf), nil
	}

	fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	_, err := fork.Execute(caseIf)
	if err != nil {
		return false, err
	}

	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return false, err
	}

	return val == string(utils.CrLfTrim(b)), err
}

func executeConditional(p *lang.Process, block []rune) (bool, error) {
	fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	exitNum, err := fork.Execute(block)
	if err != nil {
		return false, err
	}

	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return false, err
	}

	result := types.IsTrue(b, exitNum)
	return result, nil
}

func executeThen(p *lang.Process, block []rune) error {
	_, err := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN).Execute(block)
	return err
}
