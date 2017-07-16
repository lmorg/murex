# Language Guide: Quick start guide for Bash developers

This is a cheat sheet reference for lazy Bash developers wishing to
accelerate their introduction into _murex_:

## Piping and redirection

_murex_ supports the `|` pipe just like Bash but the preferred pipe
token in _murex_ the arrow, `->` (those two token are interchangeable).

Redirection of stdout and stderr is very different in _murex_. There is
no support for the `>` or `<` tokens,  instead you name the pipe as the
first parameter:

    err: <!out> "error message redirected to stdout"
    out: <err> "output redirected to stderr"

You can also use named pipes this way too. See [GUIDE.syntax.md](GUIDE.syntax.md#piping)
for more details on named pipes.

## Embeddable subshells

There are two types of embeddable subshells: strings and arrays.

* string subshells, `${ command }`, take the results from the subshell
and return it as a single parameter. Equivalent to the following in bash:
`command "$(subshell command)"`.

* array subshells, `@{ command }`, take the results from the subshell
and expand it as parameters. Arrays can be multiple lines (like in Bash)
or array objects in more complex data formats like JSON. Unlike bash,
other white spaces such as tabs and space characters are not counted as
separators for walking through arrays. This is intentional to allow line
formatting and space characters in file names. Array shells are
equivalent to the following in Bash: `command $(subshell command)`

Examples:

    ls -l ${out: file name}           # works because file name contain space
    ls -l @{out: file1 file2 file3}   # fails because not an array
    ls -l @{out: file1\nfile2\nfile3} # works because output is an array

## Globbing

There isn't auto-expansion of globbing to protect against accidental
damage. Instead globbing is achieved via subshells using either:

* `g` (traditional globbing)
* `rx` (regexp matching in current directory only)
* `f` (file or directory type matching)

Examples:

    # all text files via globbing:
    ls -l @{g *.txt}

    # all text and markdown files via regexp:
    ls -l @{rx '\.(txt|md)$'}

    # all files via type matching:
    ls -l @{f +f}

You can also using type matching against globbing and regexp to filter
out types in conjunction with file name matching:

    # all directories named *.txt
    ls -l @{g *.txt -> f +d}

## Exit code

In bash the variable `$?` would store the exit code. This doesn't exist
in _murex_. Instead there a separate command `exitnum`:

    open: test/fox.txt -> grep: foobar; exitnum

## Array expansion

In bash you can expand arrays using the following syntax: `a{1..5}b`. In
_murex_ this is another subshell process: `a: a[1..5]b`. As you can see,
_murex_ also uses square brackets instead as well. There are a few other
changes, read [GUIDE.arrays-and-maps.md](GUIDE.arrays-and-maps.md#the-array-builtin)
for more on using the `array` builtin.