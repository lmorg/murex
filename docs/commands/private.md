# _murex_ Shell Docs

## Command Reference: `private`

> Define a private function block

## Description

`private` defines a function who's scope is limited to that module or source
file.

Privates cannot be called from one module to another (unless they're wrapped
around a global `function`) and nor can they be called from the interactive
command line. The purpose of a `private` is to reduce repeated code inside
a module or source file without cluttering up the global namespace.

## Usage

    private: name { code-block }

## Examples

    # The following cannot be entered via the command line. You need to write
    # it to a file and execute it from there.
    
    private hw {
        out "Hello, World!"
    }
    
    function tom {
        hw
        out "My name is Tom."
    }
    
    function dick {
        hw
        out "My name is Dick."
    }
    
    function harry {
        hw
        out "My name is Harry."
    }

## Detail

### Allowed characters

Private names can only include any characters apart from dollar (`$`).
This is to prevent functions from overwriting variables (see the order of
preference below).

### Undefining a private

Because private functions are fixed to the source file that declares them,
there isn't much point in undefining them. Thus at this point in time, it
is not possible to do so.

### Order of preference

There is an order of preference for which commands are looked up:
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

## See Also

* [commands/`alias`](../commands/alias.md):
  Create an alias for a command
* [commands/`export`](../commands/export.md):
  Define a local variable and set it's value
* [commands/`function`](../commands/function.md):
  Define a function block
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
* [commands/`source` ](../commands/source.md):
  Import _murex_ code from another file of code block