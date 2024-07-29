# `$Variable` Sigil

> Expand values as a scalar

## Description

The scalar token is used to tell Murex to expand variables and sub-shells as a
string (ie one single parameter) irrespective of the data that is stored in the
string. One handy common use case is file names where traditional POSIX shells
would treat spaces as a new file, whereas Murex treats spaces as a printable
character unless explicitly told to do otherwise.

The string token must be followed with one of the following characters: 
alpha, numeric, underscore (`_`) or a full stop / period (`.`).



## Examples

### ASCII variable names

```
» $example = "foobar"
» out $example
foobar
```

### Unicode variable names

Variable names can be non-ASCII however they have to be surrounded by
parenthesis. eg

```
» $(比如) = "举手之劳就可以使办公室更加环保，比如，使用再生纸。"
» out $(比如)
举手之劳就可以使办公室更加环保，比如，使用再生纸。
```

### Infixing inside text

Sometimes you need to denote the end of a variable and have text follow on:

```
» $partial_word = "orl"
» out "Hello w$(partial_word)d!"
Hello world!
```

### Variables are tokens

Please note the new line (`\n`) character. This is not split using `$`:

```
» $example = "foo\nbar"
```

Output as a scalar (`$`):

```
» out $example
foo
bar
```

Output as an array (`@`):

```
» out @example
foo bar
```

### Scalar and Array Sub-shells

Scalar:

```
» out ${ %[Mon..Fri] }
["Mon","Tue","Wed","Thu","Fri"]
```

Array:

```
» out @{ %[Mon..Fri] }
Mon Tue Wed Thu Fri
```

> `out` will take an array and output each element, space delimited. Exactly
> the same how `echo` would in Bash.

### Variable as a Command

If a variable is used as a commend then Murex will just print the content of
that variable.

```
» $example = "Hello World!"

» $example
Hello World!
```

## Detail

### Infixing

Strings and sub-shells can be expanded inside double quotes, brace quotes as
well as used as barewords. But they cannot be expanded inside single quotes.

```
» set example="World!"

» out Hello $example
Hello World!

» out 'Hello $example'
Hello $example

» out "Hello $example"
Hello World!

» out %(Hello $example)
Hello World!
```

However you cannot expand arrays (`@`) inside any form of quotation since
it wouldn't be clear how that value should be expanded relative to the
other values inside the quote. This is why array and object builders (`%[]`
and `%{}` respectively) support array variables but string builders (`%()`)
do not.

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`"Double Quote"`](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`out`](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`~` Home Sigil](../parser/tilde.md):
  Home directory path variable

<hr/>

This document was generated from [gen/parser/variables_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/variables_doc.yaml).