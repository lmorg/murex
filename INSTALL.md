# Install Instructions

While _murex_ aims at being cross platform, there are some known limitations on
Windows and Plan 9. Please read (docs/FAQ.supported-platforms.md) for more
information.

## From A Package Manager

Currently only Homebrew is supported. More package managers are expected to be
supported in the future however due to the numbers and variety of solutions out
there, we do ask for community support to help bring _mure_ to your preferred
platform, if it isn't already supported.

### Homebrew

    brew install murex

## Pre-Compiled Binaries (HTTP download)

If you wish to download a pre-compiled binary then head to the [DOWNLOAD](DOWNLOAD.md)
page to select your platform.

This is the recommended way to install _murex_ unless you're already familiar
with `git` and `go`, or need to enable/disable a specific builtin from what is
compiled as part of the standard build.

## From Source

> Go 1.12 or higher is required

### Prerequisites

You will need `go` (Golang) compiler, `gcc` (C compiler) and `git` installed
as well as your `$GOPATH` environmental variable set. You can check these by
running:

    which go
    which git
    which gcc
    echo $GOPATH

(each of those commands should return a non-zero length string aside from the
`echo` statement).

These should be easy to install on most operating systems however Windows is a
lot more tricky with regards to `gcc`. Please check with your operating systems
package manager first but see further reading below if you get stuck.

#### Further Reading:

- [How to install Go](https://golang.org/doc/install)
- [How to install git](https://github.com/git-guides/install-git)
- [How to install gcc](https://gcc.gnu.org/install/)
- [How to set GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Installation From Source Steps

#### Importing the source code

> At present, _murex_ depends on being in a specific directory hierarchy for
> the tests to work and packages to import correctly. These instructions will
> talk you through creating that initial structure ready to import the source
> into. Experienced users in Go may opt to ignore some of these steps and run
> `go get -u github.com/lmorg/murex` instead. While this _should_ work in most
> cases, it is difficult to run automated tests to ensure any updates doesn't
> break the `go get` import tool. And thus that approach is not officially
> supported. If you are in any doubt, please follow the `git clone` process
> below.

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

    ./murex --run-tests
    ./murex -c 'source: ./flags_test.mx; try {test: run *}'

#### Start the shell

    ./murex

or, on Windows,...

    murex.exe

## Inside Docker

If you don't have nor want to install Go and already have `docker` (and
`docker-compose` installed), then you can install _murex_ using the CI/CD
pipeline scripts.

### Docker Hub

Due to licensing changes from Docker, Docker Hub images are no longer up to
date. However you can still build your own container.

### Building Your Own Container

From the project root (the location of this INSTALL.md file) run the following:

    docker-compose up --build murex

## Including Optional Builtins

Some optional builtins will be included by default, however there may be others
you wish to include which are not part of the default build (such as `select`).
To add them, copy (or symlink) the applicable include file from
`builtins/import_src` to `builtins/import_build`.

A tool will be introduced in a later version to automate this.

## Supported Platforms

Most popular operating systems and CPU types are supported. More details
can be read at (docs/FAQ.supported-platforms.md).

Please note that Windows support is experimental and there are bugs specific to
Windows due to the differences in how commands are executed on Windows. In some
instances these bugs are significant to the user experience of _murex_ and
cannot be worked around. The recommended approach for running _murex_ on
Windows is to use a Linux port of the shell on a POSIX compatibility layer such
as WSL or Cygwin. Please see (docs/FAQ.supported-platforms.md) for more details.

## Recommended Terminal Typeface

This is obviously just a subjective matter and everyone will have their own
personal preference. However if I was asked what my preference was then that
would be [Hasklig](https://github.com/i-tu/Hasklig). It's a clean typeface
based off Source Code Pro but with a few added ligatures - albeit subtle ones
designed to make Haskell more readable. Those ligatures also suite _murex_
pretty well. So the overall experience is a clean and readable terminal.
