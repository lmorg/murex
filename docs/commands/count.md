# Count (`count`)

> Count items in a map, list or array

## Description

Counts the number of items in a structure, be that a list, map or other object
type.

`count` has several modes ranging from updating values in place, returning new
structures, or just outputting totals.

## Usage

```
<stdin> -> count [ --duplications | --unique | --total ] -> <stdout>
```

## Examples

Count number of items in a map, list or array:

```
» tout json (["a", "b", "c"]) -> count 
3
```

## Flags

* `--bytes`
    Count the total number of bytes read from stdin
* `--duplications`
    Output a JSON map of items and the number of their occurrences in a list or array
* `--runes`
    Count the total number of unicode characters (_runes_) read from stdin. Zero width symbols, wide characters and other non-typical graphemes are all each treated as a single _rune_
* `--sum`
    Read an array, list or map from stdin and output the sum of all the values (ignore non-numeric values)
* `--sum-strict`
    Read an array, list or map from stdin and output the sum of all the values (error on non-numeric values)
* `--total`
    Read an array, list or map from stdin and output the length for that array (default behaviour)
* `--unique`
    Print the number of unique elements in a list or array
* `-b`
    
Alias for `--bytes`
* `-d`
    Alias for `--duplications`
* `-r`
    Alias for `--runes`
* `-s`
    Alias for `--sum`
* `-t`
    Alias for `--total`
* `-u`
    Alias for `--unique`

## Detail

If no flags are set, `count` will default to using `--total`.

### Total: `--total` / `-t`

This will read an array, list or map from stdin and output the length for
that array.

```
» a [25-Dec-2020..05-Jan-2021] -> count 
12
```

> This also replaces the older `len` method.

Please note that this returns the length of the _array_ rather than string.
For example `out "foobar" -> count` would return `1` because an array in the
`str` data type would be new line separated (eg `out "foo\nbar" -> count`
would return `2`). If you need to count characters in a string and are
running POSIX (eg Linux / BSD / OSX) then it is recommended to use `wc`
instead. But be mindful that `wc` will also count new line characters.

```
» out "foobar" -> count
1

» out "foo\nbar" -> count
2

» out "foobar" -> wc: -c
7

» out "foo\nbar" -> wc: -c
8

» printf "foobar" -> wc: -c
6
# (printf does not print a trailing new line)
```

### Duplications: `--duplications` / `-d`

This returns a JSON map of items and the number of their occurrences in a list
or array.

For example in the quote below, only the word "the" is repeated so that entry
will have a value of `2` while ever other entry has a value of `1` because they
only appear once in the quote.

```
» out "the quick brown fox jumped over the lazy dog" -> jsplit \s -> count --duplications
{
    "brown": 1,
    "dog": 1,
    "fox": 1,
    "jumped": 1,
    "lazy": 1,
    "over": 1,
    "quick": 1,
    "the": 2
}
```

### Unique: `--unique` / `-u`

Returns the number of unique elements in a list or array.

For example in the quote below, only the word "the" is repeated, thus the
unique count should be one less than the total count:

```
» out "the quick brown fox jumped over the lazy dog" -> jsplit \s -> count --unique
8
» out "the quick brown fox jumped over the lazy dog" -> jsplit \s -> count --total
9
```

## Synonyms

* `count`
* `len`


## See Also

* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create Map (`map`)](../commands/map.md):
  Creates a map from two data sources
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Prepend To List (`prepend`)](../commands/prepend.md):
  Add data to the start of an array
* [Reverse Array (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [Sort Array (`msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

<hr/>

This document was generated from [builtins/core/datatools/count_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/datatools/count_doc.yaml).