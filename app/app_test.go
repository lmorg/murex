package app

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/lmorg/murex/test/count"
)

func TestAppName(t *testing.T) {
	count.Tests(t, 1)

	if Name == "" {
		t.Error("Name isn't valid:")
		t.Log("  Name:", Name)
	}
}

func TestVersion(t *testing.T) {
	count.Tests(t, 1)
	rx := regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+( (ALPHA|BETA|RC[0-9]))?$`)

	if !rx.MatchString(Version) {
		t.Error("Release version doesn't contain a valid string:")
		t.Log("  Version:", Version)
	}
}

func TestCopyright(t *testing.T) {
	count.Tests(t, 1)
	rx := regexp.MustCompile(fmt.Sprintf(`^(\(c\)|©) [0-9]{4}-%s .*$`, time.Now().Format("2006")))

	if !rx.MatchString(Copyright) {
		t.Error("Copyright string doesn't contain a valid string (possibly out of date?):")
		t.Log("  Copyright:", Copyright)
	}
}

func TestLicense(t *testing.T) {
	if License == "" {
		t.Error("License isn't valid:")
		t.Log("  License:", License)
	}
}
