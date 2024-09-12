# Get Item (`[ Index ]`)

> Outputs an element from an array, map or table

## Description

Outputs an element or multiple elements from an array, map or table.

Please note that indexes in Murex are counted from zero.

## Usage

```
<stdin> -> [ element ] -> <stdout>
$variable[ element ] -> <stdout>

<stdin> -> ![ element ] -> <stdout>
```

## Examples

### Multiple array indexes

Return the 2nd (1), 4th (3) and 6th (5) element in an array:

```
» ja [0..9] -> [ 1 3 5 ]
[
    "1",
    "3",
    "5"
]
```

### Multiple object keys

Return the data-type and description of **config shell syntax-highlighting**:

```
» config -> [[ /shell/syntax-highlighting ]] -> [ Data-Type Description ]
[
    "bool",
    "Syntax highlighting of murex code when in the interactive shell"
]
```

### Excluding array indexes

Return all elements _except_ for 1 (2nd), 3 (4th) and 5 (6th):

```
» a [0..9]-> ![ 1 3 5 ]
0
2
4
6
7
8
9
```

### Excluding object keys

Return all elements except for the data-type and description:

```
» config -> [[ /shell/syntax-highlighting ]] -> ![ Data-Type Description ]
{
    "Default": true,
    "Dynamic": false,
    "Global": true,
    "Value": true
}
```

### Filtering tabulated output

#### Selecting columns by name

Return the top 5 processes from `ps`, ordered by memory usage:

```
» ps aux -> [PID %MEM COMMAND] -> sort -nrk2 -> [..5]
915961  14.4  /home/lau/dev/go/bin/gopls
916184  4.4   /opt/visual-studio-code/code
108025  2.9   /usr/lib/firefox/firefox
1036    2.4   /usr/lib/baloo_file
915710  1.9   /opt/visual-studio-code/code
```

#### Selecting columns by index

Return the 1st and 5th column:

```
» ps aux -> [*A *E] -> head -n5                                                                                                                                                                                                       
USER    VSZ
root    168284
root    0
root    0
root    0
```

#### Selecting rows

Return the 1st and 30th row:

```
» ps aux -> [*1 *30]
USER    PID     %CPU    %MEM    VSZ     RSS     TTY     STAT    START   TIME    COMMAND
root    37      0.0     0.0     0       0       ?       I<      Dec18   0:00    [kworker/3:0H-events_highpri]
```

## Detail

### Index counts from zero

Indexes in Murex behave like any other computer array in that all arrays
start from zero (`0`).

### Include vs exclude

As demonstrated in the examples above, `[` specifies elements to include
where as `![` specifies elements to exclude.

### Don't error upon missing elements

By default, **index** generates an error if an element doesn't exist. However
you can disable this behavior in `config`

```
» config -> [ foobar ]
Error in `[` ((builtin) 2,11): Key 'foobar' not found

» config set index silent true

» config -> [ foobar ]
```

## Synonyms

* `[`
* `![`
* `item-index`
* `index`


## See Also

* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Reverse Array (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

<hr/>

This document was generated from [builtins/core/index/index_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/index/index_doc.yaml).