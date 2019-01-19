# _murex_ Language Guide

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
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`len` ](../commands/len.md):
  Outputs the length of an array
* [`map` ](../commands/map.md):
  Creates a map from two data sources
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [a](../commands/a.md):
  
* [ja](../commands/ja.md):
  