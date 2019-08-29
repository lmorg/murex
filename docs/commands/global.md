# _murex_ Shell Guide

## Command Reference: `global`

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

As a method:

    » out "Hello, world!" -> global hw
    » out "$hw"
    Hello, World!
    
As a function:

    » global hw="Hello, world!"
    » out "$hw"
    Hello, World!

### Detail

#### Deallocation

You can unset variable names with the bang prefix:

    !global var_name
    
#### Scoping

Variables are only scoped inside the code block they're defined in (or any
children of that code block). For example `$foo` will return an empty string in
the following code because it's defined within a `try` block then being queried
outside of the `try` block:

    » try {
    »     set foo=bar
    » }
    » out "foo: $foo"
    foo:
    
However if we define `$foo` above the `try` block then it's value will be changed
even though it is being set inside the `try` block:

    » set foo
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

It's also worth remembering that any variable defined using `set` in the shell's
FID (ie in the interactive shell) is literally the same as using `global`

Exported variables (defined via `export`) are system environmental variables.
Inside _murex_ environmental variables behave much like `global` variables
however their real purpose is passing data to external processes. For example
`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):

    » export foo=bar
    » env -> grep foo
    foo=bar
    
#### Function Names

As a security feature function names cannot include variables. This is done to
reduce the risk of code executing by mistake due to executables being hidden
behind variable names.

Instead _murex_ will assume you want the output of the variable printed:

    » out "Hello, world!" -> set hw
    » $hw
    Hello, world!
    
On the rare occasions you want to force variables to be expanded inside a
function name, then call that function via `exec`:

    » set cmd=grep
    » ls -> exec: $cmd main.go
    main.go
    
This only works for external executables. There is currently no way to call
aliases, functions nor builtins from a variable and even the above `exec` trick
is considered bad form because it reduces the readability of your shell scripts.

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
    
#### Declaration Without Values

You can declare a global without a value. However this isn't hugely useful
aside a rare few edge cases (and in which case the script might be better
written another way). However the feature is available to use none-the-less
and thus maintains consistency with `set`.

### Synonyms

* `global`
* `!global`


### See Also

* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`=` (artithmetic evaluation)](../commands/equ.md):
  Evaluate a mathmatical function
* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`export`](../commands/export.md):
  Define a local variable and set it's value
* [`let`](../commands/let.md):
  Evaluate a mathmatical function and assign to variable
* [`set`](../commands/set.md):
  Define a local variable and set it's value