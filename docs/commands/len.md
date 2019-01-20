# _murex_ Language Guide

## Command Reference: `len` 

> Outputs the length of an array

### Description

This will read an array from STDIN and outputs the length for that array

### Usage

    <STDIN> -> len -> <stdout>

### Examples

    » tout: json (["a", "b", "c"]) -> len 
    3

### See Also

* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`a`](../commands/a.md):
  A sophisticated yet simply way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`map` ](../commands/map.md):
  Creates a map from two data sources
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [mtac](../commands/mtac.md):
  
* [range](../commands/range.md):
  