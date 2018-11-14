package defaults

import (
	"runtime"
	"strings"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

// Defaults defines the default config
func Defaults(c *config.Config, isInteractive bool) {

	// --- shell ---

	c.Define("shell", "prompt", config.Properties{
		Description: "Interactive shell prompt.",
		Default:     "{ out 'murex » ' }",
		DataType:    types.CodeBlock,
	})

	c.Define("shell", "prompt-multiline", config.Properties{
		Description: "Shell prompt when command line string spans multiple lines.",
		Default:     `{ out "$linenum » " }`,
		DataType:    types.CodeBlock,
	})

	c.Define("shell", "max-suggestions", config.Properties{
		Description: "Maximum number of lines with auto-completion suggestions to display.",
		Default:     4,
		DataType:    types.Integer,
	})

	c.Define("shell", "history", config.Properties{
		Description: "Write shell history (interactive shell) to disk.",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("shell", "color", config.Properties{
		Description: "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc",
		Default:     (runtime.GOOS != "windows" && isInteractive),
		DataType:    types.Boolean,
	})

	c.Define("shell", "syntax-highlighting", config.Properties{
		Description: "Syntax highlighting of murex code when in the interactive shell.",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("shell", "show-exts", config.Properties{
		Description: "Windows only! Auto-completes file extensions. This also affects the auto-completion parameters.",
		Default:     false,
		DataType:    types.Boolean,
	})

	c.Define("shell", "show-hint-text", config.Properties{
		Description: "Display the blue hint text helper. //TODO: implement this!",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("shell", "hint-text-func", config.Properties{
		Description: "Murex function to call if the helper hint text is otherwise blank.",
		Default:     `{}`,
		DataType:    types.CodeBlock,
	})

	c.Define("shell", "show-suspend-status", config.Properties{
		Description: "Display some status information about the suspended process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs).",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("shell", "suspend-status-func", config.Properties{
		Description: "Murex function to execute when an `exec` process is suspended.",
		Default:     `{ progress $PID }`,
		DataType:    types.CodeBlock,
	})

	// TODO: Add config hooks for mime types
	c.Define("shell", "mime", config.Properties{
		Description: "Supported MIME types and their corresponding Murex data types.",
		Default:     define.GetMimes(),
		DataType:    types.Json,
	})

	// TODO: Add config hooks for mime types
	c.Define("shell", "extensions", config.Properties{
		Description: "Supported file extensions and their corresponding Murex data types.",
		Default:     define.GetFileExts(),
		DataType:    types.Json,
	})

	// --- proc ---

	c.Define("proc", "force-tty", config.Properties{
		Description: "This is used to override the red highlighting on STDERR on a per-process level.",
		Default:     false,
		DataType:    types.Boolean,
	})

	// --- test ---

	c.Define("test", "enabled", config.Properties{
		Description: "Run test cases.",
		Default:     false,
		DataType:    types.Boolean,
	})

	c.Define("test", "auto-report", config.Properties{
		Description: "Automatically report the results from test cases ran.",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("test", "report-format", config.Properties{
		Description: "Output format of the report.",
		Default:     "table",
		Options:     []string{"table", "json"},
		DataType:    types.String,
	})

	c.Define("test", "report-pipe", config.Properties{
		Description: "Redirect the test reports to a named pipe. Empty string send to terminal's STDERR",
		Default:     "",
		DataType:    types.String,
	})

	c.Define("test", "crop-message", config.Properties{
		Description: "This is the character limit for STDOUT and STDERR fields inside the report message. Set to zero, `0`, to disable message cropping.",
		Default:     30,
		DataType:    types.Integer,
	})
}

var murexProfile []string

// AppendProfile is used as a way of creating a platform specific default
// profile generated at compile time
func AppendProfile(block string) {
	murexProfile = append(murexProfile, "\n"+block+"\n")
}

// DefaultMurexProfile is what was formally the example murex_profile but
// this has now been converted into a this format so it can not only be
// auto-loaded as part of the default murex binary ship (ie more user
// friendly), but it also allows me to write a tailored murex profile per
// target platform.
func DefaultMurexProfile() []rune {
	return []rune(strings.Join(murexProfile, "\r\n\r\n"))
}
