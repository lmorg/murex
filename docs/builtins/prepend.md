# _murex_ Language Guide

## Command reference: prepend

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

### See also

* [
* a
* [append](append.md): Add data to the end of an array
* cast
* update
