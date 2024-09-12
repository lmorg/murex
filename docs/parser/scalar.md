# `$Scalar` Sigil (eg variables)

> Expand values as a scalar

## Description

The scalar token is used to tell Murex to expand variables and sub-shells as a
string (ie one single parameter) irrespective of the data that is stored in the
string.

One common use case where Murex's approach is better is with file names.
Traditional shells would treat spaces as a new file. Whereas Murex treats
spaces as any other printable character character.

## Variable Syntax

There are two basic syntaxes. Bare an enclosed.

### Bare Syntax

Bare syntax looks like the following:

```
$scalar
```

The variable token must be followed with one of the following characters: 
alpha (`a` to `z`, upper and lower case), numeric (`0` to `1`), underscore
(`_`) and/or a full stop (`.`).

### Enclosed Syntax

Enclosed syntax looks like the following:

```
$(scalar)
```

Enclosed syntax supports any unicode characters however the variable name
needs to be surrounded by parenthesis. See examples below.



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

* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
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
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`~` Home Sigil](../parser/tilde.md):
  Home directory path variable

<hr/>

This document was generated from [gen/parser/variables_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/variables_doc.yaml).