package readline

/*func TestGetPreviewWidth(t *testing.T) {
	tests := []struct {
		Term    int
		Preview int
		Forward int
	}{
		{
			Term:    79,
			Preview: 75,
			Forward: 2,
		},
		{
			Term:    92,
			Preview: 80,
			Forward: 10,
		},
		{
			Term:    80,
			Preview: 76,
			Forward: 2,
		},
		{
			Term:    120,
			Preview: 116,
			Forward: 2,
		},
		{
			Term:    300,
			Preview: 120,
			Forward: 178,
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
}*/
