package expressions

import (
	"testing"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type testParseStatementT struct {
	Statement string
	Args      []string
	Pipes     []string
	Exec      bool
	Error     bool
}

func testParseStatement(t *testing.T, tests []testParseStatementT) {
	t.Helper()
	count.Tests(t, len(tests))

	for i, test := range tests {
		lang.InitEnv()
		defaults.Config(lang.ShellProcess.Config, false)
		p := lang.NewTestProcess()
		p.Name.Set("TestParseStatement")
		p.Config.Set("proc", "strict-vars", false, nil)
		p.Config.Set("proc", "strict-arrays", false, nil)
		tree := NewParser(p, []rune(test.Statement), 0)
		err := tree.ParseStatement(test.Exec)
		if err == nil {
			err = tree.statement.validate()
		}

		actual := make([]string, len(tree.statement.parameters)+1)
		actual[0] = string(tree.statement.command)
		for j := range tree.statement.parameters {
			actual[j+1] = string(tree.statement.parameters[j])
		}

		if (err != nil) != test.Error ||
			json.LazyLogging(test.Args) != json.LazyLogging(actual) ||
			json.LazyLogging(test.Pipes) != json.LazyLogging(tree.statement.namedPipes) {
			t.Errorf("Parser error in test %d", i)
			t.Logf("  Statement: %s", test.Statement)
			t.Logf("  Exec:      %v", test.Exec)
			t.Logf("  Expected:  %s", json.LazyLoggingPretty(test.Args))
			t.Logf("  Actual:    %s", json.LazyLoggingPretty(actual))
			t.Logf("  length:    %d", len(tree.statement.parameters))
			t.Logf("  pipe exp:  %s", json.LazyLogging(test.Pipes))
			t.Logf("  pipe act:  %s", json.LazyLogging(tree.statement.namedPipes))
			t.Logf("  pipe len:  %d", len(tree.statement.namedPipes))
			t.Logf("  exp err:   %v", test.Error)
			t.Logf("  act err:   %v", err)
		}
	}
}

func TestParseStatement(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `echo hello world`,
			Args: []string{
				"echo", "hello", "world",
			},
			Exec: false,
		},
		{
			Statement: `echo hello world`,
			Args: []string{
				"echo", "hello", "world",
			},
			Exec: true,
		},
		{
			Statement: `echo 'hello world'`,
			Args: []string{
				"echo", "'hello world'",
			},
			Exec: false,
		},
		{
			Statement: `echo 'hello world'`,
			Args: []string{
				"echo", "hello world",
			},
			Exec: true,
		},
		{
			Statement: `echo "hello world"`,
			Args: []string{
				"echo", `"hello world"`,
			},
			Exec: false,
		},
		{
			Statement: `echo "hello world"`,
			Args: []string{
				"echo", `hello world`,
			},
			Exec: true,
		},
		{
			Statement: `echo (hello world)`,
			Args: []string{
				"echo", "(hello world)",
			},
			Exec: false,
		},
		{
			Statement: `echo (hello world)`,
			Args: []string{
				"echo", "hello world",
			},
			Exec: true,
		},
		{
			Statement: `echo h(ello worl)d`,
			Args: []string{
				"echo", "h(ello", "worl)d",
			},
			Exec: false,
		},
		{
			Statement: `echo h(ello worl)d`,
			Args: []string{
				"echo", "h(ello", "worl)d",
			},
			Exec: true,
		},
		{
			Statement: `echo %(hello world)`,
			Args: []string{
				"echo", "%(hello world)",
			},
			Exec: false,
		},
		{
			Statement: `echo %(hello world)`,
			Args: []string{
				"echo", "hello world",
			},
			Exec: true,
		},
		{
			Statement: `echo {hello world}`,
			Args: []string{
				"echo", "{hello world}",
			},
			Exec: true,
		},
		{
			Statement: `echo {hello world}`,
			Args: []string{
				"echo", "{hello world}",
			},
			Exec: true,
		},
		/////
		{
			Statement: `echo ${out bob}`,
			Args: []string{
				"echo", "${out bob}",
			},
			Exec: false,
		},
		{
			Statement: `echo ${out bob}`,
			Args: []string{
				"echo", "bob",
			},
			Exec: true,
		},
		{
			Statement: `echo "${out bob}"`,
			Args: []string{
				"echo", `"${out bob}"`,
			},
			Exec: false,
		},
		{
			Statement: `echo "${out bob}"`,
			Args: []string{
				"echo", "bob",
			},
			Exec: true,
		},
		{
			Statement: `echo -${out bob}-`,
			Args: []string{
				"echo", "-bob-",
			},
			Exec: true,
		},
		/////
		{
			Statement: `echo @{ja: [1..3]}`,
			Args: []string{
				"echo", "@{ja: [1..3]}",
			},
			Exec: false,
		},
		{
			Statement: `echo @{ja: [1..3]}`,
			Args: []string{
				"echo", "1", "2", "3",
			},
			Exec: true,
		},
		{
			Statement: `echo "@{ja: [1..3]}"`,
			Args: []string{
				"echo", `"@{ja: [1..3]}"`,
			},
			Exec: false,
		},
		{
			Statement: `echo "@{ja: [1..3]}"`,
			Args: []string{
				"echo", `@{ja: [1..3]}`,
			},
			Exec: true,
		},
		{
			Statement: `echo -@{ja: [1..3]}-`,
			Args: []string{
				"echo", `-@{ja: [1..3]}-`,
			},
			Exec: true,
		},
		{
			Statement: `echo - @{ja: [1..3]}-`,
			Args: []string{
				"echo", `-`, `1`, `2`, `3`, `-`,
			},
			Exec: true,
		},
		/////
		{
			Statement: `echo $bob`,
			Args: []string{
				"echo", "$bob",
			},
			Exec: false,
		},
		{
			Statement: `echo $bob`,
			Args: []string{
				"echo", "",
			},
			Exec: true,
		},
		{
			Statement: `echo "$bob"`,
			Args: []string{
				"echo", `"$bob"`,
			},
			Exec: false,
		},
		{
			Statement: `echo "$bob"`,
			Args: []string{
				"echo", "",
			},
			Exec: true,
		},
		{
			Statement: `echo '$bob'`,
			Args: []string{
				"echo", "'$bob'",
			},
			Exec: false,
		},
		{
			Statement: `echo '$bob'`,
			Args: []string{
				"echo", "$bob",
			},
			Exec: true,
		},
		{
			Statement: `echo -$bob-`,
			Args: []string{
				"echo", "--",
			},
			Exec: true,
		},
		/////
		{
			Statement: `echo @bob`,
			Args: []string{
				"echo", "@bob",
			},
			Exec: false,
		},
		{
			Statement: `echo @bob`,
			Args: []string{
				"echo",
			},
			Exec: true,
		},
		{
			Statement: `echo "@bob"`,
			Args: []string{
				"echo", `"@bob"`,
			},
			Exec: false,
		},
		{
			Statement: `echo "@bob"`,
			Args: []string{
				"echo", `@bob`,
			},
			Exec: true,
		},
		{
			Statement: `echo -@bob-`,
			Args: []string{
				"echo", `-@bob-`,
			},
			Exec: true,
		},
		{
			Statement: `echo - @bob-`,
			Args: []string{
				"echo", `-`, `-`,
			},
			Exec: true,
		},
		/////
		{
			Statement: `echo { "bob" }`,
			Args: []string{
				"echo", `{ "bob" }`,
			},
			Exec: false,
		},
		{
			Statement: `echo { 'bob' }`,
			Args: []string{
				"echo", `{ 'bob' }`,
			},
			Exec: false,
		},
	}

	testParseStatement(t, tests)
}

func TestParseStatementNamedPipe(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `echo <123> hello world`,
			Args: []string{
				"echo", "hello", "world",
			},
			Pipes: []string{
				"123",
			},
			Exec: false,
		},
		{
			Statement: `echo <123> hello world`,
			Args: []string{
				"echo", "hello", "world",
			},
			Pipes: []string{
				"123",
			},
			Exec: true,
		},
		{
			Statement: `echo <1<23> hello world`,
			Args: []string{
				"echo", "<1<23>", "hello", "world",
			},
			Exec: false,
		},
		{
			Statement: `echo <1<23> hello world`,
			Args: []string{
				"echo", "<1<23>", "hello", "world",
			},
			Exec: true,
		},
		{
			Statement: `echo "<123>" hello world`,
			Args: []string{
				"echo", `"<123>"`, "hello", "world",
			},
			Exec: false,
		},
		{
			Statement: `echo "<123>" hello world`,
			Args: []string{
				"echo", "<123>", "hello", "world",
			},
			Exec: true,
		},
	}

	testParseStatement(t, tests)
}

func TestParseStatementExistingCode(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `
				test unit private autocomplete.variables {
					"PreBlock": ({ global MUREX_UNIT_TEST=foobar }),
					"PostBlock": ({ !global MUREX_UNIT_TEST }),
					"StdoutRegex": (^([_a-zA-Z0-9]+\n)+),
					"StdoutType":  "str",
					"StdoutBlock": ({
						-> len -> set len;
						if { = len>0 } then {
							out "Len greater than 0"
						} else {
							err "No elements returned"
						}
					}),
					"StdoutIsArray": true
				}`,
			Args: []string{
				"test", "unit", "private", "autocomplete.variables", `{
					"PreBlock": ({ global MUREX_UNIT_TEST=foobar }),
					"PostBlock": ({ !global MUREX_UNIT_TEST }),
					"StdoutRegex": (^([_a-zA-Z0-9]+\n)+),
					"StdoutType":  "str",
					"StdoutBlock": ({
						-> len -> set len;
						if { = len>0 } then {
							out "Len greater than 0"
						} else {
							err "No elements returned"
						}
					}),
					"StdoutIsArray": true
				}`,
			},
			Exec: false,
		},

		/////

		{
			Statement: `
				test unit private autocomplete.variables {
					"PreBlock": ({ global MUREX_UNIT_TEST=foobar }),
					"PostBlock": ({ !global MUREX_UNIT_TEST }),
					"StdoutRegex": (^([_a-zA-Z0-9]+\n)+),
					"StdoutType":  "str",
					"StdoutBlock": ({
						-> len -> set len;
						if { = len>0 } then {
							out "Len greater than 0"
						} else {
							err "No elements returned"
						}
					}),
					"StdoutIsArray": true
				}`,
			Args: []string{
				"test", "unit", "private", "autocomplete.variables", `{
					"PreBlock": ({ global MUREX_UNIT_TEST=foobar }),
					"PostBlock": ({ !global MUREX_UNIT_TEST }),
					"StdoutRegex": (^([_a-zA-Z0-9]+\n)+),
					"StdoutType":  "str",
					"StdoutBlock": ({
						-> len -> set len;
						if { = len>0 } then {
							out "Len greater than 0"
						} else {
							err "No elements returned"
						}
					}),
					"StdoutIsArray": true
				}`,
			},
			Exec: true,
		},

		/////
	}

	testParseStatement(t, tests)
}

func TestParseStatementStrings(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `echo 'hello world'`,
			Args: []string{
				"echo", `'hello world'`,
			},
			Exec: false,
		},
		{
			Statement: `echo 'hello world'`,
			Args: []string{
				"echo", `hello world`,
			},
			Exec: true,
		},
		/////
		{
			Statement: `echo "hello world"`,
			Args: []string{
				"echo", `"hello world"`,
			},
			Exec: false,
		},
		{
			Statement: `echo "hello world"`,
			Args: []string{
				"echo", `hello world`,
			},
			Exec: true,
		},
	}

	testParseStatement(t, tests)
}

func TestParseStatementObjCreators(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `echo %(hello world)`,
			Args: []string{
				"echo", `%(hello world)`,
			},
			Exec: false,
		},
		{
			Statement: `echo %(hello world)`,
			Args: []string{
				"echo", `hello world`,
			},
			Exec: true,
		},
		/////
		{
			Statement: `echo %[hello world]`,
			Args: []string{
				"echo", "%[hello world]",
			},
			Exec: false,
		},
		{
			Statement: `echo %[hello world]`,
			Args: []string{
				"echo", `["hello","world"]`,
			},
			Exec: true,
		},
		/////
		{
			Statement: `echo %{hello: world}`,
			Args: []string{
				"echo", "%{hello: world}",
			},
			Exec: false,
		},
		{
			Statement: `echo %{hello: world}`,
			Args: []string{
				"echo", `{"hello":"world"}`,
			},
			Exec: true,
		},
	}

	testParseStatement(t, tests)
}

func TestParseStatementEscCrLf(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: "echo 1\\\n2\\\n3\n",
			Args: []string{
				"echo", "1", "2", "3",
			},
			Exec: false,
		},

		/////

		{
			Statement: "echo 1\\\n2\\\n3\n",
			Args: []string{
				"echo", "1", "2", "3",
			},
			Exec: true,
		},
	}

	testParseStatement(t, tests)
}
