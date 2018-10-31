# _murex_ Language Guide

## Command reference: global

> Define a global variable and set it's value

### Description

Defines, updates or deallocates a global variable.

### Usage

    # Assume data type and value from STDIN
    <stdin> -> global var_name

    # Assume value from STDIN, define the data type manually
    <stdin> -> global datatype var_name

    # Define value manually (data type defaults to string; `str`)
    global var_name=data

    # Define value and data type manually
    global datatype var_name=data

    # Define a variable but don't set any value
    global var_name
    global datatype var_name

### Examples

    » out "Hello, world!" -> global hw
    » out "$hw"
    Hello, World!

    » global hw="Hello, world!"
    » out "$hw"
    Hello, World!

### Details

#### Deallocation

You can unset variable names with the bang prefix:

    !global var_name

#### Scoping

Variables are generally only scoped inside the code block they're defined in
(ie when defined via `set`). For example `$foo` will return an empty string in
the following code because it's defined within a `try` block then being queried
outside of the `try` block:

    » try {
    »     set foo=bar
    » }
    » out "foo: $foo"
    foo:


However if we define `$foo` above the `try` block then it's value will be changed
even though it's being set inside the `try` block:

    » set foo=""
    » try {
    »     set foo=bar
    » }
    » out "foo: $foo"
    foo: bar

So unlike the previous example, this will return `bar`.

Where `global` differs from `set` is that the variables defined with `global`
will scoped at the global shell level (please note this is not the same as
environmental variables!) so will cascade down through all scoped code-blocks
including those running in other threads.

#### Function names

As a security feature function names cannot include variables. This is done to
reduce the risk of code executing by mistake due to executables being hidden
behind variable names.

Instead _murex_ will assume you want the output of the variable printed:

    » out "Hello, world!" -> export hw
    » $hw
    Hello, world!

On the rare occasions you want to force variables to be expanded inside a
function name, then call that function via `exec`:

    » export cmd=grep
    » ls -> exec: $cmd main.go
    main.go

#### Usage Inside Quotation Marks

Like with Bash, Perl and PHP: _murex_ will expand the variable when it is used
inside a double quotes but will escape the variable name when used inside single
quotes:

    » out "$foo"
    bar

    » out '$foo'
    $foo

    » out ($foo)
    bar

#### Declaration without values

You can

### Synonyms

* !global

### See also

* `eval`
* [`export`](export.md): Define an environmental variable and set it's value
* `let`
* [`set`](set.md): Define a local variable and set it's value
