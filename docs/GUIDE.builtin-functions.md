# _murex_ Shell Docs

## Command Reference

This section is a glossary of _murex_ builtin commands.

Because _murex_ is loosely modelled on the functional paradigm, it means
all language constructs are exposed via functions and those are typically
builtins because they can share the _murex_ runtime virtual machine.
However any executable command can also be called from within _murex_;
be that either via the `exec` builtin or natively like you would from any
Linux, UNIX, or even Windows command prompt.

## Other Reference Material

### Language Guides

1. [GUIDE.control-structures.md](./GUIDE.control-structures.md), which
contains builtins required for building logic.

### _murex_'s Source Code

In _murex_'s source under the `lang/builtins` path of the project files
is several directories, each hosting different categories of _murex_
builtins. From core commands through to data-types and methods.

Each package will include a README.md file with a basic summary of what
that package is used for and all you to enable or disable builtins, should
you decide to compile the shell from source.

### Shell Commands For Querying Builtins

From the shell itself: run `builtins` to list the builtin command.

If you require a manual on any of those commands, you can run `murex-docs`
to return the same markdown-formatted document as those listed below. eg

    murex-docs trypipe

## Pages

* [`!` (not)](commands/not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [`(` (brace quote)](commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`2darray` ](commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [`<>` (read pipe)](commands/readpipe.md):
  Reads from a _murex_ named pipe
* [`<stdin>` ](commands/stdin.md):
  Read the STDIN belonging to the parent code block
* [`=` (arithmetic evaluation)](commands/equ.md):
  Evaluate a mathematical function
* [`>>` (append file)](commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (truncate file)](commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`@[` (range) ](commands/range.md):
  Outputs a ranged subset of data from STDIN
* [`[[` (element)](commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](commands/index.md):
  Outputs an element from an array, map or table
* [`a` (mkarray)](commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`alias`](commands/alias.md):
  Create an alias for a command
* [`alter`](commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [`and`](commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`append`](commands/append.md):
  Add data to the end of an array
* [`autocomplete`](commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`bg`](commands/bg.md):
  Run processes in the background
* [`cast`](commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`catch`](commands/catch.md):
  Handles the exception code raised by `try` or `trypipe` 
* [`cd`](commands/cd.md):
  Change (working) directory
* [`config`](commands/config.md):
  Query or define _murex_ runtime settings
* [`cpuarch`](commands/cpuarch.md):
  Output the hosts CPU architecture
* [`cpucount`](commands/cpucount.md):
  Output the number of CPU cores available on your host
* [`debug`](commands/debug.md):
  Debugging information
* [`die`](commands/die.md):
  Terminate murex with an exit number of 1
* [`err`](commands/err.md):
  Print a line to the STDERR
* [`escape`](commands/escape.md):
  Escape or unescapes input 
* [`esccli`](commands/esccli.md):
  Escapes an array so output is valid shell code
* [`eschtml`](commands/eschtml.md):
  Encode or decodes text for HTML
* [`escurl`](commands/escurl.md):
  Encode or decodes text for the URL
* [`event`](commands/event.md):
  Event driven programming for shell scripts
* [`exec`](commands/exec.md):
  Runs an executable
* [`exit`](commands/exit.md):
  Exit murex
* [`exitnum`](commands/exitnum.md):
  Output the exit number of the previous process
* [`export`](commands/export.md):
  Define a local variable and set it's value
* [`f`](commands/f.md):
  Lists objects (eg files) in the current working directory
* [`false`](commands/false.md):
  Returns a `false` value
* [`fg`](commands/fg.md):
  Sends a background process into the foreground
* [`for`](commands/for.md):
  A more familiar iteration loop to existing developers
* [`foreach`](commands/foreach.md):
  Iterate through an array
* [`format`](commands/format.md):
  Reformat one data-type into another data-type
* [`function`](commands/function.md):
  Define a function block
* [`g`](commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [`get-type`](commands/get-type.md):
  Returns the data-type of a variable or pipe
* [`get`](commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](commands/getfile.md):
  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.
* [`global`](commands/global.md):
  Define a global variable and set it's value
* [`history`](commands/history.md):
  Outputs murex's command history
* [`if`](commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`ja`](commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`jsplit` ](commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`len` ](commands/len.md):
  Outputs the length of an array
* [`let`](commands/let.md):
  Evaluate a mathematical function and assign to variable
* [`lockfile`](commands/lockfile.md):
  Create and manage lock files
* [`man-summary`](commands/man-summary.md):
  Outputs a man page summary of a command
* [`map` ](commands/map.md):
  Creates a map from two data sources
* [`match`](commands/match.md):
  Match an exact value in an array
* [`msort` ](commands/msort.md):
  Sorts an array - data type agnostic
* [`mtac`](commands/mtac.md):
  Reverse the order of an array
* [`murex-docs`](commands/murex-docs.md):
  Displays the man pages for _murex_ builtins
* [`murex-package`](commands/murex-package.md):
  _murex_'s package manager
* [`murex-update-exe-list`](commands/murex-update-exe-list.md):
  Forces _murex_ to rescan $PATH looking for exectables
* [`null`](commands/devnull.md):
  null function. Similar to /dev/null
* [`open-image` ](commands/open-image.md):
  Renders bitmap image data on your terminal
* [`or`](commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`os`](commands/os.md):
  Output the auto-detected OS name
* [`out`](commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [`post`](commands/post.md):
  HTTP POST request with a JSON-parsable return
* [`prepend` ](commands/prepend.md):
  Add data to the start of an array
* [`pretty`](commands/pretty.md):
  Prettifies JSON to make it human readable
* [`private`](commands/private.md):
  Define a private function block
* [`pt`](commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`read`](commands/read.md):
  `read` a line of input from the user and store as a variable
* [`regexp`](commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [`runtime`](commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`rx`](commands/rx.md):
  Regexp pattern matching for file system objects (eg '.*\.txt')
* [`set`](commands/set.md):
  Define a local variable and set it's value
* [`source` ](commands/source.md):
  Import _murex_ code from another file of code block
* [`struct-keys`](commands/struct-keys.md):
  Outputs all the keys in a structure
* [`summary` ](commands/summary.md):
  Defines a summary help text for a command
* [`switch`](commands/switch.md):
  Blocks of cascading conditionals
* [`swivel-datatype`](commands/swivel-datatype.md):
  Converts tabulated data into a map of values for serialised data-types such as JSON and YAML
* [`swivel-table`](commands/swivel-table.md):
  Rotates a table by 90 degrees
* [`ta`](commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [`time` ](commands/time.md):
  Returns the execution run time of a command or block
* [`tout`](commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [`tread`](commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable
* [`true`](commands/true.md):
  Returns a `true` value
* [`try`](commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error
* [`version` ](commands/version.md):
  Get _murex_ version
* [`while`](commands/while.md):
  Loop until condition false