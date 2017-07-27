# FAQ: Why do we need another shell?

Simple answer is we don't.

The longer answer is I wanted to try something different. _murex_ takes
a slightly different approach to conventional shells and to any of the
modern re-imagined shells I've tried.

Most shell rewrites I've found fall into one or more of the following
categories:

* A focus on Bash compatibility without trying to bring anything new to
the table (for example better error handling)

* A focus on shell scripting, neglecting basic requirements required for
a functional `$SHELL` (for example spawning new processes in a pseudo-TTY)

* Or a focus on implementing new ideas without taking into account
muscle-memory from the millions of existing Bash et al users.

