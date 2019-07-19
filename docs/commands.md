# _murex_ Language Guide

## Command Reference

This section is a catalogue of _murex_ functions (builtin commands). Because
_murex_ is loosely modelled on the functional paradigm, it means all language
constructs are exposed via functions and those are typically builtins since
they can share the _murex_ runtime virtual machine.

However any executable command can also be called from within _murex_; be
that either via the `exec` builtin or natively like you would from any Linux,
UNIX, or even Windows command prompt.

### Pages

* [`!` (not)](commands/not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [`(` (brace quote)](commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`2darray` ](commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [`>>` (append file)](commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (truncate file)](commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`@[` (range) ](commands/range.md):
  Outputs a ranged subset of data from STDIN
* [`[` (index)](commands/index.md):
  Outputs an element from an array, map or table
* [`a`](commands/a.md):
  A sophisticated yet simply way to build an array or list
* [`alter`](commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [`and`](commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`append`](commands/append.md):
  Add data to the end of an array
* [`cast`](commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`catch`](commands/catch.md):
  Handles the exception code raised by `try` or `trypipe
* [`cd`](commands/cd.md):
  Change (working) directory
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
* [`esccli`](commands/esccli.md):
  Escapes an array so output is valid shell code
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
* [`g`](commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
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
* [`lockfile`](commands/lockfile.md):
  Create and manage lock files
* [`man-summary`](commands/man-summary.md):
  Outputs a man page summary of a command
* [`map` ](commands/map.md):
  Creates a map from two data sources
* [`msort` ](commands/msort.md):
  Sorts an array - data type agnostic
* [`murex-docs`](commands/murex-docs.md):
  Displays the man pages for _murex_ builtins
* [`murex-update-exe-list`](commands/murex-update-exe-list.md):
  Forces _murex_ to rescan $PATH looking for exectables
* [`null`](commands/devnull.md):
  null function. Similar to /dev/null
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
* [`pt`](commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`read`](commands/read.md):
  `read` a line of input from the user and store as a variable
* [`rx`](commands/rx.md):
  Regexp pattern matching for file system objects (eg '.*\.txt')
* [`set`](commands/set.md):
  Define a local variable and set it's value
* [`swivel-datatype`](commands/swivel-datatype.md):
  Converts tabulated data into a map of values for serialised data-types such as JSON and YAML
* [`swivel-table`](commands/swivel-table.md):
  Rotates a table by 90 degrees
* [`ta`](commands/ta.md):
  A sophisticated yet simply way to build an array of a user defined data type
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