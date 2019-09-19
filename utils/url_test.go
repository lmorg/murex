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
