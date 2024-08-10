{{ if env "DOCGEN_TARGET=vuepress" }}
home: true
icon: home
heroImage: murex-logo-shell.svg?v={{ env "COMMITHASHSHORT" }}
heroText: Murex.Rocks
tagline: An intuitive and content aware shell for a modern command line
head:
  - [meta, { property: "og-description", content: "An intuitive and content aware shell for a modern command line" }]
actions:
  - text: " Language Tour"
    icon: "plane-departure"
    link: tour/
    type: primary
  - text: " Rosetta Stone"
    icon: table
    link: user-guide/rosetta-stone/
  - text: " Interactive Shell"
    icon: keyboard
    link: user-guide/interactive-shell/
  - text: " Install"
    icon: download
    link: install/ 


highlights:
  - header: A Modern shell for the rest of us
    description: Murex carries tons of unique features. Some highlights include...
    features:
      - title: Content Aware
        icon: file-contract
        details: |-
          Native support for manipulating data formats such as JSON, YAML, CSV, and others. This allows for seamless integration and manipulation of data in various formats.
          <br/><br/>
          <strong>Data types can be explicitly cast and reformatted, but also inferred if preferred.</strong>
        link: types/

      - title: Expressions
        icon: hashtag
        details: |-
          Smarter handling of variables and expressions to avoid accidental bugs caused by spaces or incorrect syntax. Resulting in a more reliable and predictable scripting experience.
          <br/><br/>
          <strong>Never worry about file names with weird characters, nor running equations in "bc" again.</strong>
        link: parser/

      - title: Smartly Interactive
        icon: wand-magic-sparkles
        details: |-
          A uniquely intuitive interactive shell. With command line hints pulled from man pages, AI LLMs, and other intelligent integrations.
          <br/><br/>
          <strong>Navigating the command line is faster, more intuitive and efficient than ever before.</strong>
        link: user-guide/interactive-shell

      - title: Easily Extended
        icon: "puzzle-piece"
        details: |-
          The built-in package manager makes it very easy to share your configuration, import other peoples namespaced modules, and port your environment between different machines.
          <br/><br/>
          <strong>Configure once, use everywhere.</strong>
        link: user-guide/modules

copyright: false
footer: GPLv2 Licensed, Copyright © 2017-present Laurence Morgan
---
## Getting Started

* Read the [language tour](/tour.html) to get started.

* The [Rosetta Stone](/user-guide/rosetta-stone.html) is a
great cheatsheet for those wishing to skip the tutorials and jump straight in.
This guide includes comparisons with Bash.

* The [Interactive Shell](/user-guide/interactive-shell.html)
guide walks you through using Murex as a command line as opposed to a scripting
language.

## Screenshots

<!-- markdownlint-disable -->

<div class="image-preview">
  <img src="/screenshot-kill-autocomplete.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-open-foreach.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-spellchecker.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-autocomplete-git.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-error-messages.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-hint-text-rsync.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-preview-man-page.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-preview-command-line.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-paste-safety.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-autocomplete-context-sensitive.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-history.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/screenshot-ps-select.png?v={{ env "COMMITHASHSHORT" }}" />
</div>

<style>
  .image-preview {
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    flex-wrap: wrap;
  }

  .image-preview > img {
     box-sizing: border-box;
     width: 33.3% !important;
     padding: 9px;
     border-radius: 16px;
  }

  @media (max-width: 719px) {
    .image-preview > img {
      width: 50% !important;
    }
  }

  @media (max-width: 419px) {
    .image-preview > img {
      width: 100% !important;
    }
  }
</style>

<!-- markdownlint-restore -->

Check out the [Language Tour](/tour.html) and [Interactive Shell](user-guide/interactive-shell.html) guides!

## Easy to Install

Install `murex` from your favorite package manager:

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

:::

More options are available in the [INSTALL](install/) document.

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