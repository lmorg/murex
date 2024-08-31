# Decimal Ranges

> Create arrays of decimal integers

## Description

This document describes how to create arrays of decimals using mkarray (`a` et
al).

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
» a [1..3]
1
2
3
```

```
» a [3..1]
3
2
1
```

```
» a [01..03]
01
02
03
```

## Detail

### Floating Point Numbers

If you do need a range of fixed floating point numbers generated then you can
do so by merging two decimal integer ranges together. For example

```
» a [0..5].[0..9]
0.0
0.1
0.2
0.3
0.4
0.5
0.6
0.7
0.8
0.9
1.0
1.1
1.2
1.3
...
4.8
4.9
5.0
5.1
5.2
5.3
5.4
5.5
5.6
5.7
5.8
5.9
```

### Everything Is A String

Please note that all arrays are created as strings. Even when using typed
arrays such as JSON (`ja`).

```
» ja [0..5]
[
    "0",
    "1",
    "2",
    "3",
    "4",
    "5"
] 
```

## See Also

* [Character arrays](../mkarray/character.md):
  Making character arrays (a to z)
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
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