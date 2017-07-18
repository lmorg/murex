package config

import (
	"github.com/lmorg/murex/lang/types"
)

func defaults(config *Config) {
	config.Define("shell", "Prompt", Properties{
		Description: "Shell prompt",
		//Default:     "{ exitnum->set: x; if { = x!=`0` } { set: prompt='\033[31m»\033[0m' } { set: prompt='\033[31m»\033[0m' }; out: murex $prompt }",
		Default:  "{ out 'murex » ' }",
		DataType: types.CodeBlock,
	})

	config.Define("shell", "Csv-Separator", Properties{
		Description: "The delimiter for records in a CSV file",
		Default:     `,`,
		DataType:    types.String,
	})

	config.Define("shell", "Csv-Comment", Properties{
		Description: "The prefix token for comments in a CSV table",
		Default:     `#`,
		DataType:    types.String,
	})

	config.Define("shell", "Csv-Headings", Properties{
		Description: "CSV files include headings in their first line",
		Default:     true,
		DataType:    types.Boolean,
	})
}
