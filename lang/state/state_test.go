package state_test

import (
	"testing"

	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/test/count"
)

func TestStateVar(t *testing.T) {
	count.Tests(t, 4)

	var s state.State
	fs := s.Get()
	if fs != state.Undefined {
		t.Errorf("Invalid initialised value: %s", fs.String())
	}

	s.Set(state.Executing)

	fs = s.Get()
	if fs != state.Executing {
		t.Errorf("Invalid executing value: %s", fs.String())
	}

	if s.String() != fs.String() {
		t.Errorf("Stringer mismatch: '%s' != '%s'", s.String(), fs.String())
	}
}

func TestStateNew(t *testing.T) {
	count.Tests(t, 4)

	s := new(state.State)
	fs := s.Get()
	if fs != state.Undefined {
		t.Errorf("Invalid initialised value: %s", fs.String())
	}

	s.Set(state.Executing)

	fs = s.Get()
	if fs != state.Executing {
		t.Errorf("Invalid executing value: %s", fs.String())
	}

	if s.String() != fs.String() {
		t.Errorf("Stringer mismatch: '%s' != '%s'", s.String(), fs.String())
	}
}
