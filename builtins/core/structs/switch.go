package structs

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/humannumbers"
)

func init() {
	lang.GoFunctions["switch"] = cmdSwitch
}

const (
	errSwitchParameters = "%s parameters supplied. Please read the `switch` docs for how to use. eg `murex-docs switch`"
	errReferToDocs      = "Please read the `switch` docs for how to use. eg `murex-docs switch`"
)

func cmdSwitch(p *lang.Process) error {
	switch p.Parameters.Len() {
	case 0:
		return fmt.Errorf(errSwitchParameters, "No")
	case 1:
		return switchLogic(p, false, "")
	case 2:
		return switchLogic(p, true, p.Parameters.Params[0])
	default:
		return fmt.Errorf(errSwitchParameters, "Too many")
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

	ast, pErr := lang.ParseBlock(block)
	if pErr.Code != 0 {
		return errors.New(pErr.Message)
	}

	var prevIfPassed bool

	for i := range *ast {
		params := &parameters.Parameters{Tokens: (*ast)[i].ParamTokens}
		lang.ParseParameters(p, params)

		switch (*ast)[i].Name {
		case "if", "case":
			err, caseIf, thenBlock := validateStatementParameters(ast, params, i, byVal)
			if err != nil {
				return err
			}

			var pass bool
			if byVal {
				pass, err = compareConditional(p, val, caseIf)
				if err != nil {
					return fmt.Errorf("Error comparing %s statement, %s conditional: %s",
						humannumbers.Ordinal(i+1), (*ast)[i].Name, err.Error())
				}
			} else {
				pass, err = executeConditional(p, caseIf)
				if err != nil {
					return fmt.Errorf("Error executing %s statement, %s conditional: %s",
						humannumbers.Ordinal(i+1), (*ast)[i].Name, err.Error())
				}
			}

			if pass {
				err = executeThen(p, thenBlock)
				if err != nil {
					return fmt.Errorf("Error executing %s statement, then block: %s",
						humannumbers.Ordinal(i+1), err.Error())
				}

				switch (*ast)[i].Name {
				case "if":
					prevIfPassed = true
					continue
				case "case":
					return nil
				}
			}

		case "catch":
			if prevIfPassed {
				return nil
			}

			err, _, thenBlock := validateStatementParameters(ast, params, i, byVal)
			if err != nil {
				return err
			}

			err = executeThen(p, thenBlock)
			if err != nil {
				return fmt.Errorf("Error executing %s statement, catch block: %s",
					humannumbers.Ordinal(i+1), err.Error())
			}

			return nil
		}

	}

	if !prevIfPassed {
		p.ExitNum = 1
	}

	return nil
}

func validateStatementParameters(ast *lang.AstNodes, params *parameters.Parameters, i int, byVal bool) (error, []rune, []rune) {
	var adjust int

	switch (*ast)[i].Name {
	case "if", "case":
		switch params.Len() {
		case 0:
			return fmt.Errorf("Missing parameters for %s statement (%s)\n%s",
				humannumbers.Ordinal(i+1), (*ast)[i].Name, errReferToDocs), nil, nil
		case 1:
			return fmt.Errorf("Too few parameters for %s statement (%s)\nExpected: conditional then { code block }\nFound: %s\n%s",
				humannumbers.Ordinal(i+1), (*ast)[i].Name, params.StringAll(), errReferToDocs), nil, nil

		case 3:
			if params.Params[1] != "then" {
				return fmt.Errorf("Too many parameters for %s statement (%s) or typo in statements. Expecting 'then' statement but found: '%s'\n%s",
					humannumbers.Ordinal(i+1), (*ast)[i].Name, params.StringAll(), errReferToDocs), nil, nil
			}
			adjust = 1
			fallthrough

		case 2:
			thenBlock, err := params.Block(1 + adjust)
			if err != nil {
				return fmt.Errorf("Cannot compile %s statement (%s): %s\nExpecting code block, found: '%s'",
					humannumbers.Ordinal(i+1), (*ast)[i].Name, err.Error(), params.Params[1+adjust]), nil, nil
			}

			if byVal {
				return nil, []rune(params.Params[0]), thenBlock
			}

			caseIf, err := params.Block(0)
			if err != nil {
				return fmt.Errorf("Cannot compile %s statement (%s): %s\nExpecting %s conditional block, found: '%s'",
					humannumbers.Ordinal(i+1), (*ast)[i].Name, err.Error(), (*ast)[i].Name, params.Params[0]), nil, nil
			}
			return nil, caseIf, thenBlock

		default:
			return fmt.Errorf("Too many parameters for %s statement (%s)\nFound: '%s'\n%s",
				humannumbers.Ordinal(i+1), (*ast)[i].Name, params.StringAll(), errReferToDocs), nil, nil
		}

	case "catch":
		switch params.Len() {
		case 0:
			return fmt.Errorf("Missing parameters for %s statement (%s)\n%s",
				humannumbers.Ordinal(i+1), (*ast)[i].Name, errReferToDocs), nil, nil

		case 1:
			thenBlock, err := params.Block(0)
			if err != nil {
				return fmt.Errorf("Cannot compile %s statement (%s): %s\nExpecting code block, found: '%s'",
					humannumbers.Ordinal(i+1), (*ast)[i].Name, err.Error(), params.Params[0]), nil, nil
			}
			return nil, nil, thenBlock

		default:
			return fmt.Errorf("Too many parameters for %s statement (%s)\nFound: '%s'\n%s",
				humannumbers.Ordinal(i+1), (*ast)[i].Name, params.StringAll(), errReferToDocs), nil, nil
		}

	default:
		return fmt.Errorf("Invalid %s statement '%s'", humannumbers.Ordinal(i+1), (*ast)[i].Name), nil, nil
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
