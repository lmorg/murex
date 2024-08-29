# table.add.heading

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
* `table.new.heading`


## See Also

* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`[[ Element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [`cast`](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [list.append](../commands/append.md):
  Add data to the end of an array
* [list.new.str (`a`)](../commands/a.md):
  A sophisticated yet simple way to build an array or list (mkarray)
* [list.prepend](../commands/prepend.md):
  Add data to the start of an array
* [list.regex](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [list.reverse (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [list.sort](../commands/msort.md):
  Sorts an array - data type agnostic
* [list.str (`match`)](../commands/match.md):
  Match an exact value in an array
* [struct.count](../commands/count.md):
  Count items in a map, list or array

<hr/>

This document was generated from [builtins/core/arraytools/addheading_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/addheading_doc.yaml).