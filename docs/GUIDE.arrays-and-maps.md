# Language Guide: Arrays And Maps

## Working with structured data

Firstly this shell doesn't have support for arrays as a native data type
however since _murex_ is aware of the structure of various data formats
it is possible to use these formats to maintain complex structured data
natively within _murex_. For example a `days.json` file might look like

    [
            "monday",
            "tuesday",
            "wednesday",
            "thursday",
            "friday",
            "saturday",
            "sunday"
    ]

...which can be queried directly within _murex_ via a variety of builtins.

To iterate through the array and print each element and print the value:

    » open: days.json -> foreach: day { $day }

    monday
    tuesday
    wednesday
    thursday
    friday
    saturday
    sunday

To iterate through the map or array and print each index and its value:

    » open: days.json -> formap: key value { echo: "$key: $value" }

    0: "monday"
    1: "tuesday"
    2: "wednesday"
    3: "thursday"
    4: "friday"
    5: "saturday"
    6: "sunday"

To return a specific element within an array or map you can query it
directly by its key using the `index` builtin:

    » open: days.json -> [ 0 ]

    monday

Or multiple elements in the data set:

    » open: days.json -> [ 0 2 5 6 ]

    ["monday","wednesday","saturday","sunday"]

The `index` builtin returned the values in JSON format because the input
format was JSON. If the input format was a CSV then it would return the
selected columns of that CSV. Or if it's just a new line separated list
of strings then it would return a the rows in the list.

## The `array` builtin

_murex_ has a pretty sophisticated builtin for generating arrays. Think
like bash's `{1..9}` syntax:

    a: [1..9]

You can also specify an alternative number base by using an `x` or `.`
in the end range:

    a: [00..ffx16]
    a: [00..ff.16]

All number bases from 2 (binary) to 36 (0-9 plus a-z) are supported.
Please note that the start and end range are written in the target base
while the base identifier is written in decimal: `[hex..hex.dec]`

Also note that the additional zeros denotes padding (ie the results will
start at `00`, `01`, etc rather than `0`, `1`...

### Character arrays

You can select a range of letters (a to z):

    a: [a..z]
    a: [z..a]
    a: [A..Z]
    a: [Z..A]

...or any characters within that range.

### Special ranges

Unlike bash, _murex_ also supports some special ranges:

    a: [mon..sun]
    a: [monday..sunday]
    a: [jan..dec]
    a: [janurary..december]
    a: [spring..winter]

It is also case aware. If the ranges are uppercase then the return will
be uppercase. If the ranges are title case (capital first letter) then
the return will be in title case:

    » a: [Monday..Sunday]

    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday

Where the special ranges differ from a regular range is they cannot
cannot down. eg `a: [3..1]` would output

    3
    2
    1

however a negative range in special ranges will cycle through to the end
of the range and then loop back from the start:

    » a: [Thursday..Wednesday]

    Thursday
    Friday
    Saturday
    Sunday
    Monday
    Tuesday
    Wednesday

This decision was made because generally with ranges of this type, you
would more often prefer to cycle through values rather than iterate
backwards through the list.

If you did want to reverse then just pipe the output into another UNIX
tool:

    » a: [Monday..Friday] -> tac                         # Linux
    » a: [Monday..Friday] -> tail -r                     # BSD / OS X
    » a: [Monday..Friday] -> perl -e "print reverse <>"  # Multiplaform

    Friday
    Thurday
    Wednesday
    Tuesday
    Monday

(I may build a reverse builtin to standardise this and make _murex_ more
accessible to Windows users)

### Advanced `array` syntax

The syntax for `array` is a comma separated list of parameters with
expansions stored in square brackets. You can have an expansion embedded
inside a parameter or as it's own parameter. Expansions can also have
multiple parameters.

    » a: 01,02,03,05,06,07

    01
    02
    03
    05
    06
    07

    » a: 0[1..3],0[5..7]

    01
    02
    03
    05
    06
    07

    » a: 0[1..3,5..7]

    01
    02
    03
    05
    06
    07

    » a: b[o,i]b

    bob
    bib

You can also have multiple expansion blocks in a single parameter:

    » a: a[1..3]b[5..7]

    a1b5
    a1b6
    a1b7
    a2b5
    a2b6
    a2b7
    a3b5
    a3b6
    a3b7

`array` will cycle through each iteration of the last expansion, moving
itself backwards through the string; behaving like an normal counter:

    » ja: [0..2][0..9] -> format: str ","

    00,01,02,03,04,05,06,07,08,09,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29

(`format` used here for readability)

### Creating JSON arrays with `ja`

As you can see from the previous examples, `a` returns the array as a
list of strings. This is so you can stream excessively long arrays, for
example every IPv4 address: `a: [0..254].[0..254].[0..254].[0..254]`
(this kind of array expansion would hang bash).

However if you needed a JSON string then you can use all the same syntax
as `a` but forgo the streaming capability:

    » ja: [Monday..Sunday]

    [
            "Monday",
            "Tuesday",
            "Wednesday",
            "Thursday",
            "Friday",
            "Saturday",
            "Sunday"
    ]
