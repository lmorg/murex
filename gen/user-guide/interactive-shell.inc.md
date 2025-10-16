Murex's interactive shell is also built around productivity. To achieve this we
wrote our own state-of-the-art readline library.

Below are just some of the features you can enjoy.

{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>

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

{{ end }}

## Advanced Autocompletion

Autocompletion happen when you press `[tab]` and will differ slightly depending
on what is defined in `autocomplete` and whether you use the traditional
[POSIX pipe token](../parser/pipe-posix.md), `|`, or the [arrow pipe](../parser/pipe-arrow.md),
`->`.

The `|` token will behave much like any other shell however `->` will offer
suggestions with matching data types. Which makes working working with data
quick and easy while still intelligent and readable.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img class="vhs-autocompletion">
<!-- markdownlint-restore -->
{{ else }}![autocomplete preview](/images/vhs-autocompletion-dark.gif){{ end }}

## Hint Text

{{ file "gen/user-guide/hint-text-overview.inc.md" }}

{{link "Read more about Hint Text" "hint-text"}}.

## Spellchecker

Murex supports inline spellchecking, where errors are underlined. For example

[![asciicast](https://asciinema.org/a/408024.svg)](https://asciinema.org/a/408024)

This might require some manual steps to enable, please see the [spellcheck user guide](spellcheck.md)
for more details.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<figure>
    <img src="/screenshot-spellchecker.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
    <figcaption>Inline spellchecking</figcaption>
</figure>
<!-- markdownlint-restore -->
{{ end }}

## Preview Autocompletions

> Enabled via `[f1]`

This displays a more detailed view of each parameter you're about to pass to a
command, without you having to run that command nor leave the half-completed
command line.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img class="vhs-preview-autocomplete">
<!-- markdownlint-restore -->
{{ else }}![autocomplete preview](/images/vhs-preview-autocomplete-dark.gif){{ end }}

It can display:
* {{link "`man` pages" "man-pages"}}
* custom guides like {{link "https://cheat.sh" "cheat.sh"}} and {{link "AI generated docs" "chatgpt"}}
* information about binary files
* contents of text files
* and even images too!


## Preview Command Lines

> Enabled via `[f9]`

The Command Line Preview allows you to view the output of a command line while
you're still writing it. This interactivity removes the trial-and-error from
working with complicated command line incantations. For example parsing parsing
complex documents like machine generated JSON becomes very easy.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img class="vhs-preview-commandline">
<!-- markdownlint-restore -->
{{ else }}![autocomplete preview](/images/vhs-preview-commandline-dark.gif){{ end }}

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
<figure>
    <img src="/screenshot-paste-safety.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
    <figcaption>Murex presets a menu to double check what you're pasting before the shell executes it</figcaption>
</figure>
<!-- markdownlint-restore -->
{{ end }}

This gives you piece-of-mind that you are executing the right clipboard content
rather than something else you copied hours ago and forgotten about.

## Smarter Error Messages

Errors messages in most shells suck. That's why Murex has taken extra care to
give you as much useful detail as it can.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img class="vhs-better-errors-errmsg">
<!-- markdownlint-restore -->
{{ else }}![autocomplete preview](/images/vhs-better-errors-errmsg-dark.png){{ end }}

## Hotkeys

A full breakdown of supported hotkeys is available in the [terminal-keys](terminal-keys.md)
guide.