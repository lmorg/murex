# `alias` - Command Reference

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

1. `runmode`: this is executed before the rest of the script. It is invoked by
   the pre-compiler forking process and is required to sit at the top of any
   scripts.

1. `test` and `pipe` functions also alter the behavior of the compiler and thus
   are executed ahead of any scripts.

4. private functions - defined via `private`. Private's cannot be global and
   are scoped only to the module or source that defined them. For example, You
   cannot call a private function directly from the interactive command line
   (however you can force an indirect call via `fexec`).

2. Aliases - defined via `alias`. All aliases are global.

3. _murex_ functions - defined via `function`. All functions are global.

5. Variables (dollar prefixed) which are declared via `global`, `set` or `let`.
   Also environmental variables too, declared via `export`.

6. globbing: however this only applies for commands executed in the interactive
   shell.

7. _murex_ builtins.

8. External executable files

You can override this order of precedence via the `fexec` and `exec` builtins.

## Synonyms

* `alias`
* `!alias`


## See Also

* [`exec`](../commands/exec.md):
  Runs an executable
* [`export`](../commands/export.md):
  Define an environmental variable and set it's value
* [`fexec` ](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [`function`](../commands/function.md):
  Define a function block
* [`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`method`](../commands/method.md):
  Define a methods supported data-types
* [`private`](../commands/private.md):
  Define a private function block
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`source` ](../commands/source.md):
  Import _murex_ code from another file of code block