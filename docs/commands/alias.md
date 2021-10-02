# _murex_ Shell Docs

## Command Reference: `alias`

> Create an alias for a command

## Description

`alias` defines an alias for global usage

## Usage

    alias: alias=command parameter parameter
    
    !alias: command

## Examples

Because aliases are parsed into an array of parameters, you cannot put the
entire alias within quotes. For example:

    # bad :(
    » alias hw="out Hello, World!"
    » hw
    exec: "out\\ Hello,\\ World!": executable file not found in $PATH
    
    # good :)
    » alias hw=out "Hello, World!"
    » hw
    Hello, World!
    
Notice how only the command `out "Hello, World!"` is quoted in `alias` the
same way you would have done if you'd run that command "naked" in the command
line? This is how `alias` expects it's parameters and where `alias` on _murex_
differs from `alias` in POSIX shells.

In some ways this makes `alias` a little less flexible than it might
otherwise be. However the design of this is to keep `alias` focused on it's
core objective. For any more advanced requirements you can use a `function`
instead.

## Detail

### Allowed characters

Alias names can only include alpha-numeric characters, hyphen and underscore.
The following regex is used to validate the `alias`'s parameters:
`^([-_a-zA-Z0-9]+)=(.*?)$`

### Undefining an alias

Like all other definable states in _murex_, you can delete an alias with the
bang prefix:

    » alias hw=out "Hello, World!"
    » hw
    Hello, World!
    
    » !alias hw
    » hw
    exec: "hw": executable file not found in $PATH
    
### Order of preference

There is an order of precedence for which commands are looked up:
1. `test` and `pipe` functions because they alter the behavior of the compiler
2. Aliases - defined via `alias`. All aliases are global
3. _murex_ functions - defined via `function`. All functions are global
4. private functions - defined via `private`. Private's cannot be global and
   are scoped only to the module or source that defined them. For example, You
   cannot call a private function from the interactive command line
5. variables (dollar prefixed) - declared via `set` or `let`
6. auto-globbing prefix: `@g`
7. murex builtins
8. external executable files

## Synonyms

* `alias`
* `!alias`


## See Also

* [commands/`exec`](../commands/exec.md):
  Runs an executable
* [commands/`export`](../commands/export.md):
  Define an environmental variable and set it's value
* [commands/`fexec` ](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [commands/`global`](../commands/global.md):
  Define a global variable and set it's value
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable
* [commands/`method`](../commands/method.md):
  Define a methods supported data-types
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value
* [commands/`source` ](../commands/source.md):
  Import _murex_ code from another file of code block