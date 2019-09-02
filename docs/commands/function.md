# _murex_ Shell Guide

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

#### Order of preference

There is an order of preference for which commands are looked up:
1. Aliases - defined via `alias`. All aliases are global
2. _murex_ functions - defined via `function`. All functions are global
3. private functions - defined via `private`. Private's cannot be global and
   are scoped only to the module or source that defined them. For example, You
   cannot call a private function from the interactive command line
4. variables (dollar prefixed) - declared via `set` or `let`
5. auto-globbing prefix: `@g`
6. murex builtins
7. external executable files

### Synonyms

* `function`
* `!function`
* `func`
* `!func`


### See Also

* [commands/`alias`](../commands/alias.md):
  Create an alias for a command
* [commands/`export`](../commands/export.md):
  Define a local variable and set it's value
* [commands/`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [commands/`global`](../commands/global.md):
  Define a global variable and set it's value
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value
* [commands/source](../commands/source.md):
  