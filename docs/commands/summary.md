# Set Command Summary Hint (`summary`)

> Defines a summary help text for a command

## Description

`summary` define help text for a command. This is effectively like a tooltip
message that appears, by default, in blue in the interactive shell.

Normally this text is populated from the `man` pages or `murex-docs`, however
if neither exist or if you wish to override their text, then you can use
`summary` to define that text.

## Usage

### Define a commands summary

```
summary command description
```

### Undefine a summary

```
!summary command
```

## Examples

### Define a commands summary

```
» summary foobar "Hello, world!"
» runtime --summaries -> [ foobar ]
Hello, world! 
```

### Undefine a summary

```
» !summary foobar
```

## Synonyms

* `summary`
* `!summary`


## See Also

* [Check Builtin Exists (`bexists`)](../commands/bexists.md):
  Check which builtins exist
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Murex's Offline Documentation (`murex-docs`)](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [Re-Scan $PATH For Executables](../commands/murex-update-exe-list.md):
  Forces Murex to rescan $PATH looking for executables
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Shell Runtime (`builtins`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex

<hr/>

This document was generated from [builtins/core/management/shell_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/shell_doc.yaml).