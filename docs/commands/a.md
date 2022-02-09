# _murex_ Shell Docs

## Command Reference: `a` (mkarray)

> A sophisticated yet simple way to build an array or list

## Description

Pronounced "make array", like `mkdir` (etc), _murex_ has a pretty sophisticated
builtin for generating arrays. Think like bash's `{1..9}` syntax:

    a: [1..9]
    
Except _murex_ also supports other sets of ranges like dates, days of the week,
and alternative number bases.

## Usage

    a: [start..end] -> <stdout>
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

### Advanced Array Syntax

The syntax for `a` is a comma separated list of parameters with expansions
stored in square brackets. You can have an expansion embedded inside a
parameter or as it's own parameter. Expansions can also have multiple
parameters.

    » a: 01,02,03,05,06,07
    01
    02
    03
    05
    06
    07
    
    » a: 0[1..3],0[5..7]
    01
    02
    03
    05
    06
    07
    
    » a: 0[1..3,5..7]
    01
    02
    03
    05
    06
    07
    
    » a: b[o,i]b
    bob
    bib
    
You can also have multiple expansion blocks in a single parameter:

    » a: a[1..3]b[5..7]
    a1b5
    a1b6
    a1b7
    a2b5
    a2b6
    a2b7
    a3b5
    a3b6
    a3b7
    
`a` will cycle through each iteration of the last expansion, moving itself
backwards through the string; behaving like an normal counter.

### Creating JSON arrays with `ja`

As you can see from the previous examples, `a` returns the array as a
list of strings. This is so you can stream excessively long arrays, for
example every IPv4 address: `a: [0..254].[0..254].[0..254].[0..254]`
(this kind of array expansion would hang bash).

However if you needed a JSON string then you can use all the same syntax
as `a` but forgo the streaming capability:

    » ja: [Monday..Sunday]
    [
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
        "Sunday"
    ]
    
This is particularly useful if you are adding formatting that might break
under `a`'s formatting (which uses the `str` data type).

### Smart arrays

_murex_ supports a number of different formats that can be used to generate
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

## See Also

* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [types/`str` (string) ](../types/str.md):
  string (primitive)
* [commands/`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [commands/len](../commands/len.md):
  