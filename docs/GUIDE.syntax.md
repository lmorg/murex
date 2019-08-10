# Language Guide: Syntax

## Structure

The following examples all produce the same output but demonstrate the
way the language is structured.

standard `if` / `else` block:

    if { echo: foo\nbar -> match: bar } then {
        echo: "bar found"
    } else {
        echo: "no bar found"
    }


`if` evaluation pulled from stdin:

    echo: foo\nbar -> match: bar -> if {
        echo: "bar found"
    }

`foreach` line check value:

    echo: foo\nbar -> foreach: line {
        if { echo: $line } then {
            echo: "bar found"
        } else {
            echo: "no bar found"
        }
    }

However since the language is optimised for one liners, the idiomatic
way of writing this code would be:

    out: foo\nbar -> match: bar -> if { out: "bar found" }

The syntax is designed to be fast for writing logic while also readable.
Which positions it between Bash and many other scripting languages (for
example Perl and Python).

It's easier to write more complex logic than one would normally do in a
shell scripting language (plus has a fewer hidden traps than traditional
shells) and yet you still have the ability to quickly draft up Bash-
style one liners in a way that many more advanced scripting languages
normally struggle.

### Processes

Processes are called by name then parameters. The process name can be
separated by either a colon (:) and/or a white space character (\s, \t,
\r, or \n). eg `out: "hello world"`.

Process names can be quoted (eg `"out": "hello world"`) however this is
not necessary unless you are calling an external executable with white
space characters or a colon. Additionally characters can be escaped in
the process name (eg: `o\ut: "hello world"`).

### Piping

To pass streams into processes you use pipes like you would in your standard
shell. However _murex_ pipes are typically written as arrows, `->`, because
that helps illustrate the direction of the stream and thus improves readability.
However for compatibility reasons the standard pipe character is also supported

    # Idiomatic murex:
    out: "hello world" -> grep: "world"
    
    # POSIX (Bourne, Bash, Zsh, etc) shells and also supported in murex:
    out "hello world" | grep "world"

So while the idiomatic way to write _murex_ code would be using arrow pipes,
there isn't any danger in using the traditional pipe character since those two
tokens are interchangeable.

#### Redirection

Another important difference in piping is the way redirection is handled. In
_murex_ you define redirection as the first parameter(s). For example:

    err: <!out> "error message redirected to stdout"
    out: <err> "output redirected to stderr"

The redirection must be the first parameter surrounded by less than / greater
than tokens and can only be alpha and numeric characters. The prefixed bang
(exclamation mark) denotes whether you are redirecting the stdout or stderr
stream (stderr contains the `!`).

The advantage of this method is that you can create more meaningful named pipes
for tunneling information between concurrent processes that might not otherwise
sit on the same code pipeline.

    pipe: foobar

    # Create background process that outputs from the named pipe,
    # then pipes that into a regexp matcher.
    # <foobar> will stay running until it is closed.

    bg {
        <foobar> -> regexp: m/00$/
        out: "pipe closed, exiting `bg`"
    }

    # Lets send some data to our named pipe, then close it.

    a: <foobar> [1..1000]
    !pipe: foobar

(**PLEASE NOTE** that _murex_ named pipes are not file system FIFO objects)

#### `null` pipe

There is a builtin `null` device for forwarding output into a black hole. This
`null` device is analogous to `/dev/null` on POSIX systems:

    try <!null> {
        err "raise an error to fail `try`."
    } -> catch {
        out "An error was raised but the message was dumped into `null`."
    }

The `null` device doesn't need to be created.

#### File writer pipe

You can also use named pipes for writing files:

    pipe: log --file error.log
    try <!log> {
        err "Do something bad."
    } -> catch {
        out "An error was raised. See error.log for details."
    }

**PLEASE NOTE** that the <pipe> parameter cannot be populated by variables.
This is a security design to protect against a $variable containing `<>` and
causing unexpected behaviour. Alternatively you can use <pipe> as a function:

    pipe: foobar
    bg { <foobar> -> cat }
    out: "writing to foobar..." -> <foobar>
    !pipe: foobar

(this example includes a redundant usage of `cat` to demonstrate named-pipes writing to a function)

#### Networking pipes

There are 4 networking pipes:

1. `tcp-dial`: makes an outbound TCP connection
2. `udp-dial`: makes an outbound UDP connection
3. `tcp-listen`: listens for an incoming TCP connection request
4. `udp-listen`: listens for an incoming UDP connection request

These are used in the same way as the other named-pipes described above.

    pipe: google --tcp-dial google.com:80
    out: <google> "GET /"
    <google>
    !pipe: google

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

## Bang Prefix

Some builtins support an optional bang (!) prefix. These are similar to the
bang operator in some languages where bang will perform a boolean `not`.
However in _murex_, builtins prefixed with a bang will perform the opposite functional task rather than logical operation.

For example `out: hello world -> !match: world` would return no results as the
`!match` builtin will look for strings that **don't**  match "world" (rather
than the typical operation of `match` looking for strings that **do** match).
Some encoding / compression routines also support a bang prefix to decode or deflate their input.

## End of chain semi-colons

Like on Linux/UNIX and cmd.exe, you can terminate a pipeline chain with
a semi-colon (;).

## To colon or not to colon?

It might seem weird to support a colon as a delimiter separating the function
name and parameters but making it optional. The rational is that longer and
more complex scripts look more readable with a colon since it clearly delimits
what is a command from it's parameters. However this piece of punctuation
remains optional to retain support for "Bash muscle memory".

The idiomatic style guide would be that any commands which have parameters that
follow should contain a colon. eg:

    out: "hello world" -> match: "world"

However if the parameter if a code block or JSON, then you drop the colon since
the delimitation is already clear:

    if { = 1==2 } then { out: "The laws of the universe are broken" }

Commands without a parameter don't need a colon and some of _murex_'s primitive
builtins (`[`, `[[`, `(`, `=`) also look cleaner without a colon.

However these are only recommendations for writing readable code rather than
enforced standards.

## Code golfing

Code "golfing" is the process of writing source code in the fewest number of
characters possible (much like the sport of Golf). While code golfing isn't
something sane people would advocate for normal programming practices (ie
anything that requires any level of maintainability), sometimes when in the
interactive command line you might wish to minify your code when typing a
one-liner for the sake of convenience.

The syntax of _murex_ is designed to be flexible enough to write readable and
maintainable multi-line scripts but also to be terse enough to write "golfed"
one-liners.

For example, one of the earlier code examples could be written like this:

    out:foo\nbar|match:bar|if{(bar found)}

It is suggested that you don't write reusable code in this style and even the
autocompletion tools built into _murex_'s interactive command line will guide
you towards writing readable code (albeit the autocompletion tools are built
around intuitive and non-intrusive suggestions so are easy to ignore / override
should you wish to)
