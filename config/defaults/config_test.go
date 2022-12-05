package defaults_test

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/test/count"
)

func TestDefaultProfileNotEmpty(t *testing.T) {
	count.Tests(t, 1)

	if len(defaults.DefaultProfiles) == 0 {
		t.Error("Empty default profile")
	}
}

func TestDefaultConfigLoads(t *testing.T) {
	count.Tests(t, 1)

	c := config.InitConf
	defaults.Config(c, false)

	if len(c.DumpConfig()) == 0 {
		t.Errorf("config not loading")
	}
}
