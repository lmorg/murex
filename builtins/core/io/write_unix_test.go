//go:build !windows
// +build !windows

package io

import (
	"fmt"
	"testing"

	"github.com/lmorg/murex/test"
)

func TestWriteFilePipelineFlags(t *testing.T) {
	file := t.TempDir()
	file += "/TestWriteFilePipelineFlags"

	tests := []test.MurexTest{
		{
			Block: fmt.Sprintf(`a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp m/0/ |> %[1]s.%[2]d; open %[1]s.%[2]d`,
				file, 0),
			Stdout: "^0\n10\n20\n$",
			Stderr: "^warning",
		},
		{
			Block: fmt.Sprintf(`a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp m/0/ |> %[3]s %[1]s.%[2]d; open %[1]s.%[2]d`,
				file, 1, _WAIT_EOF_SHORT),
			Stdout: "^0\n10\n20\n$",
			Stderr: "^$",
		},
		{
			Block: fmt.Sprintf(`a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp m/0/ |> %[3]s %[1]s.%[2]d; open %[1]s.%[2]d`,
				file, 2, _WAIT_EOF_LONG),
			Stdout: "^0\n10\n20\n$",
			Stderr: "^$",
		},
		{
			Block: fmt.Sprintf(`a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s %[1]s.%[2]d; g <!null> %[3]s`,
				file, 3, _IGNORE_PIPELINE_SHORT),
			Stdout:  "^$",
			Stderr:  "^$",
			ExitNum: 1, // just because of ending g
		},
		{
			Block: fmt.Sprintf(`a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s %[1]s.%[2]d; g <!null> %[3]s`,
				file, 4, _IGNORE_PIPELINE_LONG), Stdout: "^$",
			Stderr:  "^$",
			ExitNum: 1, // just because of ending g
		},

		// two tests here because the regexp is being quirky and I want to check beginning and end of string
		{
			Block: fmt.Sprintf(`g %[3]s; a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s; rm -- %[3]s`,
				file, 5, _WAIT_EOF_SHORT),
			Stdout: "^$",
			Stderr: `^Error in .g.`,
		},
		{
			Block: fmt.Sprintf(`g %[3]s; a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s; rm -- %[3]s`,
				file, 6, _WAIT_EOF_SHORT),
			Stdout: "^$",
			Stderr: `Error: no data returned\n$`,
		},
		// two tests here because the regexp is being quirky and I want to check beginning and end of string
		{
			Block: fmt.Sprintf(`g %[3]s; a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s; rm -- %[3]s`,
				file, 5, _WAIT_EOF_LONG),
			Stdout: "^$",
			Stderr: `^Error in .g.`,
		},
		{
			Block: fmt.Sprintf(`g %[3]s; a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s; rm -- %[3]s`,
				file, 6, _WAIT_EOF_LONG),
			Stdout: "^$",
			Stderr: `Error: no data returned\n$`,
		},
		// two tests here because the regexp is being quirky and I want to check beginning and end of string
		{
			Block: fmt.Sprintf(`g %[3]s; a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s; rm -- %[3]s`,
				file, 7, _IGNORE_PIPELINE_SHORT),
			Stdout: "^$",
			Stderr: `^Error in .g.`,
		},
		{
			Block: fmt.Sprintf(`g %[3]s; a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s; rm -- %[3]s`,
				file, 8, _IGNORE_PIPELINE_SHORT),
			Stdout: "^$",
			Stderr: `Error: no data returned\n$`,
		},
		// two tests here because the regexp is being quirky and I want to check beginning and end of string
		{
			Block: fmt.Sprintf(`g %[3]s; a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s; rm -- %[3]s`,
				file, 5, _IGNORE_PIPELINE_LONG),
			Stdout: "^$",
			Stderr: `^Error in .g.`,
		},
		{
			Block: fmt.Sprintf(`g %[3]s; a [0..20] |> %[1]s.%[2]d; open %[1]s.%[2]d -> regexp s/0// |> %[3]s; rm -- %[3]s`,
				file, 6, _IGNORE_PIPELINE_LONG),
			Stdout: "^$",
			Stderr: `Error: no data returned\n$`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestWriteEmptyFile(t *testing.T) {
	file := t.TempDir()
	file += "/TestWriteFilePipelineFlags"

	tests := []test.MurexTest{
		// two tests here because the regexp is being quirky and I want to check beginning and end of string
		{
			Block: fmt.Sprintf(`g %[1]s.%[2]d; > %[1]s.%[2]d; g %[1]s.%[2]d`,
				file, 0),
			Stdout: fmt.Sprintf(`^\[\"%[1]s.%[2]d\"\]\n$`,
				file, 0),
			Stderr:  `^Error in .g.`,
			ExitNum: 0, // no error because last command succeeded
		},
		{
			Block: fmt.Sprintf(`g %[1]s.%[2]d; > %[1]s.%[2]d; g %[1]s.%[2]d`,
				file, 1),
			Stdout: fmt.Sprintf(`^\[\"%[1]s.%[2]d\"\]\n$`,
				file, 1),
			Stderr:  `Error: no data returned\n$`,
			ExitNum: 0, // no error because last command succeeded
		},
	}

	test.RunMurexTestsRx(tests, t)
}
