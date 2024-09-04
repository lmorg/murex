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
Murex: 6.3.0687 (863/job-control)
Built: 2024-09-04 19:43:47
License: GPL v2
Copyright: 2018-2024 Laurence Morgan
```

### --no-app-name

```
» version --no-app-name
6.3.0688 (863/job-control)
```

### --short

```
» version --short
6.3
```

## Flags

* `--branch`
    The source code branch used in this build. This will typically be `master`
* `--build-date`
    Date of last code generation. This usually happens are part of the compilation process
* `--copyright`
    Prints copyright holder(s)
* `--license`
    Just print the license name
* `--license-full`
    Prints the full license terms
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