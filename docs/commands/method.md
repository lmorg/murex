# _murex_ Shell Docs

## Command Reference: `method`

> Define a methods supported data-types

## Description

`method` defines what the typical data type would be for a function's STDIN
and STDOUT.

## Usage

    method: define name { code-block }

## Examples

    method: define name {
        "Stdin": "@Any",
        "Stdout": "json"
    }

## See Also

* [commands/`alias`](../commands/alias.md):
  Create an alias for a command
* [commands/`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_