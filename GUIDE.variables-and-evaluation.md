# Language Guide: Variables And Evaluation

## Defining a variable

It's easier to think of variables as three major classes:

1. Numeric data types:
 * `number` (default type for numeric data. Always stored as a `float`)
 * `integer`
 * `float`

2. Textual data types:
 * `string`
 * `json` (including objects)
 * `binary` (bytes that are not intended to be parsed by text utilities)
 * `code blocks` (code that will be parsed by a subshell)

3. System types:
 * `generic` (an support any data type)
 * `null` (no data supported or returned)
 * `boolean` (more of a "data state" than a type as both numeric and
   textual types can be described as `boolean`)
 * `die` (this type kills the shells. Generally only used to stop
   processing in the event of a serious and unhandled error)

## Declaring a variable

System variables cannot be declared (albeit `boolean` and `generic` types
can be handled. More on that later.

Numbers are declared with `let` and strings with `set`.

### Declaring a number

`let` supports mathematical operations, for example

    let percent=(7/10)*100

will declare a variable called `percent` and assign it the value of the
formula `(7/10)*100`.

### Declaring a string

Strings are declared via `set` with the input being added without any
evaluation. For example

    set formula=(7/10)*100

will create a variable called `formula` with the value being the string
`(7/10)*100` (ie the formula stored as text rather than the calculation
of the formula).

## Variable usage in common functions and methods

Like with many other shells and some scripting languages (eg AWK, Bash,
Perl and PHP), this shell supports inlining variables via the dollar
prefix. eg

    set a=world
    echo "Hello $a" # outputs: Hello world

Currently this is only supported in function parameters (ie you cannot
use variables as function names).

## Variable usage in `let` and `eval`

TODO: write me