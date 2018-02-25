# _murex_ Language Guide

## Command reference: append

> Add data to the end of an array

### Description

`append` a data to the end of an array.

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

### See also

* [
* cast
* [prepend](prepend.md): Add data to the start of an array
* update
* a
