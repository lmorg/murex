# _murex_ Language Guide

## Command reference: read

> `read` a line of input from the user and store as a variable

### Description

A readline function to allow a line of data inputed from the terminal.

### Usage

    read: "prompt" var_name

    <stdin> -> read: var_name

### Examples

    read: "What is your name? " name
    out: "Hello $name"

    out: What is your name? -> read: name
    out: "Hello $name"

### Detail

If `read` is called as a method then the prompt string is taken from STDIN.
Otherwise the prompt string will be the first parameter. However if no prompt
string is given then `read` will not write a prompt.

The last parameter will be the variable name to store the string read by `read`.
This variable cannot be prefixed by dollar, `$`, otherwise the shell will write
the output of that variable as the last parameter rather than the name of the
variable.

The data type the `read` line will be stored as is `str` (string). If you
require this to be different then please use `tread` (typed read).

### See also

* [`(`]((.md): Write a string to the STDOUT without new line
* [`>`](>.md): Writes STDIN to disk - overwriting contents if file already exists
* [`>>`](>>.md): Writes STDIN to disk - appending contents if file already exists
* `cast`
* [`err`](err.md): `echo` a string to the STDERR
* [`out`](out.md): `echo` a string to the STDOUT with a trailing new line character
* `sprintf`
* [`tout`](tout.md): `echo` a string to the STDOUT and set it's data-type
* [`tread`](tread.md): `read` a line of input from the user and store as a user defined typed variable
* [`ttyfd`](ttyfd.md): Returns the TTY device of the parent.
