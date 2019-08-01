# _murex_ Language Guide

## Command Reference: `match`

> Match an exact value in an array

### Description

`match` takes input from STDIN and returns any array items / lines which
contain an exact match of the parameters supplied.

When multiple parameters are supplied they are concatenated into the search
string and white space delimited. eg all three of the below are the same:
    match "a b c"
    match a\sb\sc
    match a b c
    match a    b    c

### Usage

    <stdin> -> match search string -> <stdout>

### Examples

    » ja: [Monday..Friday] -> match Wed
    [
        "Wednesday"
    ]

### Detail

`match` is data-type aware so will work against lists or arrays of whichever
_murex_ data-type is passed to it via STDIN and return the output in the
same data-type.

### See Also

* [`2darray` ](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [`a` (make array)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`len` ](../commands/len.md):
  Outputs the length of an array
* [`map` ](../commands/map.md):
  Creates a map from two data sources
* [`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [`ta`](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [prefix](../commands/prefix.md):
  
* [suffix](../commands/suffix.md):
  