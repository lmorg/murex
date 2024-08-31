# Add Heading (`addheading`)

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

* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Match String (`match`)](../commands/match.md):
  Match an exact value in an array
* [Prepend To List (`prepend`)](../commands/prepend.md):
  Add data to the start of an array
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [Reverse Array (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [Sort Array (`msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

<hr/>

This document was generated from [builtins/core/arraytools/addheading_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/addheading_doc.yaml).