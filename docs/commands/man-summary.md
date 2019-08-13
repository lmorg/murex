# _murex_ Language Guide

## Command Reference: `man-summary`

> Outputs a man page summary of a command

### Description

`man-summary` reads the man pages for a given command and outputs it's
summary (if one exists).

### Usage

    man-summary command -> <stdout>

### Examples

    » man-summary: man 
    man - an interface to the on-line reference manuals

### See Also

* [`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for _murex_ builtins
* [summary](../commands/summary.md):
  