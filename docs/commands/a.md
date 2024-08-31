# Stream New List (`a`)

> A sophisticated yet simple way to stream an array or list (mkarray)

## Description

_mkarray_, pronounced "make array" like `mkdir` (etc), is Murex's sophisticated
syntax for generating arrays. Think like bash's `{1..9}` syntax:

```
a [1..9]
```

Except Murex also supports other sets of ranges like dates, days of the week,
and alternative number bases.

This builtin streams arrays as a list of strings (`str`).

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

» a [3..1]
3
2
1

» a [01..03]
01
02
03
```

## Detail

### Advanced Array Syntax

The syntax for `a` is a comma separated list of parameters with expansions
stored in square brackets. You can have an expansion embedded inside a
parameter or as it's own parameter. Expansions can also have multiple
parameters.

```
» a 01,02,03,05,06,07
01
02
03
05
06
07
```

```
» a 0[1..3],0[5..7]
01
02
03
05
06
07
```

```
» a 0[1..3,5..7]
01
02
03
05
06
07
```

```
» a b[o,i]b
bob
bib
```

You can also have multiple expansion blocks in a single parameter:

```
» a a[1..3]b[5..7]
a1b5
a1b6
a1b7
a2b5
a2b6
a2b7
a3b5
a3b6
a3b7
```

`a` will cycle through each iteration of the last expansion, moving itself
backwards through the string; behaving like an normal counter.

### Creating JSON arrays with `ja`

As you can see from the previous examples, `a` returns the array as a
list of strings. This is so you can stream excessively long arrays, for
example every IPv4 address: `a: [0..254].[0..254].[0..254].[0..254]`
(this kind of array expansion would hang bash).

However if you needed a JSON string then you can use all the same syntax
as `a` but forgo the streaming capability:

```
» ja [Monday..Sunday]
[
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday"
]
```

This is particularly useful if you are adding formatting that might break
under `a`'s formatting (which uses the `str` data type).

### Smart arrays

Murex supports a number of different formats that can be used to generate
arrays. For more details on these please refer to the documents for each format

* [Calendar Date Ranges](../mkarray/date.md):
  Create arrays of dates
* [Character arrays](../mkarray/character.md):
  Making character arrays (a to z)
* [Decimal Ranges](../mkarray/decimal.md):
  Create arrays of decimal integers
* [Non-Decimal Ranges](../mkarray/non-decimal.md):
  Create arrays of integers from non-decimal number bases
* [Special Ranges](../mkarray/special.md):
  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)

## Synonyms

* `a`
* `mkarray`


## See Also

* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Reverse Array (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays
* [`str` (string)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [builtins/core/mkarray/array_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/mkarray/array_doc.yaml).