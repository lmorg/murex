# _murex_ Language Guide

## Command Reference: `prepend` 

> Add data to the start of an array

### Description

`prepend` a data to the start of an array.

### Usage

    <stdin> -> prepend: value -> <stdout>

### Examples

    » a: [January..December] -> prepend: 'New Year'
    New Year
    January
    February
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

### See Also

* [`append`](../commands/append.md):
  Add data to the end of an array
* [a](../commands/a.md):
  
* [cast](../commands/cast.md):
  
* [square-bracket-open](../commands/square-bracket-open.md):
  
* [update](../commands/update.md):
  