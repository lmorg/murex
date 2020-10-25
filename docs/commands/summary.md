# _murex_ Shell Docs

## Command Reference: `summary` 

> Defines a summary help text for a command

## Description

`summary` define help text for a command. This is effectively like a tooltip
message that appears, by default, in blue in the interactive shell.

Normally this text is populated from the `man` pages or `murex-docs`, however
if neither exist or if you wish to override their text, then you can use
`summary` to define that text.

## Usage

Define a commands summary

    summary command description
    
Undefine a summary

    !summary command

## Examples

Define a commands summary

    » summary: foobar "Hello, world!"
    » runtime: --summaries -> [ foobar ]
    Hello, world! 
    
Undefine a summary

    » !summary: foobar

## Synonyms

* `summary`
* `!summary`


## See Also

* [commands/`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [commands/`exec`](../commands/exec.md):
  Runs an executable
* [commands/`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [commands/`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for _murex_ builtins
* [commands/`murex-update-exe-list`](../commands/murex-update-exe-list.md):
  Forces _murex_ to rescan $PATH looking for exectables
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/bexists](../commands/bexists.md):
  
* [commands/builtins](../commands/builtins.md):
  