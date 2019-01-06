# _murex_ Language Guide

## Command Reference: `swivel-table`

> Rotates a table by 90 degrees

### Description

`swivel-table` rotates a table by 90 degrees so the _x_ axis becomes the _y_.

### Usage

    <stdin> -> swivel-table -> <stdout>

### Examples

    » ps: aux -> head: -n5
    USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
    root         1  0.0  0.1 233996  8736 ?        Ss   Feb19   0:02 /sbin/init
    root         2  0.0  0.0      0     0 ?        S    Feb19   0:00 [kthreadd]
    root         4  0.0  0.0      0     0 ?        I<   Feb19   0:00 [kworker/0:0H]
    root         6  0.0  0.0      0     0 ?        I<   Feb19   0:00 [mm_percpu_wq]
    
That data swivelled would look like the following:

    » ps: aux -> head: -n5 -> swivel-table
    0       USER    root    root    root    root
    1       PID     1       2       4       6
    2       %CPU    0.0     0.0     0.0     0.0
    3       %MEM    0.1     0.0     0.0     0.0
    4       VSZ     233996  0       0       0
    5       RSS     8736    0       0       0
    6       TTY     ?       ?       ?       ?
    7       STAT    Ss      S       I<      I<
    8       START   Feb19   Feb19   Feb19   Feb19
    9       TIME    0:02    0:00    0:00    0:00
    10      COMMAND /sbin/init      [kthreadd]      [kworker/0:0H]  [mm_percpu_wq]
    
Please note that column one is numbered because by default _murex_ couldn't
guess whether the first line of generic output is a title or data. However if we
format that as a CSV, which by default does have a title row (configurable via
`config`), then you would see titles as column one:

    » ps: aux -> head: -n5 -> format: csv
    "USER","PID","%CPU","%MEM","VSZ","RSS","TTY","STAT","START","TIME","COMMAND"
    "root","1","0.0","0.1","233996","8736","?","Ss","Feb19","0:02","/sbin/init"
    "root","2","0.0","0.0","0","0","?","S","Feb19","0:00","[kthreadd]"
    "root","4","0.0","0.0","0","0","?","I<","Feb19","0:00","[kworker/0:0H]"
    "root","6","0.0","0.0","0","0","?","I<","Feb19","0:00","[mm_percpu_wq]"
    
    » ps: aux -> head: -n5 -> format: csv -> swivel-table
    "USER","root","root","root","root"
    "PID","1","2","4","6"
    "%CPU","0.0","0.0","0.0","0.0"
    "%MEM","0.1","0.0","0.0","0.0"
    "VSZ","233996","0","0","0"
    "RSS","8736","0","0","0"
    "TTY","?","?","?","?"
    "STAT","Ss","S","I<","I<"
    "START","Feb19","Feb19","Feb19","Feb19"
    "TIME","0:02","0:00","0:00","0:00"
    "COMMAND","/sbin/init","[kthreadd]","[kworker/0:0H]","[mm_percpu_wq]"

### See Also

* [`alter`](../docs/commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [`append`](../docs/commands/append.md):
  Add data to the end of an array
* [`prepend` ](../docs/commands/prepend.md):
  Add data to the start of an array
* [`swivel-datatype`](../docs/commands/swivel-datatype.md):
  Converts tabulated data into a map of values for serialised data-types such as JSON and YAML
* [cast](../docs/commands/commands/cast.md):
  
* [format](../docs/commands/commands/format.md):
  
* [square-bracket-open](../docs/commands/commands/square-bracket-open.md):
  