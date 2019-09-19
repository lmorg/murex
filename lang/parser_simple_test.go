package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/proc/parameters"
)

func TestParserColon(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example --flag`},
		{Expected: nodes, Block: `example       --flag`},
		{Expected: nodes, Block: `example:--flag`},
		{Expected: nodes, Block: `example: --flag`},
		{Expected: nodes, Block: `example:      --flag`},
	}

	testParser(t, tests)
}

func TestParserSpace1(t *testing.T) {
	params := [][]parameters.ParamToken{
		{{
			Key:  "--flag1",
			Type: parameters.TokenTypeValue,
		}},
		{{
			Key:  "--flag2",
			Type: parameters.TokenTypeValue,
		}},
	}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `example --flag1 --flag2`},
		{Expected: nodes, Block: `example       --flag1 --flag2`},
		{Expected: nodes, Block: `example:--flag1 --flag2`},
		{Expected: nodes, Block: `example: --flag1 --flag2`},
		{Expected: nodes, Block: `example:      --flag1 --flag2`},

		{Expected: nodes, Block: `example --flag1  --flag2`},
		{Expected: nodes, Block: `example       --flag1  --flag2`},
		{Expected: nodes, Block: `example:--flag1  --flag2`},
		{Expected: nodes, Block: `example: --flag1  --flag2`},
		{Expected: nodes, Block: `example:      --flag1  --flag2`},

		{Expected: nodes, Block: `example --flag1    --flag2`},
		{Expected: nodes, Block: `example       --flag1    --flag2`},
		{Expected: nodes, Block: `example:--flag1    --flag2`},
		{Expected: nodes, Block: `example: --flag1    --flag2`},
		{Expected: nodes, Block: `example:      --flag1    --flag2`},

		{Expected: nodes, Block: `example --flag1	--flag2`},
		{Expected: nodes, Block: `example       --flag1	--flag2`},
		{Expected: nodes, Block: `example:--flag1	--flag2`},
		{Expected: nodes, Block: `example: --flag1	--flag2`},
		{Expected: nodes, Block: `example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserSpace2(t *testing.T) {
	params := [][]parameters.ParamToken{
		{{
			Key:  "--flag1",
			Type: parameters.TokenTypeValue,
		}},
		{{
			Key:  "--flag2",
			Type: parameters.TokenTypeValue,
		}},
	}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: ` example --flag1 --flag2`},
		{Expected: nodes, Block: ` example       --flag1 --flag2`},
		{Expected: nodes, Block: ` example:--flag1 --flag2`},
		{Expected: nodes, Block: ` example: --flag1 --flag2`},
		{Expected: nodes, Block: ` example:      --flag1 --flag2`},

		{Expected: nodes, Block: ` example --flag1  --flag2`},
		{Expected: nodes, Block: ` example       --flag1  --flag2`},
		{Expected: nodes, Block: ` example:--flag1  --flag2`},
		{Expected: nodes, Block: ` example: --flag1  --flag2`},
		{Expected: nodes, Block: ` example:      --flag1  --flag2`},

		{Expected: nodes, Block: ` example --flag1    --flag2`},
		{Expected: nodes, Block: ` example       --flag1    --flag2`},
		{Expected: nodes, Block: ` example:--flag1    --flag2`},
		{Expected: nodes, Block: ` example: --flag1    --flag2`},
		{Expected: nodes, Block: ` example:      --flag1    --flag2`},

		{Expected: nodes, Block: ` example --flag1	--flag2`},
		{Expected: nodes, Block: ` example       --flag1	--flag2`},
		{Expected: nodes, Block: ` example:--flag1	--flag2`},
		{Expected: nodes, Block: ` example: --flag1	--flag2`},
		{Expected: nodes, Block: ` example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserSpace3(t *testing.T) {
	params := [][]parameters.ParamToken{
		{{
			Key:  "--flag1",
			Type: parameters.TokenTypeValue,
		}},
		{{
			Key:  "--flag2",
			Type: parameters.TokenTypeValue,
		}},
	}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `  example --flag1 --flag2`},
		{Expected: nodes, Block: `  example       --flag1 --flag2`},
		{Expected: nodes, Block: `  example:--flag1 --flag2`},
		{Expected: nodes, Block: `  example: --flag1 --flag2`},
		{Expected: nodes, Block: `  example:      --flag1 --flag2`},

		{Expected: nodes, Block: `  example --flag1  --flag2`},
		{Expected: nodes, Block: `  example       --flag1  --flag2`},
		{Expected: nodes, Block: `  example:--flag1  --flag2`},
		{Expected: nodes, Block: `  example: --flag1  --flag2`},
		{Expected: nodes, Block: `  example:      --flag1  --flag2`},

		{Expected: nodes, Block: `  example --flag1    --flag2`},
		{Expected: nodes, Block: `  example       --flag1    --flag2`},
		{Expected: nodes, Block: `  example:--flag1    --flag2`},
		{Expected: nodes, Block: `  example: --flag1    --flag2`},
		{Expected: nodes, Block: `  example:      --flag1    --flag2`},

		{Expected: nodes, Block: `  example --flag1	--flag2`},
		{Expected: nodes, Block: `  example       --flag1	--flag2`},
		{Expected: nodes, Block: `  example:--flag1	--flag2`},
		{Expected: nodes, Block: `  example: --flag1	--flag2`},
		{Expected: nodes, Block: `  example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserSpace4(t *testing.T) {
	params := [][]parameters.ParamToken{
		{{
			Key:  "--flag1",
			Type: parameters.TokenTypeValue,
		}},
		{{
			Key:  "--flag2",
			Type: parameters.TokenTypeValue,
		}},
	}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `	example --flag1 --flag2`},
		{Expected: nodes, Block: `	example       --flag1 --flag2`},
		{Expected: nodes, Block: `	example:--flag1 --flag2`},
		{Expected: nodes, Block: `	example: --flag1 --flag2`},
		{Expected: nodes, Block: `	example:      --flag1 --flag2`},

		{Expected: nodes, Block: `	example --flag1  --flag2`},
		{Expected: nodes, Block: `	example       --flag1  --flag2`},
		{Expected: nodes, Block: `	example:--flag1  --flag2`},
		{Expected: nodes, Block: `	example: --flag1  --flag2`},
		{Expected: nodes, Block: `	example:      --flag1  --flag2`},

		{Expected: nodes, Block: `	example --flag1    --flag2`},
		{Expected: nodes, Block: `	example       --flag1    --flag2`},
		{Expected: nodes, Block: `	example:--flag1    --flag2`},
		{Expected: nodes, Block: `	example: --flag1    --flag2`},
		{Expected: nodes, Block: `	example:      --flag1    --flag2`},

		{Expected: nodes, Block: `	example --flag1	--flag2`},
		{Expected: nodes, Block: `	example       --flag1	--flag2`},
		{Expected: nodes, Block: `	example:--flag1	--flag2`},
		{Expected: nodes, Block: `	example: --flag1	--flag2`},
		{Expected: nodes, Block: `	example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserSpace5(t *testing.T) {
	params := [][]parameters.ParamToken{
		{{
			Key:  "--flag1",
			Type: parameters.TokenTypeValue,
		}},
		{{
			Key:  "--flag2",
			Type: parameters.TokenTypeValue,
		}},
	}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `		example --flag1 --flag2`},
		{Expected: nodes, Block: `		example       --flag1 --flag2`},
		{Expected: nodes, Block: `		example:--flag1 --flag2`},
		{Expected: nodes, Block: `		example: --flag1 --flag2`},
		{Expected: nodes, Block: `		example:      --flag1 --flag2`},

		{Expected: nodes, Block: `		example --flag1  --flag2`},
		{Expected: nodes, Block: `		example       --flag1  --flag2`},
		{Expected: nodes, Block: `		example:--flag1  --flag2`},
		{Expected: nodes, Block: `		example: --flag1  --flag2`},
		{Expected: nodes, Block: `		example:      --flag1  --flag2`},

		{Expected: nodes, Block: `		example --flag1    --flag2`},
		{Expected: nodes, Block: `		example       --flag1    --flag2`},
		{Expected: nodes, Block: `		example:--flag1    --flag2`},
		{Expected: nodes, Block: `		example: --flag1    --flag2`},
		{Expected: nodes, Block: `		example:      --flag1    --flag2`},

		{Expected: nodes, Block: `		example --flag1	--flag2`},
		{Expected: nodes, Block: `		example       --flag1	--flag2`},
		{Expected: nodes, Block: `		example:--flag1	--flag2`},
		{Expected: nodes, Block: `		example: --flag1	--flag2`},
		{Expected: nodes, Block: `		example:      --flag1	--flag2`},
	}

	testParser(t, tests)
}

func TestParserCmdQuotes1(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `'example' --flag`},
		{Expected: nodes, Block: `'example'       --flag`},
		{Expected: nodes, Block: `'example':--flag`},
		{Expected: nodes, Block: `'example': --flag`},
		{Expected: nodes, Block: `'example':      --flag`},

		{Expected: nodes, Block: `"example" --flag`},
		{Expected: nodes, Block: `"example"       --flag`},
		{Expected: nodes, Block: `"example":--flag`},
		{Expected: nodes, Block: `"example": --flag`},
		{Expected: nodes, Block: `"example":      --flag`},
	}

	testParser(t, tests)
}

func TestParserCmdQuotes2(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example ",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `'example ' --flag`},
		{Expected: nodes, Block: `'example '       --flag`},

		{Expected: nodes, Block: `"example " --flag`},
		{Expected: nodes, Block: `"example "       --flag`},
	}

	testParser(t, tests)
}

func TestParserCmdQuotes3(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example:",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `'example:' --flag`},
		{Expected: nodes, Block: `'example:'      --flag`},

		{Expected: nodes, Block: `"example:" --flag`},
		{Expected: nodes, Block: `"example:"       --flag`},
	}

	testParser(t, tests)
}

func TestParserCmdQuotes4(t *testing.T) {
	params := [][]parameters.ParamToken{{{
		Key:  "--flag2",
		Type: parameters.TokenTypeValue,
	}}}

	nodes := astNodes{{
		NewChain:    true,
		Name:        "example:--flag1",
		ParamTokens: params,
	}}

	var tests = []parserTestConditions{
		{Expected: nodes, Block: `'example:'--flag1 --flag2`},
		{Expected: nodes, Block: `"example:"--flag1 --flag2`},
	}

	testParser(t, tests)
}
