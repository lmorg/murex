package readline

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestGetPreviewWidth(t *testing.T) {
	tests := []struct {
		Term    int
		Preview int
		Forward int
	}{
		{
			Term:    92,
			Preview: 32,
			Forward: 58,
		},
		{
			Term:    80,
			Preview: 28,
			Forward: 50,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		preview, forward := getPreviewWidth(test.Term)
		if preview != test.Preview || forward != test.Forward {
			t.Errorf("Maths fail in test %d", i)
			t.Logf("  Term Width:  %d", test.Term)
			t.Logf("  Exp Preview: %d", test.Preview)
			t.Logf("  Act Preview: %d", preview)
			t.Logf("  Exp Forward: %d", test.Forward)
			t.Logf("  Act Forward: %d", forward)
		}
	}
}
