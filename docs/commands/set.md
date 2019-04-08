# _murex_ Language Guide

## Command Reference: `set`

> Define a local variable and set it's value

### Description

Defines, updates or deallocates a local variable.

### Usage

    <stdin> -> set var_name
    
    # Assume value from STDIN, define the data type manually
    <stdin> -> set datatype var_name
    
    # Define value manually (data type defaults to string; `str`)
    set var_name=data
    
    # Define value and data type manually
    set datatype var_name=data
    
    # Define a variable but don't set any value
    set var_name
    set datatype var_name

### Examples

As a method:

    » out "Hello, world!" -> set hw
    » out "$hw"
    Hello, World!
    
As a function:

    » set hw="Hello, world!"
    » out "$hw"
    Hello, World!

### Detail

#### Deallocation

You can unset variable names with the bang prefix:

    !set var_name
    
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

It's also worth remembering that any variable defined in the shell's FID (ie
naked in the interactive shell or otherwise outside of a function or method) is
literally the same as using `global`

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

You can declare a variable without a value. This is largely only of use when
you want to overide the scoping of a variable inside a nested code-block.
(see the text above about variable scoping).

### Synonyms

* `set`
* `!set`


### See Also

* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`exec`](../commands/exec.md):
  Runs an executable
* [`export`](../commands/export.md):
  Define a local variable and set it's value
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [equ](../commands/equ.md):
  
* [let](../commands/let.md):
  
* [square-bracket-open](../commands/square-bracket-open.md):
  