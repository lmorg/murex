# Murex: A Smarter Shell

[![Version](version.svg?v=undef)](DOWNLOAD.md)
[![Murex Tests](https://github.com/lmorg/murex/actions/workflows/murex-tests.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/murex-tests.yaml)
[![Deploy Docs](https://github.com/lmorg/murex/actions/workflows/deploy-docs.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/deploy-docs.yaml)
[![Official Website](images/website-badge.svg?v=undef)](https://murex.rocks)

[![Official Website](images/murex-term-light.png?v=undef)](https://murex.rocks)

> A smarter, more intuitive, and readable shell. You'll get more done, and more easily, with Murex

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


## Smart Data

Murex has native support for data formats such as JSON, YAML, XML, CSV, and others.

![video demonstrating Murex's data capabilities](/images/vhs-clever-data-dark.gif)

## Extremely Expressive

Murex has a flexible syntax that is both succinct enough to allow for fast typing
in the command line, but also readable shell scripts.

![video demonstrating various different syntactic features of Murex](/images/vhs-expressive-dark.gif)

## Better Error Handling

Shell scripts are notorious for having leaky failure modes. Murex fixes this
with familiar features like error handling and unit tests.

![vhs-better-errors centre-image](/images/vhs-better-errors-dark.gif)

# Getting Started

Visit our [official website](https://murex.rocks) for easier browsing of the documentation.

## Learn About The Command Line

Murex features a state-of-the-art [interactive command line](/user-guide/interactive-shell.md).
Read more about it's unique features.

![banner](/images/banner-interactive-light.png)

## Learn The Syntax

Read the [language tour](/tour.md) to learn about the syntax and how
shell scripting is easier in Murex.

![banner](/images/banner-tour-light.png)

## Cheat Sheet

The [Rosetta Stone](/user-guide/rosetta-stone.md) is a great cheat sheet for
those wishing to skip the tutorials and jump straight in.
This guide provides comparisons with Bash.

![banner](/images/banner-rosetta-light.png)

# Easy to Install

Install `murex` from your favorite package manager or directly from source:


### ArchLinux

From AUR: [https://aur.archlinux.org/packages/murex](https://aur.archlinux.org/packages/murex)

```bash
wget -O PKGBUILD 'https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=murex'
makepkg --syncdeps --install 
```

### FreeBSD Ports

```bash
pkg install murex
```

### Homebrew

```bash
brew install murex
```

### MacPorts

```bash
port install murex
```


More details are available in the [INSTALL](INSTALL.md) document.


## Discuss Murex

Discussions presently happen in [Github discussions](https://github.com/lmorg/murex/discussions).

## Compatibility Commitment

Murex is committed to backwards compatibility. While we do want to continue to
grow and improve the shell, this will not come at the expense of long term
usability. [Read more](compatibility.md)

## Issue Tracking

Murex is considered stable, however if you do run into problems then please
raise them on the project's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)

<hr/>

This document was generated from [gen/root/README_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/root/README_doc.yaml).