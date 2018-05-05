package structs

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
)

func init() {
	proc.GoFunctions["switch"] = cmdSwitch
}

func cmdSwitch(p *proc.Process) error {
	compLeft, block, err := getParameters(p)
	if err != nil {
		return err
	}

	ast, pErr := lang.ParseBlock(block)
	if pErr.Code != 0 {
		return errors.New(pErr.Message)
	}

	paramLen := 2

	for i := range ast {
		if ast[i].Name != "case" && ast[i].Name != "if" && ast[i].Name != "catch" {
			return fmt.Errorf("Missing `if`, `case` or `catch` statement after statement %d", i)
		}

		if ast[i].Name == "catch" {
			paramLen--
		}

		params := &parameters.Parameters{Tokens: ast[i].ParamTokens}
		lang.ParseParameters(p, params)

		if params.Len() < paramLen {
			return fmt.Errorf("`%s` %d is missing a comparison or execution block.", ast[i].Name, i+1)
		}
		if params.Len() > paramLen {
			return fmt.Errorf("`%s` %d has too many parameters.", ast[i].Name, i+1)
		}

		var result bool
		switch {
		case ast[i].Name == "catch":
			result = true
		case compLeft == "":
			result, err = switchCompByBlock(p, params)
		default:
			result, err = switchCompByVal(p, params, compLeft)
		}
		if err != nil {
			return err
		}

		if !result {
			continue
		}

		err = switchBlock(p, params)
		if ast[i].Name == "case" || ast[i].Name == "catch" {
			return err
		} else if err != nil {
			message := fmt.Sprintln("Error in block %d: %s", i+1, err.Error())
			ansi.Stderrln(p, ansi.FgRed, message)
		}
	}

	p.ExitNum = 1
	return nil
}

func getParameters(p *proc.Process) (compLeft string, block []rune, err error) {
	switch p.Parameters.Len() {
	case 0:
		err = errors.New("Too few parameters.")
		return

	case 1:
		block, err = p.Parameters.Block(0)
		if err != nil {
			return
		}

	case 2:
		compLeft, err = p.Parameters.String(0)
		if err != nil {
			return
		}

		if types.IsBlock([]byte(compLeft)) {
			compLeft, err = getCompLeftFromBlock(p)
			if err != nil {
				return
			}
		}

		block, err = p.Parameters.Block(1)
		if err != nil {
			return
		}

	default:
		err = errors.New("Too many parameters.")
		return
	}

	return
}

func getCompLeftFromBlock(p *proc.Process) (string, error) {
	compBlock, err := p.Parameters.Block(0)
	if err != nil {
		return "", err
	}
	stdout := streams.NewStdin()
	_, err = lang.RunBlockExistingConfigSpace(compBlock, nil, stdout, nil, p)
	if err != nil {
		return "", err
	}
	b, err := stdout.ReadAll()
	return string(b), err
}

func switchCompByVal(p *proc.Process, params *parameters.Parameters, compLeft string) (bool, error) {
	compRight, err := params.String(0)
	if err != nil {
		return false, err
	}

	if types.IsBlock([]byte(compRight)) {
		return switchCompByBlock(p, params)
	}

	return compLeft == compRight, nil
}

func switchCompByBlock(p *proc.Process, params *parameters.Parameters) (bool, error) {
	block, err := params.Block(0)
	if err != nil {
		return false, err
	}

	stdout := streams.NewStdin()
	exitNum, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, nil, p)
	if err != nil {
		return false, err
	}

	b, err := stdout.ReadAll()
	if err != nil {
		return false, err
	}

	result := types.IsTrue(b, exitNum)
	return result, nil
}

func switchBlock(p *proc.Process, params *parameters.Parameters) error {
	block, err := params.Block(params.Len() - 1)
	if err != nil {
		return err
	}

	_, err = lang.RunBlockExistingConfigSpace(block, nil, p.Stdout, p.Stderr, p)
	return err
}
