# _murex_ Language Guide

## Command reference

| Command                   | Description |
| ------------------------- | ----------- |
|               [`>`](>.md) | Writes STDIN to disk - overwriting contents if file already exists |
|             [`>>`](>>.md) | Writes STDIN to disk - appending contents if file already exists |
|       [`alter`](alter.md) | Change a value within a structured data-type and pass that change along the pipeline without altering the original source input |
|           [`and`](and.md) | Returns `true` or `false` depending on whether multiple conditions are met |
|     [`append`](append.md) | Add data to the end of an array |
| [`brace-quote`](brace-quote.md) | Write a string to the STDOUT without new line |
|       [`catch`](catch.md) | Handles the exception code raised by `try` or `trypipe` |
|           [`err`](err.md) | Print a line to the STDERR |
|       [`event`](event.md) | Event driven programming for shell scripts |
|               [`f`](f.md) | Lists objects (eg files) in the current working directory |
|               [`g`](g.md) | Glob pattern matching for file system objects (eg *.txt) |
|           [`get`](get.md) | Makes a standard HTTP request and returns the result as a JSON object |
|   [`getfile`](getfile.md) | Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines. |
|             [`if`](if.md) | Conditional statement to execute different blocks of code depending on the result of the condition |
| [`murex-docs`](murex-docs.md) | Displays the man pages for _murex_ builtins |
|             [`or`](or.md) | Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful. |
|           [`out`](out.md) | `echo` a string to the STDOUT with a trailing new line character |
|         [`post`](post.md) | HTTP POST request with a JSON-parsable return |
|   [`prepend`](prepend.md) | Add data to the start of an array |
|             [`pt`](pt.md) | Pipe telemetry. Writes data-types and bytes written |
|         [`read`](read.md) | `read` a line of input from the user and store as a variable |
|             [`rx`](rx.md) | Regexp pattern matching for file system objects (eg '.*\.txt') |
|           [`set`](set.md) | Define a variable and set it's value |
| [`swivel-datatype`](swivel-datatype.md) | Converts tabulated data into a map of values for serialised data-types such as JSON and YAML |
| [`swivel-table`](swivel-table.md) | Rotates a table by 90 degrees |
|         [`tout`](tout.md) | Print a string to the STDOUT and set it's data-type |
|       [`tread`](tread.md) | `read` a line of input from the user and store as a user defined typed variable |
|           [`try`](try.md) | Handles errors inside a block of code |
|   [`trypipe`](trypipe.md) | Checks state of each function in a pipeline and exits block on error |
|       [`ttyfd`](ttyfd.md) | Returns the TTY device of the parent. |
|       [`unset`](unset.md) | Deallocates an environmental variable (aliased to `!export`) |
