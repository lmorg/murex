package config_test

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

func TestConfigScoped(t *testing.T) {
	count.Tests(t, 2)

	orig := config.InitConf.Copy()

	orig.Define("shell", "prompt", config.Properties{
		Description: "test",
		Default:     "test",
		DataType:    types.String,
		Global:      false,
	})

	copy := orig.Copy()

	err := copy.Set("shell", "prompt", "copy")
	if err != nil {
		t.Error("Error setting string: " + err.Error())
	}

	old, err := orig.Get("shell", "prompt", types.String)
	if err != nil {
		t.Error("Unable to get old config: " + err.Error())
	}

	new, err := copy.Get("shell", "prompt", types.String)
	if err != nil {
		t.Error("Unable to get new config: " + err.Error())
	}

	if old.(string) == new.(string) {
		t.Error("old and new configs have synced when they should have drifted:")
		t.Logf("  old: %s", old)
		t.Logf("  new: %s", new)
	}
}

func TestConfigGlobal(t *testing.T) {
	count.Tests(t, 2)

	orig := config.InitConf.Copy()

	orig.Define("shell", "prompt", config.Properties{
		Description: "test",
		Default:     "test",
		DataType:    types.String,
		Global:      true,
	})

	copy := orig.Copy()

	err := copy.Set("shell", "prompt", "copy")
	if err != nil {
		t.Error("Error setting string: " + err.Error())
	}

	old, err := orig.Get("shell", "prompt", types.String)
	if err != nil {
		t.Error("Unable to get old config: " + err.Error())
	}

	new, err := copy.Get("shell", "prompt", types.String)
	if err != nil {
		t.Error("Unable to get new config: " + err.Error())
	}

	if old.(string) != new.(string) {
		t.Error("old and new configs have drifted when they should have synced:")
		t.Logf("  old: %s", old)
		t.Logf("  new: %s", new)
	}
}
