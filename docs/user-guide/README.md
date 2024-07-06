<h1>User Guide</h1>

This section contains miscellaneous documents on using and configuring the
shell and Murex's numerous features.

<h2>Table of Contents</h2>

<div id="toc">

- [Language Tour](#language-tour)
- [User Guides](#user-guides)
- [Integrations](#integrations)
- [Operators And Tokens](#operators-and-tokens)
- [Builtin Commands](#builtin-commands)
  - [Standard Builtins](#standard-builtins)
  - [Optional Builtins](#optional-builtins)
- [Data Types](#data-types)
- [Events](#events)
- [Integrations](#integrations-1)
- [API Reference](#api-reference)

</div>

## Language Tour

The [Language Tour](/tour.md) is a great introduction into the Murex language.

## User Guides

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Bang Prefix](../user-guide/bang-prefix.md):
  Bang prefixing to reverse default actions
* [Code Block Parsing](../user-guide/code-block.md):
  Overview of how code blocks are parsed
* [FileRef](../user-guide/fileref.md):
  How to track what code was loaded and from where
* [Hint Text](../user-guide/hint-text.md):
  A status bar for your shell
* [Integrations](../user-guide/integrations.md):
  Default integrations shipped with Murex
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Job Control](../user-guide/job-control.md):
  How to manage jobs with Murex
* [Modules and Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Named Pipes](../user-guide/namedpipes.md):
  A detailed breakdown of named pipes in Murex
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Profile Files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Rosetta Stone](../user-guide/rosetta-stone.md):
  A tabulated list of Bashism's and their equivalent Murex syntax
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex

## Operators And Tokens

* [Array (`@`) Token](parser/array.md):
  Expand values as an array
* [Tilde (`~`) Token](parser/tilde.md):
  Home directory path variable
* [`!` (not)](parser/not-func.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [`"Double Quote"`](parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`$Variable`](parser/scalar.md):
  Expand values as a scalar
* [`%(Brace Quote)`](parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%[]` Create Array](parser/create-array.md):
  Quickly generate arrays
* [`%{}` Create Object](parser/create-object.md):
  Quickly generate objects (dictionaries / maps)
* [`&&` And Logical Operator](parser/logical-and.md):
  Continues next operation if previous operation passes
* [`'Single Quote'`](parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](parser/brace-quote-func.md):
  Write a string to the STDOUT without new line (deprecated)
* [`*=` Multiply By Operator](parser/multiply-by.md):
  Multiplies a variable by the right hand value (expression)
* [`*` Multiplication Operator](parser/multiplication.md):
  Multiplies one numeric value with another (expression)
* [`+=` Add With Operator](parser/add-with.md):
  Adds the right hand value to a variable (expression)
* [`+` Addition Operator](parser/addition.md):
  Adds two numeric values together (expression)
* [`-=` Subtract By Operator](parser/subtract-by.md):
  Subtracts a variable by the right hand value (expression)
* [`->` Arrow Pipe](parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`-` Subtraction Operator](parser/subtraction.md):
  Subtracts one numeric value from another (expression)
* [`/=` Divide By Operator](parser/divide-by.md):
  Divides a variable by the right hand value (expression)
* [`/` Division Operator](parser/division.md):
  Divides one numeric value from another (expression)
* [`<pipe>` Read Named Pipe](parser/namedpipe.md):
  Reads from a Murex named pipe
* [`=>` Generic Pipe](parser/pipe-generic.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [`=` (arithmetic evaluation)](parser/equ.md):
  Evaluate a mathematical function (deprecated)
* [`>>` Append File](parser/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>>` Append Pipe](parser/pipe-append.md):
  Redirects STDOUT to a file and append its contents
* [`?:` Elvis Operator](parser/elvis.md):
  Returns the right operand if the left operand is falsy (expression)
* [`??` Null Coalescing Operator](parser/null-coalescing.md):
  Returns the right operand if the left operand is empty / undefined (expression)
* [`?` STDERR Pipe](parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command (DEPRECATED)
* [`[ ..Range ]`](parser/range.md):
  Outputs a ranged subset of data from STDIN
* [`[ Index ]`](parser/item-index.md):
  Outputs an element from an array, map or table
* [`[[ Element ]]`](parser/element.md):
  Outputs an element from a nested structure
* [`[{ Lambda }]`](parser/lambda.md):
  Iterate through structured data
* [`{ Curly Brace }`](parser/curly-brace.md):
  Initiates or terminates a code block
* [`|>` Truncate File](parser/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`|` POSIX Pipe](parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`||` Or Logical Operator](parser/logical-or.md):
  Continues next operation only if previous operation fails

## Builtin Commands

### Standard Builtins

* [`2darray` ](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [`<stdin>`](../commands/stdin.md):
  Read the STDIN belonging to the parent code block
* [`@g` (autoglob) ](../commands/autoglob.md):
  Command prefix to expand globbing (deprecated)
* [`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`addheading` ](../commands/addheading.md):
  Adds headings to a table
* [`alias`](../commands/alias.md):
  Create an alias for a command
* [`alter`](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`args` ](../commands/args.md):
  Command line flag parser for Murex shell scripting
* [`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`bexists`](../commands/bexists.md):
  Check which builtins exist
* [`bg`](../commands/bg.md):
  Run processes in the background
* [`break`](../commands/break.md):
  Terminate execution of a block within your processes scope
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [`cd`](../commands/cd.md):
  Change (working) directory
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`continue`](../commands/continue.md):
  Terminate process of a block within a caller function
* [`count`](../commands/count.md):
  Count items in a map, list or array
* [`cpuarch`](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [`cpucount`](../commands/cpucount.md):
  Output the number of CPU cores available on your host
* [`datetime`](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [`debug`](../commands/debug.md):
  Debugging information
* [`die`](../commands/die.md):
  Terminate murex with an exit number of 1
* [`err`](../commands/err.md):
  Print a line to the STDERR
* [`escape`](../commands/escape.md):
  Escape or unescape input
* [`esccli`](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [`eschtml`](../commands/eschtml.md):
  Encode or decodes text for HTML
* [`escurl`](../commands/escurl.md):
  Encode or decodes text for the URL
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`exec`](../commands/exec.md):
  Runs an executable
* [`exit`](../commands/exit.md):
  Exit murex
* [`exitnum`](../commands/exitnum.md):
  Output the exit number of the previous process
* [`export`](../commands/export.md):
  Define an environmental variable and set it's value
* [`expr`](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [`f`](../commands/f.md):
  Lists or filters file system objects (eg files)
* [`false`](../commands/false.md):
  Returns a `false` value
* [`fexec` ](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [`fid-kill`](../commands/fid-kill.md):
  Terminate a running Murex function
* [`fid-killall`](../commands/fid-killall.md):
  Terminate _all_ running Murex functions
* [`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [`for`](../commands/for.md):
  A more familiar iteration loop to existing developers
* [`foreach`](../commands/foreach.md):
  Iterate through an array
* [`formap`](../commands/formap.md):
  Iterate through a map or other collection of data
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`function`](../commands/function.md):
  Define a function block
* [`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [`get-type`](../commands/get-type.md):
  Returns the data-type of a variable or pipe
* [`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [`history`](../commands/history.md):
  Outputs murex's command history
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`is-null`](../commands/is-null.md):
  Checks if a variable is null or undefined
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`key-code`](../commands/key-code.md):
  Returns character sequences for any key pressed (ie sent from the terminal)
* [`left`](../commands/left.md):
  Left substring every item in a list
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`lockfile`](../commands/lockfile.md):
  Create and manage lock files
* [`man-get-flags` ](../commands/man-get-flags.md):
  Parses man page files for command line flags 
* [`man-summary`](../commands/man-summary.md):
  Outputs a man page summary of a command
* [`map`](../commands/map.md):
  Creates a map from two data sources
* [`match`](../commands/match.md):
  Match an exact value in an array
* [`method`](../commands/method.md):
  Define a methods supported data-types
* [`msort`](../commands/msort.md):
  Sorts an array - data type agnostic
* [`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [`murex-cache`](../commands/murex-cache.md):
  Management interface for Murex's cache database
* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [`murex-package`](../commands/murex-package.md):
  Murex's package manager
* [`murex-parser` ](../commands/murex-parser.md):
  Runs the Murex parser against a block of code 
* [`murex-update-exe-list`](../commands/murex-update-exe-list.md):
  Forces Murex to rescan $PATH looking for executables
* [`null`](../commands/devnull.md):
  null function. Similar to /dev/null
* [`open-image`](../commands/open-image.md):
  Renders bitmap image data on your terminal
* [`open`](../commands/open.md):
  Open a file with a preferred handler
* [`openagent`](../commands/openagent.md):
  Creates a handler function for `open`
* [`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`os`](../commands/os.md):
  Output the auto-detected OS name
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`pipe`](../commands/pipe.md):
  Manage Murex named pipes
* [`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return
* [`prefix`](../commands/prefix.md):
  Prefix a string to every item in a list
* [`prepend`](../commands/prepend.md):
  Add data to the start of an array
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`private`](../commands/private.md):
  Define a private function block
* [`pt`](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`rand`](../commands/rand.md):
  Random field generator
* [`read`](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [`return`](../commands/return.md):
  Exits current function scope
* [`right`](../commands/right.md):
  Right substring every item in a list
* [`round`](../commands/round.md):
  Round a number by a user defined precision
* [`runmode`](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [`rx`](../commands/rx.md):
  Regexp pattern matching for file system objects (eg `.*\\.txt`)
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`signal`](../commands/signal.md):
  Sends a signal RPC
* [`source`](../commands/source.md):
  Import Murex code from another file of code block
* [`struct-keys`](../commands/struct-keys.md):
  Outputs all the keys in a structure as a file path
* [`suffix`](../commands/suffix.md):
  Prefix a string to every item in a list
* [`summary` ](../commands/summary.md):
  Defines a summary help text for a command
* [`switch`](../commands/switch.md):
  Blocks of cascading conditionals
* [`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [`tabulate`](../commands/tabulate.md):
  Table transformation tools
* [`test`](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts
* [`time`](../commands/time.md):
  Returns the execution run time of a command or block
* [`tmp`](../commands/tmp.md):
  Create a temporary file and write to it
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [`tread`](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)
* [`true`](../commands/true.md):
  Returns a `true` value
* [`try`](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`tryerr`](../commands/tryerr.md):
  Handles errors inside a block of code
* [`trypipe`](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [`trypipeerr`](../commands/trypipeerr.md):
  Checks state of each function in a pipeline and exits block on error
* [`type`](../commands/type.md):
  Command type (function, builtin, alias, etc)
* [`unsafe`](../commands/unsafe.md):
  Execute a block of code, always returning a zero exit number
* [`version`](../commands/version.md):
  Get Murex version
* [`which`](../commands/which.md):
  Locate command origin
* [`while`](../commands/while.md):
  Loop until condition false

### Optional Builtins

These builtins are optional. `select` is included as part of the default build
but can be disabled without breaking functionality. The other optional builtins
are only included by default on Windows.

* [`!bz2`](../optional/bz2.md):
  Decompress a bz2 file
* [`base64` ](../optional/base64.md):
  Encode or decode a base64 string
* [`gz`](../optional/gz.md):
  Compress or decompress a gzip file
* [`qr`](../optional/qr.md):
  Creates a QR code from STDIN
* [`select`](../optional/select.md):
  Inlining SQL into shell pipelines
* [`sleep`](../optional/sleep.md):
  Suspends the shell for a number of seconds

## Data Types

* [`*` (generic)](../types/generic.md):
  generic (primitive)
* [`bool`](../types/bool.md):
  Boolean (primitive)
* [`commonlog`](../types/commonlog.md):
  Apache httpd "common" log format
* [`csv`](../types/csv.md):
  CSV files (and other character delimited tables)
* [`float` (floating point number)](../types/float.md):
  Floating point number (primitive)
* [`hcl`](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`jsonc`](../types/jsonc.md):
  Concatenated JSON
* [`jsonl`](../types/jsonl.md):
  JSON Lines
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
* [`path`](../types/path.md):
  Structured object for working with file and directory paths
* [`paths`](../types/paths.md):
  Structured array for working with `$PATH` style data
* [`str` (string)](../types/str.md):
  string (primitive)
* [`toml`](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml`](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [mxjson](../types/mxjson.md):
  Murex-flavoured JSON (deprecated)

## Events

* [`onCommandCompletion`](../events/oncommandcompletion.md):
  Trigger an event upon a command's completion
* [`onFileSystemChange`](../events/onfilesystemchange.md):
  Add a filesystem watch
* [`onKeyPress`](../events/onkeypress.md):
  Custom definable key bindings and macros
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation
* [`onPrompt`](../events/onprompt.md):
  Events triggered by changes in state of the interactive shell
* [`onSecondsElapsed`](../events/onsecondselapsed.md):
  Events triggered by time intervals
* [`onSignalReceived`](../events/onsignalreceived.md):
  Trap OS signals

## Integrations

* [ChatGPT](../integrations/chatgpt.md):
  How to enable ChatGPT hints
* [Cheat.sh](../integrations/cheatsh.md):
  Cheatsheets provided by cheat.sh
* [Kitty Integrations](../integrations/kitty.md):
  Get more out of Kitty terminal emulator
* [Makefiles / `make`](../integrations/make.md):
  `make` integrations
* [Man Pages (POSIX)](../integrations/man-pages.md):
  Linux/UNIX `man` page integrations
* [Spellcheck](../integrations/spellcheck.md):
  How to enable inline spellchecking
* [Terminology Integrations](../integrations/terminology.md):
  Get more out of Terminology terminal emulator
* [`direnv` Integrations](../integrations/direnv.md):
  Directory specific environmental variables
* [`yarn` Integrations](../integrations/yarn.md):
  Working with `yarn` and `package.json`
* [iTerm2 Integrations](../integrations/iterm2.md):
  Get more out of iTerm2 terminal emulator

## API Reference

These API docs are provided for any developers wishing to write their own builtins.

* [`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [`ReadArrayWithType()` (type)](../apis/ReadArrayWithType.md):
  Read from a data type one array element at a time and return the elements contents and data type
* [`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time
* [`lang.ArrayTemplate()` (template API)](../apis/lang.ArrayTemplate.md):
  Unmarshals a data type into a Go struct and returns the results as an array
* [`lang.ArrayWithTypeTemplate()` (template API)](../apis/lang.ArrayWithTypeTemplate.md):
  Unmarshals a data type into a Go struct and returns the results as an array with data type included
* [`lang.IndexTemplateObject()` (template API)](../apis/lang.IndexTemplateObject.md):
  Returns element(s) from a data structure
* [`lang.IndexTemplateTable()` (template API)](../apis/lang.IndexTemplateTable.md):
  Returns element(s) from a table
* [`lang.MarshalData()` (system API)](../apis/lang.MarshalData.md):
  Converts structured memory into a Murex data-type (eg for stdio)
* [`lang.UnmarshalData()` (system API)](../apis/lang.UnmarshalData.md):
  Converts a Murex data-type into structured memory
