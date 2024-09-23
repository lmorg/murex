# Regex Operations (`regexp`)

> Regexp tools for arrays / lists of strings

## Description

`regexp` provides a few tools for text matching and manipulation against an
array or list of strings - thus `regexp` is Murex data-type aware.

## Usage

```
<stdin> -> regexp expression -> <stdout>
```

Where _expression_ consists of the following

```
function separator pattern [ separator parameter2 ]
```

* _function_: single alphabetic character (eg, `m`, `s` or `f`)

* _separator_: typically a single unicode character (eg, `/`, `#`, `▷`, `🙂`)
  however `regexp` _parameters_ can also be passed as shell parameters (eg
  `s search replace`)

* _pattern_: regexp pattern

* _parameter2_: any extra operations to perform. For example a string to
  replace matched patterns when using the regexp substitution function

## Examples

### Find elements

```
» ja [monday..sunday] -> regexp 'f/^([a-z]{3})day/'
[
    "mon",
    "fri",
    "sun"
]
```

This returns only 3 days because only 3 days match the expression (where
the days have to be 6 characters long) and then it only returns the first 3
characters because those are inside the parenthesis.

### Match elements

#### Elements containing

```
» ja [monday..sunday] -> regexp 'm/(mon|fri|sun)day/'
[
    "monday",
    "friday",
    "sunday"
]
```

#### Elements excluding

```
» ja [monday..sunday] -> !regexp 'm/(mon|fri|sun)day/'
[
    "tuesday",
    "wednesday",
    "thursday",
    "saturday"
]
```

#### Include heading

```
» ps -fe -> regexp 'M/murex/'
UID   PID  PPID   C STIME   TTY           TIME CMD
501 39631 39630   0  6:48pm ??         0:04.31 -murex
501 57496 17220   0 11:59pm ??         0:02.90 ./murex
501 41982 17219   0 10:53pm ttys000    0:39.73 -murex
501 17220 17219   0  2:09pm ttys002    1:44.06 -murex 
```

### Substitute expression

```
» ja [monday..sunday] -> regexp 's/day/night/'
[
    "monnight",
    "tuesnight",
    "wednesnight",
    "thursnight",
    "frinight",
    "saturnight",
    "sunnight"
]
```

## Flags

* `M`
    output first element (eg table headings), followed by any elements that match (supports bang prefix)
* `f`
    output found expressions (doesn't support bang prefix)
* `m`
    output elements that match expression (supports bang prefix)
* `s`
    output all elements, substituting elements that match expression (doesn't support bang prefix)

## Detail

### Data Types

`regexp` is data-type aware so will work against lists or arrays of whichever
Murex data-type is passed to it via stdin and return the output in the
same data-type.

### Inverse Matches

If you want to exclude any matches based on wildcards, rather than include
them, then you can use the bang prefix. For example if you wanted to exclude
any days of the week that contained the letter `s`:

```
» %[Monday..Friday] -> !regexp m/s/
[
    "Monday",
    "Friday"
]
```

## Expression Syntax

Murex regex expressions are based on [Go's stdlib regexp library](https://pkg.go.dev/regexp).

> The syntax of the regular expressions accepted is the same general syntax
> used by Perl, Python, and other languages. More precisely, it is the syntax
> accepted by RE2 and described at https://golang.org/s/re2syntax, except for
> `\C`.
>
> The regexp implementation provided by this package is guaranteed to run in
> time linear in the size of the input. (This is a property not guaranteed by
> most open source implementations of regular expressions.) For more
> information about this property, see https://swtch.com/~rsc/regexp/regexp1.html
> or any book about automata theory.
>
> All characters are UTF-8-encoded code points. Each byte of an invalid UTF-8
> sequence is treated as if it encoded as **U+FFFD**. 

### Single characters

    .              any character, possibly including newline (flag s=true)
    [xyz]          character class
    [^xyz]         negated character class
    \d             Perl character class
    \D             negated Perl character class
    [[:alpha:]]    ASCII character class
    [[:^alpha:]]   negated ASCII character class
    \pN            Unicode character class (one-letter name)
    \p{Greek}      Unicode character class
    \PN            negated Unicode character class (one-letter name)
    \P{Greek}      negated Unicode character class

### Composites

    xy             x followed by y
    x|y            x or y (prefer x)

### Repetitions

    x*             zero or more x, prefer more
    x+             one or more x, prefer more
    x?             zero or one x, prefer one
    x{n,m}         n or n+1 or ... or m x, prefer more
    x{n,}          n or more x, prefer more
    x{n}           exactly n x
    x*?            zero or more x, prefer fewer
    x+?            one or more x, prefer fewer
    x??            zero or one x, prefer zero
    x{n,m}?        n or n+1 or ... or m x, prefer fewer
    x{n,}?         n or more x, prefer fewer
    x{n}?          exactly n x

#### Implementation restriction

The counting forms `x{n,m}`, `x{n,}`, and `x{n}` reject forms that create a
minimum or maximum repetition count above 1000. Unlimited repetitions are not
subject to this restriction.

### Grouping

    (re)           numbered capturing group (submatch)
    (?P<name>re)   named & numbered capturing group (submatch)
    (?<name>re)    named & numbered capturing group (submatch)
    (?:re)         non-capturing group
    (?flags)       set flags within current group; non-capturing
    (?flags:re)    set flags during re; non-capturing

    Flag syntax is xyz (set) or -xyz (clear) or xy-z (set xy, clear z). The flags are:

    i              case-insensitive (default false)
    m              multi-line mode: ^ and $ match begin/end line in addition to begin/end text (default false)
    s              let . match \n (default false)
    U              ungreedy: swap meaning of x* and x*?, x+ and x+?, etc (default false)

### Empty strings

    ^              at beginning of text or line (flag m=true)
    $              at end of text (like \z not \Z) or line (flag m=true)
    \A             at beginning of text
    \b             at ASCII word boundary (\w on one side and \W, \A, or \z on the other)
    \B             not at ASCII word boundary
    \z             at end of text

### Escape sequences

    \a             bell (== \007)
    \f             form feed (== \014)
    \t             horizontal tab (== \011)
    \n             newline (== \012)
    \r             carriage return (== \015)
    \v             vertical tab character (== \013)
    \*             literal *, for any punctuation character *
    \123           octal character code (up to three digits)
    \x7F           hex character code (exactly two digits)
    \x{10FFFF}     hex character code
    \Q...\E        literal text ... even if ... has punctuation

### Character class elements

    x              single character
    A-Z            character range (inclusive)
    \d             Perl character class
    [:foo:]        ASCII character class foo
    \p{Foo}        Unicode character class Foo
    \pF            Unicode character class F (one-letter name)

### Named character classes as character class elements

    [\d]           digits (== \d)
    [^\d]          not digits (== \D)
    [\D]           not digits (== \D)
    [^\D]          not not digits (== \d)
    [[:name:]]     named ASCII class inside character class (== [:name:])
    [^[:name:]]    named ASCII class inside negated character class (== [:^name:])
    [\p{Name}]     named Unicode property inside character class (== \p{Name})
    [^\p{Name}]    named Unicode property inside negated character class (== \P{Name})

### Perl character classes (all ASCII-only)

    \d             digits (== [0-9])
    \D             not digits (== [^0-9])
    \s             whitespace (== [\t\n\f\r ])
    \S             not whitespace (== [^\t\n\f\r ])
    \w             word characters (== [0-9A-Za-z_])
    \W             not word characters (== [^0-9A-Za-z_])

### ASCII character classes

    [[:alnum:]]    alphanumeric (== [0-9A-Za-z])
    [[:alpha:]]    alphabetic (== [A-Za-z])
    [[:ascii:]]    ASCII (== [\x00-\x7F])
    [[:blank:]]    blank (== [\t ])
    [[:cntrl:]]    control (== [\x00-\x1F\x7F])
    [[:digit:]]    digits (== [0-9])
    [[:graph:]]    graphical (== [!-~] == [A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\\]^_`{|}~])
    [[:lower:]]    lower case (== [a-z])
    [[:print:]]    printable (== [ -~] == [ [:graph:]])
    [[:punct:]]    punctuation (== [!-/:-@[-`{-~])
    [[:space:]]    whitespace (== [\t\n\v\f\r ])
    [[:upper:]]    upper case (== [A-Z])
    [[:word:]]     word characters (== [0-9A-Za-z_])
    [[:xdigit:]]   hex digit (== [0-9A-Fa-f])

## Synonyms

* `regexp`
* `!regexp`
* `list.regex`
* `!list.regex`


## See Also

* [Add Prefix (`prefix`)](../commands/prefix.md):
  Prefix a string to every item in a list
* [Add Suffix (`suffix`)](../commands/suffix.md):
  Prefix a string to every item in a list
* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create 2d Array (`2darray`)](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create Map (`map`)](../commands/map.md):
  Creates a map from two data sources
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Match String (`match`)](../commands/match.md):
  Match an exact value in an array
* [Prepend To List (`prepend`)](../commands/prepend.md):
  Add data to the start of an array
* [Prettify JSON](../commands/pretty.md):
  Prettifies data documents to make it human readable
* [Sort Array (`msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

<hr/>

This document was generated from [builtins/core/lists/regexp_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/regexp_doc.yaml).