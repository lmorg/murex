package autocomplete_test

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/murex/utils/readline"
)

type testAutocompleteFlagsT struct {
	CmdLine  string
	ExpItems []string
	ExpDefs  map[string]string
}

func initAutocompleteFlagsTest(exe string, acJson string) {
	lang.InitEnv()
	defaults.Defaults(lang.ShellProcess.Config, false)
	//debug.Enabled = true

	err := lang.ShellProcess.Config.Set("shell", "autocomplete-soft-timeout", 3000)
	if err != nil {
		panic(err.Error())
	}

	err = lang.ShellProcess.Config.Set("shell", "autocomplete-hard-timeout", 10000)
	if err != nil {
		panic(err.Error())
	}

	var flags []autocomplete.Flags

	err = json.UnmarshalMurex([]byte(acJson), &flags)
	if err != nil {
		panic(err.Error())
	}

	for i := range flags {
		// So we don't have nil values in JSON
		if len(flags[i].Flags) == 0 {
			flags[i].Flags = make([]string, 0)
		}

		sort.Strings(flags[i].Flags)
	}

	autocomplete.ExesFlags[exe] = flags
	autocomplete.ExesFlagsFileRef[exe] = &ref.File{Source: &ref.Source{Module: "test/test"}}
}

func testAutocompleteFlags(t *testing.T, tests []testAutocompleteFlagsT) {
	t.Helper()
	count.Tests(t, len(tests))

	for i, test := range tests {
		var err error
		errCallback := func(e error) { err = e }

		pt, _ := parser.Parse([]rune(test.CmdLine), 0)

		dtc := readline.DelayedTabContext{}
		dtc.Context, _ = context.WithCancel(context.Background())

		act := autocomplete.AutoCompleteT{
			Definitions:       make(map[string]string),
			ErrCallback:       errCallback,
			DelayedTabContext: dtc,
			ParsedTokens:      pt,
		}

		var prefix string
		if len(pt.Parameters) > 0 {
			prefix = pt.Parameters[len(pt.Parameters)-1]
		}

		pIndex := 0
		autocomplete.MatchFlags(autocomplete.ExesFlags[pt.FuncName], prefix, pt.FuncName, pt.Parameters, &pIndex, &act)

		if err != nil {
			t.Errorf("Error in test %d: %s", i, err.Error())
			continue
		}

		sort.Strings(test.ExpItems)
		sort.Strings(act.Items)

		expectedItems := jsonOutput(test.ExpItems)
		actualItems := jsonOutput(act.Items)
		expectedDefs := jsonOutput(test.ExpDefs)
		actualDefs := jsonOutput(act.Definitions)

		if len(test.ExpItems) != len(act.Items) || len(test.ExpDefs) != len(act.Definitions) {
			t.Errorf("Item count != expected item count")
			t.Logf("  Test number:      %d", i)
			t.Logf("  Command line:     %s", test.CmdLine)
			t.Logf("  Expected n Items: %d", len(test.ExpItems))
			t.Logf("  Actual n Items:   %d", len(act.Items))
			t.Logf("  Expected n Defs:  %d", len(test.ExpDefs))
			t.Logf("  Actual n Defs:    %d", len(act.Definitions))
			t.Logf("  exp json items:   %s", expectedItems)
			t.Logf("  act json items:   %s", actualItems)
			t.Logf("  exp json defs:    %s", expectedDefs)
			t.Logf("  act json defs:    %s", actualDefs)
			goto nextTest
		}

		for item := range test.ExpItems {
			if test.ExpItems[item] != act.Items[item] {
				t.Errorf("test.ExpItems[item] != act.Items[item]")
				t.Logf("  Test number:      %d", i)
				t.Logf("  Command line:     %s", test.CmdLine)
				t.Logf("  Item index:       %d", item)
				t.Logf("  Expected Item:    %s", test.ExpItems[item])
				t.Logf("  Actual Item:      %s", act.Items[item])
				t.Logf("  exp json items:   %s", expectedItems)
				t.Logf("  act json items:   %s", actualItems)
				t.Logf("  exp json defs:    %s", expectedDefs)
				t.Logf("  act json defs:    %s", actualDefs)
				goto nextTest
			}
		}

		for k := range test.ExpDefs {
			if test.ExpDefs[k] != act.Definitions[k] {
				t.Errorf("test.ExpDefs[k] != act.Definitions[k]")
				t.Logf("  Test number:      %d", i)
				t.Logf("  Command line:     %s", test.CmdLine)
				t.Logf("  Definition key:   %s", k)
				t.Logf("  Expected def:     %s", test.ExpDefs[k])
				t.Logf("  Actual def:       %s", act.Definitions[k])
				t.Logf("  exp json items:   %s", expectedItems)
				t.Logf("  act json items:   %s", actualItems)
				t.Logf("  exp json defs:    %s", expectedDefs)
				t.Logf("  act json defs:    %s", actualDefs)
				goto nextTest
			}
		}
	nextTest:
	}
}

func jsonOutput(v interface{}) string {
	b, err := json.Marshal(v, false)
	if err != nil && !strings.Contains(err.Error(), "o data returned") {
		panic(err.Error())
	}
	return string(b)
}

func TestAutocompleteDocgen(t *testing.T) {
	json := `
		[
			{
				"AllowMultiple": true,
				"Optional": true,
				"FlagsDesc": {
					"-panic": "panic",
					"-readonly": "ro",
					"-verbose": "v",
					"-version": "ver",
					"-warning": "warn"
				}
			},
			{
				"FlagsDesc": {
					"-config": "conf"
				},
				"FlagValues": {
					"-config": [{
						"IncFiles": true
					}]
				}
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s -`, t.Name()),
			ExpItems: []string{
				`panic`,
				"readonly",
				"verbose",
				"version",
				"warning",
				"config",
			},
			ExpDefs: map[string]string{
				"config":   "conf",
				"panic":    "panic",
				"readonly": "ro",
				"verbose":  "v",
				"version":  "ver",
				"warning":  "warn",
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s -p`, t.Name()),
			ExpItems: []string{
				`anic`,
			},
			ExpDefs: map[string]string{
				`anic`: `panic`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s -panic -w`, t.Name()),
			ExpItems: []string{
				`arning`,
			},
			ExpDefs: map[string]string{
				`arning`: `warn`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s -panic -c`, t.Name()),
			ExpItems: []string{
				`onfig`,
			},
			ExpDefs: map[string]string{
				`onfig`: `conf`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s -panic -config R`, t.Name()),
			ExpItems: []string{
				`EADME.md`,
			},
			ExpDefs: map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteDocgenBug(t *testing.T) {
	json := `
		[
			{
				"AllowMultiple": true,
				"Optional": true,
				"FlagsDesc": {
					"-panic": "panic",
					"-readonly": "ro",
					"-verbose": "v",
					"-version": "ver",
					"-warning": "warn",
					"-config": "conf"
				},
				"FlagValues": {
					"-config": [{
						"IncFiles": true
					}]
				}
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s -`, t.Name()),
			ExpItems: []string{
				`panic`,
				"readonly",
				"verbose",
				"version",
				"warning",
				"config",
			},
			ExpDefs: map[string]string{
				"config":   "conf",
				"panic":    "panic",
				"readonly": "ro",
				"verbose":  "v",
				"version":  "ver",
				"warning":  "warn",
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s -p`, t.Name()),
			ExpItems: []string{
				`anic`,
			},
			ExpDefs: map[string]string{
				`anic`: `panic`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s -panic -w`, t.Name()),
			ExpItems: []string{
				`arning`,
			},
			ExpDefs: map[string]string{
				`arning`: `warn`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s -panic -c`, t.Name()),
			ExpItems: []string{
				`onfig`,
			},
			ExpDefs: map[string]string{
				`onfig`: `conf`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s -panic -config R`, t.Name()),
			ExpItems: []string{
				`EADME.md`,
			},
			ExpDefs: map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteDynamic(t *testing.T) {
	json := `
		[
			{
				"AllowMultiple": true,
				"Dynamic": ({
					a: [Monday..Friday]
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteDynamicDesc(t *testing.T) {
	json := `
		[
			{
				"AllowMultiple": true,
				"DynamicDesc": ({
					map { a: [Monday..Friday] } { a: [1..5] }
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
			},
			ExpDefs: map[string]string{
				`Monday`:    `1`,
				`Tuesday`:   `2`,
				`Wednesday`: `3`,
				`Thursday`:  `4`,
				`Friday`:    `5`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{
				`uesday`:  `2`,
				`hursday`: `4`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
			},
			ExpDefs: map[string]string{
				`Monday`:    `1`,
				`Tuesday`:   `2`,
				`Wednesday`: `3`,
				`Thursday`:  `4`,
				`Friday`:    `5`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{
				`uesday`:  `2`,
				`hursday`: `4`,
			},
		},
	}

	testAutocompleteFlags(t, tests)
}

/////

func TestAutocompleteDynamicArrayChain(t *testing.T) {
	json := `
		[
			{
				"Dynamic": ({
					a: [Monday..Friday]
				})
			},
			{
				"Dynamic": ({
					a: [Jan..Mar]
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteDynamicArrayChainOptional(t *testing.T) {
	json := `
		[
			{
				"Optional": true,
				"Dynamic": ({
					a: [Monday..Friday]
				})
			},
			{
				"Dynamic": ({
					a: [Jan..Mar]
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s F`, t.Name()),
			ExpItems: []string{
				`riday`,
				`eb`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s J`, t.Name()),
			ExpItems: []string{
				`an`,
			},
			ExpDefs: map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteDynamicArrayChainOptionalMultiple(t *testing.T) {
	json := `
		[
			{
				"Optional": true,
				"AllowMultiple": true,
				"Dynamic": ({
					a: [Monday..Friday]
				})
			},
			{
				"Dynamic": ({
					a: [Jan..Mar]
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s F`, t.Name()),
			ExpItems: []string{
				`riday`,
				`eb`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s J`, t.Name()),
			ExpItems: []string{
				`an`,
			},
			ExpDefs: map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

/////

func TestAutocompleteDynamicDescArrayChain(t *testing.T) {
	json := `
		[
			{
				"DynamicDesc": ({
					map { a: [Monday..Friday] } { a: [Monday..Friday] }
				})
			},
			{
				"DynamicDesc": ({
					map { a: [Jan..Mar] } { a: [Jan..Mar] }
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
			},
			ExpDefs: map[string]string{
				`Monday`:    `Monday`,
				`Tuesday`:   `Tuesday`,
				`Wednesday`: `Wednesday`,
				`Thursday`:  `Thursday`,
				`Friday`:    `Friday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{
				`uesday`:  `Tuesday`,
				`hursday`: `Thursday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{
				`Jan`: `Jan`,
				`Feb`: `Feb`,
				`Mar`: `Mar`,
			},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteDynamicDescArrayChainOptional(t *testing.T) {
	json := `
		[
			{
				"Optional": true,
				"DynamicDesc": ({
					map { a: [Monday..Friday] } { a: [Monday..Friday] }
				})
			},
			{
				"DynamicDesc": ({
					map { a: [Jan..Mar] } { a: [Jan..Mar] }
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{
				`Monday`:    `Monday`,
				`Tuesday`:   `Tuesday`,
				`Wednesday`: `Wednesday`,
				`Thursday`:  `Thursday`,
				`Friday`:    `Friday`,
				`Jan`:       `Jan`,
				`Feb`:       `Feb`,
				`Mar`:       `Mar`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{
				`uesday`:  `Tuesday`,
				`hursday`: `Thursday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{
				`Jan`: `Jan`,
				`Feb`: `Feb`,
				`Mar`: `Mar`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s F`, t.Name()),
			ExpItems: []string{
				`riday`,
				`eb`,
			},
			ExpDefs: map[string]string{
				`riday`: `Friday`,
				`eb`:    `Feb`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s J`, t.Name()),
			ExpItems: []string{
				`an`,
			},
			ExpDefs: map[string]string{
				`an`: `Jan`,
			},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteDynamicDescArrayChainOptionalMultiple(t *testing.T) {
	json := `
		[
			{
				"Optional": true,
				"AllowMultiple": true,
				"DynamicDesc": ({
					map { a: [Monday..Friday] } { a: [Monday..Friday] }
				})
			},
			{
				"DynamicDesc": ({
					map { a: [Jan..Mar] } { a: [Jan..Mar] }
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{
				`Monday`:    `Monday`,
				`Tuesday`:   `Tuesday`,
				`Wednesday`: `Wednesday`,
				`Thursday`:  `Thursday`,
				`Friday`:    `Friday`,
				`Jan`:       `Jan`,
				`Feb`:       `Feb`,
				`Mar`:       `Mar`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{
				`uesday`:  `Tuesday`,
				`hursday`: `Thursday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{
				`Monday`:    `Monday`,
				`Tuesday`:   `Tuesday`,
				`Wednesday`: `Wednesday`,
				`Thursday`:  `Thursday`,
				`Friday`:    `Friday`,
				`Jan`:       `Jan`,
				`Feb`:       `Feb`,
				`Mar`:       `Mar`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s F`, t.Name()),
			ExpItems: []string{
				`riday`,
				`eb`,
			},
			ExpDefs: map[string]string{
				`riday`: `Friday`,
				`eb`:    `Feb`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s J`, t.Name()),
			ExpItems: []string{
				`an`,
			},
			ExpDefs: map[string]string{
				`an`: `Jan`,
			},
		},
	}

	testAutocompleteFlags(t, tests)
}

/////

func TestAutocompleteNested(t *testing.T) {
	json := `
		[
			{
				"Flags": [ "Sunday", "Monday", "Happy Days" ],
				"FlagValues": {
					"Sunday": [{
						"Dynamic": ({ a: [1..3] })
					}],
					"Monday": [{
						"Flags": [ "a", "b", "c" ]
					}]
				}
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Sunday`,
				`Monday`,
				`Happy Days`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s S`, t.Name()),
			ExpItems: []string{
				`unday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Sunday `, t.Name()),
			ExpItems: []string{
				`1`,
				`2`,
				`3`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Monday `, t.Name()),
			ExpItems: []string{
				`a`,
				`b`,
				`c`,
			},
			ExpDefs: map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

/////

func TestAutocompleteComplexNestedDynamic(t *testing.T) {
	json := `
		[
			{
				"Optional": true,
				"Dynamic": ({
					out: optional
				})
			},
			{
				"Dynamic": ({
					a: [Sunday..Friday]
				}),
				"FlagValues": {
					"Sunday": [{
						"Dynamic": ({ a: [1..3] })
					}],
					"Monday": [{
						"Flags": [ "a", "b", "c" ]
					}]
				}
			},
			{
				"Dynamic": ({
					a: [Jan..Mar]
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`optional`,
				`Sunday`,
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Sunday `, t.Name()),
			ExpItems: []string{
				`1`,
				`2`,
				`3`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s Monday `, t.Name()),
			ExpItems: []string{
				`a`,
				`b`,
				`c`,
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine:  fmt.Sprintf(`%s z`, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
		{
			CmdLine:  fmt.Sprintf(`%s z `, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
		{
			CmdLine:  fmt.Sprintf(`%s Sunday z`, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
		{
			CmdLine:  fmt.Sprintf(`%s Sunday z `, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteComplexNestedDynamicDesc(t *testing.T) {
	json := `
		[
			{
				"Optional": true,
				"DynamicDesc": ({
					map { out "Optional" } { out "optional" }
				})
			},
			{
				"DynamicDesc": ({
					map { a: [Sunday..Friday] } { a: [sunday..friday] }
				}),
				"FlagValues": {
					"Sunday": [{
						"DynamicDesc": ({
							map { a: [1..3] } { a: [1..3] }
						})
					}],
					"Monday": [{
						"DynamicDesc": ({
							map { a: [a..c] } { a: [a..c] }
						})
					}]
				}
			},
			{
				"DynamicDesc": ({
					map { a: [Jan..Mar] } { a: [jan..mar] }
				})
			}
		]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: fmt.Sprintf(`%s`, t.Name()),
			ExpItems: []string{
				`Optional`,
				`Sunday`,
				`Monday`,
				`Tuesday`,
				`Wednesday`,
				`Thursday`,
				`Friday`,
			},
			ExpDefs: map[string]string{
				`Optional`:  `optional`,
				`Sunday`:    `sunday`,
				`Monday`:    `monday`,
				`Tuesday`:   `tuesday`,
				`Wednesday`: `wednesday`,
				`Thursday`:  `thursday`,
				`Friday`:    `friday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				`uesday`,
				`hursday`,
			},
			ExpDefs: map[string]string{
				`uesday`:  `tuesday`,
				`hursday`: `thursday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s Tuesday `, t.Name()),
			ExpItems: []string{
				`Jan`,
				`Feb`,
				`Mar`,
			},
			ExpDefs: map[string]string{
				`Jan`: `jan`,
				`Feb`: `feb`,
				`Mar`: `mar`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s Sunday `, t.Name()),
			ExpItems: []string{
				`1`,
				`2`,
				`3`,
			},
			ExpDefs: map[string]string{
				`1`: `1`,
				`2`: `2`,
				`3`: `3`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s Monday `, t.Name()),
			ExpItems: []string{
				`a`,
				`b`,
				`c`,
			},
			ExpDefs: map[string]string{
				`a`: `a`,
				`b`: `b`,
				`c`: `c`,
			},
		},
		{
			CmdLine:  fmt.Sprintf(`%s z `, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
		{
			CmdLine:  fmt.Sprintf(`%s z`, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
		{
			CmdLine:  fmt.Sprintf(`%s z `, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
		{
			CmdLine:  fmt.Sprintf(`%s Sunday z`, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
		{
			CmdLine:  fmt.Sprintf(`%s Sunday z `, t.Name()),
			ExpItems: []string{},
			ExpDefs:  map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}
