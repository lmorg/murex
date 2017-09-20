package sexp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanonicalLexer(t *testing.T) {
	l := NewLexer([]byte("(3:foo3:bar(3:baz))"))

	item := l.Next()
	assert.Equal(t, "(", string(item.Value))
	assert.Equal(t, ItemBracketLeft, item.Type)

	item = l.Next()
	assert.Equal(t, "foo", string(item.Value))
	assert.Equal(t, ItemVerbatim, item.Type)

	item = l.Next()
	assert.Equal(t, "bar", string(item.Value))
	assert.Equal(t, ItemVerbatim, item.Type)

	item = l.Next()
	assert.Equal(t, "(", string(item.Value))
	assert.Equal(t, ItemBracketLeft, item.Type)

	item = l.Next()
	assert.Equal(t, "baz", string(item.Value))
	assert.Equal(t, ItemVerbatim, item.Type)

	item = l.Next()
	assert.Equal(t, ")", string(item.Value))
	assert.Equal(t, ItemBracketRight, item.Type)

	item = l.Next()
	assert.Equal(t, ")", string(item.Value))
	assert.Equal(t, ItemBracketRight, item.Type)

	item = l.Next()
	assert.Equal(t, ItemEOF, item.Type)
}

func TestLexer(t *testing.T) {
	l := NewLexer([]byte(`(foo "bar baz" (6:foobar 1 000000000 #fff P4M5S))`))

	item := l.Next()
	assert.Equal(t, "(", string(item.Value))
	assert.Equal(t, ItemBracketLeft, item.Type)

	item = l.Next()
	assert.Equal(t, "foo", string(item.Value))
	assert.Equal(t, ItemToken, item.Type)

	item = l.Next()
	assert.Equal(t, "bar baz", string(item.Value))
	assert.Equal(t, ItemQuote, item.Type)

	item = l.Next()
	assert.Equal(t, "(", string(item.Value))
	assert.Equal(t, ItemBracketLeft, item.Type)

	item = l.Next()
	assert.Equal(t, "foobar", string(item.Value))
	assert.Equal(t, ItemVerbatim, item.Type)

	item = l.Next()
	assert.Equal(t, "1", string(item.Value))
	assert.Equal(t, ItemToken, item.Type)

	item = l.Next()
	assert.Equal(t, "000000000", string(item.Value))
	assert.Equal(t, ItemToken, item.Type)

	item = l.Next()
	assert.Equal(t, "#fff", string(item.Value))
	assert.Equal(t, ItemToken, item.Type)

	item = l.Next()
	assert.Equal(t, "P4M5S", string(item.Value))
	assert.Equal(t, ItemToken, item.Type)

	item = l.Next()
	assert.Equal(t, ")", string(item.Value))
	assert.Equal(t, ItemBracketRight, item.Type)

	item = l.Next()
	assert.Equal(t, ")", string(item.Value))
	assert.Equal(t, ItemBracketRight, item.Type)

	item = l.Next()
	assert.Equal(t, ItemEOF, item.Type)
}
