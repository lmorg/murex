# murex
(I'm not sold on that name either. However I am open to suggestions)

## Description

Murex is a cross-platform shell like Bash but with greater emphasis on
writing shell scripts and powerful one-liners while maintaining
readability.

To achieve this the language employs a relatively simple syntax modelled
loosely on functional and stack-based programming paradigms (albeit
without the LISP-style nested parentheses that scare a lot of developers.
For example, a program structure could look like the following:
```
command -> command -> if { then_command } -> else { else_command }
```

The language supports multiple data types, with JSON (and later XML)
support as a native data type. Which makes passing data through the
pipeline easier when dealing with more complex arrangements of data than
a simple byte stream when compared to standard shells like Bash.

However despite the amount of features added to shell, I have  tried to
keep the amount of "magic" to a minimum and follow a pretty standard
structure so the language is predictable and guessable.
 
I will also be working on hardening the shell to make it more robust for
writing shell scripts. Bash, for all it's power, is littered with hidden
traps. I'm hoping to address as many of them as I can without taking
much flexibility nor power away from the command line.

## Dependencies
```
go get github.com/chzyer/readline
go get github.com/kr/pty
go get github.com/Knetic/govaluate
```

Explanation behind these dependencies:
* `readline` is used for the REPL (interactive mode)
* `pty` is used for spawning pseudo-terminals for shell processes
* `govaluate` evaluates the math formulas (exposed via `eval` and `let`)

## Build
```
go build github.com/lmorg/murex
```

Test the binary (requires Bash):
```
test/regression_test.sh
```

## Language guides

Please read the following guides:

1. [GUIDE.syntax.md](./GUIDE.syntax.md) - this is recommended first as
it gives an overview if the shell scripting languages syntax and data
types.

2. [GUIDE.control-structures.md](./GUIDE.control-structures.md) - this
will list how to use if statements and iteration like for loops.

3. [GUIDE.builtin-functions.md](./GUIDE.builtin-functions.md) - lastly
this will list some of the builtin functions available for this shell.

## Known bugs / TODO

* _Currently no support for interactive commands._ This will need to be
addressed.

* _Interactive shell auto-completion is unreliable._ I have a nasty
feeling I may need to fork the readline package or even create my own
one.

* _Interactive shell does not support multiline scripts._ Related to
previous issue.

* _`foreach` only supports line splitting - not JSON objects._ This is a
TODO rather than bug.

* _No support for piping scripts to the shell executable._ This will be
supported via a `--stdin` flag. It's an easy thing to implement but
wasn't considered necessary for the MVP (minimum viable product).

* _Subshells._ At some point I am planning on implementing support for
subshells like `$( command )` (or back ticked) stuff are in Bash.
However I'm considering this a low priority as it has a scope to cause
more damage if not implemented correctly.

* _No support for file globbing._ This is one feature I'm most nervous
about implementing because of the issues it can create eg with filename
containing spaces or prefixed with hyphens. I have a few ideas on how to
do this safely but will welcome any better suggestions. I was also
tempted by the idea to support regex pattern matching as well as (or
instead of) globbing. Regex, being more complex to write expressions,
does offer a possibility for greater user error so this might just be an
optional feature that is disabled by default.
