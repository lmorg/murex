{{ if env "DOCGEN_TARGET=vuepress" }}
home: true
icon: home
#heroImage: murex-logo-shell.svg?v={{ env "COMMITHASHSHORT" }}
heroImage: {{ .DocumentMeta.Logo }}?v={{ env "COMMITHASHSHORT" }}
heroText: {{ .Title }}
tagline: {{ .Summary }}
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
{{ else }}# Murex: A Smarter Shell

[![Version](version.svg?v={{env "COMMITHASHSHORT"}})](DOWNLOAD.md)
[![Murex Tests](https://github.com/lmorg/murex/actions/workflows/murex-tests.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/murex-tests.yaml)
[![Deploy Docs](https://github.com/lmorg/murex/actions/workflows/deploy-docs.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/deploy-docs.yaml)
[![Official Website](images/website-badge.svg?v={{ env "COMMITHASHSHORT" }})](https://murex.rocks)

[![Official Website](images/{{ .DocumentMeta.Logo }}?v={{ env "COMMITHASHSHORT" }})](https://murex.rocks)

> {{ .Summary }}

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
{{ end }}

## Smart Data

Murex has native support for data formats such as JSON, YAML, XML, CSV, and others.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img class="vhs-clever-data centre-image" alt="video demonstrating Murex's data capabilities">
<!-- markdownlint-restore -->
{{ else }}![video demonstrating Murex's data capabilities](/images/vhs-clever-data-dark.gif){{ end }}

## Extremely Expressive

Murex has a flexible syntax that is both succinct enough to allow for fast typing
in the command line, but also readable shell scripts.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img class="vhs-expressive centre-image" alt="video demonstrating various different syntactic features of Murex">
<!-- markdownlint-restore -->
{{ else }}![video demonstrating various different syntactic features of Murex](/images/vhs-expressive-dark.gif){{ end }}

## Better Error Handling

Shell scripts are notorious for having leaky failure modes. Murex fixes this
with familiar features like error handling and unit tests.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img class="vhs-better-errors centre-image" alt="video demonstrating error handling">
<!-- markdownlint-restore -->
{{ else }}![vhs-better-errors centre-image](/images/vhs-better-errors-dark.gif){{ end }}

# Getting Started

{{ if env "DOCGEN_TARGET=" }}Visit our [official website](https://murex.rocks) for easier browsing of the documentation.
{{ end }}
## Learn About The Command Line

Murex features a state-of-the-art [interactive command line](/user-guide/interactive-shell.html).
Read more about it's unique features.

{{ if env "DOCGEN_TARGET=vuepress" }}<!-- markdownlint-disable -->
<img class="banner-interactive centre-image" alt="banner">
<!-- markdownlint-restore -->
{{ else }}![banner](/images/banner-interactive-light.png){{ end }}

## Learn The Syntax

Read the [language tour](/tour.html) to learn about the syntax and how
shell scripting is easier in Murex.

{{ if env "DOCGEN_TARGET=vuepress" }}<!-- markdownlint-disable -->
<img class="banner-tour centre-image" alt="banner">
<!-- markdownlint-restore -->
{{ else }}![banner](/images/banner-tour-light.png){{ end }}

## Cheat Sheet

The [Rosetta Stone](/user-guide/rosetta-stone.html) is a great cheat sheet for
those wishing to skip the tutorials and jump straight in.
This guide provides comparisons with Bash.

{{ if env "DOCGEN_TARGET=vuepress" }}<!-- markdownlint-disable -->
<img class="banner-rosetta centre-image" alt="banner">
<!-- markdownlint-restore -->
{{ else }}![banner](/images/banner-rosetta-light.png){{ end }}

# Easy to Install

Install `murex` from your favorite package manager or directly from source:

{{ tmpl (file "gen/includes/install-package.inc.md") .Ptr }}

More details are available in the [INSTALL]({{ if env "DOCGEN_TARGET=vuepress" }}install/{{ else }}INSTALL.md{{ end }}) document.

{{ if env "DOCGEN_TARGET=" }}
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