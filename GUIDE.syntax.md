# Language Guide: Syntax

## Processes

There are 3 types of processes:

1. Functions
2. Methods
3. External

Like with all programming languages, methods are functions attached to a
particular object.

### Functions

These are processes that only accept no originating stream data via STDIN

### Methods

These are processes that accept stream data via STDIN.

### External

These are external processes. They can be called in exactly the same way
as functions or via `exec`

## Types

Because this is a shell language optimised for stream processing,
all types are compiled down to strings. However there are a few types
defined for special purposes:

* Generic   (defined: *)
* Null      (defined: null)
* Die       (defined: die)
* Binary    (defined: bin) // Not yet implemented
* String    (defined: str)
* Boolean   (defined: bool)
* Integer   (defined: int) // Implemented by currently unused
* Float     (defined: float) // Implemented by currently unused
* Code Block (defined: block)
* JSON      (defined: json)
* XML       (defined: xml) // Not yet implemented

### Generic

This is used by functions and methods to state they can accept any data
type. Both methods and functions can accept `generic` input.

### Null

This states that no data expected. Use `null` input to define functions
and/or `null` output to state the process doesn't write to STDOUT.

### Die

If a `die` object is created it kills the shell.

### Binary

String output that's not expected to be analyzed by text manipulation
tools.

### Boolean

True or False. Generic input can be translated to boolean:

* 0 == False, none zero numbers == True
* "false" == False, "true" == True
* "no" == False, "yes" == True
* "off" == False, "on" == True
* "fail" == False, "pass" == True
* "failed" == False, "passed" == True
* "" == False, all other strings == True

Strings are not case sensitive when converted to boolean.

### Integer

As with normal programming languages, any number that doesn't have a
decimal point.

### Float

A number which does have a decimal point.

### Code Block

A sub-shell. This is used to inline code. eg for loops. Blocks are
encapsulated by curly brackets, `{}`.

### JSON

A JSON object.

### XML

An XML object.

## Structure

The following examples all produce the same output but demonstrate the
way the language is structured.

```
if: { echo: foo\nbar -> match: bar } {
    echo: "bar found"
} {
    echo: "no bar found"
}
```

```
echo: foo\nbar -> match: bar -> if {
    echo: "bar found"
}
```

```
echo: foo\nbar -> foreach: line {
    if: {echo: $line } {
        echo: "bar found"
    } {
        echo: "no bar found"
    }
}
```

However since the language is optimised for one liners, the idiomatic
way of writing this code would be:

```
out foo\nbar -> match bar -> if { out "bar found" }
```

The syntax is designed to be fast for writing logic while also readable.
Which positions it between Bash and many other scripting languages (for
example Perl and Python).

It's easier to write more complex logic than one would normally do in a
shell scripting language (plus has a fewer hidden traps than traditional
shells) and yet you still have the ability to quickly draft up Bash-
style one liners in a way that many more advanced scripting languages
normally struggle. It also supports forking processes in PTYs while many
scripting REPLs do not.

### Functions

Functions are called by name then parameters. The function name can be
separated by either a colon (:) and/or a white space character (\s, \t, \r,
or \n). eg `out: "hello world"`.

Function name can be quoted (eg `"out": "hello world"`) however this is
not necessary unless you are calling an external executable with white
space characters or a colon. Additionally characters can be escaped in
the function name (eg: `o\ut: "hello world"`).

### Methods

Methods are prefix by an ASCII arrow (->). eg `->match("world")`.

Since methods require data input from the pipeline, you would use a
method like so:
```
out: "hello world" -> match: "world"
```

### Piping to external processes

To pass streams into external processes you can use pipes like you would
in your standard shell. eg `out: "hello world" | grep: "world"`.

However where this shell differs from the likes of Bash is STDOUT and
STDERR redirection. With this shell there are two pipe operators:

* | == pipe STDOUT to next processes STDIN and STDERR to parents STDERR
* ? == pipe STDERR to next processes STDIN and STDOUT to parents STDERR
(essentially swapping the STDOUT and STDERR streams in that process)

An example of error redirection:
```
exec: sh -c "echo hello world 1>&2" ? grep: "world"
```

Currently there is no method to merge STDOUT and STDERR streams.

### Parameters

As you have probably guessed from the above examples, parameters are
space delimited (much like with Bash) and support single and double
quotations and escaping.

The quotes and escaping works in the same was as they do for function
names:

#### No quotes

Parameters are space delimited.

`out: a b c 1 2 3;` translates to 6 parameters: "a", "b", "c", "1", "2",
"3".

#### Single quotes and double quotes

This auto-escapes all characters expect the escape character (\\) and
variables (eg `$variableName`)

`out: 'a b c' "1 2 3";` translates to 2 parameters: "a b c"s, "1 2 3".

## Anti-aliases: bang prefix

Some functions support an optional bang (!) prefix. These are called
"anti-aliases" and similarly to how a bang can `not` a boolean state,
anti-aliases are aliases that perform opposite functions to their
default behaviour. For example `out: hello world -> !match: world` would
return no results as the `!match` anti-alias with look for strings that
don't match "world". Some encoding / compression routines also have an
anti-alias to decode or deflate their input.

## Anonymous pipes

As already discussed earlier in this document, there are 3 types of
anonymous pipes:

1. `|`: This works exactly the same as in Linux/UNIX and cmd.exe. It
pipes STDOUT to the STDIN of an external process (eg grep). 

2. `?`: This works similarly to the pipe (|) character except it pipes
STDERR to the STDIN if ab external process.

3. `->`: This denotes the next process is a method, then pipes STDOUT to
the STDIN of that method.

## End of chain semi-colons

Like on Linux/UNIX and cmd.exe, you can terminate a pipeline chain with
a semi-colon (;).

## To colon or not to colon?

It might seem weird to support a colon as a delimiter separating the
function name and parameters but making it optional. The rational is
that I think longer and more complex scripts look more readable with a
colon (it also looks "prettier" in my personal opinion). However I
wanted to retain support for "Bash muscle memory" as it is pretty
annoying typing a one-liner then having that failing because of a
missing colon.

If I was to impose my own style guidelines then I would argue that the
idiomatic way to write code would be with a colon. eg:
```
out: "hello world" -> match: "world"
```
but support for dropping the colon is allowed to ease the learning
curve of using this new shell. eg:
```
echo hello world | grep world
```
(which would work both in this shell and in Bash)

## Back ticks

Back ticks have no special function and thus are treated like a regular,
printable, character.

## Code golfing

For those who are unaware, 'code golfing' is the process of writing a
piece of code in the fewest number of characters possible (much like the
sport of Golf). While code golfing isn't something sane people would
advocate for normal programming practices (ie anything that requires any
level of maintainability), sometimes command line users would "golf"
their code when typing a one-liner in their interactive shell purely out
of convenience / laziness.

The syntax of this shell is designed to be flexible enough to write
readable and maintainable multi-line scripts but also to be terse enough
to write "golfed" one liners.

One of the earlier code examples could be written like this:
```
out:foo\nbar->match:bar->if{out:bar found}
```
