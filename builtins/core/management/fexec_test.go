package management_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestFexec(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `fexec help -> [help]`,
			Stdout: "help",
		},
		{
			Block:  `fexec builtin out "foobar"`,
			Stdout: "foobar",
		},
		{
			Block: fmt.Sprintf(`
				function pub.%s { out "foobar" }
				fexec function pub.%s`,
				t.Name(), t.Name(),
			),
			Stdout: "foobar",
		},
		{
			Block: fmt.Sprintf(`
				private pvt.%s { out "foobar" }
				fexec private /murex/%s/pvt.%s; runtime --privates`,
				t.Name(), t.Name(), t.Name(),
			),
			Stdout: "foobar",
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestFexecErrors(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `fexec khlkwajfhldskjfhasdlkjfhaskdljfhasd`,
			Stderr:  "invalid flag",
			ExitNum: 1,
		},
		{
			Block:   `fexec builtin dslahfaksdljhfkasdjhflsdjahf`,
			Stderr:  "no builtin",
			ExitNum: 1,
		},
		{
			Block:   `fexec function dslahfaksdljhfkasdjhflsdjahf`,
			Stderr:  "cannot locate function",
			ExitNum: 1,
		},
		{
			Block:   `fexec private dslahfaksdljhfkasdjhflsdjahf`,
			Stderr:  "no private functions",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
