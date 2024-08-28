# list.append: `append`

> Add data to the end of an array

## Description

`append` data to the end of an array.

## Usage

```
<stdin> -> append value -> <stdout>
```

## Examples

```
» a [Monday..Sunday] -> append Funday
Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
Sunday
Funday
```

## Detail

`prepend` and `append` are data type aware:

```
» tout json [1,2,3] -> append 4 5 6 bob
Error in `append` (1,22): cannot convert 'bob' to a floating point number: strconv.ParseFloat: parsing "bob": invalid syntax
```

## Synonyms

* `append`
* `list.append`


## See Also

* [`[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`[[ Element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [`cast`](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [list.new.str: `a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [list.prepend: `prepend`](../commands/prepend.md):
  Add data to the start of an array
* [list.regex: `regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [list.reverse: `mtac`](../commands/mtac.md):
  Reverse the order of an array
* [list.sort: `msort`](../commands/msort.md):
  Sorts an array - data type agnostic
* [list.str: `match`](../commands/match.md):
  Match an exact value in an array
* [struct.count: `count`](../commands/count.md):
  Count items in a map, list or array
* [table.new.heading: `addheading` ](../commands/addheading.md):
  Adds headings to a table

<hr/>

This document was generated from [builtins/core/lists/append_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/append_doc.yaml).