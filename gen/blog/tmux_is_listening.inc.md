## What Is Tmux

(You can skip this section if you're already know familiar with tmux)

[Tmux](https://github.com/tmux/tmux) is a _terminal multiplexer_. Which probably means nothing to a lot of people.

[Wikipedia](https://en.wikipedia.org/wiki/Terminal_multiplexer) describes terminal multiplexers as:

> A terminal multiplexer is a software application that can be used to multiplex several separate pseudoterminal-based login sessions inside a single terminal display, terminal emulator window, PC/workstation system console, or remote login session, or to detach and reattach sessions from a terminal. It is useful for dealing with multiple programs from a command line interface, and for separating programs from the session of the Unix shell that started the program, particularly so a remote process continues running even when the user is disconnected.

...which is equally confounding. But thankfully it's [description of `tmux`](https://en.wikipedia.org/wiki/Terminal_multiplexer) is less like death from jargon-soup:

> tmux is an open-source terminal multiplexer for Unix-like operating systems. It allows multiple terminal sessions to be accessed simultaneously in a single window. It is useful for running more than one command-line program at the same time. It can also be used to detach processes from their controlling terminals, allowing remote sessions to remain active without being visible.

So basically, it allows you to do more and concurrently inside a single terminal interface.

## What Is This Vulnerability?

For any running tmux sessions, you can attach a listener which records the terminal session, for all terminals and applications running inside tmux.

### That Sounds Scary!

It does, but it isn't.

An attacker would need you to run their executable on your local machine. And if they can already do that then it is already game over for you.

### So Why Report This?

Because it surprised 