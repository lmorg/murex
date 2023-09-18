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

* [Decimal Ranges](../mkarray/decimal.md):
  Create arrays of decimal integers
* [Non-Decimal Ranges](../mkarray/non-decimal.md):
  Create arrays of integers from non-decimal number bases
* [a](../mkarray/a.md):
  
* [count](../mkarray/count.md):
  
* [element](../mkarray/element.md):
  
* [index](../mkarray/index.md):
  
* [ja](../mkarray/ja.md):
  
* [range](../mkarray/range.md):
  
* [ta](../mkarray/ta.md):
  

<hr/>

This document was generated from [builtins/core/mkarray/ranges_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/mkarray/ranges_doc.yaml).