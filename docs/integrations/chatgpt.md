# ChatGPT

> How to enable ChatGPT hints

## Description

While `man` pages are an invaluable resource, they can sometimes be verbose and
in a language that is hard to follow if you're unfamiliar with either that tool
and/or Linux / UNIX man pages in general.

The rise of AI tools like ChatGPT can offer a lifeline for people who want a
more tailored guide to using command line tools. So Murex has a ChatGPT
integration to support those use cases directly from the command line.

It is disabled by default (see below) however once enabled, you can view
ChatGPT's response in the preview screen, by pressing `[f1]`.

## Enabling ChatGPT preview

To enable, you must add your OpenAI key to your profile, which is typically
`~/.murex_profile` (see [profile](/docs/user-guide/profile.md) for more details):

```
config set preview-exec-api xxxxx
```

(where **xxxxx** is replaced with your OpenAI API key)

That API key must have write access to `/v1/chat/completions`. If you're
uncertain what this means, it will become more clear when you create that API
on OpenAI's website.

> By default, ChatGPT integrations are disabled.
> 
> If enabled then the default settings are to only send the command name to
> ChatGPT, rather than the whole command line.

If you're not working with sensitive command lines, then you might also wish
to enable sending the whole command line rather than just the command name.
See [chatgpt.md#working-with-sensitive-data](Working with sensitive data) for more
details. 

### Why config

The reason Murex requires this `config` parameter set rather than reading
directly from an environmental variable (typically `$OPENAI_API_KEY`) is so
that your OpenAI API key doesn't get unintentionally leaked to any other
programs, and also so that you explicitly have to opt into enabling this
feature.

### Working with sensitive data

By default, only the command name is sent to ChatGPT. So if your command line
reads something like `rsync medical-records.csv secret-server:/CSVs/`, the only
data leaked if `preview-exec-api` is enabled would be `rsync`.

However you might want more context specific help based on the entirety of your
command line in situations where you know you aren't going to include sensitive
data within your context.

To enable that context to be included in your ChatGPT prompt, add the following
to your profile:

```
config set openai preview-exec-send-cmdline true
```

> Please note that command lines are still stored in `~/.murex_history`! The
> default of disabling or enabling command line context (ie only sending the
> command name) via **preview-exec-send-cmdline** only applies it to OpenAI.
> And even then, that only happens if **preview-exec-api** is _also_ set, which
> it isn't by default.

## GPT model

The **gpt-3.5-turbo** model is used by default because it is quick and cheap.
However you can change that via **preview-exec-model**. For example

```
config set preview-exec-model gpt-4-turbo
```

You may need to clear the cache to see any changes:

```
runtime --clear-cache
```

## Customising the prompt

The default prompt passed to ChatGPT can be viewed via `config get openai
preview-exec-prompt`.

This can be changed should you want something more specific to your needs. For
example:

```
config set openai preview-exec-prompt "You are a medieval pirate transported into the future and gifted the skills of the ${os} command line. Describe how to use UNIX commands in your own, old English pirate, language"
```

...might produce output like the following:

> Arrr matey! To use the powerful `reboot` command in the land of Linux, ye
> must open thy terminal and type in `reboot`. This command will set sail the
> process of restarting thy device, just like hoisting the anchor and setting
> off on a new adventure on the high seas. But beware, for all unsaved
> documents will walk the plank if not properly stashed away afore using this
> command. Fair winds to ye!

You may need to clear the cache to see any changes:

```
runtime --clear-cache
```

## Caching

A successful ChatGPT response is cached for 50 days. You can clear that cache
with `runtime --clear-cache`.

## Source Code

The source code is available on [Github](https://github.com/lmorg/murex/blob/master/integrations/chatgpt_any.mx)
under `/integrations`.

## See Also

* [Man Pages (POSIX)](../integrations/man-pages.md):
  Linux/UNIX `man` page integrations
* [Murex's Offline Documentation (`murex-docs`)](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [Profile Files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
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

This document was generated from [gen/integrations/chatgpt_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/integrations/chatgpt_doc.yaml).