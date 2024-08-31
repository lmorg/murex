# Man Pages (POSIX)

> Linux/UNIX `man` page integrations

## Description

`man` page parsing is used extensively throughout Murex. From providing default
autocompletions through to quick access of those manuals via the `[f1]` preview
pane.

This document describes the various different integrations around those `man`
pages.

> Please note that this is only supported on Linux, macOS, BSD and other
> UNIX-like platforms. Windows and Plan 9 do not currently support these
> specific integrations.

## Autocomplete

Autocomplete, sometimes referred to as "tab-completion", are command and
parameter suggestions offered when you press `[tab]`.

A lot of commands will have a bespoke [autocomplete config](/docs/commands/autocomplete.md)
defined for them. However writing autocomplete rules for every command out
there would be hugely time consuming and impractical. Thus **Murex can
automatically generate autocompletions** based on the contents of a commands
`man` page.

## Preview

If you want to quickly read a `man` page without disrupting your flow in the
command line, then you can press `[f1]` to preview it.

This preview will allow you to keep typing out your command line while also
presenting the `man` page in an easy to read layout.

## Summary

When you type a command in, you will see a brief description of what that
command is, in your [hint text](/docs/user-guide/interactive-shell.md#hint-text).
That command will generally be pulled from its accompanying `man` page.

## See Also

* [ChatGPT](../integrations/chatgpt.md):
  How to enable ChatGPT hints
* [Murex's Offline Documentation (`murex-docs`)](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation
* [cheat.sh](../integrations/cheatsh.md):
  Cheatsheets provided by cheat.sh

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

This document was generated from [gen/integrations/man_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/integrations/man_doc.yaml).