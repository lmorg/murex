# Builtins Reference

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

* [`!` (not)](../commands/not-func.md):
  Reads the stdin and exit number from previous process and not's it's condition
* [`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`args` ](../commands/args.md):
  Command line flag parser for Murex shell scripting
* [`break`](../commands/break.md):
  Terminate execution of a block within your processes scope
* [`cast`](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [`catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [`cd`](../commands/cd.md):
  Change (working) directory
* [`continue`](../commands/continue.md):
  Terminate process of a block within a caller function
* [`die`](../commands/die.md):
  Terminate murex with an exit number of 1 (deprecated)
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`exit`](../commands/exit.md):
  Exit murex
* [`expr`](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [`false`](../commands/false.md):
  Returns a `false` value
* [`for`](../commands/for.md):
  A more familiar iteration loop to existing developers
* [`foreach`](../commands/foreach.md):
  Iterate through an array
* [`formap`](../commands/formap.md):
  Iterate through a map or other collection of data
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`get-type`](../commands/get-type.md):
  Returns the data-type of a variable or pipe
* [`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`is-null`](../commands/is-null.md):
  Checks if a variable is null or undefined
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`key-code`](../commands/key-code.md):
  Returns character sequences for any key pressed (ie sent from the terminal)
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [`null`](../commands/devnull.md):
  null function. Similar to /dev/null
* [`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return
* [`rand`](../commands/rand.md):
  Random field generator
* [`return`](../commands/return.md):
  Exits current function scope
* [`runmode`](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [`signal`](../commands/signal.md):
  Sends a signal RPC
* [`switch`](../commands/switch.md):
  Blocks of cascading conditionals
* [`tabulate`](../commands/tabulate.md):
  Table transformation tools
* [`test`](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts
* [`time`](../commands/time.md):
  Returns the execution run time of a command or block
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
* [`which`](../commands/which.md):
  Locate command origin
* [`while`](../commands/while.md):
  Loop until condition false
* [escape.cli](../commands/esccli.md):
  Escapes an array so output is valid shell code
* [escape.html](../commands/eschtml.md):
  Encode or decodes text for HTML
* [escape.quote](../commands/escape.md):
  Escape or unescape input
* [escape.url](../commands/escurl.md):
  Encode or decodes text for the URL
* [exec.* (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [exec.file: `exec`](../commands/exec.md):
  Runs an executable
* [exec.include (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [fs.files (`f`)](../commands/f.md):
  Lists or filters file system objects (eg files)
* [fs.glob (`g`)](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [fs.lockfile](../commands/lockfile.md):
  Create and manage lock files
* [fs.open](../commands/open.md):
  Open a file with a preferred handler
* [fs.open.image](../commands/open-image.md):
  Renders bitmap image data on your terminal
* [fs.regex (`rx`)](../commands/rx.md):
  Regexp pattern matching for file system objects (eg `.*\\.txt`)
* [fs.tmpfile (`tmp`)](../commands/tmp.md):
  Create a temporary file and write to it
* [fs.truncate (`>`)](../commands/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists
* [help.man.flags](../commands/man-get-flags.md):
  Parses man page files for command line flags 
* [help.man.summary](../commands/man-summary.md):
  Outputs a man page summary of a command
* [io.err](../commands/err.md):
  Print a line to the stderr
* [io.in (`<stdin>`)](../commands/stdin.md):
  Read the stdin belonging to the parent code block
* [io.new.pipe](../commands/pipe.md):
  Manage Murex named pipes
* [io.out](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [io.out.type (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [io.pipe (`<pipe>`)](../commands/namedpipe.md):
  Reads from a Murex named pipe
* [io.read](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [io.status (`pt`)](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [list.append](../commands/append.md):
  Add data to the end of an array
* [list.case](../commands/list.case.md):
  Changes the character case of a string or all elements in an array
* [list.join](../commands/mjoin.md):
  Joins a list or array into a single string
* [list.left](../commands/left.md):
  Left substring every item in a list
* [list.new.str (`a`)](../commands/a.md):
  A sophisticated yet simple way to build an array or list (mkarray)
* [list.new.type: `ta`](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [list.prefix](../commands/prefix.md):
  Prefix a string to every item in a list
* [list.prepend](../commands/prepend.md):
  Add data to the start of an array
* [list.regex](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [list.reverse (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [list.right](../commands/right.md):
  Right substring every item in a list
* [list.sort](../commands/msort.md):
  Sorts an array - data type agnostic
* [list.str (`match`)](../commands/match.md):
  Match an exact value in an array
* [list.suffix](../commands/suffix.md):
  Prefix a string to every item in a list
* [num.round: `round`](../commands/round.md):
  Round a number by a user defined precision
* [proc.bg](../commands/bg.md):
  Run processes in the background
* [proc.exitnum](../commands/exitnum.md):
  Output the exit number of the previous process
* [proc.fg](../commands/fg.md):
  Sends a background process into the foreground
* [proc.kill](../commands/fid-kill.md):
  Terminate a running Murex function
* [proc.kill.all](../commands/fid-killall.md):
  Terminate _all_ running Murex functions
* [proc.list](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [shell.alias](../commands/alias.md):
  Create an alias for a command
* [shell.autocomplete](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [shell.builtins.exist](../commands/bexists.md):
  Check which builtins exist
* [shell.config](../commands/config.md):
  Query or define Murex runtime settings
* [shell.debug](../commands/debug.md):
  Debugging information
* [shell.function](../commands/function.md):
  Define a function block
* [shell.history](../commands/history.md):
  Outputs murex's command history
* [shell.method](../commands/method.md):
  Define a methods supported data-types
* [shell.open: `openagent`](../commands/openagent.md):
  Creates a handler function for `open`
* [shell.packages (`murex-package`)](../commands/murex-package.md):
  Murex's package manager
* [shell.private](../commands/private.md):
  Define a private function block
* [shell.rescan.path](../commands/murex-update-exe-list.md):
  Forces Murex to rescan $PATH looking for executables
* [shell.runtime](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [shell.summary](../commands/summary.md):
  Defines a summary help text for a command
* [shell.version](../commands/version.md):
  Get Murex version
* [str.datetime: `datetime`](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [str.split](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [struct.alter](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [struct.count](../commands/count.md):
  Count items in a map, list or array
* [struct.json.pretty](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [struct.keys](../commands/struct-keys.md):
  Outputs all the keys in a structure as a file path
* [struct.new.2darray](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [struct.new.map (`map`)](../commands/map.md):
  Creates a map from two data sources
* [sys.cpu.arch](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [sys.cpu.count](../commands/cpucount.md):
  Output the number of CPU cores available on your host
* [sys.os](../commands/os.md):
  Output the auto-detected OS name
* [table.add.heading](../commands/addheading.md):
  Adds headings to a table
* [var.env: `export`](../commands/export.md):
  Define an environmental variable and set it's value
* [var.global: `global`](../commands/global.md):
  Define a global variable and set it's value
* [var.set: `set`](../commands/set.md):
  Define a local variable and set it's value

## Optional Builtins

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
* [table.select: `select`](../optional/select.md):
  Inlining SQL into shell pipelines
