# List Filesystem Objects (`f`)

> Lists or filters file system objects (eg files)

## Description

`f` lists or filters lists of file system objects, such as files, directories,
etc. `f` provides a quick way to output file system items that confirm to
specific criteria.

You define this criteria by using +flags (eg `+s` for all symlinks) and
optionally then restricting that criteria with -flags (eg `-x` to remove all
executable items). All flags supported as `+` are also supported as a `-`.

By default `f` will return no results. You need to include +flags.

Output is a JSON array as this format preserves whitespace in file names.

## Usage

```
f options -> <stdout>

<stdin> -> f options -> <stdout>
```

## Examples

### Return only directories

```
f +d
```

### Mixing inclusion and exclusions

Return file and directories but exclude symlinks:

```
f +fd -s
```

### As a method

Filter out files in a list (eg created by `g`) using conditions set by `f`:

```
g '*.go' -> f +f

rx '\.(txt|md)' -> f +fw
```

## Flags

* `+`
    include files (pair this with any other flag apart from `-`)
* `-`
    exclude files (pair this with any other flag apart from `+`)
* `?`
    deprecated -- use `i` instead
* `D`
    regular directories
* `E`
    other read permissions
* `F`
    regular files (exc symlinks, devices, sockets, named pipes, etc)
* `Q`
    other write permissions
* `R`
    user read permissions
* `S`
    sockets
* `W`
    user write permissions
* `X`
    user execute permissions
* `Z`
    other execute permissions
* `b`
    block devices
* `c`
    character devices
* `d`
    all directories (inc symlinks)
* `e`
    group read permissions
* `f`
    all files (inc symlinks, devices, sockets, name pipes, etc)
* `i`
    irregular files (nothing else is known about these files)
* `l`
    symlinks
* `p`
    POSIX named pipes (FIFO)
* `q`
    group write permissions
* `r`
    read permissions (user, group, or other)
* `s`
    symlinks
* `w`
    write permissions (user, group, or other)
* `x`
    execute permissions (user, group, or other)
* `z`
    group execute permissions

## Detail

`+` flags are always matched first. Then the `-` flags are used to filter out
any matches from the `+` flags.

## Synonyms

* `f`


## See Also

* [Globbing (`g`)](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [Regex Matches (`rx`)](../commands/rx.md):
  Regexp pattern matching for file system objects (eg `.*\\.txt`)
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)

<hr/>

This document was generated from [builtins/core/io/f_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/f_doc.yaml).