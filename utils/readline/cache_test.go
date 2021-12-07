package readline

import (
	"fmt"
	"testing"

	"github.com/lmorg/murex/test/count"
)

// string

func TestCacheStringMethods(t *testing.T) {
	rl := NewInstance()
	rl.MaxCacheSize = 10
	nTests := 20
	c := new(cacheString)
	c.Init(rl)

	count.Tests(t, (2*nTests)+2)

	for i := 0; i < 20; i++ {
		sLine := fmt.Sprintf("line %d", i)
		rLine := []rune(sLine)
		sValue := fmt.Sprintf("value %d", i)

		r := c.Get(rLine)
		if string(r) == sValue {
			t.Fatalf(`Pre-Append error: c.Get(%s) == "%s"`, sLine, string(r))
		}

		c.Append(rLine, sValue)

		r = c.Get(rLine)
		if string(r) != sValue {
			t.Fatalf(`Post-Append error: c.Get(%s) == "%s"`, sLine, string(r))
		}
	}

	c.Init(rl)
	if c.size != 0 || len(c.values["line 1"]) > 0 {
		t.Error("cacheString failed to reinitialize correctly")
		t.Logf("  size:                %d", c.size)
		t.Logf(`  c.values["line 1"]: "%s"`, string(c.values["line 1"]))
	}
}

func TestCacheStringOutOfBounds(t *testing.T) {
	count.Tests(t, 1)

	rl := NewInstance()
	rl.MaxCacheSize = 10

	c := new(cacheString)
	c.Init(rl)

	for i := 0; i < 20; i++ {
		line := []rune(fmt.Sprintf("line %d", i))
		value := fmt.Sprintf("value %d", i)
		c.Append(line, value)
	}
}

// []rune

func TestCacheSliceRuneMethods(t *testing.T) {
	rl := NewInstance()
	rl.MaxCacheSize = 10
	nTests := 20
	c := new(cacheSliceRune)
	c.Init(rl)

	count.Tests(t, (2*nTests)+2)

	for i := 0; i < 20; i++ {
		sLine := fmt.Sprintf("line %d", i)
		rLine := []rune(sLine)
		sValue := fmt.Sprintf("value %d", i)
		rValue := []rune(sValue)

		r := c.Get(rLine)
		if string(r) == sValue {
			t.Fatalf(`Pre-Append error: c.Get(%s) == "%s"`, sLine, string(r))
		}

		c.Append(rLine, rValue)

		r = c.Get(rLine)
		if string(r) != sValue {
			t.Fatalf(`Post-Append error: c.Get(%s) == "%s"`, sLine, string(r))
		}
	}

	c.Init(rl)
	if c.size != 0 || len(c.values["line 1"]) > 0 {
		t.Error("cacheString failed to reinitialize correctly")
		t.Logf("  size:                %d", c.size)
		t.Logf(`  c.values["line 1"]: "%s"`, string(c.values["line 1"]))
	}
}

func TestCacheSlineRuneOutOfBounds(t *testing.T) {
	count.Tests(t, 1)

	rl := NewInstance()
	rl.MaxCacheSize = 10

	c := new(cacheSliceRune)
	c.Init(rl)

	for i := 0; i < 20; i++ {
		line := []rune(fmt.Sprintf("line %d", i))
		value := []rune(fmt.Sprintf("value %d", i))
		c.Append(line, value)
	}
}
