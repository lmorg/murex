# _murex_ Shell Docs

## mkarray: Decimal Ranges

> Create arrays of decimal integers

## Description

This document describes how to create arrays of decimals using mkarray (`a` et
al).

Please refer to [a (mkarray)](../commands/a.md) for more detailed usage of mkarray.

## Usage

    a: [start..end] -> <stdout>
    a: [start..end.base] -> <stdout>
    a: [start..end,start..end] -> <stdout>
    a: [start..end][start..end] -> <stdout>
    
All usages also work with `ja` and `ta` as well:

    ja: [start..end] -> <stdout>
    ta: data-type [start..end] -> <stdout>

## Examples

    » a: [1..3]
    1
    2
    3
    
    » a: [3..1]
    3
    2
    1
    
    » a: [01..03]
    01
    02
    03

## Detail

### Floating Point Numbers

If you do need a range of fixed floating point numbers generated then you can
do so by merging two decimal integer ranges together. For example

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
    
### Everything Is A String

Please note that all arrays are created as strings. Even when using typed
arrays such as JSON (`ja`).

    » ja [0..5]
    [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5"
    ] 

## See Also

* [mkarray/Character arrays](../mkarray/character.md):
  Making character arrays (a to z)
* [mkarray/Non-Decimal Ranges](../mkarray/non-decimal.md):
  Create arrays of integers from non-decimal number bases
* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`len` ](../commands/len.md):
  Outputs the length of an array
* [commands/`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type