package httpclient

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestExtractFileName(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{
			url:      `https://example.com`,
			expected: `example.com`,
		},
		{
			url:      `https://example.com/`,
			expected: `example.com`,
		},
		{
			url:      `https://example.com/example.txt`,
			expected: `example.txt`,
		},
		{
			url:      `https://example.com/example.txt/`,
			expected: `example.txt`,
		},
		///// ?
		{
			url:      `https://example.com?`,
			expected: `example.com`,
		},
		{
			url:      `https://example.com/?`,
			expected: `example.com`,
		},
		{
			url:      `https://example.com/example.txt?`,
			expected: `example.txt`,
		},
		{
			url:      `https://example.com/example.txt/?`,
			expected: `example.txt`,
		},
		{
			url:      `https://example.com/example.txt?key=value`,
			expected: `example.txt`,
		},
		{
			url:      `https://example.com/example.txt/?key=value`,
			expected: `example.txt`,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := extractFileName(test.url)
		if test.expected != actual {
			t.Errorf("mismatch in test %d", i)
			t.Logf("  URL:      '%s'", test.url)
			t.Logf("  Expected: '%s'", test.expected)
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}
