package main

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"testing"
)

// TestMurex tests murex runtime environment can be initialised and and simple
// command line can exexute
func TestMurex(t *testing.T) {
	proc.InitEnv()

	block := []rune("a [Mon..Fri]->regexp <null> m/^T/")

	_, err := lang.RunBlockShellNamespace(block, nil, nil, nil)

	if err != nil {
		t.Error(err.Error())
	}
}
