package autocomplete_test

import (
	"fmt"
	"testing"
)

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
			CmdLine: fmt.Sprintf(`%s `, t.Name()),
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
			CmdLine: t.Name(),
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
			CmdLine: t.Name(),
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
			CmdLine: t.Name(),
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
			CmdLine: t.Name(),
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
			CmdLine: t.Name(),
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
			CmdLine: t.Name(),
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
			CmdLine: t.Name(),
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

func TestAutocompleteDynamicAllowSubstring(t *testing.T) {
	json := `
		[{
			"Dynamic": "{a [Monday..Friday]}",
			"AllowSubstring": true
		}]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: t.Name(),
			ExpItems: []string{
				"\x02Monday",
				"\x02Tuesday",
				"\x02Wednesday",
				"\x02Thursday",
				"\x02Friday",
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				"\x02Tuesday",
				"\x02Thursday",
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s ay`, t.Name()),
			ExpItems: []string{
				"\x02Monday",
				"\x02Tuesday",
				"\x02Wednesday",
				"\x02Thursday",
				"\x02Friday",
			},
			ExpDefs: map[string]string{},
		},
		{
			CmdLine: fmt.Sprintf(`%s n`, t.Name()),
			ExpItems: []string{
				"\x02Monday",
				"\x02Wednesday",
			},
			ExpDefs: map[string]string{},
		},
	}

	testAutocompleteFlags(t, tests)
}

func TestAutocompleteDynamicDescAllowSubstring(t *testing.T) {
	json := `
		[{
			"DynamicDesc": "{ map {a [Monday..Friday]} {a [Monday..Friday]} }",
			"AllowSubstring": true
		}]`

	initAutocompleteFlagsTest(t.Name(), json)

	tests := []testAutocompleteFlagsT{
		{
			CmdLine: t.Name(),
			ExpItems: []string{
				"\x02Monday",
				"\x02Tuesday",
				"\x02Wednesday",
				"\x02Thursday",
				"\x02Friday",
			},
			ExpDefs: map[string]string{
				"\x02Monday":    `Monday`,
				"\x02Tuesday":   `Tuesday`,
				"\x02Wednesday": `Wednesday`,
				"\x02Thursday":  `Thursday`,
				"\x02Friday":    `Friday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s T`, t.Name()),
			ExpItems: []string{
				"\x02Tuesday",
				"\x02Thursday",
			},
			ExpDefs: map[string]string{
				"\x02Tuesday":  `Tuesday`,
				"\x02Thursday": `Thursday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s ay`, t.Name()),
			ExpItems: []string{
				"\x02Monday",
				"\x02Tuesday",
				"\x02Wednesday",
				"\x02Thursday",
				"\x02Friday",
			},
			ExpDefs: map[string]string{
				"\x02Monday":    `Monday`,
				"\x02Tuesday":   `Tuesday`,
				"\x02Wednesday": `Wednesday`,
				"\x02Thursday":  `Thursday`,
				"\x02Friday":    `Friday`,
			},
		},
		{
			CmdLine: fmt.Sprintf(`%s n`, t.Name()),
			ExpItems: []string{
				"\x02Monday",
				"\x02Wednesday",
			},
			ExpDefs: map[string]string{
				"\x02Monday":    `Monday`,
				"\x02Wednesday": `Wednesday`,
			},
		},
	}

	testAutocompleteFlags(t, tests)
}
