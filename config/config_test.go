package config

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

// TestConfig tests the config structure
func TestConfig(t *testing.T) {
	count.Tests(t, 2, "TestConfig")

	conf := NewConfiguration()

	conf.Define("shell", "prompt", Properties{
		Description: "test",
		Default:     "test",
		DataType:    types.String,
	})

	copy := conf.Copy()

	err := copy.Set("shell", "prompt", "copy")
	if err != nil {
		t.Error("Error setting string: " + err.Error())
	}

	if conf.values["shell"]["prompt"].(string) != conf.properties["shell"]["prompt"].Default.(string) {
		t.Error("Original struct should retain it's state: " + conf.values["shell"]["prompt"].(string) + "!=" + conf.properties["shell"]["prompt"].Default.(string))
	}

	if copy.values["shell"]["prompt"].(string) != "copy" {
		t.Error("Copy struct should have new state: " + copy.values["shell"]["prompt"].(string))
	}

	v, err := copy.Get("shell", "prompt", types.String)
	if err != nil {
		t.Error("Unable to get config: " + err.Error())
	}

	if v.(string) != "copy" {
		t.Error("config.Get returns invalid string: " + v.(string))
	}
}
