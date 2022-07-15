package lang

import (
	"strconv"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestFuncId(t *testing.T) {
	f := newFuncID()

	bad := uint32(1337)
	err := f.Executing(bad)
	if err == nil {
		t.Errorf("fid %d: err should NOT be nil", bad)
	}

	_, err = f.Proc(bad)
	if err == nil {
		t.Errorf("fid %d: err should NOT be nil", bad)
	}

	bad = 0
	_, err = f.Proc(0)
	if err == nil {
		t.Errorf("fid %d: err should NOT be nil", bad)
	}

	count.Tests(t, 3)

	var tests [5]*Process

	count.Tests(t, len(tests)*6)

	for i := range tests {
		tests[i] = NewTestProcess()
		tests[i].Name.Set(strconv.Itoa(i))
		fid := f.Register(tests[i])

		if fid != tests[i].Id {
			t.Errorf("test %d: fid %d: fid != p.Id (%d)", i, fid, tests[i].Id)
		}

		_, err = f.Proc(fid)
		if err == nil {
			t.Errorf("test %d: fid %d: err should NOT be nil", i, fid)
		}

		err := f.Executing(fid)
		if err != nil {
			t.Errorf("test %d: fid %d: err should be nil: %s", i, fid, err.Error())
		}

		p, err := f.Proc(fid)
		if err != nil {
			t.Errorf("test %d: fid %d: unexpected error: %s", i, fid, err.Error())
		}

		if p.Id != fid {
			t.Errorf("test %d: fid %d: fid != p.Id (%d)", i, fid, p.Id)
		}

		f.Deregister(fid)
		err = f.Executing(fid)
		if err == nil {
			t.Errorf("test %d: fid %d: err should NOT be nil", i, fid)
		}
	}
}

func TestFuncIdListAllEmpty(t *testing.T) {
	f := newFuncID()

	count.Tests(t, 1)

	c := 100

	for i := 0; i < c; i++ {
		p := NewTestProcess()
		f.Register(p)
	}

	list := f.ListAll()
	if len(list) != 0 {
		t.Error("list includes non-executed processes")
	}
}

func TestFuncIdListAllFull(t *testing.T) {
	f := newFuncID()

	count.Tests(t, 2)

	c := 100

	for i := 0; i < c; i++ {
		p := NewTestProcess()
		fid := f.Register(p)
		f.Executing(fid)
	}

	list := f.ListAll()
	if len(list) != c {
		t.Error("list != c")
	}

	for i := 0; i < c-1; i++ {
		if list[i].Id >= list[i+1].Id {
			t.Errorf("list not sorted")
		}
	}
}
