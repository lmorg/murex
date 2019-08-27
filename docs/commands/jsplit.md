# _murex_ Shell Guide

## Command Reference: `jsplit` 

> Splits STDIN into a JSON array based on a regex parameter

### Description

`jsplit` will read from STDIN and split it based on a regex parameter. It outputs a JSON array.

### Usage

    <STDIN> -> jsplit: regex -> <stdout>

### Examples

    » (hello, world) -> jsplit: l+ 
    [
        "he",
        "o, wor",
        "d"
    ]

### See Also

* [`2darray` ](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`len` ](../commands/len.md):
  Outputs the length of an array
* [`map` ](../commands/map.md):
  Creates a map from two data sources
* [`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [element](../commands/element.md):
  