package virtualterm_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/ansi/codes"
	"github.com/lmorg/murex/utils/virtualterm"
)

func TestWriteSgrFgRedExportHtml(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(120, 1)
	test := fmt.Sprintf("Normal%sBold%sRed%sReset", codes.Bold, codes.FgRed, codes.Reset)
	exp1 := `<span class="">Normal</span><span class="sgr-bold">Bold</span><span class="sgr-bold sgr-red">Red</span><span class="">Reset</span><span class="">                                                                                                      
</span>`
	exp2 := `<span class="">Normal</span><span class="sgr-bold">Bold</span><span class="sgr-bold sgr-red">Red</span><span class="">Reset</span><span class="">                                                                                                      
</span>`

	term.Write([]rune(test))
	act := strings.TrimSpace(term.ExportHtml())

	if exp1 != act && exp2 != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp1)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp1))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestWriteSgrFgColoursExportHtml(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(120, 1)
	test := fmt.Sprintf("%sRed%sGreen%sBlue", codes.FgRed, codes.FgGreen, codes.FgBlue)
	exp := `<span class=""></span><span class="sgr-red">Red</span><span class="sgr-green">Green</span><span class="sgr-blue">Blue</span><span class="">                                                                                                            
</span>`

	term.Write([]rune(test))
	act := strings.TrimSpace(term.ExportHtml())

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}
