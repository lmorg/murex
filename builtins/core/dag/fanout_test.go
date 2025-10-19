package dag_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

func TestFanoutAsFunctionDefault(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `fanout {
				out 1
			} {
				out 2
			} {
				out 3
			}`,
			Stdout: "1\n2\n3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestFanoutAsFunctionConcatenate(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `fanout --concat {
				out 1
			} {
				out 2
			} {
				out 3
			}`,
			Stdout: "1\n2\n3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestFanoutAsFunctionConcatenateAlias(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `fanout -c {
				out 1
			} {
				out 2
			} {
				out 3
			}`,
			Stdout: "1\n2\n3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestFanoutAsMethodDefault(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `%[1..3] -> fanout {
				get-type stdin -> :str: format json
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> fanout {
				%[ ${ <stdin> -> debug -> [[/Data-Type/Murex]] } ]
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> fanout {
				-> regexp m/1/
			} {
				-> regexp m/2/
			} {
				-> regexp m/3/
			}`,
			Stdout: `["1","2","3"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestFanoutAsMethodConcatenate(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `%[1..3] -> fanout --concat {
				get-type stdin -> :str: format json
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> fanout --concat {
				%[ ${ <stdin> -> debug -> [[/Data-Type/Murex]] } ]
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> fanout --concat {
				-> regexp m/1/
			} {
				-> regexp m/2/
			} {
				-> regexp m/3/
			}`,
			Stdout: `["1"]["2"]["3"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestFanoutAsMethodConcatenateAlias(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `%[1..3] -> fanout -c {
				get-type stdin -> :str: format json
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> fanout -c {
				%[ ${ <stdin> -> debug -> [[/Data-Type/Murex]] } ]
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> fanout -c {
				-> regexp m/1/
			} {
				-> regexp m/2/
			} {
				-> regexp m/3/
			}`,
			Stdout: `["1"]["2"]["3"]`,
		},
	}

	test.RunMurexTests(tests, t)
}
