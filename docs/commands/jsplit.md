# Split String (`jsplit`)

> Splits stdin into a JSON array based on a regex parameter

## Description

`jsplit` will read from stdin and split it based on a regex parameter. It outputs a JSON array.

## Usage

```
<stdin> -> jsplit regex -> <stdout>
```

## Examples

```
» (hello, world) -> jsplit l+ 
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
* `str.split`


## See Also

* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create 2d Array (`2darray`)](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create Map (`map`)](../commands/map.md):
  Creates a map from two data sources
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Prepend To List (`prepend`)](../commands/prepend.md):
  Add data to the start of an array
* [Reverse Array (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [Sort Array (`msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

<hr/>

This document was generated from [builtins/core/lists/jsplit_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/jsplit_doc.yaml).