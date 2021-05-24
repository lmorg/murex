package io_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins"
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

func TestWriteFile(t *testing.T) {
	file, err := test.TempDir()
	if err != nil {
		t.Fatalf(err.Error())
	}
	file += "/TestWriteFile.txt"

	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`out foo | > %s; out bar | > %s; open %s`, file, file, file),
			Stdout: "bar\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestAppendFile(t *testing.T) {
	file, err := test.TempDir()
	if err != nil {
		t.Fatalf(err.Error())
	}
	file += "/TestAppendFile.txt"

	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`out foo | > %s; out bar | >> %s; open %s`, file, file, file),
			Stdout: "foo\nbar\n",
		},
	}

	test.RunMurexTests(tests, t)
}
