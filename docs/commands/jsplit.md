# `jsplit`  - Command Reference

> Splits STDIN into a JSON array based on a regex parameter

## Description

`jsplit` will read from STDIN and split it based on a regex parameter. It outputs a JSON array.

## Usage

```
<STDIN> -> jsplit: regex -> <stdout>
```

## Examples

```
» (hello, world) -> jsplit: l+ 
[
    "he",
    "o, wor",
    "d"
]
```

## Detail

`jsplit` will trim trailing carriage returns and line feeds from each element
as well as any trailing empty elements (zero length strings) in the JSON array.
However any empty elements will be retained and any other whitespace characters
- or carriage returns and/or line feeds in the middle of an element - will be
retained.

This is so that the formatting of (multiline) text is retained as much as
possible to ensure the `jsplit` is accurate while at the same time any commonly
unwanted "noise" is stripped from the output.

## Synonyms

* `jsplit`
* `list.split`


## See Also

* [`2darray` ](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`count`](../commands/count.md):
  Count items in a map, list or array
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`map` ](../commands/map.md):
  Creates a map from two data sources
* [`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array