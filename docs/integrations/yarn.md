# `yarn` Integrations

> Working with `yarn` and `package.json`

## Description

`yarn` is a common tool for working with `package.json`. Some developers use
this almost like a `Makefile`.

Murex comes with some handy autocompletions out-of-the-box for working with
`yarn` in the command line.

## Autocompletions

Custom [autocomplete](/docs/commands/autocomplete.md) rules exist for `yarn` which will
not only include any `yarn` specific flags, but also include any parameters
defined in your `package.json` too.

## Source Code

The source code is available on [Github](https://github.com/lmorg/murex/blob/master/integrations/yarn_any.mx)
under `/integrations`.

## See Also

* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation

## Other Integrations

* [ChatGPT](../integrations/chatgpt.md):
  How to enable ChatGPT hints
* [Cheat.sh](../integrations/cheatsh.md):
  Cheatsheets provided by cheat.sh
* [Kitty Integrations](../integrations/kitty.md):
  Get more out of Kitty terminal emulator
* [Makefiles / `make`](../integrations/make.md):
  `make` integrations
* [Man Pages (POSIX)](../integrations/man-pages.md):
  Linux/UNIX `man` page integrations
* [Spellcheck](../integrations/spellcheck.md):
  How to enable inline spellchecking
* [Terminology Integrations](../integrations/terminology.md):
  Get more out of Terminology terminal emulator
* [`direnv` Integrations](../integrations/direnv.md):
  Directory specific environmental variables
* [`yarn` Integrations](../integrations/yarn.md):
  Working with `yarn` and `package.json`
* [iTerm2 Integrations](../integrations/iterm2.md):
  Get more out of iTerm2 terminal emulator


<hr/>

This document was generated from [gen/integrations/yarn_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/integrations/yarn_doc.yaml).