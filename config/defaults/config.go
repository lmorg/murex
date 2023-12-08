package defaults

import (
	"os"
	"runtime"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/noglob"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/cache"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/murex/utils/spellcheck/userdictionary"
)

// Config defines the default config
func Config(c *config.Config, isInteractive bool) {

	// --- shell ---

	c.Define("shell", "prompt", config.Properties{
		Description: "Interactive shell prompt",
		Default:     "{ out 'murex » ' }",
		DataType:    types.CodeBlock,
		Global:      true,
	})

	c.Define("shell", "prompt-multiline", config.Properties{
		Description: "Shell prompt when command line string spans multiple lines",
		Default:     `{ out "$linenum » " }`,
		DataType:    types.CodeBlock,
		Global:      true,
	})

	c.Define("shell", "get-cursor-pos", config.Properties{
		Description: "Toggle support for the getCursorPos ANSI escape sequence to detect when prompt needs to write a new line (more accurate than zsh's wrap mode but also supported by fewer terminal emulators)",
		Default:     false,
		DataType:    types.Boolean,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  shell.ConfigReadGetCursorPos,
			Write: shell.ConfigWriteGetCursorPos,
		},
	})

	c.Define("shell", "max-suggestions", config.Properties{
		Description: "Maximum number of lines with auto-completion suggestions to display",
		Default:     12,
		DataType:    types.Integer,
		Global:      true,
	})

	c.Define("shell", "recursive-enabled", config.Properties{
		Description: "Enable a recursive scan through the directory hierarchy when using autocomplete against a file or directory parameter",
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "recursive-max-depth", config.Properties{
		Description: "Maximum depth to scan through when compiling the recursive list for auto-completion",
		Default:     5,
		DataType:    types.Integer,
		Global:      true,
	})

	c.Define("shell", "autocomplete-soft-timeout", config.Properties{
		Description: "Number of milliseconds (1/1000th second) to wait when running autocompletins before the task is backgrounded and the results appended to the existing completions (longer reduces responsiveness, shorter means autocompletion fields aren't sized to the longest suggestion)",
		Default:     150,
		DataType:    types.Integer,
		Global:      true,
	})

	c.Define("shell", "autocomplete-hard-timeout", config.Properties{
		Description: "Number of milliseconds (1/1000th second) to wait when running long autocompletions before the shell gives up. When timeout is reached on recursive directory lookups, only the results it had up to that point are returned",
		Default:     5000,
		DataType:    types.Integer,
		Global:      true,
	})

	c.Define("shell", "history", config.Properties{
		Description: "Write shell history (interactive shell) to disk",
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "color", config.Properties{
		Description: "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, variables, etc",
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "syntax-highlighting", config.Properties{
		Description: "Syntax highlighting of murex code when in the interactive shell",
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "hint-text-enabled", config.Properties{
		Description: "Display the interactive shell's hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions",
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "hint-text-func", config.Properties{
		Description: "Murex function to call if the helper hint text is otherwise blank",
		Default:     `{}`,
		DataType:    types.CodeBlock,
		Global:      true,
	})

	c.Define("shell", "hint-text-formatting", config.Properties{
		Description: "Any additional ANSI formatting for the hint text (typically color)",
		Default:     "{BLUE}",
		DataType:    types.String,
		Global:      true,
	})

	c.Define("shell", "pre-cache-hint-summaries", config.Properties{
		Description: "Run the command hint summary pre-cache",
		Default:     "on-start",
		Options:     []string{"on-start", "on-tab", "false"},
		DataType:    types.String,
		Global:      true,
	})

	c.Define("shell", "cache.db-enabled", config.Properties{
		Description: "Enable or disable the persistent cache.db. Typically located in: " + profile.ModulePath(),
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  cache.ReadStatus,
			Write: cache.WriteStatus,
		},
	})

	var defaultTitleBarFunc string
	if runtime.GOOS != "windows" {
		defaultTitleBarFunc = `{ out "$(USER)\@$(HOSTNAME):$(PWD)" }`
	}
	c.Define("shell", "titlebar-func", config.Properties{
		Description: "Linux and UNIX only! Murex function to define your terminal emulators title bar text while you're sat on a prompt. Carriage returns and tabs are replaced with spaces",
		Default:     defaultTitleBarFunc,
		DataType:    types.CodeBlock,
		Global:      true,
	})

	c.Define("shell", "stop-status-enabled", config.Properties{
		Description: "Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)",
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "stop-status-func", config.Properties{
		Description: "Murex function to execute when an `exec` process is stopped",
		Default:     `{}`,
		DataType:    types.CodeBlock,
		Global:      true,
	})

	c.Define("shell", "mime-types", config.Properties{
		Description: "Supported MIME types and their corresponding Murex data types",
		Default:     lang.GetMimes(),
		DataType:    types.Json,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  lang.ReadMimes,
			Write: lang.WriteMimes,
		},
	})

	c.Define("shell", "extensions", config.Properties{
		Description: "Supported file extensions and their corresponding Murex data types",
		Default:     lang.GetFileExts(),
		DataType:    types.Json,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  lang.ReadFileExtensions,
			Write: lang.WriteFileExtensions,
		},
	})

	c.Define("shell", "safe-commands", config.Properties{
		Description: "Commands whitelisted for being safe to automatically execute in the commandline pipe",
		Default:     parser.GetSafeCmds(),
		DataType:    types.Json,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  parser.ReadSafeCmds,
			Write: parser.WriteSafeCmds,
		},
	})

	c.Define("shell", "spellcheck-enabled", config.Properties{
		Description: "Enable spellchecking in the interactive prompt",
		Default:     false,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "spellcheck-func", config.Properties{
		Description: "Code block to run as part of the spellchecker (STDIN the line, STDOUT is array for misspelt words)",
		Default:     "{ -> aspell list }",
		DataType:    types.CodeBlock,
		Global:      true,
	})

	c.Define("shell", "spellcheck-user-dictionary", config.Properties{
		Description: "An array of words not to count as misspellings",
		Default:     userdictionary.Get(),
		DataType:    types.Json,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  userdictionary.Read,
			Write: userdictionary.Write,
		},
	})

	_, WSLENV := os.LookupEnv("WSLENV")
	_, WSL_DISTRO_NAME := os.LookupEnv("WSL_DISTRO_NAME")
	wsl := WSLENV || WSL_DISTRO_NAME

	c.Define("wsl", "windows-mounts", config.Properties{
		Description: "Windows mount points when running WSL (this improves autocompletion suggestions)",
		Default:     "",
		DataType:    types.String,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  autocomplete.WslMountsGet,
			Write: autocomplete.WslMountsSet,
		},
	})

	c.Define("shell", "extensions-enabled", config.Properties{
		Description: "Windows and WSL only! Auto-completes file extensions. This also affects the auto-completion parameters. This is best left `true` for WSL. You may need to run `murex-update-exe-list` to apply the changes",
		Default:     wsl || runtime.GOOS == "windows",
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "expand-globs", config.Properties{
		Description: "Expand globs in the REPL shell (globs in functions / modules will still need to be invoked via the `g` function)",
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "expand-glob-unsafe-commands", config.Properties{
		Description: "Commands blacklisted for being unsafe to glob",
		Default:     noglob.GetNoGlobCmds(),
		DataType:    types.Json,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  noglob.ReadNoGlobCmds,
			Write: noglob.WriteNoGlobCmds,
		},
	})

	c.Define("shell", "start-directory", config.Properties{
		Description: "If set, this is the default working directory for each new instance of murex. If unset murex will default to the current working directory",
		Default:     "",
		DataType:    types.String,
		Global:      true,
	})

	c.Define("shell", "preview-images", config.Properties{
		Description: "If set, file previews will display images as ANSI art rendered graphics rather than text descriptions",
		Default:     true,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "auto-cd", config.Properties{
		Description: "If set, `cd` is assumed when directory path supplied as a command",
		Default:     false,
		DataType:    types.Boolean,
		Global:      true,
	})

	// --- proc ---

	c.Define("proc", "force-tty", config.Properties{
		Description: "This is used to override the red highlighting on STDERR on a per-process level",
		Default:     false,
		DataType:    types.Boolean,
	})

	c.Define("proc", "strict-types", config.Properties{
		Description: "Enables or disables automatic type conversions in expressions. If enabled you might need to liberally use type tagging to ensure strings from random sources are treated as numbers",
		Default:     false,
		DataType:    types.Boolean,
	})

	c.Define("proc", "strict-vars", config.Properties{
		Description: "Return an error if an unset variable is used. Enabling this means all variables will need to be `set` before than can be used",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("proc", "strict-arrays", config.Properties{
		Description: "Return an error if an array is empty (applies to array subshells as well as variables)",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("proc", "echo-tmux", config.Properties{
		Description: "Echo shell function names to the tmux window name",
		Default:     false,
		DataType:    types.Boolean,
	})

	// --- test ---

	c.Define("test", "enabled", config.Properties{
		Description: "Run test cases",
		Default:     false,
		DataType:    types.Boolean,
	})

	c.Define("test", "verbose", config.Properties{
		Description: "Report all pass conditions for a defined test rather than just the pass summary",
		Default:     false,
		DataType:    types.Boolean,
	})

	c.Define("test", "auto-report", config.Properties{
		Description: "Automatically report the results from test cases ran",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("test", "report-format", config.Properties{
		Description: "Output format of the report",
		Default:     "table",
		Options:     []string{"table", "json", "csv"},
		DataType:    types.String,
	})

	c.Define("test", "report-pipe", config.Properties{
		Description: "Redirect the test reports to a named pipe. Empty string send to terminal's STDERR",
		Default:     "",
		DataType:    types.String,
	})

	c.Define("test", "crop-message", config.Properties{
		Description: "This is the character limit for the report message when the report is set to `table`. Set to zero, `0`, to disable message cropping",
		Default:     0,
		DataType:    types.Integer,
	})
}
