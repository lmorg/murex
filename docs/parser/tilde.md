# Tilde (`~`) Token

> Home directory path variable

## Description

The tilde token is used as a lazy reference to the users home directory.

## Examples

```
» out ~
/home/bob

» out ~joe
/home/joe
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

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [`"Double Quote"`](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`$variable`](../parser/string.md):
  Expand values as a string
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](../parser/brace-quote.md):
  Write a string to the STDOUT without new line
* [`(brace quote)`](../parser/brace-quote.md):
  Write a string to the STDOUT without new line
* [ja](../parser/ja.md):
  
* [out](../parser/out.md):
  
* [reserved-vars](../parser/reserved-vars.md):
  
* [set](../parser/set.md):
  

<hr/>

This document was generated from [gen/parser/variables_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/variables_doc.yaml).