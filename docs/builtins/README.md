# _murex_ Language Guide

## Command reference

| Command                   | Description |
| ------------------------- | ----------- |
|               [`>`](>.md) | Writes STDIN to disk - overwriting contents if file already exists |
|             [`>>`](>>.md) | Writes STDIN to disk - appending contents if file already exists |
|       [`alter`](alter.md) | Change a value within a structured data-type and pass that change along the pipeline without altering the original source input |
|     [`append`](append.md) | Add data to the end of an array |
|       [`catch`](catch.md) | Handles the exception code raised by `try` or `trypipe` |
|           [`err`](err.md) | `echo` a string to the STDERR |
|       [`event`](event.md) | Event driven programming for shell scripts |
|               [`f`](f.md) | Lists objects (eg files) in the current working directory |
|               [`g`](g.md) | Glob pattern matching for file system objects (eg *.txt) |
|           [`get`](get.md) | Makes a standard HTTP request and returns the result as a JSON object |
|   [`getfile`](getfile.md) | Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines. |
|             [`if`](if.md) | Conditional statement to execute different blocks of code depending on the result of the condition |
| [`murex-docs`](murex-docs.md) | Displays the man pages for _murex_ builtins |
|           [`out`](out.md) | `echo` a string to the STDOUT |
|         [`post`](post.md) | HTTP POST request with a JSON-parsable return |
|   [`prepend`](prepend.md) | Add data to the start of an array |
|       [`print`](print.md) | Write a string to the OS STDOUT (bypassing _murex_ pipelines) |
|             [`pt`](pt.md) | Pipe telemetry. Writes data-types and bytes written |
|             [`rx`](rx.md) | Regexp pattern matching for file system objects (eg '.*\.txt') |
|           [`set`](set.md) | Define a variable and set it's value |
| [`swivel-datatype`](swivel-datatype.md) | Converts tabulated data into a map of values for serialised data-types such as JSON and YAML |
| [`swivel-table`](swivel-table.md) | Rotates a table by 90 degrees |
|         [`tout`](tout.md) | `echo` a string to the STDOUT and set it's data-type |
|           [`try`](try.md) | Handles errors inside a block of code |
|   [`trypipe`](trypipe.md) | Checks state of each function in a pipeline and exits block on error |
|       [`ttyfd`](ttyfd.md) | Returns the TTY device of the parent. |
|       [`unset`](unset.md) | Deallocates an environmental variable (aliased to `!export`) |
