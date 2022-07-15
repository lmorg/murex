package parser_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/parser"
)

func TestPipeTokenString(t *testing.T) {
	count.Tests(t, 2)

	var pt parser.ParsedTokens

	if pt.PipeToken.String() != "PipeTokenNone" {
		t.Errorf("Missing stringer definitions")
	}

	pt.PipeToken = 99
	if pt.PipeToken.String() != "PipeToken(99)" {
		t.Errorf("Possibly hardcoded stringer definitions?")
	}

}
