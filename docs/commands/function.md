# _murex_ Language Guide

## Command Reference: `function`

> Define a function block

### Description

`function` defines a block of code as a function

### Usage

    function: name { code-block }
    
    !function: command

### Examples

    » function hw { out "Hello, World!" }
    » hw
    Hello, World!
    
    » !function hw
    » hw
    exec: "hw": executable file not found in $PATH

### Detail

#### Allowed characters

Function names can only include any characters apart from dollar (`$`).
This is to prevent functions from overwriting variables (see the order of
preference below).

#### Undefining a function

Like all other definable states in _murex_, you can delete a function with
the bang prefix (see the example above).

#### Order of preferece

There is an order of preference for which commands are looked up:
1. aliases (all aliases are global)
2. murex functions (all `functions`s are global)
3. private functions (`privates` cannot be global and are scoped only to
   the module or source that defined them. You cannot call a private
   function from the interactive command line)
4. variables (dollar prefixed)
5. auto-globbing prefix: `@g`
6. murex builtins
7. external executables

### Synonyms

* `function`
* `!function`
* `func`
* `!func`


### See Also

* [`alias`](../commands/alias.md):
  Create an alias for a command
* [`export`](../commands/export.md):
  Define a local variable and set it's value
* [`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [`private`](../commands/private.md):
  Define a private function block
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [source](../commands/source.md):
  