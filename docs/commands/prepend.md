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

* [`append`](../commands/append.md):
  Add data to the end of an array
* [`len` ](../commands/len.md):
  Outputs the length of an array
* [a](../commands/a.md):
  
* [cast](../commands/cast.md):
  
* [ja](../commands/ja.md):
  
* [square-bracket-open](../commands/square-bracket-open.md):
  
* [update](../commands/update.md):
  