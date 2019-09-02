# _murex_ Shell Guide

## Command Reference: `@[` (range) 

> Outputs a ranged subset of data from STDIN

### Description

This will read from STDIN and output a subset of data in a defined range.

The range can be defined as a number of different range types - such as the
content of the array or it's index / row number. You can also omit either
the start or the end of the search criteria to cover all items before or
after the remaining search criteria.

### Usage

    <stdin> -> @[start..end]flags -> <stdout>

### Examples

Range over all months after March:

    » a: [January..December] -> @[March..]se
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December
    
Range from the 6th to the 10th month (indexes start from zero, `0`):

    » a: [January..December] -> @[5..9]
    June
    July
    August
    September
    October

### Flags

* `e`
    exclude the start and end search criteria from the range
* `n`
    array index
* `r`
    regexp match
* `s`
    exact string match

### Synonyms

* `@[`


### See Also

* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`alter`](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [commands/`append`](../commands/append.md):
  Add data to the end of an array
* [commands/`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [commands/`len` ](../commands/len.md):
  Outputs the length of an array
* [commands/`prepend` ](../commands/prepend.md):
  Add data to the start of an array