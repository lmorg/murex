package history

import (
	"testing"

	"github.com/lmorg/murex/lang"
	_ "github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/readline"
)

func newReadlineInstance() *readline.Instance {
	rl := readline.NewInstance()
	rl.History = NewTestHistory()

	return rl
}

func test(function func(string, *readline.Instance) (string, error), t *testing.T, tests, expected []string, rl *readline.Instance) {
	t.Helper()
	count.Tests(t, len(tests))

	lang.InitEnv()

	for i := range tests {
		actual, err := function(tests[i], rl)
		if actual != expected[i] {
			t.Errorf("Output does not match expected in test %d:", i)
			t.Logf("  Original: %s", tests[i])
			t.Logf("  Expected: %s", expected[i])
			t.Logf("  Actual:   %s", actual)
			t.Log("  eo bytes:  ", []byte(expected[i]))
			t.Log("  ao bytes:  ", []byte(actual))
			t.Logf("  Error:    %s", err)
		}
	}
}
func TestBangBang(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"^!!",
	}

	expected := []string{
		"out the lazy dog",
	}

	test(expandHistBangBang, t, tests, expected, rl)
}

func TestHistPrefix(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"^^out",
		"^^fox",
	}

	expected := []string{
		"out the lazy dog",
		"",
	}

	test(expandHistPrefix, t, tests, expected, rl)
}

func TestHistIndex(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"^-1",
		"^0",
		"^1",
		"^2",
		"^3",
		"^4",
	}

	expected := []string{
		"^-1",
		"out the quick brown #fox",
		"out jumped over",
		"out the lazy dog",
		"",
		"",
	}

	test(expandHistIndex, t, tests, expected, rl)
}

func TestHistRegex(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"^m/quick/",
		"^m/over/",
		"^m/dog/",
		"^m/cat/",
	}

	expected := []string{
		"out the quick brown #fox",
		"out jumped over",
		"out the lazy dog",
		"",
	}

	test(expandHistRegex, t, tests, expected, rl)
}

func TestHistHashTag(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"^#f",
		"^#fo",
		"^#fox",
		"^#over",
		"^#dog",
		"^#cat",
	}

	expected := []string{
		"",
		"",
		"out the quick brown ",
		"",
		"",
		"",
	}

	test(expandHistHashtag, t, tests, expected, rl)
}

/*func TestHistAllPs(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"^[-2][-2]",
		"^[-2][-1]",
		"^[-2][0]",
		"^[-2][1]",
		"^[-2][2]",
		"^[-2][3]",
		"^[-1][-2]",
		"^[-1][-1]",
		"^[-1][0]",
		"^[-1][1]",
		"^[-1][2]",
		"^[-1][3]",
		"^[0][-2]",
		"^[0][-1]",
		"^[0][0]",
		"^[0][1]",
		"^[0][2]",
		"^[0][3]",
		"^[1][-2]",
		"^[1][-1]",
		"^[1][0]",
		"^[1][1]",
		"^[1][2]",
		"^[1][3]",
		"^[2][-2]",
		"^[2][-1]",
		"^[2][0]",
		"^[2][1]",
		"^[2][2]",
		"^[2][3]",
	}

	expected := []string{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}

	test(expandHistAllPs, t, tests, expected, rl)
}*/

func TestHistParam(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"^[-5]",
		"^[-4]",
		"^[-3]",
		"^[-2]",
		"^[-1]",
		"^[-0]",
		"^[1]",
		"^[2]",
		"^[3]",
		"^[4]",
		"^[5]",
	}

	expected := []string{
		"",
		"out",
		"the",
		"lazy",
		"dog",
		"out",
		"the",
		"lazy",
		"dog",
		"",
		"",
	}

	test(expandHistParam, t, tests, expected, rl)
}

func TestHistReplace(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"out: the quick brown fox",
		"^s/dog/cat/",
		"out: the quick brown fox ^s/quick/slow/",
		"out: the quick brown fox ^s/fox/wolf/",
		"out: the quick brown fox ^s/out/err/",
	}

	expected := []string{
		"out: the quick brown fox",
		"",
		"out: the slow brown fox ",
		"out: the quick brown wolf ",
		"err: the quick brown fox ",
	}

	test(expandHistReplace, t, tests, expected, rl)
}

/*func TestHistRepParam(t *testing.T) {
	rl := newReadlineInstance()

	tests := []string{
		"out: the quick brown fox",
		"^s-1/dog/cat/",
		"out: the quick brown fox ^s0/quick/slow/",

	}

	expected := []string{
		"out: the quick brown fox",
		"out: the lazy cat",
		"out: the quick brown fox out: the slow brown fox",
	}

	test(expandHistRepParam, t, tests, expected, rl)
}*/
