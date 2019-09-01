# _murex_ Shell Guide

## Command Reference: `swivel-datatype`

> Converts tabulated data into a map of values for serialised data-types such as JSON and YAML

### Description

`swivel-datatype` rotates a table by 90 degrees then exports the output as a
series of maps to be marshalled by a serialised datatype such as JSON or YAML.

### Usage

    <stdin> -> swivel-datatype: data-type -> <stdout>

### Examples

Lets take the first 5 entries from `ps`:

    » ps: aux -> head: -n5 -> format: csv
    "USER","PID","%CPU","%MEM","VSZ","RSS","TTY","STAT","START","TIME","COMMAND"
    "root","1","0.0","0.1","233996","8736","?","Ss","Feb19","0:02","/sbin/init"
    "root","2","0.0","0.0","0","0","?","S","Feb19","0:00","[kthreadd]"
    "root","4","0.0","0.0","0","0","?","I<","Feb19","0:00","[kworker/0:0H]"
    "root","6","0.0","0.0","0","0","?","I<","Feb19","0:00","[mm_percpu_wq]"
    
That data swivelled would look like the following:

    » ps: aux -> head: -n5 -> format: csv -> swivel-datatype: yaml
    '%CPU':
    - "0.0"
    - "0.0"
    - "0.0"
    - "0.0"
    '%MEM':
    - "0.1"
    - "0.0"
    - "0.0"
    - "0.0"
    COMMAND:
    - /sbin/init
    - '[kthreadd]'
    - '[kworker/0:0H]'
    - '[mm_percpu_wq]'
    PID:
    - "1"
    - "2"
    - "4"
    - "6"
    RSS:
    - "8736"
    - "0"
    - "0"
    - "0"
    START:
    - Feb19
    - Feb19
    - Feb19
    - Feb19
    STAT:
    - Ss
    - S
    - I<
    - I<
    TIME:
    - "0:02"
    - "0:00"
    - "0:00"
    - "0:00"
    TTY:
    - '?'
    - '?'
    - '?'
    - '?'
    USER:
    - root
    - root
    - root
    - root
    VSZ:
    - "233996"
    - "0"
    - "0"
    - "0"
    
Please note that for input data-types whose table doesn't define titles (such as
the `generic` datatype), the map keys are defaulted to column numbers:

    » ps: aux -> head: -n5 -> swivel-datatype: yaml
    "0":
    - USER
    - root
    - root
    - root
    - root
    "1":
    - PID
    - "1"
    - "2"
    - "4"
    - "6"
    "2":
    - '%CPU'
    - "0.0"
    - "0.0"
    - "0.0"
    - "0.0"
    "3":
    - '%MEM'
    - "0.1"
    - "0.0"
    - "0.0"
    - "0.0"
    ...

### Detail

You can check what output data-types are available via the `runtime` command:

    runtime --marshallers
    
Marshallers are enabled at compile time from the `builtins/data-types` directory.

### See Also

* commands/[`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* commands/[`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* commands/[`alter`](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* commands/[`append`](../commands/append.md):
  Add data to the end of an array
* commands/[`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* commands/[`format`](../commands/format.md):
  Reformat one data-type into another data-type
* commands/[`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* commands/[`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* commands/[`swivel-table`](../commands/swivel-table.md):
  Rotates a table by 90 degrees