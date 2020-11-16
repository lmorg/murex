package mxjson

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestQuote(t *testing.T) {
	count.Tests(t, 3)

	var q quote

	if q.IsOpen() {
		t.Error("q.IsOpen incorrectly returned true")
	}

	q.Open(1)

	if !q.IsOpen() {
		t.Error("q.IsOpen incorrectly returned false")
	}

	q.Close()

	if q.IsOpen() {
		t.Error("q.IsOpen incorrectly returned true")
	}
}

func TestPair(t *testing.T) {
	count.Tests(t, 4)

	var p pair

	if p.IsOpen() {
		t.Error("p.IsOpen incorrectly returned true")
	}

	p.Open(1)

	if !p.IsOpen() {
		t.Error("p.IsOpen incorrectly returned false")
	}

	err := p.Close()
	if err != nil {
		t.Errorf("p.Close incorrectly returned err: %s", err.Error())
	}

	if p.IsOpen() {
		t.Error("p.IsOpen incorrectly returned true")
	}

	err = p.Close()
	if err == nil {
		t.Error("p.Close incorrectly returned no err")
	}
}

func TestPairNested(t *testing.T) {
	count.Tests(t, 1)

	p := newPair()

	p.Open(1)
	p.Open(2)
	p.Open(3)
	if err := p.Close(); err != nil {
		t.Errorf("p.Close incorrectly returned err in iteration %d: %s", 0, err.Error())
	}
	if err := p.Close(); err != nil {
		t.Errorf("p.Close incorrectly returned err in iteration %d: %s", 1, err.Error())
	}
	if err := p.Close(); err != nil {
		t.Errorf("p.Close incorrectly returned err in iteration %d: %s", 2, err.Error())
	}
}

func TestPairOverflow(t *testing.T) {
	count.Tests(t, 1)

	p := newPair()
	max := cap(p.pos) * cap(p.pos)

	for i := 0; i < max; i++ {
		p.Open(i)
	}

	for i := 0; i < max; i++ {
		if err := p.Close(); err != nil {
			t.Errorf("p.Close incorrectly returned err in iteration %d: %s", 1, err.Error())
		}
	}
}

func TestLazyStringGet(t *testing.T) {
	nums := []string{
		"zero.",
		"one.",
		"two.",
		"three.",
		"four.",
		"five.",
		"six.",
		"seven.",
		"eight.",
		"nine.",
	}

	count.Tests(t, 1)

	s := newStr()
	max := cap(s.b)

	for i := 0; i < max; i++ {
		ints := strconv.Itoa(i)
		var new string
		for _, b := range []byte(ints) {
			new += nums[b-48] // 48 == ASCII code for 0
		}

		for _, b := range []byte(new) {
			s.Append(b)
		}

		b := s.Get()
		if new != string(b) {
			t.Errorf("String mismatch in test %d", i)
			t.Logf("  Expected: %s", new)
			t.Logf("  Actual:   %s", string(b))
		}
	}
}

func TestLazyStringStringer(t *testing.T) {
	nums := []string{
		"zero.",
		"one.",
		"two.",
		"three.",
		"four.",
		"five.",
		"six.",
		"seven.",
		"eight.",
		"nine.",
	}

	count.Tests(t, 1)

	s := newStr()
	max := cap(s.b)

	for i := 0; i < max; i++ {
		ints := strconv.Itoa(i)
		var new string
		for _, b := range []byte(ints) {
			new += nums[b-48] // 48 == ASCII code for 0
		}

		for _, b := range []byte(new) {
			s.Append(b)
		}

		act := s.String()
		if new != act {
			t.Errorf("String mismatch in test %d", i)
			t.Logf("  Expected: %s", new)
			t.Logf("  Actual:   %s", act)
		}
	}
}

func TestObjectsMap(t *testing.T) {
	count.Tests(t, 1)

	key := "foo"
	val := "bar"
	exp := `{"foo":"bar"}`

	obj := newObjs()
	obj.New(objMap)

	s := obj.GetKeyPtr()
	for _, b := range []byte(key) {
		s.Append(b)
	}

	obj.SetValue(val)

	obj.MergeDown()

	b, err := json.Marshal(obj.nest[0].value)
	if err != nil {
		t.Errorf("Unable to marshal Go struct, this is possibly an error with the standard library: %s", err.Error())
	}

	if string(b) != exp {
		t.Error("Output doesn't match expected:")
		t.Logf("  Exp: %s", exp)
		t.Logf("  Act: %s", string(b))
	}
}
