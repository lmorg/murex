package background_test

import (
	"testing"

	"github.com/lmorg/murex/lang/background"
	"github.com/lmorg/murex/test/count"
)

func TestBackground(t *testing.T) {
	count.Tests(t, 6)

	bg := new(background.Background)

	bg.Set(false)
	if bg.Get() {
		t.Errorf("Set and/or Get failed. Returned true but expected false")
	}
	if bg.String() != "no" {
		t.Errorf("String() didn't return 'no'")
	}

	bg.Set(true)
	if !bg.Get() {
		t.Errorf("Set and/or Get failed. Returned false but expected true")
	}
	if bg.String() != "yes" {
		t.Errorf("String() didn't return 'yes'")
	}
}
