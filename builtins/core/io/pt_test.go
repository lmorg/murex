package io_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestPipeTelemetry(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout * 12345 -> pt`,
			Stdout: `12345`,
		},
	}

	test.RunMurexTests(tests, t)
}
