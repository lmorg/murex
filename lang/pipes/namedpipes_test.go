package pipes_test

import (
	"testing"
	"time"

	_ "github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/pipes"
	"github.com/lmorg/murex/test/count"
)

func TestPipes(t *testing.T) {
	count.Tests(t, 3)
	pipes := pipes.NewNamed()

	if len(pipes.Dump()) != 1 {
		t.Errorf("Empty pipes != 1: %d", len(pipes.Dump()))
	}

	err := pipes.CreatePipe("test", "std", "")
	if err != nil {
		t.Error(err)
	}

	if len(pipes.Dump()) != 2 {
		t.Errorf("Empty pipes != 2: %d", len(pipes.Dump()))
	}

	err = pipes.Close("test")
	if err != nil {
		t.Error(err)
	}

	if len(pipes.Dump()) != 2 {
		t.Errorf("(pipe timeout exceeded too early) Empty pipes != 2: %d", len(pipes.Dump()))
	}

	time.Sleep(5 * time.Second)

	if len(pipes.Dump()) != 1 {
		t.Errorf("Empty pipes != 1: %d", len(pipes.Dump()))
	}
}
