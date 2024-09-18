# `COLUMNS` (int)

> Character width of terminal

## Description

`COLUMNS` returns the cell width of the terminal.

Some characters might be more than or less than 1 (one) cell in width, such
as Chinese logograms and zero-width joiners. Whereas one ASCII character is
the same width as one terminal cell.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`LINES` (int)](../variables/lines.md):
  Character height of terminal
* [`int`](../types/int.md):
  Whole number (primitive)

<hr/>

This document was generated from [gen/variables/COLUMNS_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/COLUMNS_doc.yaml).