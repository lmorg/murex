# Character arrays

> Making character arrays (a to z)

## Description

You can create arrays from a range of letters (a to z):

```
» a [a..z]
» a [z..a]
» a [A..Z]
» a [Z..A]
```

...or any characters within that range.

Please refer to [a (mkarray)](../commands/a.md) for more detailed usage of mkarray.

## Usage

```
a: [start..end] -> <stdout>
a: [start..end,start..end] -> <stdout>
a: [start..end][start..end] -> <stdout>
```

All usages also work with `ja` and `ta` as well, eg:

```
ja: [start..end] -> <stdout>
ta: data-type [start..end] -> <stdout>
```

You can also inline arrays with the `%[]` syntax, eg:

```
%[start..end]
```

## Examples

```
» a [a..c]
a
b
c
```

```
» a [c..a]
c
b
a
```

## See Also

* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Decimal Ranges](../mkarray/decimal.md):
  Create arrays of decimal integers
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Non-Decimal Ranges](../mkarray/non-decimal.md):
  Create arrays of integers from non-decimal number bases
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [index](../parser/item-index.md):
  Outputs an element from an array, map or table

<hr/>

This document was generated from [builtins/core/mkarray/ranges_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/mkarray/ranges_doc.yaml).