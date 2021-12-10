# _murex_ Shell Docs

## mkarray: Calendar Date Ranges

> Create arrays of dates

Unlike bash, _murex_ also supports date ranges:

```  
» a: [25-dec-2020..05-jan-2021]
» a: [..25-dec-2020]
» a: [25-dec-2020..]
```

{{ include "gen/includes/mkarray-range-description.inc copy.md" }}

## See Also

* [mkarray/Special Ranges](../mkarray/special.md):
  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)
* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`datetime` ](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`len` ](../commands/len.md):
  Outputs the length of an array
* [commands/`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [commands/`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type