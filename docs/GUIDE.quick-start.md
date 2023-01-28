# Language Guide: Quick start guide

This is a cheat sheet reference for new comers and experienced users alike. It
touches on a broad range of features at a high level without delving deep into
the mechanics nor options available for each syntax sugar nor builtin function.

## Bash / _murex_ Rosetta Stone

If you already know what you're trying to do in Bash and looking for the
equivalent syntax in _murex_, then you might find the [Rosetta Stone](user-guide/rosetta-stone.md)
a useful reference.

### Footnotes

1. Supported for compatibility with traditional shells like Bash.
2. Unlike Bash, whitespace (or the absence of) is optional.
3. Environmental variables can only be stored as a string. This is a limitation of current operating systems.
4. Path separator can be any 1 byte wide character, eg `/`. The path separator is defined by the first character in a path.

## Expressions and Statements

An **expression** is an evaluation, operation or assignment, for example:
```
» 6 > 5
» fruit = %[ apples oranges bananas ]
» 5 + 5
```

Whereas a **statement** is a shell command to execute:
```
» echo "hello world"
```

Due to the expectation of shell commands supporting bareword parameters,
expressions have to be parsed differently to statements. Thus _murex_ first
parses a command line to see if it is a valid expression, and if it is not, it
then assumes it is an statement and parses it as such.

This allow expressions and statements to be used interchangeably in a pipeline:
```
» 5 + 5 | grep 10
```

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

    err: <!out> "error message redirected to stdout"

And to redirect STDOUT to STDERR you would use `<err>`:

    out: <err> "output redirected to stderr"

Likewise you can redirect either STDOUT, or STDERR to `/dev/null` via `<null>`
or `<!null>` respectively.

    command: <!null> # ignore STDERR
    command: <null>  # ignore STDOUT

You can also create your own pipes that are files, network connections, or any
other custom data input or output endpoint. [read more](user-guide/namedpipes.md)

### Redirecting to files

    out: "message" |> truncate-file.txt
    out: "message" >> append-file.txt

## Emendable sub-shells

There are two types of emendable sub-shells: strings and arrays.

* string sub-shells, `${ command }`, take the results from the sub-shell and
  return it as a single parameter. This saves the need to encapsulate the shell
  inside quotation marks.

* array sub-shells, `@{ command }`, take the results from the sub-shell
  and expand it as parameters.

Examples:

    touch ${ %[1,2,3] } # creates a file named '[1,2,3]'
    touch @{ %[1,2,3] } # creates three files, named '1', '2' and '3'

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

    # all text files via globbing:
    ls -l @{g *.txt}

    # all text and markdown files via regexp:
    ls -l @{rx '\.(txt|md)$'}

    # all directories via type matching:
    ls -l @{f +d}

You can also using type matching against globbing and regexp to filter
out types in conjunction with file name matching:

    # all directories named *.txt
    ls -l @{g *.txt -> f +d}

## Brace expansion

In [bash you can expand lists](https://en.wikipedia.org/wiki/Bash_(Unix_shell)#Brace_expansion)
using the following syntax: `a{1..5}b`. In _murex_, like with globbing, brace
expansion is a function: `a: a[1..5]b` and supports a much wider range of lists
that can be expanded. ([read more](commands/a.md))

## Exit code

In bash the variable `$?` would store the exit code. This doesn't exist
in _murex_. Instead there a separate command `exitnum`:

    open: test/fox.txt -> grep: foobar; exitnum

