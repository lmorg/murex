package config

import (
	"github.com/lmorg/murex/lang/types"
	"runtime"
)

/*type defaults struct {
	app   []string
	key   []string
	prop  []Properties
	mutex sync.Mutex
}

var Defaults defaults

func (d *defaults) Add(app string, key string, properties Properties) {
	d.mutex.Lock()

	d.app = append(d.app, app)
	d.key = append(d.key, key)
	d.prop = append(d.prop, properties)

	d.mutex.Unlock()
}*/

// Defaults defines the default config
func Defaults(config *Config, isInteractive bool) {
	config.Define("shell", "prompt", Properties{
		Description: "Interactive shell prompt.",
		//Default:     "{ exitnum->set: x; if { = x!=`0` } { set: prompt='\033[31m»\033[0m' } { set: prompt='\033[31m»\033[0m' }; out: murex $prompt }",
		Default:  "{ out 'murex » ' }",
		DataType: types.CodeBlock,
	})

	config.Define("shell", "prompt-multiline", Properties{
		Description: "Shell prompt when command line string spans multiple lines.",
		Default:     `{ out "$linenum » " }`,
		DataType:    types.CodeBlock,
	})

	config.Define("shell", "max-suggestions", Properties{
		Description: "Maximum number of lines with auto-completion suggestions to display.",
		Default:     10,
		DataType:    types.Integer,
	})

	config.Define("shell", "history", Properties{
		Description: "Write shell history (interactive shell) to disk.",
		Default:     true,
		DataType:    types.Boolean,
	})

	config.Define("shell", "add-colour", Properties{
		Description: "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, etc.",
		Default:     (runtime.GOOS != "windows" && isInteractive),
		DataType:    types.Boolean,
	})

	config.Define("shell", "syntax-highlighting", Properties{
		Description: "Syntax highlighting of murex code when in the interactive shell.",
		Default:     true,
		DataType:    types.Boolean,
	})

	config.Define("shell", "show-exts", Properties{
		Description: "Windows only! Auto-completes file extensions. This also affects the auto-completion parameters.",
		Default:     false,
		DataType:    types.Boolean,
	})

	//config.Define("shell", "strip-colour", Properties{
	//	Description: "Strips the colour codes (ANSI escape sequences) from all output destined for the terminal.",
	//	Default:     false,
	//	DataType:    types.Boolean,
	//})

	config.Define("csv", "separator", Properties{
		Description: "The delimiter for records in a CSV file.",
		Default:     `,`,
		DataType:    types.String,
	})

	config.Define("csv", "comment", Properties{
		Description: "The prefix token for comments in a CSV table.",
		Default:     `#`,
		DataType:    types.String,
	})

	config.Define("csv", "headings", Properties{
		Description: "CSV files include headings in their first line.",
		Default:     true,
		DataType:    types.Boolean,
	})

}
