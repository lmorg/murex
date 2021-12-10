# _murex_ Shell Docs

## mkarray: Character arrays

> Making character arrays (a to z)

## Description

You can create arrays from a range of letters (a to z):

    » a: [a..z]
    » a: [z..a]
    » a: [A..Z]
    » a: [Z..A]
    
...or any characters within that range.

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

    » a: [a..c]
    a
    b
    c
    
    » a: [c..a]
    c
    b
    a

## See Also

* [mkarray/Decimal Ranges](../mkarray/decimal.md):
  Create arrays of decimal integers
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