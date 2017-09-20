package sexp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	exp, err := Unmarshal([]byte("(foo (bar) baz)"))
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, len(exp))

	bytes, ok := exp[0].([]byte)
	assert.Equal(t, true, ok)
	assert.Equal(t, "foo", string(bytes))

	inter, ok := exp[1].([]interface{})
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, len(inter))

	bytes, ok = inter[0].([]byte)
	assert.Equal(t, true, ok)
	assert.Equal(t, "bar", string(bytes))

	bytes, ok = exp[2].([]byte)
	assert.Equal(t, true, ok)
	assert.Equal(t, "baz", string(bytes))
}

func TestMarshal(t *testing.T) {
	// Casting.
	exp := make([]interface{}, 0)
	exp = append(exp, "foo", []byte("bar"), 1, "quote non tokens", true)

	bytes, err := Marshal(exp, false)
	assert.Equal(t, nil, err)
	assert.Equal(t, `(foo bar "1" "quote non tokens" true)`, string(bytes))

	// Subexpressions.
	subexp := make([]interface{}, 0)
	subexp = append(subexp, "sub")
	exp = make([]interface{}, 0)
	exp = append(subexp, subexp)
	exp = append(exp, exp)

	bytes, err = Marshal(exp, false)
	assert.Equal(t, nil, err)
	assert.Equal(t, `(sub (sub) (sub (sub)))`, string(bytes))
}

func TestMarshalCanonical(t *testing.T) {
	// Casting.
	exp := make([]interface{}, 0)
	exp = append(exp, "foo", []byte("bar"), 1, "quote non tokens", true)

	bytes, err := Marshal(exp, true)
	assert.Equal(t, nil, err)
	assert.Equal(t, `(3:foo3:bar1:116:quote non tokens4:true)`, string(bytes))

	// Subexpressions.
	subexp := make([]interface{}, 0)
	subexp = append(subexp, "sub")
	exp = make([]interface{}, 0)
	exp = append(subexp, subexp)
	exp = append(exp, exp)

	bytes, err = Marshal(exp, true)
	assert.Equal(t, nil, err)
	assert.Equal(t, `(3:sub(3:sub)(3:sub(3:sub)))`, string(bytes))
}
