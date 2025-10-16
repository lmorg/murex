# Interactive Shell

> What's different about Murex's interactive shell?

Murex's interactive shell is also built around productivity. To achieve this we
wrote our own state-of-the-art readline library.

Below are just some of the features you can enjoy.

<h2>Table of Contents</h2>

<div id="toc">

- [Advanced Autocompletion](#advanced-autocompletion)
- [Hint Text](#hint-text)
- [Spellchecker](#spellchecker)
- [Preview Autocompletions](#preview-autocompletions)
- [Preview Command Lines](#preview-command-lines)
- [Safer Pasting](#safer-pasting)
- [Smarter Error Messages](#smarter-error-messages)
- [Hotkeys](#hotkeys)

</div>



## Advanced Autocompletion

Autocompletion happen when you press `[tab]` and will differ slightly depending
on what is defined in `autocomplete` and whether you use the traditional
[POSIX pipe token](../parser/pipe-posix.md), `|`, or the [arrow pipe](../parser/pipe-arrow.md),
`->`.

The `|` token will behave much like any other shell however `->` will offer
suggestions with matching data types. Which makes working working with data
quick and easy while still intelligent and readable.

![autocomplete preview](/images/vhs-autocompletion-dark.gif)

## Hint Text

The **hint text** is a (typically) blue status line that appears directly below
your prompt. The idea behind the **hint text** is to provide clues to you as
type instructions into the prompt; but without adding distractions. It is there
to be used if you want it while keeping out of the way when you don't want it.

![hint-text](/images/vhs-hint-text-dark.gif)

[Read more about Hint Text](/docs/user-guide/hint-text.md).

## Spellchecker

Murex supports inline spellchecking, where errors are underlined. For example

[![asciicast](https://asciinema.org/a/408024.svg)](https://asciinema.org/a/408024)

This might require some manual steps to enable, please see the [spellcheck user guide](spellcheck.md)
for more details.



## Preview Autocompletions

> Enabled via `[f1]`

This displays a more detailed view of each parameter you're about to pass to a
command, without you having to run that command nor leave the half-completed
command line.

![autocomplete preview](/images/vhs-preview-autocomplete-dark.gif)

It can display:
* [`man` pages](/docs/integrations/man-pages.md)
* custom guides like [https://cheat.sh](/docs/integrations/cheatsh.md) and [AI generated docs](/docs/integrations/chatgpt.md)
* information about binary files
* contents of text files
* and even images too!


## Preview Command Lines

> Enabled via `[f9]`

The Command Line Preview allows you to view the output of a command line while
you're still writing it. This interactivity removes the trial-and-error from
working with complicated command line incantations. For example parsing parsing
complex documents like machine generated JSON becomes very easy.

![autocomplete preview](/images/vhs-preview-commandline-dark.gif)

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

![autocomplete preview](/images/vhs-better-errors-errmsg-dark.png)

## Hotkeys

A full breakdown of supported hotkeys is available in the [terminal-keys](terminal-keys.md)
guide.

## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Code Block Parsing](../user-guide/code-block.md):
  Overview of how code blocks are parsed
* [Define Method Relationships (`method`)](../commands/method.md):
  Define a methods supported data-types
* [Hint Text](../user-guide/hint-text.md):
  A status bar for your shell
* [Shell Configuration And Settings: `config`](../commands/config.md):
  Query or define Murex runtime settings
* [Shell Runtime: `runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Spellcheck](../integrations/spellcheck.md):
  How to enable inline spellchecking
* [Tab Autocompletion: `autocomplete`](../commands/autocomplete.md):
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