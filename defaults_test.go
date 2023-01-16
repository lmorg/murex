package main

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

// TestDefaultConfigExists tests the Default() function populates *config.Config
func TestDefaultConfigExists(t *testing.T) {
	count.Tests(t, 1)

	conf := config.InitConf.Copy()

	defaults.Config(conf, false)

	m := conf.DumpConfig()
	if len(m) == 0 {
		t.Error("Defaults() not populating *config.Config.")
	}
}

// TestDefaultProfileCompiles test the builtin murex_profile compiles
func TestDefaultProfileCompiles(t *testing.T) {
	count.Tests(t, 1)

	defaults.Config(config.InitConf, false)
	lang.InitEnv()
	lang.ShellProcess.Config = config.InitConf

	var block string
	for _, profile := range defaults.DefaultProfiles {
		block += "\n\n" + string(profile.Block)
	}

	fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_CREATE_STDERR)
	exitNum, err := fork.Execute([]rune(block))

	if err != nil {
		t.Error("Error compiling murex_profile:")
		t.Log(err)
	}

	b, err := fork.Stderr.ReadAll()
	if err != nil {
		t.Error("Error reading from streams.Stdin (stderr):")
		t.Log(err)
	}

	if len(b) > 0 {
		t.Error("Uncaptured stderr content:")
		t.Log(string(b))
	}

	if exitNum != 0 {
		t.Error("Non-zero exit number:")
		t.Log(exitNum)
	}
}
