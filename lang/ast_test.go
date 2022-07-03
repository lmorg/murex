package lang_test

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func TestAstLast(t *testing.T) {
	count.Tests(t, 3)

	var nodes lang.AstNodes

	nodes.Last()

	nodes = append(nodes, lang.AstNode{Name: "foo"})
	if nodes.Last().Name != "foo" {
		t.Errorf("%s != foo", nodes.Last().Name)
	}

	nodes = append(nodes, lang.AstNode{Name: "bar"})
	if nodes.Last().Name != "bar" {
		t.Errorf("%s != bar", nodes.Last().Name)
	}
	if len(nodes) != 2 {
		t.Error("append failed")
	}
}
