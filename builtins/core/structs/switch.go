package structs

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["switch"] = cmdSwitch
}

func cmdSwitch(p *lang.Process) error {
	compLeft, block, err := getParameters(p)
	if err != nil {
		return err
	}

	ast, pErr := lang.ParseBlock(block)
	if pErr.Code != 0 {
		return errors.New(pErr.Message)
	}

	for i := range ast {
		if p.HasCancelled() {
			return errors.New(errCancelled)
		}

		if ast[i].Name != "case" && ast[i].Name != "if" && ast[i].Name != "catch" {
			return fmt.Errorf("Missing `if`, `case` or `catch` statement after statement %d", i)
		}

		minParamLen := 2

		if ast[i].Name == "catch" {
			minParamLen = 1
		}

		params := &parameters.Parameters{Tokens: ast[i].ParamTokens}
		lang.ParseParameters(p, params)

		if params.Len() > 1 {
			if then, _ := params.String(params.Len() - 2); then == "then" {
				minParamLen++
			}
		}

		if params.Len() < minParamLen {
			b, _ := json.Marshal(ast, true)
			return fmt.Errorf("`%s` %d is missing a comparison or execution block:%s %s %s%sParameters found: %d%sParameters expected: %d%s%s",
				ast[i].Name, i+1, utils.NewLineString,
				ast[i].Name, params.StringAll(), utils.NewLineString,
				params.Len(), utils.NewLineString, minParamLen, utils.NewLineString,
				string(b),
			)
		}
		if params.Len() > minParamLen {
			b, _ := json.Marshal(ast, true)
			return fmt.Errorf("`%s` %d has too many parameters:%s %s %s%sParameters found: %d%sParameters expected: %d%s%s",
				ast[i].Name, i+1, utils.NewLineString,
				ast[i].Name, params.StringAll(), utils.NewLineString,
				params.Len(), utils.NewLineString, minParamLen, utils.NewLineString,
				string(b),
			)
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
			message := fmt.Sprintf("Error in block %d: %s", i+1, err.Error())
			//ansi.Stderrln(p, ansi.FgRed, message)
			p.Stderr.Writeln([]byte(message))
		}
	}

	p.ExitNum = 1
	return nil
}

func getParameters(p *lang.Process) (compLeft string, block []rune, err error) {
	switch p.Parameters.Len() {
	case 0:
		err = errors.New("Too few parameters")
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
		err = errors.New("Too many parameters")
		return
	}

	return
}

func getCompLeftFromBlock(p *lang.Process) (string, error) {
	compBlock, err := p.Parameters.Block(0)
	if err != nil {
		return "", err
	}
	//stdout := streams.NewStdin()
	//_, err = lang.RunBlockExistingConfigSpace(compBlock, nil, stdout, nil, p)
	fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	_, err = fork.Execute(compBlock)
	if err != nil {
		return "", err
	}
	b, err := fork.Stdout.ReadAll()
	return string(b), err
}

func switchCompByVal(p *lang.Process, params *parameters.Parameters, compLeft string) (bool, error) {
	compRight, err := params.String(0)
	if err != nil {
		return false, err
	}

	if types.IsBlock([]byte(compRight)) {
		return switchCompByBlock(p, params)
	}

	return compLeft == compRight, nil
}

func switchCompByBlock(p *lang.Process, params *parameters.Parameters) (bool, error) {
	block, err := params.Block(0)
	if err != nil {
		return false, err
	}

	//stdout := streams.NewStdin()
	//exitNum, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, nil, p)
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

func switchBlock(p *lang.Process, params *parameters.Parameters) error {
	block, err := params.Block(params.Len() - 1)
	if err != nil {
		return err
	}

	//_, err = lang.RunBlockExistingConfigSpace(block, nil, p.Stdout, p.Stderr, p)
	_, err = p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN).Execute(block)
	return err
}
