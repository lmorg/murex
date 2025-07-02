package streams

import (
	"context"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestAppendBytes(t *testing.T) {
	tests := []struct {
		Slice    string
		Data     string
		Expected string
	}{
		{
			Slice:    "",
			Data:     "",
			Expected: "",
		},
		{
			Slice:    "",
			Data:     "bar",
			Expected: "bar",
		},
		{
			Slice:    "foo",
			Data:     "",
			Expected: "foo",
		},
		{
			Slice:    "foo",
			Data:     "bar",
			Expected: "foobar",
		},
		{
			Slice:    "foo",
			Data:     "barbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbar",
			Expected: "foobarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbarbar",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		b := appendBytes([]byte(test.Slice), []byte(test.Data)...)
		actual := string(b)
		if actual != test.Expected {
			t.Errorf("Actual does not match expected in test %d", i)
			t.Logf("  Slice:    '%s'", test.Slice)
			t.Logf("  Data:     '%s'", test.Data)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}

func TestSetDataType(t *testing.T) {
	count.Tests(t, 2)

	stream := NewStdin()

	stream.SetDataType("foo")
	dt := stream.GetDataType()
	if dt != "foo" {
		t.Errorf("Data type unset: '%s'", dt)
	}

	stream.SetDataType("bar")
	dt = stream.GetDataType()
	if dt != "foo" {
		t.Errorf("Data type unset or overwritten: '%s'", dt)
	}
}

func TestSetNewStream(t *testing.T) {
	count.Tests(t, 3)

	stream, err := newStream("")
	if err != nil {
		t.Error(err)
	}

	stream.SetDataType("foo")
	dt := stream.GetDataType()
	if dt != "foo" {
		t.Errorf("Data type unset: '%s'", dt)
	}

	stream.SetDataType("bar")
	dt = stream.GetDataType()
	if dt != "foo" {
		t.Errorf("Data type unset or overwritten: '%s'", dt)
	}
}

func TestSetOpenClose(t *testing.T) {
	count.Tests(t, 2)

	stream, err := newStream("")
	if err != nil {
		t.Error(err)
	}

	stream.Open()
	stream.Close()
}

func TestSetCloseError1(t *testing.T) {
	count.Tests(t, 1)

	stream, err := newStream("")
	if err != nil {
		t.Error(err)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("expecting a panic due to more closed dependant than open")
		}
	}()

	stream.Close()
}

func TestSetOpenCloseError(t *testing.T) {
	count.Tests(t, 3)

	stream, err := newStream("")
	if err != nil {
		t.Error(err)
	}

	stream.Open()
	stream.Close()

	defer func() {
		if r := recover(); r == nil {
			t.Error("expecting a panic due to more closed dependant than open")
		}
	}()

	stream.Close()
}

func TestSetForceClose(t *testing.T) {
	count.Tests(t, 3)

	stream, err := newStream("")
	if err != nil {
		t.Error(err)
	}

	stream.ForceClose()
}

func TestSetOpenForceClose(t *testing.T) {
	count.Tests(t, 3)

	stream, err := newStream("")
	if err != nil {
		t.Error(err)
	}

	stream.Open()
	stream.ForceClose()
}

func TestSetOpenForceCloseWithContext(t *testing.T) {
	count.Tests(t, 3)

	ctx := context.Background()
	v := "not ran"
	cancelFunc := func() {
		v = "have run"
	}

	stream := NewStdinWithContext(ctx, cancelFunc)

	stream.Open()
	stream.ForceClose()

	if v != "have run" {
		t.Error("cancelFunc not invoked when stream was force closed")
	}
}
