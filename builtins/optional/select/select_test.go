package sqlselect_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

var table = `
["FID","Parent","Scope","State","RunMode","BG","OutPipe","ErrPipe","Command","Parameters"]
[2723,0,0,"Executing","Default",false,"out","err","exec","dog"]
[2724,0,0,"Executing","Default",false,"out","err","exec","cat"]
[2724,0,0,"Executing","Default",true,"out","err","foo",""]
[2724,0,0,"Executing","Default",false,"out","err","bar",""]
`

func TestSelectStdin(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: fmt.Sprintf(
				"tout jsonl (%s) -> select * WHERE command = `exec`",
				table),
			Stdout: string(
				`["FID","Parent","Scope","State","RunMode","BG","OutPipe","ErrPipe","Command","Parameters"]
["2723","0","0","Executing","Default","false","out","err","exec","dog"]
["2724","0","0","Executing","Default","false","out","err","exec","cat"]
`,
			),
		},
	}

	test.RunMurexTests(tests, t)
}
