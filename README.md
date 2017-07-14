# murex

Murex is a cross-platform shell like Bash but with greater emphasis on
writing safe shell scripts and powerful one-liners while maintaining
readability.

To achieve this the language employs a relatively simple syntax modelled
loosely on functional and stack-based programming paradigms (albeit
without the LISP-style nested parentheses that scare a lot of developers).
For example, a program structure could look like the following:

    command -> command -> if { then_command } -> else { else_command }

The language supports multiple data types, with JSON (and later XML)
support as a native data type. Which makes passing data through the
pipeline easier when dealing with more complex arrangements of data than
a simple byte stream when compared to standard shells like Bash.

## Concise yet predictable

However despite the amount of features added to shell, I have  tried to
keep the amount of "magic" to a minimum and follow a pretty standard
structure so the language is predictable and guessable. However there
are times when a little magic goes a long way. For example you _murex_
supports complex data objects from various formats including JSON and
CSV files and you can query those nodes on them directly. eg

    text: file.csv -> [ $column_name ] # return specific columns in CSV file
    text: file.json -> [ $index ]      # return specific items from JSON

The index function (`[`) alters it's matching depending on the piped
data type and `text` sets the data type depending on the file extension.
However sometimes you will want fewer guesswork, and on those occasions
you can remove one layer of magic by casting the data type:

    text: file.txt -> cast csv -> [ $column_name ]
    text: file.txt -> cast json -> [ $index ]

This awareness of data structures is also utilised in `foreach` (cycle
through each index in an array) and `formap` (key/value iteration against
complex objects). See [GUIDE.control-structures.md](./GUIDE.control-structures.md)
for more details on these and other control structures.

## More robust scripts / shell one liners

I will also be working on hardening the shell to make it more robust for
writing shell scripts. Bash, for all it's power, is littered with hidden
traps. I'm hoping to address as many of them as I can without taking
much flexibility nor power away from the command line.

The biggest breaking change from regular shells (or introduced annoyance
as I'm sure some might see it) is how globbing isn't auto-expanded by
the shell. This is instead done by inlining functions as arrays:

    # Bash
    ls -l *.go

    # Murex
    ls -l @{g *.go}

The advantage of _murex_'s method is that we can now offer other idiomatic
ways of matching file system objects:

    # Match files by regexp pattern
    ls -l @{rx '\.go$}

    # Match only directories
    ls -l @{f +d}

(more information on `g`, `rx` and `f` are available in [GUIDE.quick-start.md](./GUIDE.quick-start.md)).

## Dependencies

    go get github.com/chzyer/readline
    go get github.com/kr/pty
    go get github.com/Knetic/govaluate

Explanation behind these dependencies:
* `readline` is used for the REPL (interactive mode)
* `pty` is used for spawning pseudo-terminals for shell processes
* `govaluate` evaluates the math formulas (exposed via `eval` and `let`)

## Build

    go build github.com/lmorg/murex

Test the binary (requires Bash):

    test/regression_test.sh

A Dockerfile is also included for your convenience. The file is located
in test/docker and includes a [README.md](test/docker/README.md) with
more information.

## Language guides

Please read the following guides:

1. [GUIDE.syntax.md](GUIDE.syntax.md) is recommended first as it gives
an overview if the shell scripting languages syntax and data types.

2. [GUIDE.variables-and-evaluation.md](GUIDE.variables-and-evaluation.md)
describes how to define variables and how to use them.

3. [GUIDE.control-structures.md](GUIDE.control-structures.md) will
list how to use if statements and iteration like for loops.

4. [GUIDE.arrays-and-maps.md](GUIDE.arrays-and-maps.md) demonstrates how
to create arrays and return specific fields from an array or map.

5. [GUIDE.builtin-functions.md](GUIDE.builtin-functions.md) lists some
of the builtin functions available for this shell.

Or if you're already a seasoned Bash developer then you read the Quick
Start Guide, [GUIDE.quick-start.md](GUIDE.quick-start.md), to jump
straight into using _murex_.

## Known bugs / TODO

These have now been moved into Github's issue tracker: https://github.com/lmorg/murex/issues
