# Interactive Shell

> What's different about Murex's interactive shell?

<h2>Table of Contents</h2>

<div id="toc">

- [Overview](#overview)
- [readline](#readline)
- [Hotkeys](#hotkeys)
- [Autocompletion](#autocompletion)
- [Syntax Completion](#syntax-completion)
- [Syntax Highlighting](#syntax-highlighting)
- [Spellchecker](#spellchecker)
- [Hint Text](#hint-text)
- [Preview](#preview)
  - [Autocomplete Preview](#autocomplete-preview)
  - [Command Line Preview](#command-line-preview)
- [Safer Pasting](#safer-pasting)
- [Smarter Error Messages](#smarter-error-messages)

</div>



## Overview

Aside from Murex being carefully designed with scripting in mind, the
interactive shell itself is also built around productivity. To achieve this
we wrote our own readline library. Below is an example of that library in use:

[![asciicast](https://asciinema.org/a/232714.svg)](https://asciinema.org/a/232714)

The above demo includes the following features of Murex's bespoke readline
library:

* hint text - blue status text below the prompt (the colour is configurable)
* syntax highlighting (albeit there isn’t much syntax to highlight in the
  example). This can also be turned off if your preference is to have colours
  disabled
* tab-completion in gridded mode (seen when typing `cd`)
* tab-completion in list view (seen when selecting a process name to `kill`
  where the process ID was substituted when selected)
* searching through the tab-completion suggestions (seen in both `cd` and
  `kill` - enabled by pressing `[ctrl]`+`[f]`)
* line editing using $EDITOR (`vi` in the example - enabled by pressing `[esc]`
  followed by `[v]`)
* readline’s warning before pasting multiple lines of data into the buffer and
  the preview option that’s available as part of the aforementioned warning
* and VIM keys (enabled by pressing `[esc]`)

## readline

Murex uses a custom `readline` library to enable support for new features in
addition to the existing uses you'd normally expect from a shell. It is because
of this, Murex provides one of the best user experiences of any of the shells
available today.

## Hotkeys



A full breakdown of supported hotkeys is available in the [terminal-keys](terminal-keys.md)
guide.

## Autocompletion

Autocompletion happen when you press `[tab]` and will differ slightly depending
on what is defined in `autocomplete` and whether you use the traditional
[POSIX pipe token](../parser/pipe-posix.md), `|`, or the [arrow pipe](../parser/pipe-arrow.md),
`->`.

The `|` token will behave much like any other shell however `->` will offer
suggestions with matching data types (as seen in `runtime --methods`). This is
a way of helping highlight commands that naturally follow after another in a
pipeline. Which is particularly important in Murex as it introduces data
types and dozens of new builtins specifically for working with data structures
in an intelligent and readable yet succinct way.

You can add your own commands and functions to Murex as methods by defining
them with `method`. For example if we were to add `jq` as a method:

```
method define jq {
    "Stdin":  "json",
    "Stdout": "@Any"
}
```



## Syntax Completion

Like with most IDEs, Murex will auto close brackets et al.

[![asciicast](https://asciinema.org/a/408029.svg)](https://asciinema.org/a/408029)

## Syntax Highlighting

Pipelines in the interactive terminal are syntax highlighted. This is similar
to what one expects from an IDE.

Syntax highlighting can be disabled by running:

```
config set shell syntax-highlighting off
```

## Spellchecker

Murex supports inline spellchecking, where errors are underlined. For example

[![asciicast](https://asciinema.org/a/408024.svg)](https://asciinema.org/a/408024)

This might require some manual steps to enable, please see the [spellcheck user guide](spellcheck.md)
for more details.



## Hint Text

The **hint text** is a (typically) blue status line that appears directly below
your prompt. The idea behind the **hint text** is to provide clues to you as
type instructions into the prompt; but without adding distractions. It is there
to be used if you want it while keeping out of the way when you don't want it.




[Read more about Hint Text](/docs/user-guide/hint-text.md).

## Preview

Murex supports a couple of full screen preview modes:

* Autocomplete Preview ([read more](#autocomplete-preview))
* Command Line Preview ([read more](#command-line-preview))

### Autocomplete Preview

> Enabled via `[f1]`

This displays a more detailed view of each parameter you're about to pass to a
command, without you having to run that command nor leave the half-completed
command line.

It can display:
* [`man` pages](/docs/integrations/man-pages.md)
* custom guides like [https://cheat.sh](/docs/integrations/cheatsh.md) and [AI generated docs](/docs/integrations/chatgpt.md)
* information about binary files
* contents of text files
* and even images too!



### Command Line Preview

> Enabled via `[f9]`

The Command Line Preview allows you to view the output of a command line while
you're still writing it. This interactivity removes the trial-and-error from
working with complicated command line incantations. For example parsing parsing
complex documents like machine generated JSON becomes very easy.



This does come with some risks because most command line operations change you
systems state. However Murex comes with some guardrails here too:

* Each command in the pipeline is cached. So if a command's parameters are
  changed, Murex only needs to re-run the commands _from_ the changed
  parameter onwards.

* Each time there is a change in the commands themselves, for example a new
  command added to the pipeline, you are requested to press `[f9]` to re-run
  the entire pipeline.

* The only commands considered "safe" for auto-execution if any parameters do
  change are those marked as "safe" in `config`. For example:
  ```
  » config get shell safe-commands -> tail -n5
  td
  cut
  jobs
  select
  dig
  ```

## Safer Pasting

A common behaviour for command line users is to copy and paste data into the
terminal emulator. Some shells like Zsh support [Bracketed paste](https://en.wikipedia.org/wiki/Bracketed-paste)
but that does a pretty poor job of protecting you against the human error of
pasting potentially dangerous contents from an invisible clipboard.

Where Murex differs is that any multi-line text pasted will instantly display
a warning prompt with one of the options being to view the contents that you're
about to execute.



This gives you piece-of-mind that you are executing the right clipboard content
rather than something else you copied hours ago and forgotten about.

## Smarter Error Messages

Errors messages in most shells suck. That's why Murex has taken extra care to
give you as much useful detail as it can.


## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Code Block Parsing](../user-guide/code-block.md):
  Overview of how code blocks are parsed
* [Define Method Relationships (`method`)](../commands/method.md):
  Define a methods supported data-types
* [Hint Text](../user-guide/hint-text.md):
  A status bar for your shell
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Spellcheck](../integrations/spellcheck.md):
  How to enable inline spellchecking
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation
* [`{ Curly Brace }`](../parser/curly-brace.md):
  Initiates or terminates a code block
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [gen/user-guide/interactive-shell_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/interactive-shell_doc.yaml).