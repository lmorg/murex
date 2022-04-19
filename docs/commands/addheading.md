# _murex_ Shell Docs

## Command Reference: `addheading` 

> Adds headings to a table

## Description

`addheading` takes a list of parameters and adds them to the start of a table.
Where `prepend` is designed to work with arrays, `addheading` is designed to
prepend to tables.

## Usage

    <stdin> -> addheading: value value value ... -> <stdout>

## Examples

    » tout: jsonl '["Bob", 23, true]' -> addheading name age active                                                                                   
    ["name","age","active"]
    ["Bob","23","true"]

## See Also

* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`append`](../commands/append.md):
  Add data to the end of an array
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`count`](../commands/count.md):
  Count items in a map, list or array
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`match`](../commands/match.md):
  Match an exact value in an array
* [commands/`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [commands/`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [commands/`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [commands/`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings