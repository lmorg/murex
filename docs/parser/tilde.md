# `~` Home Sigil

> Home directory path variable

## Description

The tilde token is used as a lazy reference to the users home directory.



## Examples

### Current user

Assuming current username is "bob":

```
» out ~
/home/bob
```

### Alternative user

Assuming "joe" is a valid user on local system:

```
» out ~joe
/home/joe
```

### Unhappy path

If username does not exist, `~` will default to the root path.

Assuming "foobar" isn't a valid local user:

```
» out ~foobar
/
```

## Detail

Tilde can be expanded inside double quotes, brace quotes as well as used naked.
But it cannot be expanded inside single quotes.

```
» out ~
/home/bob

» out '~'
~

» out "~"
/home/bob

» out %(~)
/home/bob
```

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
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/parser/variables_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/variables_doc.yaml).