# Change Text Case (`list.case`)

> Changes the character case of a string or all elements in an array

## Description

`list.case` will read an array from either stdin or it's command line
parameters, and change that list to be upper or lower case.

## Usage

```
<stdin> -> list.case operation                           -> <stdout>
           list.case operation item1 [ item2 item3 ... ] -> <stdout>
```

## Examples

### As a method

```
» %[Monday..Friday] -> list.case upper
[
    "MONDAY",
    "TUESDAY",
    "WEDNESDAY",
    "THURSDAY",
    "FRIDAY"
]
```

### As a function

```
» list.case lower @{ %[Monday..Friday] }
[
    "monday",
    "tuesday",
    "wednesday",
    "thursday",
    "friday"
]
```

## Synonyms

* `list.case`


## See Also

* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array

<hr/>

This document was generated from [builtins/core/lists/case_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/case_doc.yaml).