# Install Instructions

## Pre-Compiled Binaries

If you wish to download a pre-compiled binary then head to the [DOWNLOAD](DOWNLOAD.md)
page to select your platform.

This is the recommended way to install _murex_ unless you're already familiar
with `git` and `go`, or need to enable/disable a specific builtin from what is
compiled as part of the standard build.

## From Source

> Go 1.12 or higher is required

### Prerequisites

You will need `go` (Golang) compiler and `git` installed, and your `$GOPATH`
environmental variable set. You can check these by running:

    which go
    which git
    echo $GOPATH

(each of those commands should return a non-zero length string).

#### Further Reading:

- [How to install Go](https://golang.org/doc/install)
- [How to install git](https://github.com/git-guides/install-git)
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

    mkdir -p $GOPATH/lmorg/murex
    git clone github.com/lmorg/murex $GOPATH/lmorg/murex

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

_murex_ provides two prebuilt images on Docker Hub:

* `lmorg/murex:develop` - this is the latest build of the `develop` branch,
  as such it might contain unstable code
* `lmorg/murex:latest` - this is the latest build of the `master` branch and
  is the recommended image to use

### Building Your Own Container

From the project root (the location of this INSTALL.md file) run the following:

    docker-compose up --build murex

## Including Optional Builtins

Some optional builtins will be included by default, however there may be others
you wish to include which are not part of the default build (such as `select`).
To add them, copy (or symlink) the optional file from `builtins/import_src` to
`builtins/import_build`.

A tool will be introduced in a later version to automate this.

## Supported Platforms

Most popular operating systems and CPU types are supported. More details
can be read at (docs/FAQ.supported-platforms.md).

## Recommended Terminal Typeface

This is obviously just a subjective matter and everyone will have their own
personal preference. However if I was asked what my preference was then that
would be [Hasklig](https://github.com/i-tu/Hasklig). It's a clean typeface
based off Source Code Pro but with a few added ligatures - albeit subtle ones
designed to make Haskell more readable. Those ligatures also suite _murex_
pretty well. So the overall experience is a clean and readable terminal.
