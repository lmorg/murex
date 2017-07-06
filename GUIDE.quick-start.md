# Language Guide: Quick start guide for Bash developers

This is a cheat sheet reference for lazy Bash developers to accelerate
your introduction into Murex:

## Piping and redirection

Murex supports the `|` pipe just like Bash. However piping Murex methods
is done via `->`. For more information on methods read [GUIDE.syntax.md](./GUIDE.syntax.md).

Redirection of stdout and stderr is very different in Murex. There is no
support for the `>` or `<` tokens. Instead you use `?` to pipe stderr to
the stdout stream.

Example:

    sh -c 'echo "stderr" 2>&1' ? grep stderr

## Subshells

There are two types of subshells, strings and arrays:

* string subshells, `${ command }`, take the results from the subshell
and return it as a single parameter. Equivalent to the following in bash:
`command "$(subshell command)"`.

* array subshells, `@{ command }`, take the results from the subshell
and expand it as parameters. Arrays are defined either as multiple lines
or a JSON array. Unlike bash, other white spaces such as tabs and space
characters are not counted as separators for walking through arrays.
This is intentional to allow line formatting and space characters in
file names. Array shells are equivalent to the following in Bash:
`command $(subshell command)`

Examples:

    ls -l ${echo: file name}
    ls -l @{echo: file1 file2 file3}

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

In Bash the variable `$?` would store the exit code. This doesn't exit in
_murex_. Instead there a separate command `exitnum`:

    open test/fox.txt | grep foobar; exitnum
