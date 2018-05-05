package main

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
)

// TestMurex tests murex runtime environment can be initialised and and simple
// command line can exexute
func TestMurex(t *testing.T) {
	proc.InitEnv()

	block := []rune("a [Mon..Fri]->regexp m/^T/")

	_, err := lang.RunBlockShellConfigSpace(block, nil, nil, nil)

	if err != nil {
		t.Error(err.Error())
	}
}
