# Add Heading: `addheading`

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

## Synonyms

* `addheading`


## See Also

* [Array Append: `append`](../commands/append.md):
  Add data to the end of an array
* [Array Prepend: `prepend`](../commands/prepend.md):
  Add data to the start of an array
* [Array Reverse (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [Array Sort: `msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Count: `count`](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array: `ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create Streamable Array `a`](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [Define Type: `cast`](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Get Item Property: `[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element: `[[ Element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [Regex Patterns: `regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [String Match: `match`](../commands/match.md):
  Match an exact value in an array

<hr/>

This document was generated from [builtins/core/arraytools/addheading_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/addheading_doc.yaml).