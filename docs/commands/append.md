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

### See Also

* [`prepend` ](../docs/commands/prepend.md):
  Add data to the start of an array
* [a](../docs/commands/commands/a.md):
  
* [cast](../docs/commands/commands/cast.md):
  
* [square-bracket-open](../docs/commands/commands/square-bracket-open.md):
  
* [update](../docs/commands/commands/update.md):
  