# Generate Random Sequence (`rand`)

> Random field generator

## Description

`rand` can generate random numbers, strings and other data types.

## Usage

```
rand data-type [ max-value ]
```

## Examples

### Specifying data-type

Random integer (64-bit on 64-bit machines):

```
rand int
```

### Specifying custom range

Random integer between 0 and 9:

```
rand int 9
```

Random floating point number between 0 and 1:

```
rand float
```

Random string, fixed to 5 characters long:

```
rand str 5
```

## Flags

* `*`
    Alias for **str** (string)
* `float`
    Random floating point number between 0 and 1. **max-value** not supported
* `int`
    Random integer, **max-value** specifies highest value for integer
* `num`
    Alias for **int** in this builtin
* `str`
    Random string of printable ASCII characters. **max-value** specifies the fixed length of the string

## Detail

### Security

WARNING: is should be noted that while `rand` can produce random numbers and
strings which might be useful for password generation, neither the RNG nor the
the random string generator (which is ostensibly the same RNG but applied to an
array of bytes within the range of printable ASCII characters) are considered
cryptographically secure.

## See Also

* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)

<hr/>

This document was generated from [builtins/core/random/rand_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/random/rand_doc.yaml).