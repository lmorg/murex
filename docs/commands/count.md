# `count`

> Count items in a map, list or array

## Description



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

* `--duplications`
    Output a JSON map of items and the number of their occurrences in a list or array
* `--total`
    Read an array, list or map from STDIN and output the length for that array (default behaviour)
* `--unique`
    Print the number of unique elements in a list or array
* `-d`
    Alias for `--duplications`
* `-t`
    Alias for `--total`
* `-u`
    Alias for `--unique`

## Detail

### Modes

If no flags are set, `count` will default to using `--total`.

#### Total: `--total` / `-t`

This will read an array, list or map from STDIN and output the length for
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

#### Duplications: `--duplications` / `-d`

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

#### Unique: `--unique` / `-u`

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

* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (range)](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`map`](../commands/map.md):
  Creates a map from two data sources
* [`msort`](../commands/msort.md):
  Sorts an array - data type agnostic
* [`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [`prepend`](../commands/prepend.md):
  Add data to the start of an array
* [`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [index](../commands/item-index.md):
  Outputs an element from an array, map or table