// +build ignore

package defaults

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/builtins/pipes/streams"
)

// TestDefaults tests the Default() function populates *config.Config
func TestDefaults(t *testing.T) {
	conf := config.NewConfiguration()

	Defaults(conf, false)

	m := conf.Dump()
	if len(m) == 0 {
		t.Error("Defaults() not populating *config.Config.")
	}
}

// TestProfile test the builtin murex_profile compiles
func TestProfile(t *testing.T) {
	Defaults(proc.InitConf, false)
	proc.InitEnv()

	stderr := streams.NewStdin()

	exitNum, err := lang.RunBlockShellConfigSpace(DefaultMurexProfile(), nil, nil, stderr)

	if err != nil {
		t.Error("Error compiling murex_profile: " + err.Error())
	}

	b, err := stderr.ReadAll()
	if err != nil {
		t.Error("Error reading from streams.Stdin (stderr): " + err.Error())
	}

	if len(b) > 0 {
		t.Error("Uncaptured stderr content: " + string(b))
	}

	if exitNum != 0 {
		t.Error("Non-zero exit number")
	}
}
