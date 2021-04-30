package config_test

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

func TestExistsAndGlobal(t *testing.T) {
	count.Tests(t, 3)
	conf := config.InitConf.Copy()

	exists, global := conf.ExistsAndGlobal("app", "key")
	if exists || global {
		t.Errorf("Invalid response:")
		t.Logf("  Exists (expected): %v (%v)", exists, false)
		t.Logf("  Global (expected): %v (%v)", global, false)
	}

	conf.Define("app", "key", config.Properties{
		Description: "test",
		Default:     "test",
		DataType:    types.String,
		Global:      false,
	})

	exists, global = conf.ExistsAndGlobal("app", "key")
	if exists || global {
		t.Errorf("Invalid response:")
		t.Logf("  Exists (expected): %v (%v)", exists, true)
		t.Logf("  Global (expected): %v (%v)", global, false)
	}

	conf.Define("app", "key", config.Properties{
		Description: "test",
		Default:     "test",
		DataType:    types.String,
		Global:      true,
	})

	exists, global = conf.ExistsAndGlobal("app", "key")
	if exists || global {
		t.Errorf("Invalid response:")
		t.Logf("  Exists (expected): %v (%v)", exists, true)
		t.Logf("  Global (expected): %v (%v)", global, true)
	}
}

func TestConfigScoped(t *testing.T) {
	count.Tests(t, 3)

	orig := config.InitConf.Copy()

	orig.Define("app", "key", config.Properties{
		Description: "test",
		Default:     "test",
		DataType:    types.String,
		Global:      false,
	})

	copy := orig.Copy()

	err := copy.Set("app", "key", "copy")
	if err != nil {
		t.Error("Error setting string: " + err.Error())
	}

	old, err := orig.Get("app", "key", types.String)
	if err != nil {
		t.Error("Unable to get old config: " + err.Error())
	}

	new, err := copy.Get("app", "key", types.String)
	if err != nil {
		t.Error("Unable to get new config: " + err.Error())
	}

	if old.(string) == new.(string) {
		t.Error("old and new configs have synced when they should have drifted:")
		t.Logf("  old: %s", old)
		t.Logf("  new: %s", new)
	}

	///// Default

	err = copy.Default("app", "key")
	if err != nil {
		t.Error("Error defaulting string: " + err.Error())
	}

	old, err = orig.Get("app", "key", types.String)
	if err != nil {
		t.Error("Unable to get old config: " + err.Error())
	}

	new, err = copy.Get("app", "key", types.String)
	if err != nil {
		t.Error("Unable to get new config: " + err.Error())
	}

	if old.(string) != new.(string) {
		t.Error("old and new configs should have synced again:")
		t.Logf("  old: %s", old)
		t.Logf("  new: %s", new)
	}
}

func TestConfigGlobal(t *testing.T) {
	count.Tests(t, 3)

	orig := config.InitConf.Copy()

	orig.Define("app", "key", config.Properties{
		Description: "test",
		Default:     "test",
		DataType:    types.String,
		Global:      true,
	})

	copy := orig.Copy()

	err := copy.Set("app", "key", "copy")
	if err != nil {
		t.Error("Error setting string: " + err.Error())
	}

	old, err := orig.Get("app", "key", types.String)
	if err != nil {
		t.Error("Unable to get old config: " + err.Error())
	}

	new, err := copy.Get("app", "key", types.String)
	if err != nil {
		t.Error("Unable to get new config: " + err.Error())
	}

	if old.(string) != new.(string) {
		t.Error("old and new configs have drifted when they should have synced:")
		t.Logf("  old: %s", old)
		t.Logf("  new: %s", new)
	}

	///// Default

	err = copy.Default("app", "key")
	if err != nil {
		t.Error("Error defaulting string: " + err.Error())
	}

	old, err = orig.Get("app", "key", types.String)
	if err != nil {
		t.Error("Unable to get old config: " + err.Error())
	}

	new, err = copy.Get("app", "key", types.String)
	if err != nil {
		t.Error("Unable to get new config: " + err.Error())
	}

	if old.(string) != new.(string) {
		t.Error("old and new configs should have synced again:")
		t.Logf("  old: %s", old)
		t.Logf("  new: %s", new)
	}
}

// TestConfigDumps purely checks that changes don't introduce panics
func TestConfigDumps(t *testing.T) {
	count.Tests(t, 3)

	conf := config.InitConf.Copy()

	conf.Define("app", "key", config.Properties{
		Description: "test",
		Default:     "test",
		DataType:    types.String,
		Global:      false,
	})

	conf.DumpConfig()
	conf.DumpRuntime()
}
