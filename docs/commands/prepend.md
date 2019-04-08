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

### Detail

It's worth noting that `prepend` and `append` are not data type aware. So 
any integers in data type aware structures will be converted into strings:

    » tout: json [1,2,3] -> prepend: new 
    [
        "new",
        "1",
        "2",
        "3"
    ]

### See Also

* [`a`](../commands/a.md):
  A sophisticated yet simply way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`len` ](../commands/len.md):
  Outputs the length of an array
* [`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [mtac](../commands/mtac.md):
  
* [square-bracket-open](../commands/square-bracket-open.md):
  
* [update](../commands/update.md):
  