package config

import (
	"github.com/lmorg/murex/lang/types"
)

func defaults(config *Config) {
	config.Define("shell", "prompt", Properties{
		Description: "Shell prompt",
		//Default:     "{ exitnum->set: x; if { = x!=`0` } { set: prompt='\033[31m»\033[0m' } { set: prompt='\033[31m»\033[0m' }; out: murex $prompt }",
		Default:  "{ out 'murex » ' }",
		DataType: types.CodeBlock,
	})

	config.Define("shell", "prompt-multiline", Properties{
		Description: "Shell prompt when command line string spans multiple lines",
		Default:     `{ out "$linenum » " }`,
		DataType:    types.CodeBlock,
	})

	config.Define("shell", "history", Properties{
		Description: "Save shell history",
		Default:     true,
		DataType:    types.Boolean,
	})

	config.Define("shell", "add-colour", Properties{
		Description: "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, etc",
		Default:     true,
		DataType:    types.Boolean,
	})

	//config.Define("shell", "strip-colour", Properties{
	//	Description: "Strips the colour codes (ANSI escape sequences from all output destined for the terminal",
	//	Default:     true,
	//	DataType:    types.Boolean,
	//})

	config.Define("csv", "separator", Properties{
		Description: "The delimiter for records in a CSV file",
		Default:     `,`,
		DataType:    types.String,
	})

	config.Define("csv", "comment", Properties{
		Description: "The prefix token for comments in a CSV table",
		Default:     `#`,
		DataType:    types.String,
	})

	config.Define("csv", "headings", Properties{
		Description: "CSV files include headings in their first line",
		Default:     true,
		DataType:    types.Boolean,
	})
}
