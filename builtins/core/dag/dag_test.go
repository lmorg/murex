package dag_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

func TestDagAsFunctionDefault(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `dag {
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

func TestDagAsFunctionAppend(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `dag --append {
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

func TestDagAsFunctionAppendAlias(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `dag -a {
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

func TestDagAsMethodDefault(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `%[1..3] -> dag {
				get-type stdin -> :str: format json
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> dag {
				%[ ${ <stdin> -> debug -> [[/Data-Type/Murex]] } ]
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> dag {
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

func TestDagAsMethodAppend(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `%[1..3] -> dag --append {
				get-type stdin -> :str: format json
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> dag --append {
				%[ ${ <stdin> -> debug -> [[/Data-Type/Murex]] } ]
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> dag --append {
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

func TestDagAsMethodAppendAlias(t *testing.T) {
	count.Tests(t, 1)

	tests := []test.MurexTest{
		{
			Block: `%[1..3] -> dag -a {
				get-type stdin -> :str: format json
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> dag -a {
				%[ ${ <stdin> -> debug -> [[/Data-Type/Murex]] } ]
			}`,
			Stdout: fmt.Sprintf(`["%s"]`, types.Json),
		},
		{
			Block: `%[1..3] -> dag -a {
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
