[![Version](version.svg)](DOWNLOAD.md)
[![CodeBuild](https://codebuild.eu-west-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoib3cxVnoyZUtBZU5wN1VUYUtKQTJUVmtmMHBJcUJXSUFWMXEyc2d3WWJldUdPTHh4QWQ1eFNRendpOUJHVnZ5UXBpMXpFVkVSb3k2UUhKL2xCY2JhVnhJPSIsIml2UGFyYW1ldGVyU3BlYyI6Im9QZ2dPS3ozdWFyWHIvbm8iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](DOWNLOAD.md)
[![CircleCI](https://circleci.com/gh/lmorg/murex/tree/master.svg?style=svg)](https://circleci.com/gh/lmorg/murex/tree/master)
[![codecov](https://codecov.io/gh/lmorg/murex/branch/master/graph/badge.svg)](https://codecov.io/gh/lmorg/murex)
[![Tests](testcases.svg)](https://github.com/lmorg/murex)

## About _murex_

_murex_ is a shell, like bash / zsh / fish / etc. It follows a similar syntax
to POSIX shells like Bash however supports more advanced features than you'd
typically expect from a $SHELL.

It aims to be similar enough to traditional shells that you can retain most of
your muscle memory, while not being afraid to make breaking changes where
"bash-isms" lead to unreadable, hard to maintain, or unsafe code.

_murex_ is designed for DevOps productivity so it isn't suited for high
performance workloads beyond what you'd typically run in Bash (eg pipelines
forked as concurrent processes).

A non-exhaustive list features would include:

* Support for typed pipelines - which can be used to work with complex data
  formats like JSON natively. But also the ability to override or even ignore
  typed data entirely so it works transparently with standard UNIX tools too.
  This means you can use a common set of commands for manipulating any form of
  data file.

* Usability improvements such as a smarter `readline` API, in-line spell
  checking, hint text detailing a commands behavior before you hit return, and auto-parsing man pages for auto-completions on commands that don't have any
  auto-completion config set.
  
* Smarter handling of errors and debugging tools. For example try/catch blocks,
  line numbers included in error messages, errors optionally highlighted in
  red, and script testing and debugging frameworks baked right into the
  language itself.

## Type system

_murex_ supports multiple data types natively; such as JSON, YAML, CSV, 
S-Expressions and even loosely tabulated terminal output (eg `ps`, `ls -l`).
This makes passing data through the pipeline and parsing output easier when
dealing with more complex arrangements of data.

For example a traditional pipeline might look like the following:

    curl -s https://api.github.com/repos/lmorg/murex/issues | jq -r  '.[] | [(.number|tostring), .title] | join(": ")'

Because traditional shells send everything as dumb byte streams, working with
any structured data means learning a multitude of additional languages like
awk, sed, tr, Perl, jq and so on and so forth.

The same pipeline in _murex_ might look like the following:

    open https://api.github.com/repos/lmorg/murex/issues -> foreach issue { out "$issue[number]: $issue[title]" }

The aims of _murex_ is to create a single consistant language that can work for
any data structure, be readable but also as terse and quick to write as Bash.

A big part of that ambition is realized via the interactive shell.

## Interactive shell

Aside from _murex_ being carefully designed with scripting in mind, the
interactive shell itself is also built around productivity. To achieve this
we wrote our own readline library. Below is an example of that library in use:

[![asciicast](https://asciinema.org/a/232714.svg)](https://asciinema.org/a/232714)

See [the interactive shell user guide](/docs/user-guide/interactive-shell.md)
for details on all tricks supported by _murex_'s interactive terminal.

## Pipe tokens: `->` vs `|`

_murex_ supports multiple different pipe tokens. The main two being `|` and
`->`.

* `|` works exactly the same as in any normal shell

* `->` displays all of the supported methods (commands that support the output
  of the previous command). Think of it a little like object orientated
  programming where an object will have functions (methods) attached.

In _murex_ scripts you can use `|` and `->` interchangeably, so there's no need
to remember which commands are methods and which are not. The difference only
applies in the interactive shell where `->` can be used with tab-autocompletion
to display a shortlist of supported functions that can manipulate the data from
the previous command. It's purely a clue to the parser to generate different
autocompletion suggestions to help with your discovery of different commandline
tools.

You can [read more about the _murex_ parser](/docs/GUIDE.parser.md) and the
different supported tokens in the [docs](https://murex.rocks/).

## Concise yet predictable

Despite the amount of features added to shell, we have tried to keep the
amount of "magic" to a minimum and follow a pretty standard structure so
the language is predictable. However there are times when a little magic
goes a long way. For example _murex_'s support for complex data objects
of differing formats is managed in the pipeline so you don't need to think
about the data format when querying data from them.

    open: file.csv  -> [ column_name ] # returns specific columns (or rows) in CSV file
    open: file.json -> [ index ]       # returns specific items from JSON

The index function (`[`) alters its matching algorithm depending on the
piped data type and `open` sets the data type depending on the file
extension or MIME type.

Sometimes you will want less guesswork or just the robustness of a forced
behavior. On those occasions you can remove one layer of magic by
casting the data type:

    open: file.txt -> cast csv  -> [ column_name ]
    open: file.txt -> cast json -> [ index ]

This awareness of data structures is also utilised in `foreach` (which
will cycle through each index in an array) and `formap` (key/value
iteration against complex objects). See [GUIDE.control-structures](docs/GUIDE.control-structures.md)
for more details on these and other control structures.

## More robust scripts / shell one liners

_murex_ employs a few methods to make shell scripting more robust:

Bash, for all it's power, is littered with hidden traps. The aim of _murex_ is
to address as many of them as we can without taking the flexibility or power
away from the interactive command line. This is achieved through a couple of
key concepts:

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

(more information on `g`, `rx` and `f` are available in [GUIDE.quick-start](docs/GUIDE.quick-start.md)).

However there will be occasions when you just want an inlined expansion
(eg when using an interactive shell) and that can be achieved via the `@g`
command prefix:

    @g ls -l *.go

### Powerful autocompletion

_murex_ takes a slightly different approach to command line autocompletion,
both from a usability perspective as well as defining completion rules.

Inspired by IDEs, _murex_ queries man pages directly for flags as well as
"tooltip" descriptions. Custom completions are defined via JSON meaning
simple commands are much easier to define and complex commands can still
fallback to using dynamic shell code just like they are in other shells.

This makes it easier to write completion rules as well as making the code
more readable. An example of `git`s autocompletion definition:

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
"popups" to cater for different scenarios (demo above) as well as built in
support for jumping to files within nested directories quickly and easily:

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

    # compare two strings
    if { = `foo`==`bar` } then {
        out: "`foo` matched `bar`"
    }

    # check if command ran successfully
    !if { foobar } then {
        err: "`foobar` could not be run"
    }

### Test and debugging frameworks

Unlike traditional shells, _murex_ is designed with a test and debugging modes
baked into the shell language. This means you can write tests against your
shell scripts as part of the shell scripts themselves.

For example:

    function: hello-world {
        test: define example {
            "StdoutRegex": (^Hello World$)
        }

        out: <test_example> "Hello Earth"
    }

    test: run { hello-world }

...will output:

    Hello Earth
     Status  Definition Function                                           Line Col. Message
    [FAILED] example    out                                                5    9    stdout: regexp did not match 'Hello Earth'

If test mode isn't enabled then any `test` commands are skipped without being
executed so you can liberally include test cases throughout your functions
without worrying about any performance impact.

#### _murex_ also supports unit tests

For example:

    test: unit function aliases {
        "PreBlock": ({
            alias ALIAS_UNIT_TEST=example param1 param2 param3
        }),
        "StdoutRegex": "([- _0-9a-zA-Z]+ => .*?\n)+",
        "StdoutType": "str",
        "PostBlock": ({
            !alias ALIAS_UNIT_TEST
        })
    }

    function: aliases {
        # Output the aliases in human readable format
        runtime: --aliases -> formap: name alias {
            $name -> sprintf: "%10s => ${esccli @alias}\n"
        } -> cast: str
    }

    test: run aliases

...will output:

     Status  Definition Function                                           Line Col. Message
    [PASSED] (unit)     aliases                                            13   1    All test conditions were met

## Language guides

1. [GUIDE.syntax](docs/GUIDE.syntax.md) is recommended first
   as it gives an overview if the shell scripting languages syntax and data
   types.

2. [GUIDE.type-system](docs/GUIDE.type-system.md) describes
   _murex_'s type system. Most of the time you will not need to worry about
   typing in _murex_ as the shell is designed around productivity.

3. [GUIDE.builtin-functions](docs/GUIDE.builtin-functions.md)
   lists some of the builtin functions available for this shell.

Or if you're already a seasoned Bash developer then you read the Quick
Start Guide, [GUIDE.quick-start](docs/GUIDE.quick-start.md),
to jump straight into using _murex_.

## Install instructions

There are various ways you can load _murex_ on to your system. See [INSTALL](INSTALL.md) for details.

## Known bugs / TODO

Please see GitHub's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)
