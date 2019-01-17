package main

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
)

// TestDefaultConfigExists tests the Default() function populates *config.Config
func TestDefaultConfigExists(t *testing.T) {
	conf := config.NewConfiguration()

	defaults.Defaults(conf, false)

	m := conf.Dump()
	if len(m) == 0 {
		t.Error("Defaults() not populating *config.Config.")
	}
}

// TestDefaultProfileCompiles test the builtin murex_profile compiles
func TestDefaultProfileCompiles(t *testing.T) {
	defaults.Defaults(proc.InitConf, false)
	proc.InitEnv()
	proc.ShellProcess.Config = proc.InitConf

	stderr := streams.NewStdin()

	exitNum, err := lang.RunBlockShellConfigSpace(defaults.DefaultMurexProfile(), nil, nil, stderr)

	if err != nil {
		t.Error("Error compiling murex_profile:")
		t.Log(err)
	}

	b, err := stderr.ReadAll()
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
