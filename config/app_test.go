package config

import (
	"regexp"
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestConsts Just ensures you use something sane for the applications constants
func TestConsts(t *testing.T) {
	rx := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+ ( (ALPHA|BETA|RC[0-9]))?`)

	count.Tests(t, 2, "TestConsts")

	if AppName == "" {
		t.Error("AppName isn't valid")
		t.Log("AppName:", AppName)
	}

	if !rx.MatchString(Version) {
		t.Error("Release version doesn't contain a valid string")
		t.Log("Version:", Version)
	}
}
