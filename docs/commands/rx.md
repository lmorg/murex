# `rx`

> Regexp pattern matching for file system objects (eg `.*\\.txt`)

## Description

Returns a list of files and directories that match a regexp pattern.

Output is a JSON list.

## Usage

```
rx: pattern -> <stdout>

!rx: pattern -> <stdout>

<stdin> -> rx: pattern -> <stdout>

<stdin> -> !rx: pattern -> <stdout>
```

## Examples

Inline regex file matching:

```
cat: @{ rx: '.*\.txt' }
```

Writing a list of files to disk:

```
rx: '.*\.go' |> filelist.txt
```

Checking if files exist:

```
if { rx: somefiles.* } then {
    # files exist
}
```

Checking if files do not exist:

```
!if { rx: somefiles.* } then {
    # files do not exist
}
```

Return all files apart from text files:

```
!g: '\.txt$'
```

Filtering a file list based on regexp matches file:

```
f: +f -> rx: '.*\.txt'
```

Remove any regexp file matches from a file list:

```
f: +f -> !rx: '.*\.txt'
```

## Detail

### Traversing Directories

Unlike globbing (`g`) which can traverse directories (eg `g: /path/*`), `rx` is
only designed to match file system objects in the current working directory.

`rx` uses Go (lang)'s standard regexp engine.

### Inverse Matches

If you want to exclude any matches based on wildcards, rather than include
them, then you can use the bang prefix. eg

```
» rx: READ*                                                                                                                                                              
[
    "README.md"
]

murex-dev» !rx: .*
Error in `!rx` (1,1): No data returned.
```

### When Used As A Method

`!rx` first looks for files that match its pattern, then it reads the file list
from STDIN. If STDIN contains contents that are not files then `!rx` might not
handle those list items correctly. This shouldn't be an issue with `rx` in its
normal mode because it is only looking for matches however when used as `!rx`
any items that are not files will leak through.

This is its designed feature and not a bug. If you wish to remove anything that
also isn't a file then you should first pipe into either `g: *`, `rx: .*`, or
`f +f` and then pipe that into `!rx`.

The reason for this behavior is to separate this from `!regexp` and `!match`.

## Synonyms

* `rx`
* `!rx`


## See Also

* [`f`](../commands/f.md):
  Lists or filters file system objects (eg files)
* [`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [`match`](../commands/match.md):
  Match an exact value in an array
* [`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings