# _murex_ Shell Docs

## Command Reference: `version` 

> Get _murex_ version

## Description

Returns _murex_ version number

## Usage

    version [ flags ] -> <stdout>

## Examples

Ran without any parameters

    » version
    murex: 0.51.1200 BETA
    
Ran with the `--no-app-name` parameter

    » version --no-app-name
    0.51.1200 BETA
    
Ran with the `--short` parameter

    » version --short
    0.51

## Flags

* `--no-app-name`
    Returns full version string minus app name
* `--short`
    Returns only the major and minor version as a `num` data-type

## See Also

* [commands/`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [commands/`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/`source` ](../commands/source.md):
  Import _murex_ code from another file of code block
* [commands/args](../commands/args.md):
  
* [commands/murex-parser](../commands/murex-parser.md):
  
* [commands/params](../commands/params.md):
  