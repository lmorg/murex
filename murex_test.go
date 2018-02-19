package main

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
)

// TestMurex tests murex runtime environment can be initialised and and simple
// command line can exexute
func TestMurex(t *testing.T) {
	initEnv()

	_, err := lang.ProcessNewBlock(
		[]rune("a [Mon..Fri]->regexp <null> m/^T/"),
		nil,
		nil,
		nil,
		proc.ShellProcess,
	)

	if err != nil {
		t.Error(err.Error())
	}
}
