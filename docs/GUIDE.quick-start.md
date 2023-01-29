<h1>Language Guide: Quick Tour</h1>

<div id="toc">

- [Introduction](#introduction)
  - [Barewords](#barewords)
  - [Expressions and Statements](#expressions-and-statements)
  - [Functions and Methods](#functions-and-methods)
- [Rosetta Stone](#rosetta-stone)
- [Basic Syntax](#basic-syntax)
  - [Quoting Strings](#quoting-strings)
  - [Comments](#comments)
- [Variables](#variables)
  - [Scalars](#scalars)
  - [Arrays](#arrays)
- [Piping and redirection](#piping-and-redirection)
  - [Pipes](#pipes)
  - [Redirection](#redirection)
  - [Redirecting to files](#redirecting-to-files)
- [Emendable sub-shells](#emendable-sub-shells)
- [Globbing](#globbing)
- [Brace expansion](#brace-expansion)

</div>

## Introduction

_murex_ is a typed shell. By this we mean it still passes byte streams along
POSIX pipes (and thus will work with all your existing command line tools) but
in addition will add annotations to describe the type of data that is being
written and read. This allows _murex_ to expand upon your command line tools
with some really interesting and advanced features not available in traditional
shells.

> POSIX is a set of underlying standards that Linux, macOS and various other
> operating systems support.

### Barewords

Shells need to [balance scripting with an efficient interactive terminal](blog/split_personalities.md)
interface. One of the most common approaches to solving that conflict between
readability and terseness is to make heavy use of barewords. Barewords are
ostensibly just instructions that are not quoted. In our case, command names
and command parameters.

_murex_ also makes heavy use of barewords and so that places requirements on
the choice of syntax we can use.

### Expressions and Statements

An **expression** is an evaluation, operation or assignment, for example:
```
» 6 > 5
» fruit = %[ apples oranges bananas ]
» 5 + 5
```

> Expressions are type sensitive

Whereas a **statement** is a shell command to execute:
```
» echo "Hello Murex"
» kill 1234
```

> All values in a statement are treated as strings

Due to the expectation of shell commands supporting bareword parameters,
expressions have to be parsed differently to statements. Thus _murex_ first
parses a command line to see if it is a valid expression, and if it is not, it
then assumes it is an statement and parses it as such.

This allow expressions and statements to be used interchangeably in a pipeline:
```
» 5 + 5 | grep 10
```

### Functions and Methods

A **function** is command that doesn't take data from STDIN whereas a **method**
is any command that does.
```
echo "Hello Murex" | grep "Murex"
^ a function         ^ a method
```

In practical terms, functions and methods are executed in exactly the same way
however some builtins might behave differently depending on whether values are
passed via STDIN or as parameters. Thus you will often find references to
functions and methods, and sometimes for the same command, within these
documents.

## Rosetta Stone

If you already know Bash and looking for the equivalent syntax in _murex_, then
our [Rosetta Stone](user-guide/rosetta-stone.md) reference will help you to
translate your Bash code into _murex_ code.

## Basic Syntax

### Quoting Strings

> It is important to note that all strings in expressions are quoted whereas
> strings in statements can be barewords.

There are three ways to quote a string in _murex_:

* `'single quote'`: use this for string literals    ([read more](parser/single-quote.md))
* `"double quote"`: use this for infixing variables ([read more](parser/double-quote.md))
* `%(brace quote)`: use this for nesting quotes     ([read more](parser/brace-quote.md))

### Comments

You can comment out a single like, or end of a line with `#`:
```
# this is a comment

echo Hello Murex # this is also a comment
```

Multiple lines or mid-line comments can be achieved with `/#` and `#/` tokens:
```
/#
This is
a multi-line
command
#/

echo Hello /# comment #/ Murex
```

(`/#` was chosen because it is similar to C-style comments however `/*` is a
valid glob so _murex_ has substituted the asterisks with a hash symbol instead)

## Variables

All variables are defined with one of three key words:

* `set`    - local variables         ([read more](commands/set.md))
* `global` - global variables        ([read more](commands/global.md))
* `export` - environmental variables ([read more](commands/export.md))

...or via an expression:

* `name = "bob"`
* `age = 20 * 2`
* `fruit = %[ apples oranges bananas ]`

If any variables are unset then reading from them will produce an error (under
_murex_'s default behavior):
```
» echo $foobar
Error in `echo` (1,1): variable 'foobar' does not exist
```

> Please note that when using `set` and `global` as a function, all assignments
> will be strings unless you specifically annotate your variables.

### Scalars

In traditional shells, variables are expanded in a way that results in spaces
be parsed as different command parameters. This results in numerous problems
where developers need to remember to enclose variables inside quotes.

_murex_ parses variables as tokens and expands them into the command line
arguments intuitively. So, there are no more accidental bugs due to spaces in
file names, or other such problems due to developers forgetting to quote
variables. For example:
```
» file = "file name.txt"
» touch $file # this would normally need to be quoted
» ls
'file name.txt'
```
### Arrays

Due to variables not being expanded into arrays by default, _murex_ supports an
additional variable construct for arrays. These are `@` prefixed:
```
» files = %[file1.txt, file2.txt, file3.txt]
» touch @files
» ls
file1.txt  file2.txt
```

## Piping and redirection

### Pipes

_murex_ supports multiple different pipe tokens. The main two being `|` and
`->`.

* `|` works exactly the same as in any normal shell ([read more](parser/pipe-posix.md))

* `->` displays all of the supported methods (commands that support the output
  of the previous command). Think of it a little like object orientated
  programming where an object will have functions (methods) attached. ([read more](parser/pipe-arrow.md))

In _murex_ scripts you can use `|` and `->` interchangeably, so there's no need
to remember which commands are methods and which are not. The difference only
applies in the interactive shell where `->` can be used with tab-autocompletion
to display a shortlist of supported functions that can manipulate the data from
the previous command. It's purely a clue to the parser to generate different
autocompletion suggestions to help with your discovery of different commandline
tools.

### Redirection

Redirection of stdout and stderr is very different in _murex_. There is no
support for the `2>` or `&1` tokens,  instead you name the pipe inside angle
brackets, in the first parameter(s).

`out` is that processes STDOUT (fd1), `err` is that processes STDERR (fd2), and
`null` is the equivalent of piping to `/dev/null`.

Any pipes prefixed by a bang means reading from that processes STDERR.

So to redirect STDERR to STDOUT you would use `<!out>`:

```
err: <!out> "error message redirected to stdout"
```

And to redirect STDOUT to STDERR you would use `<err>`:

```
out: <err> "output redirected to stderr"
```

Likewise you can redirect either STDOUT, or STDERR to `/dev/null` via `<null>`
or `<!null>` respectively.

```
command: <!null> # ignore STDERR
command: <null>  # ignore STDOUT
```

You can also create your own pipes that are files, network connections, or any
other custom data input or output endpoint. [read more](user-guide/namedpipes.md)

### Redirecting to files

```
out: "message" |> truncate-file.txt
out: "message" >> append-file.txt
```

## Emendable sub-shells

There are two types of emendable sub-shells: strings and arrays.

* string sub-shells, `${ command }`, take the results from the sub-shell and
  return it as a single parameter. This saves the need to encapsulate the shell
  inside quotation marks.

* array sub-shells, `@{ command }`, take the results from the sub-shell
  and expand it as parameters.

Examples:

```
touch ${ %[1,2,3] } # creates a file named '[1,2,3]'
touch @{ %[1,2,3] } # creates three files, named '1', '2' and '3'
```

The reason _murex_ breaks from the POSIX tradition of using backticks and
parentheses is because _murex_ works on the principle that everything inside
a curly bracket is considered a new block of code.

## Globbing

While glob expansion is supported in the interactive shell, there isn't
auto-expansion of globbing in shell scripts. This is to protect against
accidental damage. Instead globbing is achieved via sub-shells using either:

* `g`  - traditional globbing ([read more](commands/g.md))
* `rx` - regexp matching in current directory only ([read more](commands/rx.md))
* `f`  - file type matching ([read more](commands/f.md))

Examples:

```
# all text files via globbing:
ls -l @{g *.txt}
```
```
# all text and markdown files via regexp:
ls -l @{rx '\.(txt|md)$'}
```
```
# all directories via type matching:
ls -l @{f +d}
```

You can also using type matching against globbing and regexp to filter
out types in conjunction with file name matching:

```
# all directories named *.txt
ls -l @{g *.txt -> f +d}
```

## Brace expansion

In [bash you can expand lists](https://en.wikipedia.org/wiki/Bash_(Unix_shell)#Brace_expansion)
using the following syntax: `a{1..5}b`. In _murex_, like with globbing, brace
expansion is a function: `a: a[1..5]b` and supports a much wider range of lists
that can be expanded. ([read more](commands/a.md))

