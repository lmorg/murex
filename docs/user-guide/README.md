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
* [Modules And Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Named Pipes](../user-guide/namedpipes.md):
  A detailed breakdown of named pipes in Murex
* [Operators And Tokens](../user-guide/operators-and-tokens.md):
  A table of all supported operators and tokens
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
* [Strict Types In Expressions](../user-guide/strict-types.md):
  Expressions can auto-convert types or strictly honour data types
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex

## Operators And Tokens

* [( expression )](parser/expr-inlined.md):
  Inline expressions
* [C-style functions](parser/c-style-fun.md):
  Inlined commands for expressions and statements
* [Filter By Range `[ ..Range ]`](parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Item (`[ Index ]`)](parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](parser/element.md):
  Outputs an element from a nested structure
* [Read / Write To A Named Pipe (`<pipe>`)](parser/namedpipe.md):
  Reads from a Murex named pipe
* [Read From Stdin (`<stdin>`)](parser/stdin.md):
  Read the stdin belonging to the parent code block
* [Truncate File (`>`)](parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists
* [`"Double Quote"`](parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`$Scalar` Sigil (eg variables)](parser/scalar.md):
  Expand values as a scalar
* [`%(Brace Quote)`](parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%[]` Array Builder](parser/create-array.md):
  Quickly generate arrays
* [`%{}` Object Builder](parser/create-object.md):
  Quickly generate objects (dictionaries / maps)
* [`&&` And Logical Operator](parser/logical-and.md):
  Continues next operation if previous operation passes
* [`'Single Quote'`](parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)
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
  Pipes stdout from the left hand command to stdin of the right hand command
* [`-` Subtraction Operator](parser/subtraction.md):
  Subtracts one numeric value from another (expression)
* [`/=` Divide By Operator](parser/divide-by.md):
  Divides a variable by the right hand value (expression)
* [`/` Division Operator](parser/division.md):
  Divides one numeric value from another (expression)
* [`<~` Assign Or Merge](parser/assign-or-merge.md):
  Merges the right hand value to a variable on the left hand side (expression)
* [`=>` Generic Pipe](parser/pipe-generic.md):
  Pipes a reformatted stdout stream from the left hand command to stdin of the right hand command
* [`=` (arithmetic evaluation)](parser/equ.md):
  Evaluate a mathematical function (deprecated)
* [`>>` Append File](parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists
* [`?:` Elvis Operator](parser/elvis.md):
  Returns the right operand if the left operand is falsy (expression)
* [`??` Null Coalescing Operator](parser/null-coalescing.md):
  Returns the right operand if the left operand is empty / undefined (expression)
* [`?` stderr Pipe](parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`@Array` Sigil](parser/array.md):
  Expand values as an array
* [`[{ Lambda }]`](parser/lambda.md):
  Iterate through structured data
* [`{ Curly Brace }`](parser/curly-brace.md):
  Initiates or terminates a code block
* [`|` POSIX Pipe](parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`||` Or Logical Operator](parser/logical-or.md):
  Continues next operation only if previous operation fails
* [`~` Home Sigil](parser/tilde.md):
  Home directory path variable

## Builtin Commands

### Standard Builtins

* [ASCII And ANSI Escape Sequences (`key-code`)](../commands/key-code.md):
  Returns character sequences for any key pressed (ie sent from the terminal)
* [Add Heading (`addheading`)](../commands/addheading.md):
  Adds headings to a table
* [Add Prefix (`prefix`)](../commands/prefix.md):
  Prefix a string to every item in a list
* [Add Suffix (`suffix`)](../commands/suffix.md):
  Prefix a string to every item in a list
* [Alias Pointer (`alias`)](../commands/alias.md):
  Create an alias for a command
* [Alter Data Structure (`alter`)](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [CPU Architecture (`cpuarch`)](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [CPU Count (`cpucount`)](../commands/cpucount.md):
  Output the number of CPU cores available on your host
* [Caught Error Block (`catch`)](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [Change Directory (`cd`)](../commands/cd.md):
  Change (working) directory
* [Change Text Case (`list.case`)](../commands/list.case.md):
  Changes the character case of a string or all elements in an array
* [Check Builtin Exists (`bexists`)](../commands/bexists.md):
  Check which builtins exist
* [Command Line History (`history`)](../commands/history.md):
  Outputs murex's command history
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create 2d Array (`2darray`)](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create Map (`map`)](../commands/map.md):
  Creates a map from two data sources
* [Create Named Pipe (`pipe`)](../commands/pipe.md):
  Manage Murex named pipes
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Create Temporary File (`tmp`)](../commands/tmp.md):
  Create a temporary file and write to it
* [Date And Time Conversion (`datetime`)](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [Debugging Mode (`debug`)](../commands/debug.md):
  Debugging information
* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Define Function Arguments (`args`)](../commands/args.md):
  Command line flag parser for Murex shell scripting
* [Define Global (`global`)](../commands/global.md):
  Define a global variable and set it's value
* [Define Handlers For "`open`" (`openagent`)](../commands/openagent.md):
  Creates a handler function for `open`
* [Define Method Relationships (`method`)](../commands/method.md):
  Define a methods supported data-types
* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Disable Error Handling In Block (`unsafe`)](../commands/unsafe.md):
  Execute a block of code, always returning a zero exit number
* [Display Command Type (`type`)](../commands/type.md):
  Command type (function, builtin, alias, etc)
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Download File (`getfile`)](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Escape Command Line String (`esccli`)](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [Escape HTML (`eschtml`)](../commands/eschtml.md):
  Encode or decodes text for HTML
* [Escape URL (`escurl`)](../commands/escurl.md):
  Encode or decodes text for the URL
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [Exit Function (`return`)](../commands/return.md):
  Exits current function scope
* [Exit Murex (`exit`)](../commands/exit.md):
  Exit murex
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [False (`false`)](../commands/false.md):
  Returns a `false` value
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [For Loop (`for`)](../commands/for.md):
  A more familiar iteration loop to existing developers
* [Foreground Process (`fg`)](../commands/fg.md):
  Sends a background process into the foreground
* [Function / Module Defaults (`runmode`)](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [Generate Random Sequence (`rand`)](../commands/rand.md):
  Random field generator
* [Get Data Type (`get-type`)](../commands/get-type.md):
  Returns the data-type of a variable or pipe
* [Get Exit Code (`exitnum`)](../commands/exitnum.md):
  Output the exit number of the previous process
* [Get Pipe Status (`pt`)](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [Get Request (`get`)](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [Globbing (`g`)](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Is Value Null (`is-null`)](../commands/is-null.md):
  Checks if a variable is null or undefined
* [Join Array To String (`mjoin`)](../commands/mjoin.md):
  Joins a list or array into a single string
* [Kill All In Session (`fid-killall`)](../commands/fid-killall.md):
  Terminate all running Murex functions in current session
* [Kill Function (`fid-kill`)](../commands/fid-kill.md):
  Terminate a running Murex function
* [Left Sub-String (`left`)](../commands/left.md):
  Left substring every item in a list
* [List Filesystem Objects (`f`)](../commands/f.md):
  Lists or filters file system objects (eg files)
* [Location Of Command (`which`)](../commands/which.md):
  Locate command origin
* [Lock Files (`lockfile`)](../commands/lockfile.md):
  Create and manage lock files
* [Logic And Statements (`and`)](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [Logic Or Statements (`or`)](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [Loop While (`while`)](../commands/while.md):
  Loop until condition false
* [Man-Page Summary (`man-summary`)](../commands/man-summary.md):
  Outputs a man page summary of a command
* [Match String (`match`)](../commands/match.md):
  Match an exact value in an array
* [Murex Package Management (`murex-package`)](../commands/murex-package.md):
  Murex's package manager
* [Murex Version (`version`)](../commands/version.md):
  Get Murex version
* [Murex's Offline Documentation (`murex-docs`)](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [Next Iteration (`continue`)](../commands/continue.md):
  Terminate process of a block within a caller function
* [Not (`!`)](../commands/not-func.md):
  Reads the stdin and exit number from previous process and not's it's condition
* [Null (`null`)](../commands/devnull.md):
  null function. Similar to /dev/null
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Operating System (`os`)](../commands/os.md):
  Output the auto-detected OS name
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Parse Man-Page For Flags (`man-get-flags`)](../commands/man-get-flags.md):
  Parses man page files for command line flags 
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [Post Request (`post`)](../commands/post.md):
  HTTP POST request with a JSON-parsable return
* [Prepend To List (`prepend`)](../commands/prepend.md):
  Add data to the start of an array
* [Prettify JSON](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [Print Map / Structure Keys (`struct-keys`)](../commands/struct-keys.md):
  Outputs all the keys in a structure as a file path
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Processes Execution Time (`time`)](../commands/time.md):
  Returns the execution run time of a command or block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Quote String (`escape`)](../commands/escape.md):
  Escape or unescape input
* [Re-Scan $PATH For Executables](../commands/murex-update-exe-list.md):
  Forces Murex to rescan $PATH looking for executables
* [Read User Input (`read`)](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [Read With Type (`tread`) (removed 7.x)](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Regex Matches (`rx`)](../commands/rx.md):
  Regexp pattern matching for file system objects (eg `.*\\.txt`)
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [Render Image In Terminal (`open-image`)](../commands/open-image.md):
  Renders bitmap image data on your terminal
* [Reverse Array (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [Right Sub-String (`right`)](../commands/right.md):
  Right substring every item in a list
* [Round Number (`round`)](../commands/round.md):
  Round a number by a user defined precision
* [Set Command Summary Hint (`summary`)](../commands/summary.md):
  Defines a summary help text for a command
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Shell Script Tests (`test`)](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts
* [Sort Array (`msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [Stderr Checking In Pipes (`trypipeerr`)](../commands/trypipeerr.md):
  Checks state of each function in a pipeline and exits block on error
* [Stderr Checking In TTY (`tryerr`)](../commands/tryerr.md):
  Handles errors inside a block of code
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [Switch Conditional (`switch`)](../commands/switch.md):
  Blocks of cascading conditionals
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [Transformation Tools (`tabulate`)](../commands/tabulate.md):
  Table transformation tools
* [True (`true`)](../commands/true.md):
  Returns a `true` value
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`die`](../commands/die.md):
  Terminate murex with an exit number of 1 (deprecated)
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`murex-parser`](../commands/murex-parser.md):
  Runs the Murex parser against a block of code 
* [`signal`](../commands/signal.md):
  Sends a signal RPC

### Optional Builtins

These builtins are optional. `select` is included as part of the default build
but can be disabled without breaking functionality. The other optional builtins
are only included by default on Windows.

* [Inline SQL (`select`)](../optional/select.md):
  Inlining SQL into shell pipelines
* [`!bz2`](../optional/bz2.md):
  Decompress a bz2 file
* [`base64` ](../optional/base64.md):
  Encode or decode a base64 string
* [`gz`](../optional/gz.md):
  Compress or decompress a gzip file
* [`qr`](../optional/qr.md):
  Creates a QR code from stdin
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
