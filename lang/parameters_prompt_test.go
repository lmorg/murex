package lang

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/readline"
)

func TestAutoGlobPromptHintText(t *testing.T) {
	rl := readline.NewInstance()

	tests := []struct {
		Match []string
		Hint  string
	}{
		{
			Match: []string{},
			Hint:  warningNoGlobMatch,
		},
		{
			Match: []string{"foo", "bar"},
			Hint:  "foo bar",
		},
		{
			Match: []string{"foo bar"},
			Hint:  "foo\\ bar",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := string(autoGlobPromptHintText(rl, test.Match))

		if test.Hint != actual {
			t.Errorf("HintText doesn't match expected in test %d", i)
			t.Logf("  Match:     %s", json.LazyLogging(test.Match))
			t.Logf("  Expected: '%s'", test.Hint)
			t.Logf("  Actual:   '%s'", actual)
		}
	}

}
