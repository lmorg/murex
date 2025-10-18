{{ if env "DOCGEN_TARGET=vuepress" }}
home: true
icon: home
#heroImage: murex-logo-shell.svg?v={{ env "COMMITHASHSHORT" }}
heroImage: murex-term.png?v={{ env "COMMITHASHSHORT" }}
heroText: Murex
tagline: A smarter, more intuitive, and readable shell. You'll get more done, and more easily, with Murex
actions:
  - text: " Language Tour"
    icon: "plane-departure"
    link: tour/
    type: primary
  - text: " Cheat Sheet"
    icon: table
    link: user-guide/rosetta-stone/
  - text: " Interactive Shell"
    icon: keyboard
    link: user-guide/interactive-shell/
  - text: " Install"
    icon: download
    link: install/

copyright: false
footer: GPLv2 Licensed, Copyright Laurence Morgan
---

## Smart Data

Murex has native support for data formats such as JSON, YAML, XML, CSV, and others.

<!-- markdownlint-disable -->
<img class="vhs-clever-data centre-image" alt="video demonstrating Murex's data capabilities">
<!-- markdownlint-restore -->

## Extremely Expressive

Murex has a flexible syntax that is both succinct enough to allow for fast typing
in the command line, but also readable shell scripts.

<!-- markdownlint-disable -->
<img class="vhs-expressive centre-image" alt="video demonstrating various different syntactic features of Murex">
<!-- markdownlint-restore -->

## Better Error Handling

Shell scripts are notorious for having leaky failure modes. Murex fixes this
with familiar features like error handling and unit tests.

<!-- markdownlint-disable -->
<img class="vhs-better-errors centre-image" alt="video demonstrating error handling">
<!-- markdownlint-restore -->

## Getting Started

<!-- markdownlint-disable -->
<img class="banner-interactive centre-image" alt="banner">
<!-- markdownlint-restore -->

Murex features a state-of-the-art [interactive command line](/user-guide/interactive-shell.html).
Read more about it's unique features.

<!-- markdownlint-disable -->
<img class="banner-tour centre-image" alt="banner">
<!-- markdownlint-restore -->

Read the [language tour](/tour.html) to learn about the syntax and how
shell scripting is easier in Murex.

<!-- markdownlint-disable -->
<img class="banner-rosetta centre-image" alt="banner">
<!-- markdownlint-restore -->

The [Rosetta Stone](/user-guide/rosetta-stone.html) is a great cheat sheet for
those wishing to skip the tutorials and jump straight in.
This guide includes comparisons with Bash.

## Easy to Install

Install `murex` from your favorite package manager or directly from source:

::: code-tabs#shell

@tab macOS
```sh
# via Homebrew:
brew install murex

# via MacPorts:
port install murex
```

@tab ArchLinux
```sh
# From AUR: https://aur.archlinux.org/packages/murex
wget -O PKGBUILD 'https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=murex'
makepkg --syncdeps --install 
```

@tab FreeBSD
```sh
pkg install murex
```

@tab Powershell
```powershell
# This requires `go` (Golang) and `git` to already be installed.
$env:GOBIN="$(pwd)"; & go install -v github.com/lmorg/murex@latest
```

:::

More details are available in the [INSTALL](install/) document.

{{ else }}# Murex: A Smarter Shell

[![Version](version.svg?v={{env "COMMITHASHSHORT"}})](DOWNLOAD.md)
[![Murex Tests](https://github.com/lmorg/murex/actions/workflows/murex-tests.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/murex-tests.yaml)
[![Deploy Docs](https://github.com/lmorg/murex/actions/workflows/deploy-docs.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/deploy-docs.yaml)
[![Official Website](images/website-badge.svg?v={{ env "COMMITHASHSHORT" }})](https://murex.rocks)

<img src="https://murex.rocks/murex-logo-shell.svg?v={{ env "COMMITHASHSHORT" }}" class="no-border">

Murex is a shell, like bash / zsh / fish / etc however Murex supports improved
features and an enhanced UX.

A non-exhaustive list features would include:

* Support for **additional type information in pipelines**, which can be used
  for complex data formats like JSON or tables. Meaning all of your existing
  UNIX tools to work more intelligently and without any additional configuration.

* **Usability improvements** such as in-line spell checking, context sensitive
  hint text that details a commands behavior before you hit return, and
  auto-parsing man pages for auto-completions on commands that don't have auto-
  completions already defined.
  
* **Smarter handling of errors** and **debugging tools**. For example try/catch
  blocks, line numbers included in error messages, stdout highlighted in red
  and script testing and debugging frameworks baked into the language itself.

## Language Guides

* Read the [language tour](/docs/tour.md) to get started.

* The [Rosetta Stone](/docs/user-guide/rosetta-stone.md) is a
great cheatsheet for those wishing to skip the tutorials and jump straight in.
This guide includes comparisons with Bash.

* The [Interactive Shell](/docs/user-guide/interactive-shell.md)
guide walks you through using Murex as a command line as opposed to a scripting
language.

## Examples

### Smart data:

<img src="images/screenshot-open-foreach.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

<img src="images/screenshot-ps-select.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

### Inline spellchecking:

<img src="images/screenshot-spellchecker.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

### Autocomplete:

<img src="images/screenshot-kill-autocomplete.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

<img src="images/screenshot-autocomplete-git.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

<img src="images/screenshot-history.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

### Preview screen:

<img src="images/screenshot-preview-man-page.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

<img src="images/screenshot-preview-command-line.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

### Useful error messages:

<img src="images/screenshot-error-messages.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

<img src="images/screenshot-paste-safety.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

### Plus More!

Visit the [official website](https://murex.rocks).

## Install instructions

See [INSTALL](INSTALL.md) for details.

## Discuss Murex

Discussions presently happen in [Github discussions](https://github.com/lmorg/murex/discussions).

## Compatibility Commitment

Murex is committed to backwards compatibility. While we do want to continue to
grow and improve the shell, this will not come at the expense of long term
usability. [Read more](compatibility.md)

## Issue Tracking

Murex is considered stable, however if you do run into problems then please
raise them on the project's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)
{{ end }}