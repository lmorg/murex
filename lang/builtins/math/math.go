package math

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["math"] = proc.GoFunction{Func: cmdMath, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdMath(p *proc.Process) (err error) {

	if p.Method == true {

	}

}

func parseFormula(formula []byte) {
	type fTree struct {
		symbol string
	}
	for i, b := range formula {
		switch b {
		case ",":
			// do nothing
		case ".":
		case "/", "*", "-", "+", "(", ")":
		}
	}
}
