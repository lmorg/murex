<h1>Language Tour</h1>

<h2>Table of Contents</h2>

<div id="toc">

- [Introduction](#introduction)
  - [Read–Eval–Print Loop](#readevalprint-loop)
  - [Barewords](#barewords)
  - [Expressions and Statements](#expressions-and-statements)
  - [Functions and Methods](#functions-and-methods)
  - [The Bang Prefix](#the-bang-prefix)
- [Rosetta Stone](#rosetta-stone)
- [Basic Syntax](#basic-syntax)
  - [Quoting Strings](#quoting-strings)
  - [Code Comments](#code-comments)
- [Variables](#variables)
  - [Global variables](#global-variables)
  - [Environmental Variables](#environmental-variables)
  - [Type Inference](#type-inference)
  - [Scalars](#scalars)
  - [Arrays](#arrays)
- [Piping and Redirection](#piping-and-redirection)
  - [Pipes](#pipes)
  - [Redirection](#redirection)
  - [Redirecting to files](#redirecting-to-files)
  - [Type Conversion](#type-conversion)
    - [Cast](#cast)
    - [Format](#format)
- [Sub-Shells](#sub-shells)
- [Filesystem Wildcards (Globbing)](#filesystem-wildcards-globbing)
- [Brace expansion](#brace-expansion)
- [Executables](#executables)
  - [Aliases](#aliases)
  - [Public Functions](#public-functions)
  - [Private Functions](#private-functions)
  - [External Executables](#external-executables)
- [Control Structures](#control-structures)
  - [Using `if` Statements](#using-if-statements)
  - [Using `switch` Statements](#using-switch-statements)
  - [Using `foreach` Loops](#using-foreach-loops)
  - [Using `formap` Loops](#using-formap-loops)
- [Stopping Execution](#stopping-execution)
  - [The `continue` Statement](#the-continue-statement)
  - [The `break` Statement](#the-break-statement)
  - [The `return` Statement](#the-return-statement)
  - [The `exit` Statement](#the-exit-statement)
  - [Signal: SIGINT](#signal-sigint)
  - [Signal: SIGQUIT](#signal-sigquit)
  - [Signal: SIGTSTP](#signal-sigtstp)

</div>

## Introduction

Murex is a typed shell.

Unlike other typed shells, Murex can still work natively with existing CLI
tools without any tweaks.

Murex's unique approach to type annotations means you have the safety and
convenience of working with data formats besides just byte streams and string
variables, while still having compatibility with every tool written for Linux
and UNIX over the last 50 years.

### Read–Eval–Print Loop



Murex doesn't just aim to be a well thought out language, the interactive shell
boast an excellent out-of-the-box experience with makes other shells feel stone
age in comparison.

If you want to learn more about the interactive shell then there is a dedicated
document detailing [Murex's REPL features](/docs/user-guide/interactive-shell.md).

### Barewords

Shells need to [balance scripting with an efficient interactive terminal](/docs/blog/split_personalities.md)
interface. One of the most common approaches to solving that conflict between
readability and terseness is to make heavy use of barewords. Barewords are
ostensibly just instructions that are not quoted. In our case, command names
and command parameters.

Murex also makes heavy use of barewords and so that places restrictions on the
choice of syntax we can use.

### Expressions and Statements

An _expression_ is an evaluation, operation or assignment, for example:
```
» 6 > 5
» fruit = %[ apples oranges bananas ]
» 5 + 5
```

> Expressions are type sensitive

Whereas a _statement_ is a shell command to execute:
```
» echo "Hello Murex"
» kill 1234
```

> All values in a statement are treated as strings

Due to the expectation of shell commands supporting bareword parameters,
expressions have to be parsed differently to statements. Thus Murex first
parses a command line to see if it is a valid expression, and if it is not, it
then assumes it is an statement and parses it as such.

This allow expressions and statements to be used interchangeably in a pipeline:
```
» 5 + 5 | grep 10
```

### Functions and Methods

A _function_ is command that doesn't take data from stdin whereas a _method_
is any command that does.
```
echo "Hello Murex" | grep "Murex"
^ a function         ^ a method
```

In practical terms, functions and methods are executed in exactly the same way
however some builtins might behave differently depending on whether values are
passed via stdin or as parameters. Thus you will often find references to
functions and methods, and sometimes for the same command, within these
documents.

### The Bang Prefix

Some Murex builtins support a bang prefix. This prefix alters the behavior of
those builtins to perform the conceptual opposite of their primary role.

For example, you could grep a file with `regexp 'm/(dogs|cats)/'` but then you
might want to exclude any matches by using `!regexp 'm/(dogs|cats)/'` instead.

The details for each supported bang prefix will be in the documents for their
respective builtin.

## Rosetta Stone

If you already know Bash and looking for the equivalent syntax in Murex, then
our [Rosetta Stone](/docs/user-guide/rosetta-stone.md) reference will help you to
translate your Bash code into Murex code.

## Basic Syntax

A table for all operators and tokens can be found on the [operators and tokens cheatsheet](/docs/user-guide/operators-and-tokens.md).
However if you are new to the shell, we recommend that you read this document
first.

### Quoting Strings

> It is important to note that all strings in expressions are quoted whereas
> strings in statements can be barewords.

There are three ways to quote a string in Murex:

* `'single quote'`: use this for string literals    ([read more](/docs/parser/single-quote.md))
* `"double quote"`: use this for infixing variables ([read more](/docs/parser/double-quote.md))
* `%(brace quote)`: use this for nesting quotes     ([read more](/docs/parser/brace-quote.md))

### Code Comments

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
```
...which can also be inlined...
```
» echo Hello /# comment #/ Murex
```

(`/#` was chosen because it is similar to C-style comments however `/*` is a
valid glob so Murex has substituted the asterisks with a hash symbol instead)

## Variables

All variables can be defined as expressions and their data types are inferred:

* `name = "bob"`
* `age = 20 * 2`
* `fruit = %[ apples oranges bananas ]`

If any variables are unset then reading from them will produce an error (under
Murex's default behavior):
```
» echo $foobar
Error in `echo` (1,1): variable 'foobar' does not exist
```

### Global variables

Global variables can be defined using the `$GLOBAL` namespace:
```
» $GLOBAL.foo = "bar"
```

You can also force Murex to read the global assignment of `$foo` (ignoring
any local assignments, should they exist) using the same syntax. eg:
```
» $GLOBAL.name = "Tom"
» out $name
Tom

» $name = "Sally"
» out $GLOBAL.name
Tom
» out $name
Sally
```

### Environmental Variables

Environmental Variables are like global variables except they are copied to any
other programs that are launched from your shell session.

Environmental variables can be assigned using the `$ENV` namespace:
```
» $ENV.foo = "bar"
```
as well as using the `export` statement like with traditional shells. ([read more](/docs/commands/export.md))

Like with global variables, you can force Murex to read the environmental
variable, bypassing and local or global variables of the same name, by also
using the `$ENV` namespace prefix.

### Type Inference

In general, Murex will try to infer the data type of a variable or pipe. It
can do this by checking the `Content-Type` HTTP header, the file name
extension or just looking at how that data was constructed (when defined via
expressions). However sometimes you may need to annotate your types. ([read more](/docs/commands/set.md#type-annotations))

### Scalars

In traditional shells, variables are expanded in a way that results in spaces
be parsed as different command parameters. This results in numerous problems
where developers need to remember to enclose variables inside quotes.

Murex parses variables as tokens and expands them into the command line
arguments intuitively. So, there are no more accidental bugs due to spaces in
file names, or other such problems due to developers forgetting to quote
variables. For example:
```
» file = "file name.txt"
» touch $file   # this would normally need to be quoted
» ls
'file name.txt'
```
### Arrays

Due to variables not being expanded into arrays by default, Murex supports an
additional variable construct for arrays. These are `@` prefixed:
```
» files = %[file1.txt, file2.txt, file3.txt]
» touch @files
» ls
file1.txt  file2.txt  file3.txt
```

## Piping and Redirection

### Pipes

Murex supports multiple different pipe tokens. The main two being `|` and
`->`.

* `|` works exactly the same as in any normal shell. ([read more](/docs/parser/pipe-posix.md))

* `->` displays all of the supported methods (commands that support the output
  of the previous command). Think of it a little like object orientated
  programming where an object will have functions (methods) attached. ([read more](/docs/parser/pipe-arrow.md))

In Murex scripts you can use `|` and `->` interchangeably, so there's no need
to remember which commands are methods and which are not. The difference only
applies in the interactive shell where `->` can be used with tab-autocompletion
to display a shortlist of supported functions that can manipulate the data from
the previous command. It's purely a clue to the parser to generate different
autocompletion suggestions to help with your discovery of different command
line tools.

### Redirection

Redirection of stdout and stderr is very different in Murex. There is no
support for the `2>` or `&1` tokens,  instead you name the pipe inside angle
brackets, in the first parameter(s).

`out` is that processes stdout (fd1), `err` is that processes stderr (fd2), and
`null` is the equivalent of piping to `/dev/null`.

Any pipes prefixed by a bang means reading from that processes stderr.

So to redirect stderr to stdout you would use `<!out>`:
```
err <!out> "error message redirected to stdout"
```

And to redirect stdout to stderr you would use `<err>`:
```
out <err> "output redirected to stderr"
```

Likewise you can redirect either stdout, or stderr to `/dev/null` via `<null>`
or `<!null>` respectively.
```
command <!null> # ignore stderr
command <null>  # ignore stdout
```

You can also create your own pipes that are files, network connections, or any
other custom data input or output endpoint. ([read more](/docs/user-guide/namedpipes.md))

### Redirecting to files

```
out "message" |> truncate-file.txt
out "message" >> append-file.txt
```

### Type Conversion

Aside from annotating variables upon definition, you can also transform data
along the pipeline using `format`.

#### Cast

Casting doesn't alter the data, it simply changes the meta-information about
how that data should be read.
```
out [1,2,3] | cast json | foreach { ... }
```

There is also a little syntactic sugar to do the same:
```
out [1,2,3] | :json: foreach { ... }
```

#### Format

`format` takes the source data and reformats it into another data format:
```
» out [1,2,3] | :json: format yaml
- 1
- 2
- 3
```

## Sub-Shells

There are two main types of emendable sub-shells: strings and arrays.

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

The reason Murex breaks from the traditions of using backticks and parentheses
is because Murex works on the principle that everything inside a curly bracket
is considered a new block of code.

## Filesystem Wildcards (Globbing)

While glob expansion is supported in the interactive shell, there isn't
auto-expansion of globbing in shell scripts. This is to protect against
accidental damage. Instead globbing is achieved via sub-shells using either:

* `g`  - traditional globbing                      ([read more](/docs/commands/g.md))
* `rx` - regexp matching in current directory only ([read more](/docs/commands/rx.md))
* `f`  - file type matching                        ([read more](/docs/commands/f.md))

Examples:

All text files via globbing:
```
g *.txt
```

All text and markdown files via regexp:
```
rx '\.(txt|md)$'
```

All directories via type matching:
```
f +d
```

You can also chain them together, eg all directories named `*.txt`:
```
g *.txt | f +d
```

To use them in a shell script it could look something a like this:
```
rm @{g *.txt | f +s}
```
(this deletes any symlinks called `*.txt`)

## Brace expansion

In [bash you can expand lists](https://en.wikipedia.org/wiki/Bash_(Unix_shell)#Brace_expansion)
using the following syntax: `a{1..5}b`. In Murex, like with globbing, brace
expansion is a function: `a a[1..5]b` and supports a much wider range of lists
that can be expanded. ([read more](/docs/commands/a.md))

## Executables

### Aliases

You can create "aliases" to common commands to save you a few keystrokes. For
example:
```
alias gc=git commit
```

`alias` behaves slightly differently to Bash. ([read more](/docs/commands/alias.md))

### Public Functions

You can create custom functions in Murex using `function`. ([read more](/docs/commands/function.md))
```
function gc (message: str) {
    # shorthand for `git commit`
    
    git commit -m $message
}
```

### Private Functions

`private` functions are like [public functions](#public-functions) except they
are only available within their own modules namespace. ([read more](/docs/commands/private.md))

### External Executables

External executables (including any programs located in `$PATH`) are invoked
via the `exec` builtin ([read more](/docs/commands/exec.md)) however if a command
isn't an expression, alias, function nor builtin, then Murex assumes it is an
external executable and automatically invokes `exec`.

For example the two following statements are the same:

1. `exec uname`
2. `uname`

Thus for normal day to day usage, you shouldn't need to include `exec`.

## Control Structures

### Using `if` Statements

`if` can be used in a number of different ways, the most common being:
```
if { true } then {
    # do something
} else {
    # do something else
}
```

`if` supports a flexible variety of incarnation to solve different problems. ([read more](/docs/commands/if.md))

### Using `switch` Statements

Because `if ... else if` chains are ugly, Murex supports `switch` statements:
```
switch $USER {
    case "Tom"   { out "Hello Tom" }
    case "Dick"  { out "Howdie Richard" }
    case "Sally" { out "Nice to meet you" }

    default {
        out "I don't know who you are"
    }
}
```

`switch` supports a flexible variety of different usages to solve different
problems. ([read more](/docs/commands/switch.md))

### Using `foreach` Loops

`foreach` allows you to easily iterate through an array or list of any type: ([read more](/docs/commands/foreach.md))
```
%[ apples bananas oranges ] | foreach fruit { out "I like $fruit" }
```

### Using `formap` Loops

`formap` loops are the equivalent of `foreach` but against map objects: ([read more](/docs/commands/formap.md))
```
%{
    Bob:     {age: 10},
    Richard: {age: 20},
    Sally:   {age: 30}
} | formap name person {
    out "$name is $person[age] years old"
}
```

## Stopping Execution

### The `continue` Statement

`continue` will terminate execution of an inner block in iteration loops like
`foreach` and `formap`. Thus _continuing_ the loop from the next iteration:
```
%[1..10] | foreach i {
    if { $i == 5 } then {
        continue foreach
        # ^ jump back to the next iteration
    }

    out $i
}
```

`continue` requires a parameter to define while block to iterate on. This means
you can use `continue` within nested loops and still have readable code. ([read more](/docs/commands/continue.md))

### The `break` Statement

`break` will terminate execution of a block (eg `function`, `private`, `if`,
`foreach`, etc):
```
%[1..10] | foreach i {
    if { $i == 5 } then {
        break foreach
        # ^ exit foreach
    }

    out $i
}
```

`break` requires a parameter to define while block to end. Thus `break` can be
considered to exhibit the behavior of _return_ as well as _break_ in other
languages:
```
function example {
    if { $USER == "root" } then {
        err "Don't run this as root"
        break example
    }

    # ... do something ...
}
```

`break` cannot exit anything above it's callers scope. ([read more](/docs/commands/break.md))

### The `return` Statement

`return` ends the current scope (typically a function). ([read more](/docs/commands/return.md))

### The `exit` Statement

`exit` terminates Murex. It is not scope aware; if it is included in a function
then the whole shell will still exist and not just that function. ([read more](/docs/commands/exit.md))

### Signal: SIGINT

This can be invoked by pressing `Ctrl` + `c`.

This is functionally the same as `fid-kill`. (([read more](/docs/commands/fid-kill.md)))

### Signal: SIGQUIT

This can be invoked by pressing `Ctrl` + `\`. ([read more](/docs/user-guide/terminal-keys.md))

Sending SIGQUIT will terminate all running functions in the current Murex
session. Which is a handy escape hatch if your shell code starts misbehaving.

### Signal: SIGTSTP

This can be invoked by pressing `Ctrl` + `z`. ([read more](/docs/user-guide/job-control.md))

## See Also

* [Install](/INSTALL.md):
  Installation instructions
* [Operators And Tokens](/operators-and-tokens.md):
  A table of all supported operators and tokens

<hr/>

This document was generated from [gen/root/tour_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/root/tour_doc.yaml).