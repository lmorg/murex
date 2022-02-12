# _murex_ Shell Docs

## Command Reference: `fexec` 

> Execute a command or function, bypassing the usual order of precedence.

## Description

`fexec` allows you to execute a command or function, bypassing the usual order
of precedence.

## Usage

    fexec: flag command [ parameters... ] -> <stdout>
    ``` 

## Examples

    fexec: private /source/builtin/autocomplete.alias

## Flags

* `builtin`
    Execute a _murex_ builtin
* `function`
    Execute a _murex_ public function
* `help`
    Display help message
* `private`
    Execute a _murex_ private function

## Detail

### Order of precedence

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

## See Also

* [commands/`alias`](../commands/alias.md):
  Create an alias for a command
* [commands/`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [commands/`bg`](../commands/bg.md):
  Run processes in the background
* [commands/`builtins`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/`event`](../commands/event.md):
  Event driven programming for shell scripts
* [commands/`exec`](../commands/exec.md):
  Runs an executable
* [commands/`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`jobs`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [commands/`open`](../commands/open.md):
  Open a file with a preferred handler
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`source` ](../commands/source.md):
  Import _murex_ code from another file of code block