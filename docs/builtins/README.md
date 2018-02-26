# _murex_ Language Guide

## Command reference

| Command                   | Description |
| ------------------------- | ----------- |
|       [`alter`](alter.md) | Change a value within a structured data-type and pass that change along the pipeline without altering the original source input |
|     [`append`](append.md) | Add data to the end of an array |
|       [`catch`](catch.md) | Handles the exception code raised by `try` or `trypipe` |
|           [`err`](err.md) | `echo` a string to the STDERR |
|             [`if`](if.md) | Conditional statement to execute different blocks of code depending on the result of the condition |
|           [`out`](out.md) | `echo` a string to the STDOUT |
|   [`prepend`](prepend.md) | Add data to the start of an array |
|       [`print`](print.md) | Write a string to the OS STDOUT (bypassing _murex_ pipelines) |
|             [`pt`](pt.md) | Pipe telemetry. Writes data-types and bytes written |
| [`swivel-datatype`](swivel-datatype.md) | Converts tabulated data into a map of values for serialised data-types such as JSON and YAML |
| [`swivel-table`](swivel-table.md) | Rotates a table by 90 degrees |
|         [`tout`](tout.md) | `echo` a string to the STDOUT and set it's data-type |
|           [`try`](try.md) | Handles errors inside a block of code |
|   [`trypipe`](trypipe.md) | Checks state of each function in a pipeline and exits block on error |
