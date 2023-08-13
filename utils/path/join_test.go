package path_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/path"
)

func TestJoin(t *testing.T) {
	slash := consts.PathSlash

	tests := []struct {
		Slice    []string
		Expected string
	}{
		{
			Slice:    []string{slash},
			Expected: slash,
		},
		{
			Slice:    []string{"foo", "bar"},
			Expected: fmt.Sprintf("foo%sbar", slash),
		},
		{
			Slice:    []string{"/", "foo", "bar"},
			Expected: fmt.Sprintf("%sfoo%sbar", slash, slash),
		},
		{
			Slice:    []string{".", "foo", "bar"},
			Expected: fmt.Sprintf(".%sfoo%sbar", slash, slash),
		},
		{
			Slice:    strings.Split(home.MyDir, slash),
			Expected: home.MyDir + slash,
		},
		{
			Slice:    strings.Split(home.MyDir, slash)[1:],
			Expected: home.MyDir[1:],
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := path.Join(test.Slice)

		if test.Expected != actual {
			t.Errorf("expected != actual in test %d", i)
			t.Logf("  Slice:    %s", json.LazyLogging(test.Slice))
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}
