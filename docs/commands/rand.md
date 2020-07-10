# _murex_ Shell Docs

## Command Reference: `rand`

> Random field generator

## Description

`rand` can generate random numbers, strings and other data types.

## Usage

    rand data-type [ max-value ]

## Examples

Random integer: 64-bit on 64-bit machines

    rand int
    
Random number between 

## Detail

### Security

WARNING: is should be noted that while `rand` can produce random numbers and
strings which might be useful for password generation, neither the RNG nor the
the random string generator (which is ostensibly the same RNG but applied to an
array of bytes within the range of printable ASCII characters) are considered
cryptographically secure.

## See Also

* [commands/`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value