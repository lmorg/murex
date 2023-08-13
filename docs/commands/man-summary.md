# `man-summary`

> Outputs a man page summary of a command

## Description

`man-summary` reads the man pages for a given command and outputs it's
summary (if one exists).

## Usage

```
man-summary command -> <stdout>
```

## Examples

```
» man-summary: man 
man - an interface to the on-line reference manuals
```

## See Also

* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`man-get-flags` ](../commands/man-get-flags.md):
  Parses man page files for command line flags 
* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [`summary` ](../commands/summary.md):
  Defines a summary help text for a command