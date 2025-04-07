package utils

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestIsURL tests the IsURL function
func TestIsURL(t *testing.T) {
	bad := []string{
		"http//domain",
		"https//domain",
		"http:/domain",
		"https:/domain",
		"http:domain",
		"https:domain",
		"ftp://domain",
		"domain/https://",
		"domain/http://",
		"domain",
	}

	good := []string{
		"http://domain",
		"https://domain",
	}

	count.Tests(t, len(good)+len(bad))

	for _, s := range bad {
		if IsURL(s) {
			t.Error("String incorrectly identified as valid URL: " + s)
		}
	}

	for _, s := range good {
		if !IsURL(s) {
			t.Error("String incorrectly identified as invalid URL: " + s)
		}
	}
}

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
		actual := ExtractFileNameFromURL(test.url)
		if test.expected != actual {
			t.Errorf("mismatch in test %d", i)
			t.Logf("  URL:      '%s'", test.url)
			t.Logf("  Expected: '%s'", test.expected)
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}
