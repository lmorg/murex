# _murex_ Shell Docs

## Command Reference: `test`

> _murex_'s test framework

## Description

`pipe` creates and destroys _murex_ named pipes.

## Usage

Create pipe

    pipe: name [ pipe-type ]
    
Destroy pipe

    !pipe: name

## Examples

Create and destroy a standard pipe

    pipe: example
    
    !pipe: example
    
Create a TCP pipe

    pipe example --tcp-dial google.com:80
    bg { <example> }
    out: "GET /" -> <example>
    
    !pipe: example

## Detail

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

* `test`
* `!test`


## See Also

* [commands/`<>` (murex named pipe)](../commands/namedpipe.md):
  Reads from a _murex_ named pipe
* [commands/`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [parser/namedpipe](../parser/namedpipe.md):
  