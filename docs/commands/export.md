# _murex_ Shell Docs

## Command Reference: `export`

> Define an environmental variable and set it's value

## Description

Defines, updates or deallocates an environmental variable.

## Usage

    <stdin> -> export var_name
    
    export var_name=data

## Examples

As a method:

    » out "Hello, world!" -> export hw
    » out "$hw"
    Hello, World!
    
As a function:

    » export hw="Hello, world!"
    » out "$hw"
    Hello, World!

## Detail

### Deallocation

You can unset variable names with the bang prefix:

    !export var_name
    
For compatibility with other shells, `unset` is also supported but it's really
not an idiomatic method of deallocation since it's name is misleading and
suggests it is a deallocator for local _murex_ variables defined via `set`.

### Exporting a local or global variable

You can also export a local or global variable of the same name by specifying
that variable name without a following value. For example

    # Create a local variable called 'foo':
    » set: foo=bar
    » env -> grep: foo
    
    # Export that local variable as an environmental variable:
    » export: foo
    » env -> grep: foo
    foo=bar
    
    # Changing the value of the local variable doesn't alter the value of the environmental variable:
    » set: foo=rab
    » env -> grep: foo
    foo=bar
    » out: $foo
    rab
    
### Scoping

Variable scoping is simplified to three layers:

1. Local variables (`set`, `!set`, `let`)
2. Global variables (`global`, `!global`)
3. Environmental variables (`export`, `!export`, `unset`)

Variables are looked up in that order of too. For example a the following
code where `set` overrides both the global and environmental variable:

    » set:    foobar=1
    » global: foobar=2
    » export: foobar=3
    » out: $foobar
    1
    
#### Local variables

These are defined via `set` and `let`. They're variables that are persistent
across any blocks within a function. Functions will typically be blocks
encapsulated like so:

    function example {
        # variables scoped inside here
    }
    
...or...

    private example {
        # variables scoped inside here
    }
    
    
...however dynamic autocompletes, events, unit tests and any blocks defined in
`config` will also be triggered as functions.

Code running inside any control flow or error handing structures will be
treated as part of the same part of the same scope as the parent function:

    » function example {
    »     try {
    »         # set 'foobar' inside a `try` block
    »         set: foobar=example
    »     }
    »     # 'foobar' exists outside of `try` because it is scoped to `function`
    »     out: $foobar
    » }
    example
    
Where this behavior might catch you out is with iteration blocks which create
variables, eg `for`, `foreach` and `formap`. Any variables created inside them
are still shared with any code outside of those structures but still inside the
function block.

Any local variables are only available to that function. If a variable is
defined in a parent function that goes on to call child functions, then those
local variables are not inherited but the child functions:

    » function parent {
    »     # set a local variable
    »     set: foobar=example
    »     child
    » }
    » 
    » function child {
    »     # returns the `global` value, "not set", because the local `set` isn't inherited
    »     out: $foobar
    » }
    » 
    » global: $foobar="not set"
    » parent
    not set
    
It's also worth remembering that any variable defined using `set` in the shells
FID (ie in the interactive shell) is localised to structures running in the
interactive, REPL, shell and are not inherited by any called functions.

#### Global variables

Where `global` differs from `set` is that the variables defined with `global`
will be scoped at the global shell level (please note this is not the same as
environmental variables!) so will cascade down through all scoped code-blocks
including those running in other threads.

#### Environmental variables

Exported variables (defined via `export`) are system environmental variables.
Inside _murex_ environmental variables behave much like `global` variables
however their real purpose is passing data to external processes. For example
`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):

    » export foo=bar
    » env -> grep foo
    foo=bar
    
### Function Names

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

### Usage Inside Quotation Marks

Like with Bash, Perl and PHP: _murex_ will expand the variable when it is used
inside a double quotes but will escape the variable name when used inside single
quotes:

    » out "$foo"
    bar
    
    » out '$foo'
    $foo
    
    » out ($foo)
    bar

## Synonyms

* `export`
* `!export`
* `unset`


## See Also

* [user-guide/Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by _murex_
* [user-guide/Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within _murex_
* [commands/`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [commands/`=` (arithmetic evaluation)](../commands/equ.md):
  Evaluate a mathematical function
* [commands/`global`](../commands/global.md):
  Define a global variable and set it's value
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value