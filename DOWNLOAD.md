<h1>Download Murex</h1>

[![Version](version.svg)](DOWNLOAD.md)

<h2>Table of Contents</h2>

<div id="toc">

- [Download Links](#download-links)
  - [Darwin (macOS)](#darwin-macos)
  - [Linux](#linux)
  - [Windows](#windows)
  - [BSD's](#bsds)
    - [DragonflyBSD](#dragonflybsd)
    - [FreeBSD](#freebsd)
    - [NetBSD](#netbsd)
    - [OpenBSD](#openbsd)
  - [Solaris](#solaris)
  - [Plan 9](#plan-9)
- [Install Instructions](#install-instructions)
  - [Linux / UNIX / macOS Instructions](#linux--unix--macos-instructions)
  - [Windows Instructions](#windows-instructions)

</div>

## Download Links

Below are the instructions to download a pre-compiled binary via HTTPS. If you
wish to install from source or use your preferred package manager, then please
refer to the [INSTALL](INSTALL.md) page for further instructions.

### Darwin (macOS)

The `arm64` builds support the ARM-based M1, M2 and M3 processors. Older Macs
will need to run `amd64`. Murex is also available on [Homebrew](INSTALL.md#homebrew) and [MacPorts](INSTALL.md#macports).


[Install instructions](#linux-unix-macos-instructions) 
can be found further down this page.

* [murex-darwin-amd64.gz](https://nojs.murex.rocks/bin/latest/murex-darwin-amd64.gz)
* [murex-darwin-arm64.gz](https://nojs.murex.rocks/bin/latest/murex-darwin-amd64.gz)

### Linux

[Install instructions](#linux-unix-macos-instructions)
can be found further down this page.

* [murex-linux-386.gz](https://nojs.murex.rocks/bin/latest/murex-linux-386.gz)
* [murex-linux-amd64.gz](https://nojs.murex.rocks/bin/latest/murex-linux-amd64.gz)
* [murex-linux-arm.gz](https://nojs.murex.rocks/bin/latest/murex-linux-arm.gz)
* [murex-linux-arm64.gz](https://nojs.murex.rocks/bin/latest/murex-linux-arm64.gz)

### Windows

[Install instructions](#windows-instructions)
can be found further down this page.

* [murex-windows-386.exe.zip](https://nojs.murex.rocks/bin/latest/murex-windows-386.exe.zip)
* [murex-windows-amd64.exe.zip](https://nojs.murex.rocks/bin/latest/murex-windows-amd64.exe.zip)

### BSD's

[Install instructions](#linux-unix-macos-instructions)
can be found further down this page.

#### DragonflyBSD

* [murex-dragonfly-amd64.gz](https://nojs.murex.rocks/bin/latest/murex-dragonfly-amd64.gz)

#### FreeBSD

Murex is also available in [FreeBSD Ports](https://murex.rocks/INSTALL.html#freebsd-ports).

* [murex-freebsd-386.gz](https://nojs.murex.rocks/bin/latest/murex-freebsd-386.gz)
* [murex-freebsd-amd64.gz](https://nojs.murex.rocks/bin/latest/murex-freebsd-amd64.gz)
* [murex-freebsd-arm.gz](https://nojs.murex.rocks/bin/latest/murex-freebsd-arm.gz)
* [murex-freebsd-arm64.gz](https://nojs.murex.rocks/bin/latest/murex-freebsd-arm64.gz)

#### NetBSD

* [murex-netbsd-386.gz](https://nojs.murex.rocks/bin/latest/murex-netbsd-386.gz)
* [murex-netbsd-amd64.gz](https://nojs.murex.rocks/bin/latest/murex-netbsd-amd64.gz)
* [murex-netbsd-arm.gz](https://nojs.murex.rocks/bin/latest/murex-netbsd-arm.gz)
* [murex-netbsd-arm64.gz](https://nojs.murex.rocks/bin/latest/murex-netbsd-arm64.gz)

#### OpenBSD

* [murex-openbsd-386.gz](https://nojs.murex.rocks/bin/latest/murex-openbsd-386.gz)
* [murex-openbsd-amd64.gz](https://nojs.murex.rocks/bin/latest/murex-openbsd-amd64.gz)
* [murex-openbsd-arm.gz](https://nojs.murex.rocks/bin/latest/murex-openbsd-arm.gz)
* [murex-openbsd-arm64.gz](https://nojs.murex.rocks/bin/latest/murex-openbsd-arm64.gz)

### Solaris

This build should be treated as experimental however unlike the other
experimental builds (Plan 9 and Windows), Solaris is at least POSIX compliant
so expect fewer issues than on the non-POSIX platforms.

* [murex-solaris-amd64.gz](https://nojs.murex.rocks/bin/latest/murex-solaris-amd64.gz) 

### Plan 9

Plan9 is untested. The code compiles and it is syscall compatible with Plan9
operating systems, however you may experience bugs using Murex on Plan9. If
you do encounter any issues then please raise them at:
[github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)

* [murex-plan9-386.gz](https://nojs.murex.rocks/bin/latest/murex-plan9-386.gz)
* [murex-plan9-amd64.gz](https://nojs.murex.rocks/bin/latest/murex-plan9-amd64.gz)
* [murex-plan9-arm.gz](https://nojs.murex.rocks/bin/latest/murex-plan9-arm.gz)

## Install Instructions

### Linux / UNIX / macOS Instructions

> macOS builds are listed as [darwin](https://en.wikipedia.org/wiki/Darwin_(operating_system))
> as per the name of Apple have given to their [underlying OS](https://en.wikipedia.org/wiki/MacOS#Architecture).

Download the appropriate `.gz` file from the list above, one that matches both
your OS and CPU architecture. Then extract it and make the resulting file
executable.

For example, in Bash, Zsh and similar shells, you can copy/paste the following
to run on any Linux or UNIX-like OS from sh/bash/zsh. 

```sh
MUREX_BUILD="murex-linux-amd64"
wget "https://nojs.murex.rocks/bin/latest/${MUREX_BUILD}.gz"
gunzip "${MUREX_BUILD}.gz"
chmod +x "$MUREX_BUILD"
```

Additionally you may wish to add Murex to `/etc/shells` if you want to expose
Murex as a optional login shell. If you do this, please ensure Murex has been
placed in a sensible location that all users can access. eg `/usr/local/bin`.

Most of these builds have received _some_ level of user acceptance testing with
Linux and macOS builds receiving the most attention since that's what we mostly
use ourselves.

### Windows Instructions

Click the Windows link that matches your CPU architecture. Unzip using your
preferred too then launch using your preferred console. Murex cannot be
started via double clicking the executable -- it requires a starting from
within an existing console session.

Please also note that Windows support is also considered experimental. In part
due to the lack of **coreutils** (as seen on Linux and UNIX) and in part due to
the different underpinning technologies behind consoles / terminal emulators.
If you do experience some wonky behavior then our recommendation is to run the
`linux-amd64` build for Linux on top of WSL. The instructions above will guide
you through installing on Linux, WSL install instructions can be found at the
following site: [docs.microsoft.com/en-us/windows/wsl/install-win10](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

## See Also

* [Compatibility Commitment](/compatibility.md):
  Hack confidence in our backwards compatibility 
* [Install](/INSTALL.md):
  Installation instructions
* [Supported Platforms](docs//supported-platforms.md):
  Operating systems and CPU architectures supported by Murex

<hr/>

This document was generated from [gen/root/DOWNLOAD_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/root/DOWNLOAD_doc.yaml).