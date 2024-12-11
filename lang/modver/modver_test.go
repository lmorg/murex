package modver_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/lang/modver"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/semver"
)

func TestModver(t *testing.T) {
	count.Tests(t, 4)

	module := fmt.Sprintf("%s/%s/%d", t.Name(), time.Now().String(), rand.Int())

	version := modver.Get(module)
	if !version.Compare(app.Semver()).IsEqualTo() {
		t.Errorf("Expecting to default to v%s", app.Semver().String())
	}

	modver.Set(module, &semver.Version{1, 2, 3})

	version = modver.Get(module)
	if !version.Compare(&semver.Version{1, 2, 3}).IsEqualTo() {
		t.Error("Expecting to version 1.2.3")
	}

	version = modver.Get("!!!" + module)
	if !version.Compare(app.Semver()).IsEqualTo() {
		t.Errorf("Expecting to default to v%s", app.Semver().String())
	}
}
