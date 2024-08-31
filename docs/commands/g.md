# Globbing (`g`)

> Glob pattern matching for file system objects (eg `*.txt`)

## Description

Returns a list of files and directories that match a glob pattern.

Output is a JSON list.

## Usage

```
g pattern -> <stdout>

!g pattern -> <stdout>

<stdin> -> g pattern -> <stdout>

<stdin> -> !g pattern -> <stdout>
```

## Examples

### Inline globbing

```
cat @{ g *.txt }
```

### Writing a JSON array of files to disk

```
g *.txt |> filelist.json
```

### Checking if a file exists

```
if { g somefile.txt } then {
    # file exists
}
```

### Checking if a file does not exist

```
!if { g somefile.txt } then {
    # file does not exist
}
```

### Return all files apart from

```
!g *.txt
```

### Filtering a file list based on glob matches

```
f +f -> g *.md
```

### Remove any glob matches from a file list

```
f +f -> !g *.md
```

### Files in directories that begin with a vowel

```
g [aeiou]*/*
```

## Detail

### Pattern Reference

Murex globbing is based on [Go's stdlib Match library](https://pkg.go.dev/path/filepath#Match)

#### pattern

    { term }

#### term

    '*'         matches any sequence of non-Separator characters
    '?'         matches any single non-Separator character
    '[' [ '^' ] { character-range } ']'
                character class (must be non-empty)
    c           matches character c (c != '*', '?', '\\', '[')
    '\\' c      matches character c

#### character-range

    c           matches character c (c != '\\', '-', ']')
    '\\' c      matches character c
    lo '-' hi   matches character c for lo <= c <= hi

### Inverse Matches

If you want to exclude any matches based on wildcards, rather than include
them, then you can use the bang prefix. eg

```
» g READ*
[
    "README.md"
]

» !g *
Error in `!g` (1,1): No data returned.
```

### When Used As A Method

`!g` first looks for files that match its pattern, then it reads the file list
from stdin. If stdin contains contents that are not files then `!g` might not
handle those list items correctly. This shouldn't be an issue with `frx` in its
normal mode because it is only looking for matches however when used as `!g`
any items that are not files will leak through.

This is its designed feature and not a bug. If you wish to remove anything that
also isn't a file then you should first pipe into either `g *`, `rx .*`, or
`f +f` and then pipe that into `!g`.

The reason for this behavior is to separate this from `!regexp` and `!match`.

## Synonyms

* `g`
* `!g`


## See Also

* [List Filesystem Objects (`f`)](../commands/f.md):
  Lists or filters file system objects (eg files)
* [Match String (`match`)](../commands/match.md):
  Match an exact value in an array
* [Regex Matches (`rx`)](../commands/rx.md):
  Regexp pattern matching for file system objects (eg `.*\\.txt`)
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings

<hr/>

This document was generated from [builtins/core/io/g_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/g_doc.yaml).