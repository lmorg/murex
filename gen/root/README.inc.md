{{ if env "DOCGEN_TARGET=vuepress" }}
home: true
icon: home
heroImage: murex-logo-shell.svg?v={{ env "COMMITHASHSHORT" }}
# bgImage: bluebg.jpg
# bgImageDark: https://theme-hope-assets.vuejs.press/bg/4-dark.svg
# bgImageStyle:
#   background-attachment: fixed
heroText: Murex.Rocks
tagline: An intuitive, typed and content aware shell for the 2020s and beyond.
actions:
  - text: Getting Started 💡
    link: tour/
    type: primary
  - text: Rosetta Stone 🪨
    link: user-guide/rosetta-stone/
  - text: Builtins 📔
    link: commands/

highlights:
  - header: A Modern shell for the rest of us
    description: Murex carries tons of unique features...
    # image: /murex.svg
    # bgImage: linesbg.svg
    # bgImageDark: https://theme-hope-assets.vuejs.press/bg/2-dark.svg
    # bgImageStyle:
    #   background-repeat: repeat
    #   background-size: initial
    features:
      - title: Content Aware
        icon: arrows-to-circle
        details: Murex has built-in support for natively manipulating various file formats such as JSON, TOML, YAML, CSV, and commonlog. This allows for seamless integration and manipulation of data in these formats.
        link: types/

      - title: Objects and Maps
        icon: table-columns
        details: Murex provides powerful data structures like maps, hashes, lists, and dictionaries, which can be used for efficient and flexible data manipulation. These data structures enable you to organize and manipulate data in a structured and intuitive way.
        link: mkarray/

      - title: Array manipulation
        icon: layer-group
        details: Murex comes with native built-in functions and features that allow for proper manipulation of arrays. These functions make it easy to perform operations like filtering, sorting, appending, and merging arrays, providing a seamless experience for working with array data
        link: mkarray/

      - title: Scalar expression
        icon: check-double
        details: Murex treats variables as expressions, allowing you to perform calculations and evaluations directly within the shell. This feature helps to avoid accidental bugs caused by spaces or incorrect syntax, providing a more reliable and predictable scripting experience.
        link: tour/#scalars

      - title: Public & Private functions
        icon: unlock-keyhole
        details: Murex supports both public and private functions. Private functions have restricted scope visibility, meaning they can only be accessed within the nearest module or source file. This allows for better encapsulation and organization of code, enhancing code readability and maintainability.
        link: commands/private

      - title: Type inference
        icon: text-height
        details: Murex employs type inference to automatically determine the data type of variables and pipelines it manages. This means that you don't always have to explicitly specify the data type, as the shell can intelligently infer it based on the context. This simplifies scripting and reduces the need for explicit type declarations.
        link: tour/#type-inference

      - title: Enhanced pipelines & redirection
        icon: puzzle-piece
        details: Murex supports sending typed information to compatible functions via redirection and pipelines. This allows for more efficient and flexible data processing.
        link: user-guide/pipeline

      - title: Type casting and formats
        icon: text-width
        details: Murex allows you to change the meta-data about how an information should be read or displayed. This can be useful for manipulating and formatting data in a desired way.
        link: tour/#type-conversion

      - title: Inline spellchecking
        icon: spell-check
        details: Murex provides inline spellchecking, which quickly identifies typing spelling errors with underlined text. This helps to catch and correct errors in real-time.
        link: user-guide/spellcheck

      - title: Smart Autocomplete
        icon: wand-magic-sparkles
        details: Murex parses man pages for command line flags and provides smart autocomplete functionality. By pressing the TAB key, you can auto-complete commands and parameters, making command line navigation faster and more efficient.
        link: commands/autocomplete

      - title: Hint text
        icon: comment
        details: Murex provides hint text, which gives clues to the user without any distractions. This can be useful for providing additional information or guidance to the user
        link: user-guide/interactive-shell#hint-text

      - title: Syntax highlighting
        icon: highlighter
        details: In the interactive terminal, Murex provides syntax highlighting, making it easier to read and understand code. Syntax highlighting can also be piped to the next built-in for further processing.
        link: user-guide/interactive-shell#syntax-highlighting

      - title: Syntax Completion
        icon: down-left-and-up-right-to-center
        details: Murex balances and auto-closes brackets and accolades, making it easier to write and edit code. This feature helps to ensure that code is properly formatted and avoids syntax errors.
        link: user-guide/interactive-shell#syntax-completion

      - title: Extension Framework
        icon: cube
        details: Murex has an extension framework that allows you to design your own modules or enjoy prebuilt extensions such as `auto-jump` or `starfish`. This allows for customization and additional functionality.
        link: user-guide/modules

      - title: Built-in Package Manager
        icon: cubes
        details: Murex comes with a built-in package manager that allows you to search and manage the lifecycle of packages. This makes it easy to install and manage dependencies.
        link: commands/murex-package

      - title: 80 builtins commands
        icon: building
        details: Murex provides 80 built-in commands allowing for fast execution and portability. These built-in commands cover a wide range of functionalities.
        link: commands/

      - title: Realtime Events
        icon: bolt
        details: Murex supports realtime events, which streamline script notifications upon elapsed time, keypress, completion, prompt, or filesystem changes. This allows for more dynamic and responsive scripts.
        link: events/

      - title: PNG Generation
        icon: image
        details: Murex can generate barcodes and images directly from scripts. This can be useful for generating visual representations of data or for creating graphical outputs.
        link: optional/qr

      - title: Multi-threaded
        icon: gears
        details: Murex uses separate threads for built-ins, rather than forking processes like in a traditional POSIX shell. This optimizes resource usage and improves performance.
        link: commands/fid-list

      - title: NOT POSIX compliant!
        icon: recycle
        details: Murex is purposely not POSIX compliant in order to be performant and allow for extended capabilities. This allows for more flexibility and advanced features.
        link: /

copyright: false
footer: GPLv2 Licensed, Copyright © 2017-present Laurence Morgan
---

## 👁‍🗨 Screenshots

<!-- markdownlint-disable -->

<div class="image-preview">
  <img src="/murex-kill-autocomplete.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/murex-open-foreach.png?v={{ env "COMMITHASHSHORT" }}" />
  <img src="/murex-spellchecker.png?v={{ env "COMMITHASHSHORT" }}" />
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

## 📦 Easy to Install

Install `Murex` from your favorite package manager

::: code-tabs#shell

@tab macOS


```bash
# via Homebrew:
brew install murex

# via MacPorts:
sudo port install murex
```

@tab ArchLinux
```bash
# From AUR: https://aur.archlinux.org/packages/murex
wget -O PKGBUILD 'https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=murex'
makepkg --syncdeps --install 
```

@tab FreeBSD
```
# Murex is available in [FreeBSD ports](https://www.freebsd.org/ports/)
```

:::

More options are available in the [INSTALL](install/) document

## 🛟 Getting Started

Take your your first steps with `Murex` by following our [Language Tutorial](tour/)

{{ else }}# Murex: A Smarter Shell

[![Version](version.svg?{{env "COMMITHASHSHORT"}})](DOWNLOAD.md)
[![CodeBuild](https://codebuild.eu-west-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoib3cxVnoyZUtBZU5wN1VUYUtKQTJUVmtmMHBJcUJXSUFWMXEyc2d3WWJldUdPTHh4QWQ1eFNRendpOUJHVnZ5UXBpMXpFVkVSb3k2UUhKL2xCY2JhVnhJPSIsIml2UGFyYW1ldGVyU3BlYyI6Im9QZ2dPS3ozdWFyWHIvbm8iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](DOWNLOAD.md)
[![Tests](https://github.com/lmorg/murex/actions/workflows/go-tests.yaml/badge.svg?branch=master)](https://github.com/lmorg/murex/actions/workflows/go-tests.yaml)

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

## Examples

**JSON wrangling:**

<img src="images/murex-open-foreach.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

**Inline spellchecking:**

<img src="images/murex-spellchecker.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

**Autocomplete descriptions, process IDs accompanied by process names:**

<img src="images/murex-kill-autocomplete.png?v={{ env "COMMITHASHSHORT" }}" class="readme">

More examples: [/examples](https://github.com/lmorg/murex/tree/master/examples)

## Install instructions

See [INSTALL](https://murex.rocks/INSTALL.html) for details.

## Language Tour

Read the [language tour](https://murex.rocks/docs/tour.html) to get started.

## Discuss Murex

Discussions presently happen in [Github discussions](https://github.com/lmorg/murex/discussions).

## Known bugs / TODO

Murex is considered stable, however if you do run into problems then please
raise them on the project's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)
{{ end }}