# _murex_ Shell Docs

## Command Reference: `@g` (autoglob) 

> Command prefix to expand globbing

## Description

By default _murex_ does not expand globbing (`*` and `?` wildcards) instead
encouraging the use of `g` (and similar) inside a subshell. While the aim of
this is to promote correctness, it can be a little annoying while working in
the interactive shell. For this reason you can prefix any command with `@g` to
enable Bash-like globbing.

## Usage

    @g command ...

## Examples

    @g echo *

## Detail

As of _murex_ `2.9` and above it is possible to enable automatic globbing in
the interactive shell without having to prefix the command with `@g` by
enabling the following `config` option:

    config: set shell auto-glob true
    
To make this setting persistent you will need to add it to your `.murex_profile

## See Also

* [user-guide/_murex_ Profile Files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [commands/`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [commands/`f`](../commands/f.md):
  Lists objects (eg files) in the current working directory
* [commands/`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [commands/`rx`](../commands/rx.md):
  Regexp pattern matching for file system objects (eg '.*\.txt')