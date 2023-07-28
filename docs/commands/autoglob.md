# `@g` (autoglob) 

> Command prefix to expand globbing (deprecated)

## Description

**This feature is now deprecated and only applies to murex version 2:**

By default Murex does not expand globbing (`*` and `?` wildcards) instead
encouraging the use of `g` (and similar) inside a subshell. While the aim of
this is to promote correctness, it can be a little annoying while working in
the interactive shell. For this reason you can prefix any command with `@g` to
enable Bash-like globbing.

## Usage

```
@g command ...
```

## Examples

```
@g echo *
```

## Detail

As of Murex `2.9` and above it is possible to enable automatic globbing in
the interactive shell without having to prefix the command with `@g` by
enabling the following `config` option:

```
config: set shell auto-glob true
```

It is enabled by default on from version 3.x onwards (and renamed to
`expand-glob`)

## See Also

* [Murex Profile Files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`f`](../commands/f.md):
  Lists or filters file system objects (eg files)
* [`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [`rx`](../commands/rx.md):
  Regexp pattern matching for file system objects (eg `.*\\.txt`)