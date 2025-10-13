# Array Append: `append`

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

* [Add Heading: `addheading`](../commands/addheading.md):
  Adds headings to a table
* [Array Prepend: `prepend`](../commands/prepend.md):
  Add data to the start of an array
* [Array Reverse: `mtac`](../commands/mtac.md):
  Reverse the order of an array
* [Array Sort: `msort`](../commands/msort.md):
  Sorts an array - data type agnostic
* [Count: `count`](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array: `ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create Streamable Array: `a`](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [Define Type: `cast`](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Filter By Range: `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Item Property: `[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Item Property: `[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element: `[[ Element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [Regex Patterns: `regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [String Match: `match`](../commands/match.md):
  Match an exact value in an array

<hr/>

This document was generated from [builtins/core/lists/append_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/append_doc.yaml).