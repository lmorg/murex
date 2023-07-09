package parameters_test

import (
	"testing"

	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/test/count"
)

func TestGetString(t *testing.T) {
	count.Tests(t, 6)

	p := new(parameters.Parameters)
	p.DefineParsed([]string{"one"})

	var err error

	_, err = p.Block(0)
	if err == nil {
		t.Error("Error should have been raised: can't import string as block")
	}

	_, err = p.Bool(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.Byte(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.String(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.Int(0)
	if err == nil {
		t.Error("Error should have been raised: can't import string as int")
	}

	_, err = p.Uint32(0)
	if err == nil {
		t.Error("Error should have been raised: can't import string as uint32")
	}
}

func TestGetInt(t *testing.T) {
	count.Tests(t, 6)

	p := new(parameters.Parameters)
	p.DefineParsed([]string{"1"})

	var err error

	_, err = p.Block(0)
	if err == nil {
		t.Error("Error should have been raised: can't import int as block")
	}

	_, err = p.Bool(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.Byte(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.String(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.Int(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.Uint32(0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetUint32(t *testing.T) {
	count.Tests(t, 6)

	p := new(parameters.Parameters)
	p.DefineParsed([]string{"-1"})

	var err error

	_, err = p.Block(0)
	if err == nil {
		t.Error("Error should have been raised: can't import int as block")
	}

	_, err = p.Bool(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.Byte(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.String(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.Int(0)
	if err != nil {
		t.Error(err)
	}

	_, err = p.Uint32(0)
	if err == nil {
		t.Error("Error should have been raised: can't import int as uint32")
	}
}

// We don't test this because if this condition arises then it's because any Go
// code for murex builtins is wrong, so a panic is absolutely the right way to
// flush out those errors
/*func TestGetBoundsNeg1(t *testing.T) {
	count.Tests(t, 6)

	p := new(parameters.Parameters)
	p.DefineParsed([]string{"1"})

	var err error

	_, err = p.Block(-1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.Bool(-1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.Byte(-1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.String(-1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.Int(-1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.Uint32(-1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}
}*/

func TestGetBounds1(t *testing.T) {
	count.Tests(t, 6)

	p := new(parameters.Parameters)
	p.DefineParsed([]string{"1"})

	var err error

	_, err = p.Block(1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.Bool(1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.Byte(1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.String(1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.Int(1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}

	_, err = p.Uint32(1)
	if err == nil {
		t.Error("Out of bounds error should have been raised")
	}
}
