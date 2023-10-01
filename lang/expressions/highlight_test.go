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

func TestSingleQuotes(t *testing.T) {
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

func TestDoubleQuotes(t *testing.T) {
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

func TestBraceQuotesExpr(t *testing.T) {
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

func TestBraceQuotesStr(t *testing.T) {
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
