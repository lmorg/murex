# Murex: A Smarter Shell

[![Version](version.svg?undef)](DOWNLOAD.md)
[![Murex Tests](https://github.com/lmorg/murex/actions/workflows/murex-tests.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/murex-tests.yaml)
[![Deploy Docs](https://github.com/lmorg/murex/actions/workflows/deploy-docs.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/deploy-docs.yaml)

<img src="https://murex.rocks/murex-logo-shell.svg?v=undef" class="readme">

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
  blocks, line numbers included in error messages, STDOUT highlighted in red
  and script testing and debugging frameworks baked into the language itself.

## Language Guides

* Read the [language tour](https://murex.rocks/tour.html) to get started.

* The [Rosetta Stone](https://murex.rocks/user-guide/rosetta-stone.html) is a
great cheatsheet for those wishing to skip the tutorials and jump straight in.
This guide includes comparisons with Bash.

* The [Interactive Shell](https://murex.rocks/user-guide/interactive-shell.html)
guide walks you through using Murex as a command line as opposed to a scripting
language.

## Examples

**Smart data:**

<img src="images/screenshot-open-foreach.png?v=undef" class="readme">

<img src="images/screenshot-ps-select.png?v=undef" class="readme">

**Inline spellchecking:**

<img src="images/screenshot-spellchecker.png?v=undef" class="readme">

**Autocomplete:**

<img src="images/screenshot-kill-autocomplete.png?v=undef" class="readme">

<img src="images/screenshot-autocomplete-git.png?v=undef" class="readme">

<img src="images/screenshot-history.png?v=undef" class="readme">

**Preview screen:**

<img src="images/screenshot-preview-man-page.png?v=undef" class="readme">

<img src="images/screenshot-preview-custom-hints.png?v=undef" class="readme">

<img src="images/screenshot-preview-image.png?v=undef" class="readme">

<img src="images/screenshot-preview-command-line.png?v=undef" class="readme">

**Useful error messages:**

<img src="images/screenshot-error-messages.png?v=undef" class="readme">

<img src="images/screenshot-paste-safety.png?v=undef" class="readme">

**Plus More!**

Check out the [Language Tour](https://murex.rocks/tour.html) and [Interactive Shell](https://murex.rocks/user-guide/interactive-shell.html) guides!

## Install instructions

See [INSTALL](https://murex.rocks/INSTALL.html) for details.

## Discuss Murex

Discussions presently happen in [Github discussions](https://github.com/lmorg/murex/discussions).

## Compatibility Commitment

Murex is committed to backwards compatibility. While we do want to continue to
grow and improve the shell, this will not come at the expense of long term
usability. [Read more](compatibility.md)

## Known bugs / TODO

Murex is considered stable, however if you do run into problems then please
raise them on the project's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)

<hr/>

This document was generated from [gen/root/README_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/root/README_doc.yaml).