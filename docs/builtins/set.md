# _murex_ Language Guide

## Command reference: set

> Define a variable and set it's value

### Description

Defines, updates or deallocates a variable

### Usage

    <stdin> -> set var_name

    set var_name=data

### Examples

    » out "Hello, world!" -> set hw
    » out "$hw"
    Hello, World!

    » set hw="Hello, world!"
    » out "$hw"
    Hello, World!

### Details

#### Deallocation

You can unset variable names with the bang prefix:

    !set var_name

#### Scoping

Variables are only scoped inside the code block they're defined in (or any
children of that code block). For example `$foo` will return an empty string in
the following code because it's defined within a `try` block then being queried
outside of the `try` block:

    try {
        set foo=bar
    }
    out $foo

However if we define foo above the `try` block then it's value will be changed
even though it's being set inside the `try` block:

    set foo=""
    try {
        set foo=bar
    }
    out $foo

So unlike the previous example, this will return `bar`.

#### Function names

As a security feature function names cannot include variables. This is done to
reduce the risk of code executing by mistake due to executables being hidden
behind Variables.

Instead _murex_ will assume you want the output of the variable printed:

    » out "Hello, world!" -> set hw
    » $hw
    Hello, world!

#### Usage Inside Quotation Marks

Like with Bash, Perl and PHP: _murex_ will expand the variable when it is used
inside a double quotes but will escape the variable name when used inside single
quotes:

    » out "$foo"
    bar

    » out '$foo'
    $foo

### See also

* `[`
* `eval`
* `export`
* `global`
* `let`
