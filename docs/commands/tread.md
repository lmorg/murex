# Read With Type (`tread`) (removed 7.x)

> `read` a line of input from the user and store as a user defined *typed* variable (deprecated)

## Description

A readline function to allow a line of data inputted from the terminal and then
store that as a typed variable.

**This builtin is now deprecated. Please use `read --datatype ...` instead**

## Usage

```
tread data-type "prompt" var_name

<stdin> -> tread data-type var_name
```

## Examples

```
tread qs "Please paste a URL: " url
out "The query string values included were:"
$url -> format json

out Please paste a URL: -> tread qs url
out "The query string values included were:"
$url -> format json
```

## Detail

If `tread` is called as a method then the prompt string is taken from stdin.
Otherwise the prompt string will be the first parameter. However if no prompt
string is given then `tread` will not write a prompt.

The last parameter will be the variable name to store the string read by `tread`.
This variable cannot be prefixed by dollar, `$`, otherwise the shell will write
the output of that variable as the last parameter rather than the name of the
variable.

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Prettify JSON](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [Read User Input (`read`)](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)

<hr/>

This document was generated from [builtins/core/io/read_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/read_doc.yaml).