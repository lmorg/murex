# `@Array` Sigil

> Expand values as an array

## Description

The array token is used to tell Murex to expand the string as multiple
parameters (an array) rather than as a single parameter string.

## Examples

**ASCII variable names:**

```
» $example = "foobar"
» out $example
foobar
```

**Unicode variable names:**

Variable names can be non-ASCII however they have to be surrounded by
parenthesis. eg

```
» $(比如) = "举手之劳就可以使办公室更加环保，比如，使用再生纸。"
» out $(比如)
举手之劳就可以使办公室更加环保，比如，使用再生纸。
```

**Infixing inside text:**

Sometimes you need to denote the end of a variable and have text follow on.

```
» $partial_word = "orl"
» out "Hello w$(partial_word)d!"
Hello world!
```

**Variables are tokens:**

Please note the new line (`\n`) character. This is not split using `$`:

```
» $example = "foo\nbar"
```

Output as a string:

```
» out $example
foo
bar
```

Output as an array:

```
» out @example
foo bar
```

The string and array tokens also works for subshells:

```
» out ${ %[Mon..Fri] }
["Mon","Tue","Wed","Thu","Fri"]

» out @{ %[Mon..Fri] }
Mon Tue Wed Thu Fri
```

> `out` will take an array and output each element, space delimited. Exactly
> the same how `echo` would in Bash.

**Variable as a command:**

If a variable is used as a commend then Murex will just print the content of
that variable.

```
» $example = "Hello World!"

» $example
Hello World!
```

## Detail

Since arrays are expanded over multiple parameters, you cannot expand an array
inside quoted strings like you can with a string variable:

```
» out "foo ${ ja [1..5] } bar"
foo ["1","2","3","4","5"] bar

» out "foo @{ ja [1..5] } bar"
foo  1 2 3 4 5  bar

» %(${ ja [1..5] })
["1","2","3","4","5"]   

» %(@{ ja: [1..5] })
@{ ja [1..5] } 
```

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
  Write a string to the STDOUT without new line (deprecated)
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`string` (stringing)](../types/str.md):
  string (primitive)
* [`~` Home Sigil](../parser/tilde.md):
  Home directory path variable

<hr/>

This document was generated from [gen/parser/variables_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/variables_doc.yaml).