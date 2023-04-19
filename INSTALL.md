<h1>Language Tour</h1>

<h2>Table of Contents</h2>
<div id="toc">

- [Supported Platforms](#supported-platforms)
- [Pre-Compiled Binaries (HTTPS download)](#pre-compiled-binaries-https-download)
- [Installing From A Package Manager](#installing-from-a-package-manager)
  - [ArchLinux](#archlinux)
  - [Homebrew](#homebrew)
  - [MacPorts](#macports)
- [Compiling From Source](#compiling-from-source)
  - [Prerequisites](#prerequisites)
    - [Further Reading:](#further-reading)
  - [Installation From Source Steps](#installation-from-source-steps)
    - [Importing the source code](#importing-the-source-code)
    - [Test the code (optional stage)](#test-the-code-optional-stage)
    - [Compile the code](#compile-the-code)
    - [Test the executable (optional stage)](#test-the-executable-optional-stage)
    - [Start the shell](#start-the-shell)
- [Including Optional Builtins](#including-optional-builtins)
- [External Dependencies (Optional)](#external-dependencies-optional)
- [Recommended Terminal Typeface](#recommended-terminal-typeface)

</div>

## Supported Platforms

Linux, BSD and macOS are fully supported, with other platforms considered
experimental and/or having known limitations.

## Pre-Compiled Binaries (HTTPS download)

[![Version](version.svg)](DOWNLOAD.md)
[![CodeBuild](https://codebuild.eu-west-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoib3cxVnoyZUtBZU5wN1VUYUtKQTJUVmtmMHBJcUJXSUFWMXEyc2d3WWJldUdPTHh4QWQ1eFNRendpOUJHVnZ5UXBpMXpFVkVSb3k2UUhKL2xCY2JhVnhJPSIsIml2UGFyYW1ldGVyU3BlYyI6Im9QZ2dPS3ozdWFyWHIvbm8iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](DOWNLOAD.md)

If you wish to download a pre-compiled binary then head to the [DOWNLOAD](DOWNLOAD.md)
page to select your platform.

## Installing From A Package Manager

[![Packaging status](https://repology.org/badge/vertical-allrepos/murex.svg)](https://repology.org/project/murex/versions)

### ArchLinux

Available as an AUR package:

* PKGBUILD file:
  [https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=murex](https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=murex)

* AUR package repository:
  [https://aur.archlinux.org/packages/murex](https://aur.archlinux.org/packages/murex)

### Homebrew

    brew install murex
    
### MacPorts

    sudo port install murex

## Compiling From Source

[![CircleCI](https://circleci.com/gh/lmorg/murex/tree/master.svg?style=svg)](https://circleci.com/gh/lmorg/murex/tree/master)

### Prerequisites

You will need `go` (Golang) compiler, `gcc` (C compiler) and `git` installed
as well as your `$GOPATH` environmental variable set. 

> Go 1.18 or higher is required.

These should be easy to install on most operating systems however Windows is a
lot more tricky with regards to `gcc`. Please check with your operating systems
package manager first but see further reading below if you get stuck.

#### Further Reading:

- [How to install Go](https://golang.org/doc/install)
- [How to install git](https://github.com/git-guides/install-git)
- [How to install gcc](https://gcc.gnu.org/install/)
- [How to set GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Installation From Source Steps

The following instructions are assuming you're compiling on a POSIX-compatible
system like Linux, BSD or macOS. Compiling from source is untested on Plan 9
(if you run into issues there then please use the pre-compiled binary for that
platform) and Windows. In the case of Windows you may run into issues with the
`gcc` installation and some of the commands below will need to be adapted (eg
`murex.exe` used instead of `./murex`).

> Compiling from source is not recommended unless you already have a strong
> understanding of compiling Go projects for your specific platform.

#### Importing the source code

At present, _murex_ depends on being in a specific directory hierarchy for
the tests to work and packages to import correctly. These instructions will
talk you through creating that initial structure ready to import the source
into. Experienced users in Go may opt to ignore some of these steps and run
`go get -u github.com/lmorg/murex` instead. While this _should_ work in most
cases, it is difficult to run automated tests to ensure any updates doesn't
break the `go get` import tool. And thus that approach is not officially
supported. If you are in any doubt, please follow the `git clone` process
below.

First create the directory path and clone the source into the appropriate
directory structure.

    mkdir -p $GOPATH/src/github.com/lmorg/murex
    cd $GOPATH/src/github.com/lmorg/murex
    git clone https://github.com/lmorg/murex .

At this point you can add and remove any optional builtins by following the
instructions on this located further down this document. This is entirely
optional as _murex_ attempts to ship with sane defaults.

#### Test the code (optional stage)

    go test ./...

#### Compile the code

    go build github.com/lmorg/murex

#### Test the executable (optional stage)

    ./murex -c 'g: behavioural/* -> foreach: f { source $f }; try {test: run *}'

#### Start the shell

    ./murex

## Including Optional Builtins

Some optional builtins will be included by default, however there may be others
you wish to include which are not part of the default build (such as `qr`). To
add them, copy (or symlink) the applicable include file from
`builtins/import_src` to `builtins/import_build`.

A tool will be introduced in a later version to automate this.

## External Dependencies (Optional)

Some of _murex_'s extended features will have additional external dependencies.

* `aspell`: This is used for spellchecking. Murex will automatically enable or
  disable spellchecking based on whether `aspell` can be found in your `$PATH`.
  [http://aspell.net](http://aspell.net)

## Recommended Terminal Typeface

This is obviously just a subjective matter and everyone will have their own
personal preference. However if I was asked what my preference was then that
would be [Hasklig](https://github.com/i-tu/Hasklig). It's a clean typeface
based off Source Code Pro but with a few added ligatures - albeit subtle ones
designed to make Haskell more readable. Those ligatures also suite _murex_
pretty well. So the overall experience is a clean and readable terminal.
