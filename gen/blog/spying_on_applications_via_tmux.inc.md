## So What Is Tmux

(You can skip this section if you're already know familiar with tmux)

[Tmux](https://github.com/tmux/tmux) is a _terminal multiplexer_. Which probably means nothing to a lot of people.

A _terminal multiplexer_ is basically just an application that sits between the backend of your terminal (the "TTY") and the frontend (your _terminal emulator_, ie iTerm2, Kitty, PuTTY, Ttyphoon, etc). It allows you to then spawn multiple new terminals within your terminal, a little like how you can open multiple windows within a single graphical display but, in the context of _terminal multiplexers_, we are talking about multiple different terminals within the same terminal.

So basically, it allows you to do more, and concurrently, inside a single terminal interface.

## What Is This Vulnerability?

"Vulnerability" is actually the wrong word here. It's not a memory safety bug that would be fixed by rewriting the software in Rust. In fact it's not a bug at all. This is intended behaviour introduced to facilitate tighter integration between iTerm2 and tmux; and later added to other terminal emulators like Ttyphoon.

This feature allows any application the ability to listen to _every_ interaction in tmux and do so without any request nor notification sent to the user.

This means 

For any running tmux sessions, you can attach a listener which records the terminal session, for all terminals and applications running inside tmux.

### That Sounds Scary!

It does, but it isn't.

An attacker would need you to run their executable on your local machine. And if they can already do that then it is already game over for you.

### So Why Report This?

Because it surprised 