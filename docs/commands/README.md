# Command Reference

This section is a glossary of Murex builtin commands.

Because Murex is loosely modelled on the functional paradigm, it means
all language constructs are exposed via functions and those are typically
builtins because they can share the Murex runtime virtual machine.
However any executable command can also be called from within Murex;
be that either via the `exec` builtin or natively like you would from any
Linux, UNIX, or even Windows command prompt.

## Other Reference Material

### Language Guides

1. [Language Tour](/docs/tour.md), which is an introduction into
    the Murex language.

2. [Rosetta Stone](/docs/user-guide/rosetta-stone.md), which is a reference
    table comparing Bash syntax to Murex's.

### Murex's Source Code

The source for each of these builtins can be found on [Github](https://github.com/lmorg/murex/tree/master/builtins/core).

### Shell Commands For Querying Builtins

From the shell itself: run `builtins` to list the builtin command.

If you require a manual on any of those commands, you can run `murex-docs`
to return the same markdown-formatted document as those listed below. eg

```
murex-docs trypipe
```

## Pages


### Input / Output Streams

Commands for managing the flow of data between different processes and/or the terminal screen.

* [Create Named Pipe (`pipe`)](../commands/pipe.md):
  Manage Murex named pipes
* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Get Pipe Status (`pt`)](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Read / Write To A Named Pipe (`<pipe>`)](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [Read From Stdin (`<stdin>`)](../parser/stdin.md):
  Read the stdin belonging to the parent code block
* [Read User Input (`read`)](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [Read With Type (`tread`) (removed 7.x)](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)
* [Render Image In Terminal (`open-image`)](../commands/open-image.md):
  Renders bitmap image data on your terminal
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)

### Filesystem Operations

Commands for working with files and/or the filesystem.

* [Create Temporary File (`tmp`)](../commands/tmp.md):
  Create a temporary file and write to it
* [Globbing (`g`)](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [List Filesystem Objects (`f`)](../commands/f.md):
  Lists or filters file system objects (eg files)
* [Lock Files (`lockfile`)](../commands/lockfile.md):
  Create and manage lock files
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Regex Matches (`rx`)](../commands/rx.md):
  Regexp pattern matching for file system objects (eg `.*\\.txt`)
* [Render Image In Terminal (`open-image`)](../commands/open-image.md):
  Renders bitmap image data on your terminal
* [Truncate File (`>`)](../parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists

### Defined by POSIX

Commands defined by POSIX.

* [Alias Pointer (`alias`)](../commands/alias.md):
  Create an alias for a command
* [Change Directory (`cd`)](../commands/cd.md):
  Change (working) directory
* [Display Command Type (`type`)](../commands/type.md):
  Command type (function, builtin, alias, etc)
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [False (`false`)](../commands/false.md):
  Returns a `false` value
* [Location Of Command (`which`)](../commands/which.md):
  Locate command origin
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Processes Execution Time (`time`)](../commands/time.md):
  Returns the execution run time of a command or block
* [Read User Input (`read`)](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [True (`true`)](../commands/true.md):
  Returns a `true` value

### List / Array Editing

Commands that operate against a list or array.

* [Add Prefix (`prefix`)](../commands/prefix.md):
  Prefix a string to every item in a list
* [Add Suffix (`suffix`)](../commands/suffix.md):
  Prefix a string to every item in a list
* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Change Text Case (`list.case`)](../commands/list.case.md):
  Changes the character case of a string or all elements in an array
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Join Array To String (`mjoin`)](../commands/mjoin.md):
  Joins a list or array into a single string
* [Left Sub-String (`left`)](../commands/left.md):
  Left substring every item in a list
* [Match String (`match`)](../commands/match.md):
  Match an exact value in an array
* [Prepend To List (`prepend`)](../commands/prepend.md):
  Add data to the start of an array
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [Reverse Array (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [Right Sub-String (`right`)](../commands/right.md):
  Right substring every item in a list
* [Sort Array (`msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

### String Manipulation

Commands for working with strings.
> All list based tools also work with strings.

* [Change Text Case (`list.case`)](../commands/list.case.md):
  Changes the character case of a string or all elements in an array
* [Date And Time Conversion (`datetime`)](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [Escape Command Line String (`esccli`)](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [Escape HTML (`eschtml`)](../commands/eschtml.md):
  Encode or decodes text for HTML
* [Escape URL (`escurl`)](../commands/escurl.md):
  Encode or decodes text for the URL
* [Generate Random Sequence (`rand`)](../commands/rand.md):
  Random field generator
* [Left Sub-String (`left`)](../commands/left.md):
  Left substring every item in a list
* [Prettify JSON](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [Quote String (`escape`)](../commands/escape.md):
  Escape or unescape input
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [Right Sub-String (`right`)](../commands/right.md):
  Right substring every item in a list
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter

### Numeric / Math Tools

Commands for working with numerical data.

* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Generate Random Sequence (`rand`)](../commands/rand.md):
  Random field generator
* [Round Number (`round`)](../commands/round.md):
  Round a number by a user defined precision

### Structured Data Management

Commands for working with structured data such as maps, tables, arrays and other data formats that are present in documents such as CSV, JSON, YAML, TOML, Sexpr, CSV, etc. 

* [Alter Data Structure (`alter`)](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create 2d Array (`2darray`)](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [Create Map (`map`)](../commands/map.md):
  Creates a map from two data sources
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Print Map / Structure Keys (`struct-keys`)](../commands/struct-keys.md):
  Outputs all the keys in a structure as a file path
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Transformation Tools (`tabulate`)](../commands/tabulate.md):
  Table transformation tools

### Table Management

Commands specifically for working with tabulated data.

* [Add Heading (`addheading`)](../commands/addheading.md):
  Adds headings to a table
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [Inline SQL (`select`)](../optional/select.md):
  Inlining SQL into shell pipelines
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Transformation Tools (`tabulate`)](../commands/tabulate.md):
  Table transformation tools

### System Inspection

Tools to inspect the host system.

* [CPU Architecture (`cpuarch`)](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [CPU Count (`cpucount`)](../commands/cpucount.md):
  Output the number of CPU cores available on your host
* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Operating System (`os`)](../commands/os.md):
  Output the auto-detected OS name

### Shell / Murex Management

Commands to manage the Murex shell session.

* [Alias Pointer (`alias`)](../commands/alias.md):
  Create an alias for a command
* [Command Line History (`history`)](../commands/history.md):
  Outputs murex's command history
* [Create Named Pipe (`pipe`)](../commands/pipe.md):
  Manage Murex named pipes
* [Debugging Mode (`debug`)](../commands/debug.md):
  Debugging information
* [Define Handlers For "`open`" (`openagent`)](../commands/openagent.md):
  Creates a handler function for `open`
* [Define Method Relationships (`method`)](../commands/method.md):
  Define a methods supported data-types
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Murex Package Management (`murex-package`)](../commands/murex-package.md):
  Murex's package manager
* [Murex Version (`version`)](../commands/version.md):
  Get Murex version
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Re-Scan $PATH For Executables](../commands/murex-update-exe-list.md):
  Forces Murex to rescan $PATH looking for executables
* [Set Command Summary Hint (`summary`)](../commands/summary.md):
  Defines a summary help text for a command
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Shell Script Tests (`test`)](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`event`](../commands/event.md):
  Event driven programming for shell scripts

### String Escaping / Character Codes

Commands to escape special characters in various different string formats.

* [ASCII And ANSI Escape Sequences (`key-code`)](../commands/key-code.md):
  Returns character sequences for any key pressed (ie sent from the terminal)
* [Escape Command Line String (`esccli`)](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [Escape HTML (`eschtml`)](../commands/eschtml.md):
  Encode or decodes text for HTML
* [Escape URL (`escurl`)](../commands/escurl.md):
  Encode or decodes text for the URL
* [Quote String (`escape`)](../commands/escape.md):
  Escape or unescape input

### Process Management

Management of system processes and Murex FIDs.

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Check Builtin Exists (`bexists`)](../commands/bexists.md):
  Check which builtins exist
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Foreground Process (`fg`)](../commands/fg.md):
  Sends a background process into the foreground
* [Get Exit Code (`exitnum`)](../commands/exitnum.md):
  Output the exit number of the previous process
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Kill All In Session (`fid-killall`)](../commands/fid-killall.md):
  Terminate all running Murex functions in current session
* [Kill Function (`fid-kill`)](../commands/fid-kill.md):
  Terminate a running Murex function
* [Location Of Command (`which`)](../commands/which.md):
  Locate command origin
* [Processes Execution Time (`time`)](../commands/time.md):
  Returns the execution run time of a command or block
* [`signal`](../commands/signal.md):
  Sends a signal RPC

### Language And Scripting

Various commands that enable control flow, error handling and other important characteristics that turn Murex into a functional programming language.

* [Define Function Arguments (`args`)](../commands/args.md):
  Command line flag parser for Murex shell scripting
* [Define Global (`global`)](../commands/global.md):
  Define a global variable and set it's value
* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [Exit Function (`return`)](../commands/return.md):
  Exits current function scope
* [Exit Murex (`exit`)](../commands/exit.md):
  Exit murex
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [For Loop (`for`)](../commands/for.md):
  A more familiar iteration loop to existing developers
* [Get Data Type (`get-type`)](../commands/get-type.md):
  Returns the data-type of a variable or pipe
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Is Value Null (`is-null`)](../commands/is-null.md):
  Checks if a variable is null or undefined
* [Logic And Statements (`and`)](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [Logic Or Statements (`or`)](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [Loop While (`while`)](../commands/while.md):
  Loop until condition false
* [Next Iteration (`continue`)](../commands/continue.md):
  Terminate process of a block within a caller function
* [Not (`!`)](../commands/not-func.md):
  Reads the stdin and exit number from previous process and not's it's condition
* [Null (`null`)](../commands/devnull.md):
  null function. Similar to /dev/null
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Switch Conditional (`switch`)](../commands/switch.md):
  Blocks of cascading conditionals

### Error Handling

Tools and control flow structures to handle errors.

* [Caught Error Block (`catch`)](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [Disable Error Handling In Block (`unsafe`)](../commands/unsafe.md):
  Execute a block of code, always returning a zero exit number
* [Function / Module Defaults (`runmode`)](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [Stderr Checking In Pipes (`trypipeerr`)](../commands/trypipeerr.md):
  Checks state of each function in a pipeline and exits block on error
* [Stderr Checking In TTY (`tryerr`)](../commands/tryerr.md):
  Handles errors inside a block of code
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code

### Help and Hint Tools

Tools for providing help and hints, useful when working inside the interactive shell.

* [Man-Page Summary (`man-summary`)](../commands/man-summary.md):
  Outputs a man page summary of a command
* [Murex's Offline Documentation (`murex-docs`)](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [Parse Man-Page For Flags (`man-get-flags`)](../commands/man-get-flags.md):
  Parses man page files for command line flags 
* [Set Command Summary Hint (`summary`)](../commands/summary.md):
  Defines a summary help text for a command

### Uncategorised

* [`die`](../commands/die.md):
  Terminate murex with an exit number of 1 (deprecated)
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`murex-parser`](../commands/murex-parser.md):
  Runs the Murex parser against a block of code 

## Optional Builtins

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
