# _murex_ Language Guide

## Command Reference: `append`

> Add data to the end of an array

### Description

`append` data to the end of an array.

### Usage

    <stdin> -> append: value -> <stdout>

### Examples

    » a: [Monday..Sunday] -> append: Funday
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
    Funday

### Detail

It's worth noting that `prepend` and `append` are not data type aware. So 
any integers in data type aware structures will be converted into strings:

    » tout: json [1,2,3] -> append: new 
    [
        "1",
        "2",
        "3",
        "new"
    ]

### See Also

* [`len` ](../commands/len.md):
  Outputs the length of an array
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [a](../commands/a.md):
  
* [cast](../commands/cast.md):
  
* [ja](../commands/ja.md):
  
* [square-bracket-open](../commands/square-bracket-open.md):
  
* [update](../commands/update.md):
  