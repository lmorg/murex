package runmode_test

import (
	"testing"

	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/test/count"
)

func TestRunModeStringer(t *testing.T) {
	count.Tests(t, 5)

	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error("Not all constants have been stringified")
		}
	}()

	t.Log(runmode.Normal.String())
	t.Log(runmode.Evil.String())
	t.Log(runmode.BlockTry.String())
	t.Log(runmode.BlockTryPipe.String())
	t.Log(runmode.FunctionTry.String())
	t.Log(runmode.FunctionTryPipe.String())
	t.Log(runmode.ModuleTry.String())
	t.Log(runmode.ModuleTryPipe.String())
}

func TestRunModeIsStrict(t *testing.T) {
	tests := []struct {
		RunMode runmode.RunMode
		Strict  bool
	}{
		{runmode.Normal, false},
		{runmode.Evil, false},
		{runmode.BlockTry, true},
		{runmode.BlockTryPipe, true},
		{runmode.FunctionTry, true},
		{runmode.FunctionTryPipe, true},
		{runmode.ModuleTry, true},
		{runmode.ModuleTryPipe, true},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.Strict != test.RunMode.IsStrict() {
			t.Errorf("Unexpected return from IsStrict in test %d", i)
			t.Logf("  RunMode:  %d (%s)", test.RunMode, test.RunMode.String())
			t.Logf("  Expected: %v", test.Strict)
			t.Logf("  Actual:   %v", test.RunMode.IsStrict())
		}
	}
}

func TestRunModeIsScopeFunction(t *testing.T) {
	tests := []struct {
		RunMode runmode.RunMode
		Strict  bool
	}{
		{runmode.Normal, false},
		{runmode.Evil, false},
		{runmode.BlockTry, true},
		{runmode.BlockTryPipe, true},
		{runmode.FunctionTry, false},
		{runmode.FunctionTryPipe, false},
		{runmode.ModuleTry, true},
		{runmode.ModuleTryPipe, true},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.Strict != test.RunMode.IsBlockOrModule() {
			t.Errorf("Unexpected return from IsScopeFunction in test %d", i)
			t.Logf("  RunMode:  %d (%s)", test.RunMode, test.RunMode.String())
			t.Logf("  Expected: %v", test.Strict)
			t.Logf("  Actual:   %v", test.RunMode.IsBlockOrModule())
		}
	}
}
