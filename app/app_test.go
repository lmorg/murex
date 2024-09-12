package app_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/semver"
)

func TestAppName(t *testing.T) {
	count.Tests(t, 1)

	if app.Name == "" {
		t.Error("Name isn't valid:")
		t.Log("  Name:", app.Name)
	}
}

func TestVersion(t *testing.T) {
	count.Tests(t, 1)
	rx := regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+( \(([-._/a-zA-Z0-9]+)\))?$`)

	if !rx.MatchString(app.Version()) {
		t.Error("Release version doesn't contain a valid string:")
		t.Log("  Version:", app.Version())
	}
}

func TestCopyright(t *testing.T) {
	count.Tests(t, 1)
	rx := regexp.MustCompile(fmt.Sprintf(`^[0-9]{4}-%s .*$`, time.Now().Format("2006")))

	if !rx.MatchString(app.Copyright) {
		t.Error("Copyright string doesn't contain a valid string (possibly out of date?):")
		t.Log("  Copyright:", app.Copyright)
	}
}

func TestLicense(t *testing.T) {
	if app.License == "" {
		t.Error("License isn't valid:")
		t.Log("  License:", app.License)
	}
}

func TestSemVer(t *testing.T) {
	_, err := semver.Parse(fmt.Sprintf("%d.%d.%d", app.Major, app.Minor, app.Revision))
	if err != nil {
		t.Error(err)
	}
}
