# `addheading` 

> Adds headings to a table

## Description

`addheading` takes a list of parameters and adds them to the start of a table.
Where `prepend` is designed to work with arrays, `addheading` is designed to
prepend to tables.

## Usage

```
<stdin> -> addheading value value value ... -> <stdout>
```

## Examples

```
» tout jsonl '["Bob", 23, true]' -> addheading name age active                                                                                   
["name","age","active"]
["Bob","23","true"]
```

## See Also

* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`count`](../commands/count.md):
  Count items in a map, list or array
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`match`](../commands/match.md):
  Match an exact value in an array
* [`msort`](../commands/msort.md):
  Sorts an array - data type agnostic
* [`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [`prepend`](../commands/prepend.md):
  Add data to the start of an array
* [`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [index](../commands/item-index.md):
  Outputs an element from an array, map or table