# `COLUMNS` (int)

> Character width of terminal

## Description

`COLUMNS` returns the cell width of the terminal.

Some characters might be more than or less than 1 (one) cell in width, such
as Chinese logograms and zero-width joiners. Whereas one ASCII character is
the same width as one terminal cell.

This is a reserved variable so it cannot be changed.



## Synonyms

* `columns`
* `COLUMNS`


## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`str` (string)](../types/str.md):
  string (primitive)
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/COLUMNS_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/COLUMNS_doc.yaml).