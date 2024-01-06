{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>

<div id="toc">

- [Overview](#overview)
- [readline](#readline)
- [Hotkeys](#hotkeys)
- [Autocompletion](#autocompletion)
- [Syntax Completion](#syntax-completion)
- [Syntax Highlighting](#syntax-highlighting)
- [Spellchecker](#spellchecker)
- [Hint Text](#hint-text)
  - [Configuring Hint Text Colour](#configuring-hint-text-colour)
  - [Custom Hint Text Statuses](#custom-hint-text-statuses)
  - [Starship Example](#starship-example)
  - [Disabling Hint Text](#disabling-hint-text)
- [Preview](#preview)
  - [Autocomplete Preview](#autocomplete-preview)
  - [Command Line Preview](#command-line-preview)
- [Safer Pasting](#safer-pasting)
- [Smarter Error Messages](#smarter-error-messages)

</div>

{{ end }}

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

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<a href="terminal-keys.html" alt="supported hotkeys"><img src="/keyboard.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/></a>
<!-- markdownlint-restore -->
{{ end }}

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

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<br/>
<img src="/screenshot-ps-select.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}

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

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img src="/screenshot-spellchecker.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}

## Hint Text

The **hint text** is a (typically) blue status line that appears directly below
your prompt. The idea behind the **hint text** is to provide clues to you as
type instructions into the prompt; but without adding distractions. It is there
to be used if you want it while keeping out of the way when you don't want it.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img src="/screenshot-hint-text-rsync.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}

### Configuring Hint Text Colour

By default the **hint text** will appear blue. This is also customizable:

```
» config get shell hint-text-formatting
{BLUE}
```

The formatting config takes a string and supports [ANSI constants](ansi.md).

It is also worth noting that if colour is disabled then the **hint text** will
not be coloured even if **hint-text-formatting** includes colour codes:

```
» config set shell color false
```

(please note that **syntax highlighting** is unaffected by the above config)

### Custom Hint Text Statuses

There is a lot of behavior hardcoded into Murex like displaying the full path
to executables and the values of variables. However if there is no status to be
displayed then Murex can fallback to a default **hint text** status. This
default is a user defined function. At time of writing this document the author
has the following function defined:

```
config set shell hint-text-func {
    trypipe <!null> {
        git status --porcelain -b -> set gitstatus
        $gitstatus -> head -n1 -> regexp 's/^## //' -> regexp 's/\.\.\./ => /'
    }
    catch {
        out "Not a git repository."
    }
}
```

...which produces a colorized status that looks something like the following:

```
develop => origin/develop
```

{{ if env "DOCGEN_TARGET=vuepress" }}
### Starship Example

The following screenshot is of a custom hint text using Starship:

<!-- markdownlint-disable -->
<img src="/screenshot-hint-starship.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->

Source: [https://github.com/orefalo/murex-module-starship](https://github.com/orefalo/murex-module-starship)
{{ end }}

### Disabling Hint Text

It is enabled by default but can be disabled if you prefer a more minimal
prompt:

```
» config set shell hint-text-enabled false
```

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
* `man` pages
* custom guides like https://cheat.sh
* information about binary files
* contents of text files
* and even images too!

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img src="/screenshot-preview-man-page.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<br/>
<img src="/screenshot-preview-image.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}

### Command Line Preview

> Enabled via `[f9]`

The Command Line Preview allows you to view the output of a command line while
you're still writing it. This interactivity removes the trial-and-error from
working with complicated command line incantations. For example parsing parsing
complex documents like machine generated JSON becomes very easy.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img src="/screenshot-preview-command-line.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}

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

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img src="/screenshot-paste-safety.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}

This gives you piece-of-mind that you are executing the right clipboard content
rather than something else you copied hours ago and forgotten about.

## Smarter Error Messages

Errors messages in most shells suck. That's why Murex has taken extra care to
give you as much useful detail as it can.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img src="/screenshot-error-messages.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}