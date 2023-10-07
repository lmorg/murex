{{ if env "DOCGEN_TARGET=vuepress" }}
{{ if env "DOCGEN_TARGET=ignore-prefix" }}
### {{ end }}icon: arrow-down

---
{{ end }}<h1>Install Murex</h1>

{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>

<div id="toc">

- [Pre-Compiled Binaries (HTTPS download)](#pre-compiled-binaries-https-download)
- [Installing From A Package Manager](#installing-from-a-package-manager)
  - [ArchLinux](#archlinux)
  - [FreeBSD Ports](#freebsd-ports)
  - [Homebrew](#homebrew)
  - [MacPorts](#macports)
- [Compiling From Source](#compiling-from-source)
  - [Installation From Source Steps](#installation-from-source-steps)
- [External Dependencies (Optional)](#external-dependencies-optional)
- [Recommended Terminal Typeface](#recommended-terminal-typeface)

</div>

{{ end }}## Supported Platforms

Linux, BSD and macOS are fully supported, with other platforms considered
experimental and/or having known limitations.

Windows is a supported platform however Murex doesn't aim to replace coreutils.
So, depending on your required use case, you may need additional 3rd party
software to provide those utilities.

There is a more detailed breakdown of known compatibility issues in the
[{{ if env "DOCGEN_TARGET=" }}docs/{{ end }}supported platforms]({{ if env "DOCGEN_TARGET=" }}docs{{ end }}/supported-platforms.md) document.

## Pre-Compiled Binaries (HTTPS download)

[![Version](version.svg)](DOWNLOAD.md)
[![Build Murex Downloads](https://github.com/lmorg/murex/actions/workflows/murex-downloads.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/murex-downloads.yaml)

If you wish to download a pre-compiled binary then head to the [DOWNLOAD](DOWNLOAD.md)
page to select your platform.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<a href="DOWNLOAD.html" alt="download murex"><img src="/download.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/></a>
<!-- markdownlint-restore -->
{{ end }}

## Installing From A Package Manager

> This is the recommended way to install Murex because you can then stay
> updated with future releases.

[![Packaging status](https://repology.org/badge/vertical-allrepos/murex.svg)](https://repology.org/project/murex/versions)

### ArchLinux

From AUR: [https://aur.archlinux.org/packages/murex(https://aur.archlinux.org/packages/murex)

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
sudo port install murex
```

## Compiling From Source

[![Test Clean Install](https://github.com/lmorg/murex/actions/workflows/clean-build.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/clean-build.yaml)

**Prerequisites:**

You will need `go` (Golang) compiler, and `git` installed.

> Go 1.18 or higher is required.

These should be easy to install on most operating systems however Windows is a
lot more tricky with regards to `gcc`. Please check with your operating systems
package manager first but see further reading below if you get stuck.

**Further Reading:**

* [How to install Go](https://golang.org/doc/install)
* [How to install git](https://github.com/git-guides/install-git)

### Installation From Source Steps

> Compiling from source is not recommended unless you already have a reasonable
> understanding of compiling Go projects for your specific platform.

Installation from source is as simple as the following one liner:

```bash
GOBIN="$(pwd)" go install github.com/lmorg/murex@latest
```

However you can change the `GOBIN` value to point to any writable location you
wish.

## External Dependencies (Optional)

Some of Murex's extended features will have additional external dependencies.

* `aspell`: This is used for spellchecking. Murex will automatically enable or
  disable spellchecking based on whether `aspell` can be found in your `$PATH`.
  [http://aspell.net](http://aspell.net)

* `git`: This is used by Murex's package manager, `murex-package`.
  [How to install git](https://github.com/git-guides/install-git)

## Recommended Terminal Typeface

This is obviously just a subjective matter and everyone will have their own
personal preference. However if I was asked what my preference was then that
would be [Hasklig](https://github.com/i-tu/Hasklig). It's a clean typeface
based off Source Code Pro but with a few added ligatures - albeit subtle ones
designed to make Haskell more readable. Those ligatures also suite Murex
pretty well. So the overall experience is a clean and readable terminal.
