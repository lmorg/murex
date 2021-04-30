# Install Instructions

## Pre-Compiled Binaries

If you wish to download a pre-compiled binary then head to the [DOWNLOAD.md](DOWNLOAD.md)
page to select your platform.

## From Source

> Go 1.11 or higher is required

Assuming you already have Go (Golang) installed, you can download the
source just by running the following from the command line

    go get -u github.com/lmorg/murex
    cd $GOPATH/src/github.com/lmorg/murex

Test the code (optional stage):

    go test ./...

Compile the code:

    go build github.com/lmorg/murex

Then to start the shell:

    ./murex

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

## Including optional builtins

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
