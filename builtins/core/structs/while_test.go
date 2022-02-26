package structs_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	_ "github.com/lmorg/murex/builtins/core/structs"
	_ "github.com/lmorg/murex/builtins/core/typemgmt"
	"github.com/lmorg/murex/test"
)

func TestWhileStdoutEvaluated(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				let: i=0
				while {
					let: i++
					= i<5
				}
				out: $i`,
			Stdout: "truetruetruetruefalse5\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestWhileConditionalEvaluated(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				let: i=0
				while { = i<5 } {
					let: i++
				}
				out: $i`,
			Stdout: "5\n",
		},
	}

	test.RunMurexTests(tests, t)
}
