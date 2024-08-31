# C-style functions

> Inlined commands for expressions and statements

## Description

The traditional way to spawn a sub-shell would be `${}` (or `$()` in Bourne
Shell and its many derivatives). This works great for inlining entire command
lines but it isn't so convenient if you want to call one command and
particularly within an expression.

This is where C-style functions can be more ergonomic. They follow the common
structure of `command(parameters...)`. For example:

```
» out("hello world")
```

The syntax is not exactly like C and its derivatives however:

* parameters are white space delimited, like with command line statements
* strings do not need to be quoted, like with command line statements

And unlike statements:

* you cannot redirect stdout nor stderr
* stdout is never a TTY. Even when it is ran directly in the terminal, it is
  still treated as a sub-shell

Ostensibly, C-style functions are just syntactic sugar for `${}`. As such,
they're not intended to be used liberally but rather just in instances it
improves readability.



## Examples

### Assignment In Expressions

As a C-style function (CSF) vs a sub-shell:

```
# CSF
» $doc = open(README.md)

# Sub-shell
» $doc = ${open README.md}
```

### Numeric Value In Expressions

As a C-style function (CSF) vs a sub-shell:

```
# CSF
» datetime(--in {now} --out {unix}) / 60
28687556.3

# Sub-shell
» ${datetime --in {now} --out {unix}} / 60
28687556.3
```

### Statement Inlining

As a C-style function (CSF) vs a sub-shell:

```
# CSF
» echo It is datetime(--in {now} --out {py}%H) o\' clock
It is 23 o' clock

# Sub-shell
» echo It is ${datetime --in {now} --out {py}%H} o\' clock
It is 23 o' clock
```

Notice in the example above, `echo`'s parameters are not quoted. This is
because C-style functions do not support infixing.

## Detail

### Valid Function Names

Please note that currently the only functions supported are ones who's names
are comprised entirely of alpha, numeric, underscore and/or exclamation marks.

### String Infixing

C-style functions do not support being infixed like sub-shells can be:

```
# CSF
» echo "It is datetime(--in {now} --out {py}%H) o\' clock"
It is datetime(--in {now} --out {py}%H) o' clock

# Sub-shell
» echo "It is ${datetime --in {now} --out {py}%H} o\' clock"
It is 23 o' clock
```

## See Also

* [Date And Time Conversion (`datetime`)](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Language Tour](../Murex/tour.md):
  Getting started with Murex
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Output String (`echo`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [sub-shell](../parser/sub-shell.md):
  

<hr/>

This document was generated from [gen/parser/c-style-functions_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/c-style-functions_doc.yaml).