package file

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSetDataType(t *testing.T) {
	count.Tests(t, 3)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.SetDataType("foo")
	dt := f.GetDataType()
	if dt != types.Null {
		t.Errorf("Data should be %s: '%s'", types.Null, dt)
	}

	f.SetDataType("bar")
	dt = f.GetDataType()
	if dt != types.Null {
		t.Errorf("Data type still should be %s: '%s'", types.Null, dt)
	}
}

func TestSetOpenClose(t *testing.T) {
	count.Tests(t, 2)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.Open()
	f.Close()
}

func TestSetCloseError1(t *testing.T) {
	count.Tests(t, 1)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	defer func() {
		if r := recover(); r == nil {
			t.Error("expecting a panic due to more closed dependant than open")
		}
	}()

	f.Close()
}

func TestSetOpenCloseError(t *testing.T) {
	count.Tests(t, 3)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.Open()
	f.Close()

	defer func() {
		if r := recover(); r == nil {
			t.Error("expecting a panic due to more closed dependant than open")
		}
	}()

	f.Close()
}

func TestSetForceClose(t *testing.T) {
	count.Tests(t, 3)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.ForceClose()
}

func TestSetOpenForceClose(t *testing.T) {
	count.Tests(t, 3)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.Open()
	f.ForceClose()
}

func TestSetWriteStat(t *testing.T) {
	count.Tests(t, 4)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.Open()
	i, err := f.Write([]byte("12345"))
	if err != nil {
		t.Error(err)
	}
	w, r := f.Stats()
	f.Close()

	if i != 5 || r != 0 || w != 5 {
		t.Error("Stats reporting incorrect values:")
		t.Logf("  i: %d, expected: 5", i)
		t.Logf("  r: %d, expected: 0", r)
		t.Logf("  w: %d, expected: 5", w)
	}
}

func TestSetWritelnStat(t *testing.T) {
	count.Tests(t, 4)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.Open()
	i, err := f.Writeln([]byte("12345"))
	if err != nil {
		t.Error(err)
	}
	w, r := f.Stats()
	f.Close()

	if i != 6 || r != 0 || w != 6 {
		t.Error("Stats reporting incorrect values:")
		t.Logf("  i: %d, expected: 6", i)
		t.Logf("  r: %d, expected: 0", r)
		t.Logf("  w: %d, expected: 6", w)
	}
}

func TestSetWriteAfterClose(t *testing.T) {
	count.Tests(t, 3)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.Open()
	f.Close()
	i, err := f.Write([]byte("12345"))
	if i != 0 || err == nil ||
		(err != nil && !strings.Contains(err.Error(), "closed pipe")) {
		t.Error("Call should error:")
		t.Logf("  i:   %d", i)
		t.Logf("  err: %v", err)
	}
}

func TestSetWritelnAfterClose(t *testing.T) {
	count.Tests(t, 3)

	filename := fmt.Sprintf("%s-%d.test", t.Name(), rand.Int())

	f, err := NewFile(filename)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.Remove(filename)
		if err != nil {
			t.Errorf("Unable to delete '%s' (please do this manually): %s", filename, err.Error())
		}
	}()

	f.Open()
	f.Close()
	i, err := f.Writeln([]byte("12345"))
	if i != 0 || err == nil ||
		(err != nil && !strings.Contains(err.Error(), "closed pipe")) {
		t.Error("Call should error:")
		t.Logf("  i:   %d", i)
		t.Logf("  err: %v", err)
	}
}
