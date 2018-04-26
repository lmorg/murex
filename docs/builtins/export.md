# _murex_ Language Guide

## Command reference: export

> Define an environmental variable and set it's value

### Description

Defines, updates or deallocates an environmental variable.

### Usage

    <stdin> -> export var_name

    export var_name=data

### Examples

    » out "Hello, world!" -> export hw
    » out "$hw"
    Hello, World!

    » export hw="Hello, world!"
    » out "$hw"
    Hello, World!

### Details

#### Deallocation

You can unset variable names with the bang prefix:

    !export var_name

For compatibility with other shells, `unset` is also supported but it's really
not an idiomatic method of deallocation since it's name is misleading and
suggests it is a deallocator for local _murex_ variables defined via `set`.

#### Scoping

Inside _murex_ environmental variables behave much like `global` variables
however their real purpose is passing data to external processes. For example
`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):

    » export foo=bar
    » env -> grep foo
    foo=bar

#### Usage Inside Quotation Marks

Like with Bash, Perl and PHP: _murex_ will expand the variable when it is used
inside a double quotes but will escape the variable name when used inside single
quotes:

    » out "$foo"
    bar

    » out '$foo'
    $foo

### Synonyms

* !export
* unset

### See also

* `eval`
* [`global`](global.md): Define a global variable and set it's value
* `let`
* [`set`](set.md): Define a local variable and set it's value
* `unset`
