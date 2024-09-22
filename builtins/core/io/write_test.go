package io_test

import (
	"fmt"
	"testing"

	"github.com/lmorg/murex/test"
)

func TestWriteFile(t *testing.T) {
	file := t.TempDir()
	file += "/TestWriteFile.txt"

	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`out foo |> %s; out bar |> %s; open %s`, file, file, file),
			Stdout: "bar\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestAppendFile(t *testing.T) {
	file := t.TempDir()
	file += "/TestAppendFile.txt"

	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`out foo |> %s; out bar >> %s; open %s`, file, file, file),
			Stdout: "foo\nbar\n",
		},
	}

	test.RunMurexTests(tests, t)
}
