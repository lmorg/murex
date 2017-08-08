# murex

## Install instructions

Install instructions have been moved into its own file: [INSTALL.md](INSTALL.md)

## About _murex_

_murex_ is a cross-platform shell like Bash but with greater emphasis on
writing safe shell scripts and powerful one-liners while maintaining
readability.

To achieve this the language employs a relatively simple syntax modelled
loosely on functional and stack-based programming paradigms (albeit
without the LISP-style nested parentheses that scare a lot of developers).
For example, a program structure could look like the following:

    command -> command -> [ index ] -> if { command }

The language supports multiple data types, with JSON (and later XML)
support as a native data type. Which makes passing data through the
pipeline easier when dealing with more complex arrangements of data than
a simple byte stream when compared to standard shells like Bash.

However for compatibility _murex_ does also support the traditional pipe
token, `|`, and can stream typed data to traditional command line
programs. This means you can use _murex_ with minimal relearning nor
retooling.

## Concise yet predictable

Despite the amount of features added to shell, I have tried to keep the
amount of "magic" to a minimum and follow a pretty standard structure so
the language is predictable. However there are times when a little magic
goes a long way. For example _murex_ supports complex data objects from
various formats including JSON and CSV files and you can query their
properties directly:

    text: file.csv -> [ $column_name ] # return specific columns in CSV file
    text: file.json -> [ $index ]      # return specific items from JSON

The index function (`[`) alters its matching algorithm depending on the
piped data type and `text` sets the data type depending on the file
extension.

Sometimes you will want fewer guesswork or just the robustness a forced
behavior. On those occasions you can remove one layer of magic by
casting the data type:

    text: file.txt -> cast csv -> [ $column_name ]
    text: file.txt -> cast json -> [ $index ]

This awareness of data structures is also utilised in `foreach` (which
will cycle through each index in an array) and `formap` (key/value
iteration against complex objects). See [GUIDE.control-structures.md](./GUIDE.control-structures.md)
for more details on these and other control structures.

## More robust scripts / shell one liners

_murex_ employs a few methods to make shell scripting more robust:

Bash, for all it's power, is littered with hidden traps. I'm aiming to
address as many of them as I can without taking the flexibility or power
away from the command line. This is achieved through a couple of key
concepts:

* Everything is a function

The biggest breaking change from regular shells (or introduced annoyance
as I'm sure some might see it) is how globbing isn't auto-expanded by
the shell. This is instead done by inlining functions as arrays:

    # Bash
    ls -l *.go

    # Murex
    ls -l @{g *.go}

The advantage of _murex_'s method is that we can now offer other ways of
matching file system objects that follows the same idiomatic pattern:

    # Match files by regexp pattern
    ls -l @{rx '\.go$}

    # Match only directories
    ls -l @{f +d}

(more information on `g`, `rx` and `f` are available in [GUIDE.quick-start.md](./GUIDE.quick-start.md)).

* Powerful autocompletion

I've modelled _murex_'s autocompletion after what I would expect if I
were to use an IDE. While _murex_'s autocompletion is a long way from
the power of, for example, IntelliJ or Visual Studio, _murex_ does go a
long way further than your traditional shells, for example it imports
command line flags from their man page.

* Error handling

Like traditional shells, _murex_ is verbose with errors by default with
options to mute them. However _murex_ also support cleaner decision
structures for working with processes you want errors captured:

    try {
        # do soemthing
    } -> catch {
        err: "Could not perform action"
    }

As well as a saner `if` syntax:

    if { = `foo`==`bar` } {
        out: "`foo` matched `bar`"
    }

    !if { foobar } {
        err: "`foobar` could not be run"
    }

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

5. [GUIDE.type-system.md](GUIDE.type-system.md) describes _murex_'s type
system. Most of the time you will not need to worry about typing in
_murex_ as the shell is designed around productivity.

6. [GUIDE.builtin-functions.md](GUIDE.builtin-functions.md) lists some
of the builtin functions available for this shell.

Or if you're already a seasoned Bash developer then you read the Quick
Start Guide, [GUIDE.quick-start.md](GUIDE.quick-start.md), to jump
straight into using _murex_.

## Known bugs / TODO

These have now been moved into Github's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)
