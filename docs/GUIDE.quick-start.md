# Language Guide: Quick start guide

This is a cheat sheet reference for Bash developers wishing to accelerate
their introduction into _murex_:

## Functions / Methods

_murex_ makes the distinction between commands that are designed to create data
(functions) and those that process data from STDIN (methods).

An example **function** might be `ls` because it doesn't take any inputs but
produces an output.

An example **method** might be `grep` because it takes an input from STDIN and
returns a result via STDOUT.

Some commands might be both functions and methods.

The reason for this distinction is to enable _murex_ to produce meaningful
autocompletion suggestions. Because _murex_ pipes are typed, _murex_ can offer
powerful suggestions based on the expected output of the previous command.

> In the _murex_ docs you might notice commands are often followed by a colon,
> for example:
> ```
> echo: bob | grep: foobar
> ```
> This colon providers the parser with hints as to whether a line of code is
> an executable or an expression. However _murex_ can make inference about the
> code without the colon being included.


## Variables

All variables are defined with one of three key words:

* set    - local variables         ([read more](commands/set.md))
* global - global variables        ([read more](commands/global.md))
* export - environmental variables ([read more](commands/export.md))

If any variables are unset then reading from them will produce an error (under
_murex_'s default behavior):

    » echo $foobar
    Error in `echo` (1,1): variable 'foobar' does not exist

### Scalars

In traditional shells, variables are expanded in a way that results in spaces
be parsed as different command parameters. This results in numerous problems
where developers need to remember to enclose variables inside quotes.

_murex_ parses variables as tokens and expands them into the command line
arguments intuitively. So, there are no more accidental bugs due to spaces in
file names, or other such problems due to developers forgetting to quote
variables:

    » set file=file name.txt
    » touch $file
    » ls
    'file name.txt'

### Arrays

Due to variables not being expanded into arrays by default, _murex_ supports an
additional variable construct for arrays. These are `@` prefixed:

    » set yaml files=[file1.txt, file2.txt]
    » touch @files
    » ls
    file1.txt  file2.txt

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

You can also create your own named pipes (not to be confused with POSIX named
pipes). These pipes could be files, network connections, or any other custom
data input or output endpoint. [read more](user-guide/namedpipes.md)

### Redirecting to files

To redirect to a file you can use the `>` or `>>` functions. They work
similarly to bash except that they are functions rather than tokens. This means
they literally work like the following:

    out: "message" -> >  new-file.txt
    out: "message" -> >> append-file.txt

However this is clearly ugly in practice. So the following syntactic sugar is
supported, `|>` for overwrite and `>>` for append:

    out: "message" |> new-file.txt
    out: "message" >> append-file.txt

## Emendable sub-shells

There are two types of emendable sub-shells: strings and arrays.

* string sub-shells, `${ command }`, take the results from the sub-shell
and return it as a single parameter. Equivalent to the following in bash:
`command "$(sub-shell command)"`.

* array sub-shells, `@{ command }`, take the results from the sub-shell
and expand it as parameters. Arrays can be multiple lines (like in Bash)
or array objects in more complex data formats like JSON. Unlike bash,
other white spaces such as tabs and space characters are not counted as
separators for walking through arrays. This is intentional to allow line
formatting and space characters in file names. Array shells are
equivalent to the following in Bash: `command $(sub-shell command)`

Examples:

    ls -l ${out: file name}           # works because file name contain space
    ls -l @{out: file1 file2 file3}   # fails because not an array
    ls -l @{out: file1\nfile2\nfile3} # works because output is an array

The reason _murex_ breaks from the POSIX tradition of using backticks and
parentheses is because _murex_ works on the principle that everything inside
a curly bracket is considered a new block of code. Typically that would mean
a sub-shell however sometimes it could be configuration code in the form of
inlined JSON.

## Globbing

There isn't auto-expansion of globbing in _murex_ shell scripts, in part due to
its functional nature but also to protect against accidental damage. Instead
globbing is achieved via sub-shells using either:

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

## Back ticks

In _murex_ back ticks do not spawn sub-shells. Back ticks are treated
like a regular, printable, character. Their only special function is
quoting strings in `=`, eg:

    if { = `quoted string`==variable } { out "do something" }
