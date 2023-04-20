# `source`  - Command Reference

> Import _murex_ code from another file of code block

## Description

`source` imports code from another file or code block. It can be used as either
an "import" / "include" directive (eg Python, Go, C, etc) or an "eval" (eg
Python, Perl, etc).

## Usage

Execute source from STDIN

    <stdin> -> source
    
Execute source from a file

    source: filename.mx
    
Execute a code block from parameter

    source: { code-block }

## Examples

Execute source from stdin:

    » tout: block { out: "Hello, world!" } -> source
    Hello, world!
    
Execute source from file:

    » tout: block { out: "Hello, world!" } |> example.mx
    » source: example.mx
    Hello, world!
    
Execute a code block from parameter

    » source { out: "Hello, world!" }
    Hello, world!

## Synonyms

* `source`
* `.`


## See Also

* [`args` ](../commands/args.md):
  Command line flag parser for _murex_ shell scripting
* [`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [`exec`](../commands/exec.md):
  Runs an executable
* [`fexec` ](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [`function`](../commands/function.md):
  Define a function block
* [`murex-parser` ](../commands/murex-parser.md):
  Runs the _murex_ parser against a block of code 
* [`private`](../commands/private.md):
  Define a private function block
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`version` ](../commands/version.md):
  Get _murex_ version