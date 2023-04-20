# Murex Named Pipes - User Guide

> A detailed breakdown of named pipes in _murex_

## Background

[Wikipedia describes](https://en.wikipedia.org/wiki/Named_pipe) a named pipe as the following:

> In computing, a named pipe (also known as a FIFO for its behavior) is an
> extension to the traditional pipe concept on Unix and Unix-like systems, and
> is one of the methods of inter-process communication (IPC). The concept is
> also found in OS/2 and Microsoft Windows, although the semantics differ
> substantially. A traditional pipe is "unnamed" and lasts only as long as the
> process. A named pipe, however, can last as long as the system is up, beyond
> the life of the process. It can be deleted if no longer used. Usually a named
> pipe appears as a file, and generally processes attach to it for IPC.

Where _murex_ differs from standard Linux/UNIX is that named pipes are not
special files but rather an object or construct within the shell runtime. This
allows for more user friendly tooling and syntactic sugar to implemented around
it while largely still having the same functionality as a more traditional file
based named pipe.

## In _murex_

In _murex_, named pipes are described in code as a value inside angle brackets.
There are four named pipes pre-configured: `<in>` (STDIN), `<out>` (STDOUT),
`<err>` (STDERR), and `<null>` (/dev/null equivalent).

You can call a named pipe as either a method, function, or parameter.

**As a method:**

    <in> -> command parameter1 parameter2 parameter3
    
**As a function:

    command parameter1 parameter2 parameter3 -> <out>
    
**As a parameter:

    command <out> <!err> parameter1 parameter2 parameter3

## See Also

* [`<>` / `read-named-pipe`](../commands/namedpipe.md):
  Reads from a _murex_ named pipe
* [`<stdin>` ](../commands/stdin.md):
  Read the STDIN belonging to the parent code block
* [`test`](../commands/test.md):
  _murex_'s test framework - define tests, run tests and debug shell scripts