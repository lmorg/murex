# Non-Decimal Ranges

> Create arrays of integers from non-decimal number bases

## Description

When making arrays you can specify ranges of an alternative number base by
using an `x` or `.` in the end range:

```
a [00..ffx16]
a [00..ff.16]
```

All number bases from 2 (binary) to 36 (0-9 plus a-z) are supported.
Please note that the start and end range are written in the target base
while the base identifier is written in decimal: `[hex..hex.dec]`

Also note that the additional zeros denotes padding (ie the results will
start at `00`, `01`, etc rather than `0`, `1`...)

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
» a [08..10x16]
08
09
0a
0b
0c
0d
0e
0f
10
```

```
» a [10..08x16]
10
f
e
d
c
b
a
9
8
```

## Detail

### Floating Point Numbers

If you do need a range of fixed floating point numbers generated then you can
do so by merging two decimal integer ranges together. For example

```
» a [05..10x8].[0..7]
05.0
05.1
05.2
05.3
05.4
05.5
05.6
05.7
06.0
06.1
06.2
...
07.5
07.6
07.7
10.0
10.1
10.2
10.3
10.4
10.5
10.6
10.7
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
* [Decimal Ranges](../mkarray/decimal.md):
  Create arrays of decimal integers
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [index](../parser/item-index.md):
  Outputs an element from an array, map or table

<hr/>

This document was generated from [builtins/core/mkarray/ranges_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/mkarray/ranges_doc.yaml).