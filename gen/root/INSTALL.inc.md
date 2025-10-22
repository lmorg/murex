{{ $doc := doct "" "INSTALL" }}
{{ if env "DOCGEN_TARGET=vuepress" }}
icon: download
title: {{ quote $doc.Title }}

---
{{ else }}<h1>{{ $doc.Title }}</h1>{{ end }}

{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>

<div id="toc">

- [Installing From A Package Manager](#installing-from-a-package-manager)
- [Pre-Compiled Binaries (HTTPS download)](#pre-compiled-binaries-https-download)
- [Compiling From Source](#compiling-from-source)
  - [Prerequisites](#prerequisites)
  - [Compiling](#compiling)
    - [Bash: Linux, MacOS, UNIX](#bash-linux-macos-unix)
    - [Powershell: Windows](#powershell-windows)
    - [Makefile: Linux, MacOS, UNIX](#makefile-linux-macos-unix)
- [External Dependencies (Optional)](#external-dependencies-optional)

</div>

{{ end }}## Supported Platforms

Linux, BSD and macOS are fully supported, with other platforms considered
experimental and/or having known limitations.

Windows is a supported platform however Murex doesn't aim to replace coreutils.
So, depending on your required use case, you may need additional 3rd party
software to provide those utilities.

There is a more detailed breakdown of known compatibility issues in the
[{{ if env "DOCGEN_TARGET=" }}docs/{{ end }}supported platforms]({{ if env "DOCGEN_TARGET=" }}docs{{ end }}/supported-platforms.md) document.

## Installing From A Package Manager

> This is the recommended way to install Murex because you can then stay
> updated with future releases.

[![Packaging status](https://repology.org/badge/vertical-allrepos/murex.svg?exclude_unsupported=1)](https://repology.org/project/murex/versions)

{{ tmpl (file "gen/includes/install-package.inc.md") .Ptr }}

## Pre-Compiled Binaries (HTTPS download)

[![Version](version.svg)](DOWNLOAD.md)
[![Build Murex Downloads](https://github.com/lmorg/murex/actions/workflows/murex-downloads.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/murex-downloads.yaml)

If you wish to download a pre-compiled binary then head to the {{ if env "DOCGEN_TARGET=vuepress" }}[DOWNLOAD](download.md){{ else }}[DOWNLOAD](DOWNLOAD.md){{ end }}
page to select your platform.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<a href="download.html" alt="download murex"><img src="/download.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/></a>
<!-- markdownlint-restore -->
{{ end }}

## Compiling From Source

[![Test Clean Install](https://github.com/lmorg/murex/actions/workflows/clean-build.yaml/badge.svg)](https://github.com/lmorg/murex/actions/workflows/clean-build.yaml)

Murex is designed to be as easy to compile as possible. However if you do not
have any prior experience with compiling software from source then the
recommended approach for installing Murex is either via a package manager (if
your platform is already supported) or downloading one of our pre-compiled
binaries.

### Prerequisites

You will need `go` (Golang) compiler, and `git` installed.

> Go 1.21 or higher is required.

These should be easy to install on most operating systems however Windows is a
lot more tricky with regards to `gcc`. Please check with your operating systems
package manager first but see further reading below if you get stuck.

**Further Reading:**

* [How to install Go](https://golang.org/doc/install)
* [How to install git](https://github.com/git-guides/install-git)

### Compiling

Installation from source is as simple as the following one liner:
{{ if env "DOCGEN_TARGET=vuepress" }}
::: code-tabs#shell

@tab Bash
```bash
GOBIN="$(pwd)" go install -v github.com/lmorg/murex@latest
```

@tab Powershell
```powershell
$env:GOBIN="$(pwd)"; & go install -v github.com/lmorg/murex@latest
```

:::
{{ else }}
#### Bash: Linux, MacOS, UNIX

```bash
GOBIN="$(pwd)" go install -v github.com/lmorg/murex@latest
```

#### Powershell: Windows

```powershell
$env:GOBIN="$(pwd)"; & go install -v github.com/lmorg/murex@latest
```
{{ end }}
#### Makefile: Linux, MacOS, UNIX

We also provide a `Makefile`.

This is convenient if you already have the source downloaded or want to taylor
the install. For example reintroduce a deprecated builtin, add optional
builtins, or utilize a different backend library for networking and/or sqlite3.

For options available in the Makefile, run `make help`.

## External Dependencies (Optional)

Some of Murex's extended features will have additional external dependencies.

* `aspell`: This is used for spellchecking. Murex will automatically enable or
  disable spellchecking based on whether `aspell` can be found in your `$PATH`.
  ([GNU Aspell](http://aspell.net))

* `git`: This is used by Murex's package manager, `murex-package`.
  ([How to install git](https://github.com/git-guides/install-git))
