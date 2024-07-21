# `version`

> Get Murex version

## Description

Returns Murex version number

## Usage

```
version [ flags ] -> <stdout>
```

## Examples

### Without parameters

```
» version
murex: 0.51.1200 BETA
```

### --no-app-name

```
» version --no-app-name
0.51.1200 BETA
```

### --short

```
» version --short
0.51
```

## Flags

* `--no-app-name`
    Returns full version string minus app name
* `--short`
    Returns only the major and minor version as a `num` data-type

## See Also

* [`args` ](../commands/args.md):
  Command line flag parser for Murex shell scripting
* [`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`function`](../commands/function.md):
  Define a function block
* [`murex-parser` ](../commands/murex-parser.md):
  Runs the Murex parser against a block of code 
* [`private`](../commands/private.md):
  Define a private function block
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [`source`](../commands/source.md):
  Import Murex code from another file of code block

<hr/>

This document was generated from [builtins/core/management/version_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/version_doc.yaml).