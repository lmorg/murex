# murex

[![Go Report Card](https://goreportcard.com/badge/github.com/lmorg/murex)](https://goreportcard.com/report/github.com/lmorg/murex)
[![GoDoc](https://godoc.org/github.com/lmorg/murex?status.svg)](https://godoc.org/github.com/lmorg/murex)

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

(for compatibility _murex_ does also support the traditional pipe token:
eg `command | command`)

The language supports multiple data types natively; such as JSON, YAML,
CSV, S-Expressions, CSV and even loosely tabulated terminal output (eg
`ps`, `ls -l`, etc). This makes passing data through the pipeline and
parsing output easier when dealing with more complex arrangements of
data than a simple byte stream in traditional shells like Bash.

## Concise yet predictable

Despite the amount of features added to shell, we have tried to keep the
amount of "magic" to a minimum and follow a pretty standard structure so
the language is predictable. However there are times when a little magic
goes a long way. For example _murex_'s support for complex data objects
of differing formats is managed in the pipeline so you don't need to think
about the data format when querying data from them.

    open: file.csv -> [ column_name ] # returns specific columns (or rows) in CSV file
    open: file.json -> [ index ]      # returns specific items from JSON

The index function (`[`) alters its matching algorithm depending on the
piped data type and `open` sets the data type depending on the file
extension or MIME type.

Sometimes you will want fewer guesswork or just the robustness of a forced
behavior. On those occasions you can remove one layer of magic by
casting the data type:

    open: file.txt -> cast csv -> [ column_name ]
    open: file.txt -> cast json -> [ index ]

This awareness of data structures is also utilised in `foreach` (which
will cycle through each index in an array) and `formap` (key/value
iteration against complex objects). See [GUIDE.control-structures.md](docs/GUIDE.control-structures.md)
for more details on these and other control structures.

## More robust scripts / shell one liners

_murex_ employs a few methods to make shell scripting more robust:

Bash, for all it's power, is littered with hidden traps. We aim to
address as many of them as we can without taking the flexibility nor power
away from the intereactive command line. This is achieved through a couple
of key concepts:

### Everything is a function

The biggest breaking change from regular shells is how globbing isn't expanded
by the shell by default. This is instead done by inlining functions as arrays:

    # Bash
    ls -l *.go

    # Murex
    ls -l @{g *.go}

The advantage of _murex_'s method is that we can now offer other ways of
matching file system objects that follows the same idiomatic pattern:

    # Match files by regexp pattern
    ls -l @{rx \.go$}

    # Match only directories
    ls -l @{f +d}

(more information on `g`, `rx` and `f` are available in [GUIDE.quick-start.md](docs/GUIDE.quick-start.md)).

However there will be occasions when you just want an inlined expansion
(eg when using an interactive shell) and that can be achieved via the `@g`
command prefix:

    @g ls -l *.go

### Powerful autocompletion

_murex_ takes a slightly different approach to command line autocompletion,
both from a usability perspective as well as defining autocompletions.

Inspired by IDEs, _murex_ queries man pages directly for flags as well as
"tooltip" descriptions. Custom completions are defined via JSON meaning
simple commands are much easier to define and complex commands can still
fallback to using dynamic shell code just like they are in other shells.

This makes it easier to write autocompletions as well as making the code
more readable. An example of `git`s autocompletion definiton:

    private git-branch {
        # returns git branches and removes the current one from the list
        git branch -> [ :0 ] -> !match *
    }

    autocomplete set git { [{
        # define the top level flags
        "Flags": [
            "clone", "init", "add", "mv", "reset", "rm", "bisect", "grep",
            "log", "show", "status", "branch", "checkout", "commit", "diff",
            "merge", "rebase", "tag", "fetch", "pull", "push", "stash"
        ],

        # specify what values those flags support
        "FlagValues": {
            "init": [{
                "Flags": [ "--bare" ]
            }],
            "add": [{
                "IncFiles": true,
                "AllowMultiple": true
            }],
            "mv": [{
                "IncFiles": true
            }],
            "rm": [{
                "IncFiles": true,
                "AllowMultiple": true
            }],
            "checkout": [{
                "Dynamic": ({ git-branch }),
                "Flags": [ "-b" ]
            }],
            "merge": [{
                "Dynamic": ({ git-branch })
            }]
        }
    }] }

_murex_ also supports several different styles of completion suggestion
"popups" to cater for different scenarios (demo below) as well as built in
support for jumping to files within nested directories quickly and easely:

    cat [ctrl+f]app.g[return]
    # same as typing: cat config/app.go

### Error handling

Like traditional shells, _murex_ is verbose with errors by default with
options to mute them. However _murex_ also supports cleaner decision
structures for when you want you want errors captured in a useful way:

    try {
        # do something
    }
    catch {
        err: "Could not perform action"
    }

As well as a saner `if` syntax:

    if { = `foo`==`bar` } then {
        out: "`foo` matched `bar`"
    }

    !if { foobar } else {
        err: "`foobar` could not be run"
    }

### Test and debugging frameworks

Unlike traditional shells, _murex_ is designed with a test and debugging modes
baked into the shell langauge. This means you can write tests against your
shell scripts as part of the shell scripts itself.

For example:

    func hello-world {
        test define example {
            "OutRegexp": (^Hello World$)
        }

        out <test_example> "Hello Earth"
    }

    test run { hello-world }

...will output:

    Hello Earth
     Status  Definition Function                                           Line Col. Message
    [FAILED] example    out                                                5    9    stdout: regexp did not match 'Hello Earth'

If test mode isn't enabled then any `test` commands are skipped without being
executed so you can liberally include test cases throughout your functions
without worrying about any performance impact.

## Interactive shell

Aside the language being designed to offer readability and more robust shell
scripting, the interactive shell itself is also designed around productivity.
To do this, we wrote our own readline library. Below is an example of that
library in use:

[![asciicast](https://asciinema.org/a/232714.svg)](https://asciinema.org/a/232714)

## Language guides

The following guides are historic and the language has been refined a little
since their creation. They are in the process of being rewritten in a format
that allows for auto-generation, however we have retain these guides for
reference in the interim.

1. [GUIDE.syntax.md](docs/GUIDE.syntax.md) is recommended first as it gives
an overview if the shell scripting languages syntax and data types.

2. [GUIDE.variables-and-evaluation.md](docs/GUIDE.variables-and-evaluation.md)
describes how to define variables and how to use them.

3. [GUIDE.control-structures.md](docs/GUIDE.control-structures.md) will
list how to use if statements and iteration like for loops.

4. [GUIDE.arrays-and-maps.md](docs/GUIDE.arrays-and-maps.md) demonstrates how
to create arrays and return specific fields from an array or map.

5. [GUIDE.type-system.md](docs/GUIDE.type-system.md) describes _murex_'s type
system. Most of the time you will not need to worry about typing in
_murex_ as the shell is designed around productivity.

6. [GUIDE.builtin-functions.md](docs/GUIDE.builtin-functions.md) lists some
of the builtin functions available for this shell.

Or if you're already a seasoned Bash developer then you read the Quick
Start Guide, [GUIDE.quick-start.md](docs/GUIDE.quick-start.md), to jump
straight into using _murex_.

## Known bugs / TODO

Please see Github's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)
