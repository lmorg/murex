package expressions_test

import (
	"regexp"
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/lang/expressions/node"
	"github.com/lmorg/murex/test/count"
)

var rxAnsiSgr = regexp.MustCompile(`\x1b\[([0-9]{1,2}(;[0-9]{1,2})*)?[m|K]`)

func testSyntaxHighlighter(t *testing.T, tests []string) {
	t.Helper()
	count.Tests(t, len(tests))

	lang.ShellProcess.Config = config.InitConf
	node.InitialiseDefaultTheme()

	for i := range tests {
		r := expressions.SyntaxHighlight([]rune(tests[i]))
		actual := rxAnsiSgr.ReplaceAllString(string(r), "")

		if tests[i] != actual {
			t.Errorf("Mismatch in test %d:", i)
			t.Logf("  Expected: '%s'", tests[i])
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}

func TestHlSingleQuotes(t *testing.T) {
	testSyntaxHighlighter(t, []string{
		`echo 'hello world`,
		`echo hello world'`,
		`echo 'hello world'`,

		`echo '$hello world`,
		`echo $hello world'`,
		`echo '$hello world'`,

		`echo 'hello $world`,
		`echo hello $world'`,
		`echo 'hello $world'`,

		`echo '~hello world`,
		`echo ~hello world'`,
		`echo '~hello world'`,

		`echo 'hello ~world`,
		`echo hello ~world'`,
		`echo 'hello ~world'`,
	})
}

func TestHlDoubleQuotes(t *testing.T) {
	testSyntaxHighlighter(t, []string{
		`echo "hello world`,
		`echo hello world"`,
		`echo "hello world"`,

		`echo "$hello world`,
		`echo $hello world"`,
		`echo "$hello world"`,

		`echo "hello $world`,
		`echo hello $world"`,
		`echo "hello $world"`,

		`echo "~hello world`,
		`echo ~hello world"`,
		`echo "~hello world"`,

		`echo "hello ~world`,
		`echo hello ~world"`,
		`echo "hello ~world"`,
	})
}

func TestHlBraceQuotesExpr(t *testing.T) {
	testSyntaxHighlighter(t, []string{
		`echo (hello world`,
		`echo hello world)`,
		`echo (hello world)`,

		`echo ($hello world`,
		`echo $hello world)`,
		`echo ($hello world)`,

		`echo (hello $world`,
		`echo hello $world)`,
		`echo (hello $world)`,

		`echo (~hello world`,
		`echo ~hello world)`,
		`echo (~hello world)`,

		`echo (hello ~world`,
		`echo hello ~world)`,
		`echo (hello ~world)`,
	})
}

func TestHlBraceQuotesStr(t *testing.T) {
	testSyntaxHighlighter(t, []string{
		`echo %(hello world`,
		`echo hello world)`,
		`echo %(hello world)`,

		`echo %($hello world`,
		`echo $hello world)`,
		`echo %($hello world)`,

		`echo %(hello $world`,
		`echo hello $world)`,
		`echo %(hello $world)`,

		`echo %(~hello world`,
		`echo ~hello world)`,
		`echo %(~hello world)`,

		`echo %(hello ~world`,
		`echo hello ~world)`,
		`echo %(hello ~world)`,
	})
}

func TestHlExpression(t *testing.T) {
	testSyntaxHighlighter(t, []string{
		`5*5`, `5 * 5`,
		`1=2+3`, `1 = 2 + 3`,
		`$foo='bar'`, `$foo = ' bar '`,
		`$foo="bar"`, `$foo = " bar "`,
		`$foo=%(bar)`, `$foo = %( bar )`,
	})
}

func TestHlArray(t *testing.T) {
	testSyntaxHighlighter(t, []string{
		`%[]`, `%[ ]`,
		`%[true]`, `%[false]`, `%[null]`,
		`%[1 2 3]`, `%[1,2,3]`, `%[1, 2, 3]`, `%[ 1 2 3 ]`,
		`%[-1 -2 -3]`, `%[-1,-2,-3]`, `%[-1, -2, -3]`, `%[ -1 -2 -3 ]`,
		`%[a b c]`, `%[a,b,c]`, `%[a, b, c]`, `%[ a b c ]`,
		`%[-a -b -c]`, `%[-a,-b,-c]`, `%[-a, -b, -c]`, `%[ -a -b -c ]`,
		`%['a' 'b' 'b']`, `%['a','b','c']`, `%['a', 'b', 'c']`, `%[ 'a' 'b' 'b' ]`,
		`%["a" "b" "b"]`, `%["a","b","c"]`, `%["a", "b", "c"]`, `%[ "a" "b" "b" ]`,
	})
}

func TestHlArrayNested(t *testing.T) {
	testSyntaxHighlighter(t, []string{
		`%[1 %[1 2 3] 3]`,
	})
}

func TestHlArrayStatement(t *testing.T) {
	testSyntaxHighlighter(t, []string{
		`echo %[]`, `echo %[ ]`,
		`echo %[1 2 3]`,
		`echo %[1 %[1 2 3] 3]`,
	})
}
