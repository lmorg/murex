package time

import (
	"testing"
	"time"

	_ "github.com/lmorg/murex/builtins/core/expressions"
	_ "github.com/lmorg/murex/builtins/optional/time"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestTimeParams(t *testing.T) {
	start := time.Now()

	lang.ShellProcess = lang.NewTestProcess()
	test.RunMethodRegexTest(
		t, cmdTime, "time", //     method
		"", types.Null, //         stdin
		[]string{"sleep", "3"}, // parameters
		``,                     // output
	)

	switch {
	case start.Add(2 * time.Second).After(time.Now()):
		t.Error("Process didn't pause")
	case start.Add(4 * time.Second).Before(time.Now()):
		t.Error("Process paused for too long")
	}
}

func TestTimeBlock(t *testing.T) {
	start := time.Now()

	lang.ShellProcess = lang.NewTestProcess()
	test.RunMethodRegexTest(
		t, cmdTime, "time", //      method
		"", types.Null, //          stdin
		[]string{"{ sleep 3 }"}, // parameters
		``,                      // output
	)

	switch {
	case start.Add(2 * time.Second).After(time.Now()):
		t.Error("Process didn't pause")
	case start.Add(4 * time.Second).Before(time.Now()):
		t.Error("Process paused for too long")
	}
}

func TestTimeBlockBg(t *testing.T) {
	start := time.Now()

	lang.ShellProcess = lang.NewTestProcess()
	test.RunMethodRegexTest(
		t, cmdTime, "time", //             method
		"", types.Null, //                 stdin
		[]string{"{ bg { sleep 3 } }"}, // parameters
		``,                             // output
	)

	switch {
	case start.Add(1 * time.Second).Before(time.Now()):
		t.Error("PRocess should not have paused")
	}
}
