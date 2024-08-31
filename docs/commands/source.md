# Include / Evaluate Murex Code (`source`)

> Import Murex code from another file or code block

## Description

`source` imports code from another file or code block. It can be used as either
an "import" / "include" directive (eg Python, Go, C, etc) or an "eval" (eg
Python, Perl, etc).

## Usage

### Execute source from stdin

```
<stdin> -> source
```

### Execute source from a file

```
source filename.mx
```

### Execute a code block from parameter

```
source { code-block }
```

## Examples

### Execute source from stdin

```
» tout block { out "Hello, world!" } -> source
Hello, world!
```

### Execute source from file

```
» tout block { out "Hello, world!" } |> example.mx
» source example.mx
Hello, world!
```

### Execute a code block from parameter

```
» source { out "Hello, world!" }
Hello, world!
```

## Synonyms

* `source`
* `.`


## See Also

* [Define Function Arguments (`args`)](../commands/args.md):
  Command line flag parser for Murex shell scripting
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Murex Version (`version`)](../commands/version.md):
  Get Murex version
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`murex-parser`](../commands/murex-parser.md):
  Runs the Murex parser against a block of code 

<hr/>

This document was generated from [builtins/core/management/source_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/source_doc.yaml).