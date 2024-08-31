# Man-Page Summary (`man-summary`)

> Outputs a man page summary of a command

## Description

`man-summary` reads the man pages for a given command(s) and outputs it's
summary (if one exists).

## Usage

```
man-summary command [ commands ] -> <stdout>
```

## Examples

```
» man-summary man 
man - an interface to the on-line reference manuals
```

## Detail

`man-summary` can take multiple parameters and will return the summary for each
command. If any commands have no summaries, then the exit number will be
incremented. In the example below, two parameters had no associated man page:

```
» man-summary aa ab ac ad ae
aa - Manipulate Apple Archives
ab - Apache HTTP server benchmarking tool
ac - connect time accounting
ad - no man page exists
ae - no man page exists

» exitnum
2
```

## Synonyms

* `man-summary`
* `help.man.summary`


## See Also

* [Murex's Offline Documentation (`murex-docs`)](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [Parse Man-Page For Flags (`man-get-flags`)](../commands/man-get-flags.md):
  Parses man page files for command line flags 
* [Set Command Summary Hint (`summary`)](../commands/summary.md):
  Defines a summary help text for a command
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings

<hr/>

This document was generated from [builtins/core/management/functions_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/functions_doc.yaml).