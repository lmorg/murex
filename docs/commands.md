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

* [`(` (brace quote)](docs/commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`>>` (write to new or appended file)](docs/commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (write to new or truncated file)](docs/commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists    
* [`alter`](docs/commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [`and`](docs/commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`append`](docs/commands/append.md):
  Add data to the end of an array
* [`catch`](docs/commands/catch.md):
  Handles the exception code raised by `try` or `trypipe
* [`err`](docs/commands/err.md):
  Print a line to the STDERR
* [`export`](docs/commands/export.md):
  Define a local variable and set it's value
* [`f`](docs/commands/f.md):
  Lists objects (eg files) in the current working directory
* [`g`](docs/commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [`get`](docs/commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](docs/commands/getfile.md):
  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.
* [`global`](docs/commands/global.md):
  Define a global variable and set it's value
* [`if`](docs/commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`murex-docs`](docs/commands/murex-docs.md):
  Displays the man pages for _murex_ builtins
* [`or`](docs/commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`out`](docs/commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [`post`](docs/commands/post.md):
  HTTP POST request with a JSON-parsable return
* [`prepend` ](docs/commands/prepend.md):
  Add data to the start of an array
* [`pt`](docs/commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`read`](docs/commands/read.md):
  `read` a line of input from the user and store as a variable
* [`rx`](docs/commands/rx.md):
  Regexp pattern matching for file system objects (eg '.*\.txt')
* [`set`](docs/commands/set.md):
  Define a local variable and set it's value
* [`swivel-datatype`](docs/commands/swivel-datatype.md):
  Converts tabulated data into a map of values for serialised data-types such as JSON and YAML
* [`swivel-table`](docs/commands/swivel-table.md):
  Rotates a table by 90 degrees
* [`tout`](docs/commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [`tread`](docs/commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable    
* [`try`](docs/commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](docs/commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error