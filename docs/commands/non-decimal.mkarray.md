# _murex_ Shell Docs

## mkarray: Non-Decimal Ranges

> Create arrays of integers from non-decimal number bases

When making arrays you can specify ranges of an alternative number base by
using an `x` or `.` in the end range:

    a: [00..ffx16]
    a: [00..ff.16]
    
All number bases from 2 (binary) to 36 (0-9 plus a-z) are supported.
Please note that the start and end range are written in the target base
while the base identifier is written in decimal: `[hex..hex.dec]`

Also note that the additional zeros denotes padding (ie the results will
start at `00`, `01`, etc rather than `0`, `1`...)

{{ include "gen/includes/mkarray-range-description.inc copy.md" }}

## See Also

* [mkarray/Character arrays](../mkarray/character.md):
  Making character arrays (a to z)
* [mkarray/Decimal Ranges](../mkarray/decimal.md):
  Create arrays of decimal integers
* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`len` ](../commands/len.md):
  Outputs the length of an array
* [commands/`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type