package expressions

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestExpressionParserOffset0(t *testing.T) {
	tests := []string{
		"1+2;other code",
		"(1+2);other code",
		"foobar=1+2;other code",
		"foobar=(1+2);other code",
		"foobar=1+2;other code",
		"foobar=(1+(2+3)+4);other code",
		"TestExpressionsBuiltin6=1+1+1+1+(1+1)+1+1+1+1;code",
		"TestExpressionsBuiltin0=1+1+1+1+1+1+1+1+1+1",
		"1+1+1+1+1",
		"true==true",
		"1 + 1",
		"3*(3+1)",
		"bob=3*(3+1)",
		"%[apples oranges grapes]",
		"fruit=%[apples oranges grapes]",
		"$a==$b",
		"%[1,2,$a]",
		"%[1,2,@a]",
		"%[1,2,$aa]",
		"%[1,2,@aa]",
		"b=%[1,2,$a]",
		"b=%[1,2,@a]",
		"b=%[1,2,$aa]",
		"b=%[1,2,@aa]",
		"foobar=%[1,2,$a]",
		"foobar=%[1,2,@a]",
		"foobar=%[1,2,$aa]",
		"foobar=%[1,2,@aa]",
	}

	count.Tests(t, len(tests))

	for j := range tests {
		expression := []rune(tests[j][0:])
		split := strings.Split(tests[j], ";")

		i, err := ExpressionParser(expression, 0, false)
		if err != nil || string(expression[:i+1]) != split[0] {
			t.Errorf("Expression did not parse correctly in test %d:", j)
			t.Log("            :           1         2         3         4         5")
			t.Log("            :  12345678901234567890123456789012345678901234567890")
			t.Logf("  Expression: '%s'", string(expression))
			t.Logf("  Error:      %v", err)
			t.Logf("  Expected:   '%s'", split[0])
			t.Logf("  Actual:     '%s'", string(expression[:i+1]))
		}
	}
}

func TestExpressionParserOffset5(t *testing.T) {
	tests := []string{
		"code;1+2;other code",
		"code;(1+2);other code",
		"code;foobar=1+2;other code",
		"code;foobar=(1+2);other code",
		"code;foobar=1+2;other code",
		"code;foobar=(1+(2+3)+4);other code",
		"code;TestExpressionsBuiltin6=1+1+1+1+(1+1)+1+1+1+1;code",
		"code;TestExpressionsBuiltin0=1+1+1+1+1+1+1+1+1+1",
		"code;3*(3+1)",
		"code;bob=3*(3+1)",
		"code;%[apples oranges grapes]",
		"code;fruit=%[apples oranges grapes]",
		"code;$a==$b",
		"code;%[1,2,$a]",
		"code;%[1,2,@a]",
		"code;%[1,2,$aa]",
		"code;%[1,2,@aa]",
		"code;b=%[1,2,$a]",
		"code;b=%[1,2,@a]",
		"code;b=%[1,2,$aa]",
		"code;b=%[1,2,@aa]",
		"code;foobar=%[1,2,$a]",
		"code;foobar=%[1,2,@a]",
		"code;foobar=%[1,2,$aa]",
		"code;foobar=%[1,2,@aa]",
	}

	count.Tests(t, len(tests))

	for j := range tests {
		expression := []rune(tests[j][5:])
		split := strings.Split(tests[j], ";")

		i, err := ExpressionParser(expression, 5, false)
		if err != nil || string(expression[:i+1]) != split[1] {
			t.Errorf("Expression did not parse correctly in test %d:", j)
			t.Log("            :           1         2         3         4         5")
			t.Log("            :  12345678901234567890123456789012345678901234567890")
			t.Logf("  Expression: '%s'", string(expression))
			t.Logf("  Error:      %v", err)
			t.Logf("  Expected:   '%s'", split[1])
			t.Logf("  Actual:     '%s'", string(expression[:i+1]))
		}
	}
}
