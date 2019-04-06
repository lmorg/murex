# _murex_ Language Guide

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

#### Function Names

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
    
However this workaround would only work for external utilities (ie executables
which are not _murex_ aliases, functions nor builtins).

#### Usage Inside Quotation Marks

Like with Bash, Perl and PHP: _murex_ will expand the variable when it is used
inside a double quotes but will escape the variable name when used inside single
quotes:

``
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
    and thus maintains consistancy with `set`.

### Synonyms

* `global`
* `!global`


### See Also

* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`export`](../commands/export.md):
  Define a local variable and set it's value
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [equ](../commands/equ.md):
  
* [let](../commands/let.md):
  