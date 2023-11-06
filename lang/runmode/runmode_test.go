package runmode_test

import (
	"testing"

	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/test/count"
)

func TestRunModeStringer(t *testing.T) {
	count.Tests(t, 16)

	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error("Not all constants have been stringified")
		}
	}()

	t.Log(runmode.Normal.String())

	t.Log(runmode.BlockUnsafe.String())
	t.Log(runmode.FunctionUnsafe.String())
	t.Log(runmode.ModuleUnsafe.String())

	t.Log(runmode.BlockTry.String())
	t.Log(runmode.BlockTryPipe.String())
	t.Log(runmode.BlockTryErr.String())
	t.Log(runmode.BlockTryPipeErr.String())

	t.Log(runmode.FunctionTry.String())
	t.Log(runmode.FunctionTryPipe.String())
	t.Log(runmode.FunctionTryErr.String())
	t.Log(runmode.FunctionTryPipeErr.String())

	t.Log(runmode.ModuleTry.String())
	t.Log(runmode.ModuleTryPipe.String())
	t.Log(runmode.ModuleTryErr.String())
	t.Log(runmode.ModuleTryPipeErr.String())
}

func TestRunModeIsStrict(t *testing.T) {
	tests := []struct {
		RunMode runmode.RunMode
		Strict  bool
	}{
		{runmode.Normal, false},

		{runmode.BlockUnsafe, false},
		{runmode.FunctionUnsafe, false},
		{runmode.ModuleUnsafe, false},

		{runmode.BlockTry, true},
		{runmode.BlockTryPipe, true},
		{runmode.BlockTryErr, true},
		{runmode.BlockTryPipeErr, true},

		{runmode.FunctionTry, true},
		{runmode.FunctionTryPipe, true},
		{runmode.FunctionTryErr, true},
		{runmode.FunctionTryPipeErr, true},

		{runmode.ModuleTry, true},
		{runmode.ModuleTryPipe, true},
		{runmode.ModuleTryErr, true},
		{runmode.ModuleTryPipeErr, true},
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
