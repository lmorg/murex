# iTerm2 Integrations

> Get more out of iTerm2 terminal emulator

## Description

iTerm2 it a macOS terminal emulator. It supports several ANSI escape codes that
are bespoke to iTerm2.

Murex can detect if it is running on iTerm2 and utilise those exclusive ANSI
escape codes, so you don't have to remember different ways of working with
different terminal emulators.

## Opening Images

Using [`open`](/docs/commands/open.md), you can render an image directly in the
terminal. Normally that would be a blocky "pixellated" representation using
block characters. But if you're running iTerm2, Murex will automatically switch
to iTerm2's ANSI escape sequences to render those images beautifully.

For example:

```
open https://murex.rocks/git-autocomplete.png
```

## See Also

* [Define Handlers For "`open`" (`openagent`)](../commands/openagent.md):
  Creates a handler function for `open`
* [Kitty Integrations](../integrations/kitty.md):
  Get more out of Kitty terminal emulator
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Terminology Integrations](../integrations/terminology.md):
  Get more out of Terminology terminal emulator

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

This document was generated from [gen/integrations/iterm2_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/integrations/iterm2_doc.yaml).