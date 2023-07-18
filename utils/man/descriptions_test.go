package man

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestParseLineFlags(t *testing.T) {
	tests := []string{
		`-a`,
		`--foobar`,
		`-a, --foobar`,
		`-f FILE`,
		`-f FILE, --file=FILE`,
		`-e PATTERNS, --regexp=PATTERNS`,
		`-E, --extended-regexp`,
		`--exclude-from=FILE`,
		`--backup[=CONTROL]`,
		`-R, -r, --recursive`,
		`--list-cmds=group[,group...]`,
		`--exec-path[=<path>]`,
		`--config-env=<name>=<envvar>`,
		`--[no]-help`,
		`--help Output a usage message and exit.`,
	}

	length := len(tests)
	for i := 0; i < length; i++ {
		tests = append(tests, tests[i]+" An autogenerated description")
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		pl := parseLineFlags([]byte(test))
		switch {
		case pl.Position == len(test):
			// success
		case pl.Description != "":
			// success
		default:
			t.Errorf("Could not match %s in test %d: len(test)==%d, result==%d", test, i, len(test), pl.Position)
			if pl.Position < len(test) {
				t.Logf("  scanned so far: '%s'", test[:pl.Position])
			}
		}
	}
}