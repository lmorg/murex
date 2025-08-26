# `~` Home Sigil

> Home directory path variable

## Description

The tilde token is used as a convenience shortcut to users home directory.

By itself, `~` will point to the current users home directory.

If a username follows, eg `~joe.bloggs`, then the home directory for that user
is returned irrespective of who is presently logged in. Characters supported by
tilde usernames are alpha upper and lower case, numeric, underscore, full stop
(period), and hyphens.



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

### Infixing

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

### Error Handling

If a username is supplied that that user doesn't exist, the tilde will raise an
error. For example:

```
» ~joe.bloggs
Error in `expr` (0,1): cannot expand variable `~joe.bloggs`: user: unknown user joe.bloggs
```

## See Also

* [Create JSON Array: `ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Define Variable: `set`](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Output String, stdout: `out`](../commands/out.md):
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
* [`HOME` (path)](../variables/home.md):
  Return the home directory for the current session user
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/parser/variables_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/variables_doc.yaml).