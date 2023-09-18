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

* [interactive-shell](../variables/interactive-shell.md):
  
* [reserved-vars](../variables/reserved-vars.md):
  
* [set](../variables/set.md):
  
* [str](../variables/str.md):
  
* [string](../variables/string.md):
  

<hr/>

This document was generated from [gen/variables/COLUMNS_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/COLUMNS_doc.yaml).