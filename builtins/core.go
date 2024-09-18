package builtins

// This is where you can add or remove built in functions. Imports here require underscoring.
// Each builtin package should include a README.md with details about its use. However the code should also be readable
// so take a look through the .go files if you're still undecided about whether to include a builtin or not.
//
// My recommendation is to keep everything in this file included, but disable any of the Go files prefixed `optional`.
import (
	// Docs
	_ "github.com/lmorg/murex/builtins/docs" // docs for builtin murex functions

	// CLI functions:
	_ "github.com/lmorg/murex/builtins/core/arraytools"   // tools for working with arrays and maps
	_ "github.com/lmorg/murex/builtins/core/autocomplete" // autocompleter cli tools
	_ "github.com/lmorg/murex/builtins/core/config"       // config cli tools
	_ "github.com/lmorg/murex/builtins/core/datatools"    // utilities for manipulating structured data
	_ "github.com/lmorg/murex/builtins/core/element"      // element function: [[ /element ]]
	_ "github.com/lmorg/murex/builtins/core/escape"       // string escaping
	_ "github.com/lmorg/murex/builtins/core/expressions"  // new expressions library
	_ "github.com/lmorg/murex/builtins/core/httpclient"   // builtins for http
	_ "github.com/lmorg/murex/builtins/core/index"        // index function: [ index ]
	_ "github.com/lmorg/murex/builtins/core/io"           // OS IO builtins
	_ "github.com/lmorg/murex/builtins/core/lists"        // tools for working with arrays and lists
	_ "github.com/lmorg/murex/builtins/core/management"   // murex management builtins
	_ "github.com/lmorg/murex/builtins/core/mkarray"      // rapidly makes arrays
	_ "github.com/lmorg/murex/builtins/core/modules"      // `murex-package`: module package management
	_ "github.com/lmorg/murex/builtins/core/open"         // `open` and default handlers
	_ "github.com/lmorg/murex/builtins/core/openimage"    // image pipelines
	_ "github.com/lmorg/murex/builtins/core/pipe"         // cli tools for named pipes
	_ "github.com/lmorg/murex/builtins/core/pretty"       // data formatting
	_ "github.com/lmorg/murex/builtins/core/processes"    // process handling
	_ "github.com/lmorg/murex/builtins/core/random"       // random data builtin
	_ "github.com/lmorg/murex/builtins/core/ranges"       // working with ranges within arrays (`@[..]`)
	_ "github.com/lmorg/murex/builtins/core/runtime"      // runtime inspection
	_ "github.com/lmorg/murex/builtins/core/structs"      // control structures
	_ "github.com/lmorg/murex/builtins/core/system"       // query the host system
	_ "github.com/lmorg/murex/builtins/core/tabulate"     // function to auto-tabulise data
	_ "github.com/lmorg/murex/builtins/core/test"         // testing framework for murex shell scripts
	_ "github.com/lmorg/murex/builtins/core/time"         // time-based utilities
	_ "github.com/lmorg/murex/builtins/core/typemgmt"     // type handling and management builtins

	// Events:
	_ "github.com/lmorg/murex/builtins/events/onCommandCompletion" // timer-based event type
	_ "github.com/lmorg/murex/builtins/events/onFileSystemChange"  // file system watcher event type
	_ "github.com/lmorg/murex/builtins/events/onKeyPress"          // readline key-press event type
	_ "github.com/lmorg/murex/builtins/events/onPreview"           // REPL preview triggers
	_ "github.com/lmorg/murex/builtins/events/onPrompt"            // REPL prompt triggers
	_ "github.com/lmorg/murex/builtins/events/onSecondsElapsed"    // triggered when a command finishes
	_ "github.com/lmorg/murex/builtins/events/onSignalReceived"    // OS signal management

	// Pipes:
	_ "github.com/lmorg/murex/builtins/pipes/file"      // writing to a file (required for history)
	_ "github.com/lmorg/murex/builtins/pipes/null"      // null interface (required!)
	_ "github.com/lmorg/murex/builtins/pipes/psuedotty" // pseudo-TTYs (required for F9 preview)
	_ "github.com/lmorg/murex/builtins/pipes/streams"   // standard interfaces for pipes (required!)
	_ "github.com/lmorg/murex/builtins/pipes/term"      // writing to the terminal / TTY (required!)

	// Data types:
	_ "github.com/lmorg/murex/builtins/types/apachelogs"  // Apache httpd "commonlog" parsing
	_ "github.com/lmorg/murex/builtins/types/boolean"     // boolean data type
	_ "github.com/lmorg/murex/builtins/types/columns"     // programmatic columns
	_ "github.com/lmorg/murex/builtins/types/csv"         // CSV data type
	_ "github.com/lmorg/murex/builtins/types/generic"     // generic data type
	_ "github.com/lmorg/murex/builtins/types/json"        // JSON data type
	_ "github.com/lmorg/murex/builtins/types/jsonconcat"  // concatenated JSON data type
	_ "github.com/lmorg/murex/builtins/types/jsonlines"   // jsonlines data type
	_ "github.com/lmorg/murex/builtins/types/null"        // null data type
	_ "github.com/lmorg/murex/builtins/types/numeric"     // formatting numeric data types (int, float, number)
	_ "github.com/lmorg/murex/builtins/types/paths"       // path and paths data type
	_ "github.com/lmorg/murex/builtins/types/querystring" // encoding values as URL query strings
	_ "github.com/lmorg/murex/builtins/types/string"      // string data type
	_ "github.com/lmorg/murex/builtins/types/toml"        // TOML data type
	_ "github.com/lmorg/murex/builtins/types/yaml"        // YAML data type

	_ "github.com/lmorg/murex/builtins/optional" // optional plugins
)
