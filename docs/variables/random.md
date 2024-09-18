# `RANDOM` (int)

> Return a random 32-bit integer (historical)

## Description

`$RANDOM` returns a number between 0 and 32767 (inclusive).

`$RANDOM` was included for POSIX support however the idiomatic way to generate
random tokens in Murex is via the [`rand` builtin](/docs/commands/rand.md).

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## See Also

* [Generate Random Sequence (`rand`)](../commands/rand.md):
  Random field generator
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`int`](../types/int.md):
  Whole number (primitive)

<hr/>

This document was generated from [gen/variables/RANDOM_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/RANDOM_doc.yaml).