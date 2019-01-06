# _murex_ Language Guide

## Command Reference: `tread`

> `read` a line of input from the user and store as a user defined *typed* variable    

### Description

A readline function to allow a line of data inputed from the terminal and then
store that as a typed variable.

### Usage

    tread: data-type "prompt" var_name
    
    <stdin> -> tread: data-type var_name

### Examples

    tread: qs "Please paste a URL: " url
    out: "The query string values included were:"
    $url -> format json
    
    out: Please paste a URL: -> tread: qs url
    out: "The query string values included were:"
    $url -> format json

### Detail

If `tread` is called as a method then the prompt string is taken from STDIN.
Otherwise the prompt string will be the first parameter. However if no prompt
string is given then `tread` will not write a prompt.

The last parameter will be the variable name to store the string read by `tread`.
This variable cannot be prefixed by dollar, `$`, otherwise the shell will write
the output of that variable as the last parameter rather than the name of the
variable.

### See Also

* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`err`](../commands/err.md):
  Print a line to the STDERR
* [`out`](../commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [`read`](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [cast](../commands/cast.md):
  
* [format](../commands/format.md):
  
* [pretty](../commands/pretty.md):
  
* [sprintf](../commands/sprintf.md):
  