<h1>Supported Operating Systems</h1>

The following is a list of platforms Murex has been tested on and the
level of support it has:

<h2>Table of Contents</h2>
<div id="toc">

- [Linux](#linux)
- [macOS (Darwin)](#macos-darwin)
- [Windows](#windows)
- [FreeBSD](#freebsd)
- [OpenBSD](#openbsd)
- [NetBSD](#netbsd)
- [DragonflyBSD](#dragonflybsd)
- [Plan 9](#plan-9)
- [Other CPU architectures](#other-cpu-architectures)

</div>

## Linux

The shell has been extensively tested across a number of distributions. There
are no known distribution specific issues.

## macOS (Darwin)

All features work aside alt-hotkeys.

Both x86 (Intel) and ARM (Apple Silicon, eg M2 et al) architectures are
supported.

## Windows

Windows is supported there are a few known issue with the way how Windows
internals are built. These cannot be easily worked around:

* Windows doesn't decouple the terminal emulator and the shell Which means you
  cannot rely upon STDIN working as expected (eg some commands don't read input
  from STDIN but instead poll the terminal emulator directly)

* Windows sends parameters as a single string rather than an array of string.
  This is to retain backwards compatibility with DOS but it breaks the way how
  quotation marks and escaping works. Murex will compile an array of
  parameters based on the quotation strings (there are 3 different types of
  quotations in Murex), infixed variables, sub-shells, etc. These would not be
  honoured by any Windows commands because every Windows application then has
  to handle how the one long string of parameters is chopped up into different
  arguments; how quotation marks are handles, spaces, escaping, etc. This means
  there is no standard so one command might handle spaces correctly but another
  wouldn't.

* Job control (`bg`, `^z`, `fg`, etc) isn't supported because Windows doesn't
  have an equivalent of the SIGSTSP (etc) POSIX signals. 

* There is also the caveat that without a broad range of command line utilities
  (eg GNU coreutils) the usefulness of Murex is seriously diminished. You can
  mitigate this by installing [MSYS2](https://www.msys2.org/) or [Cygwin](https://cygwin.com/).

## FreeBSD

FreeBSD is officially supported and tested by the community.

## OpenBSD

FreeBSD is officially supported and tested by the community.

## NetBSD

FreeBSD is officially supported and tested by the community.

## DragonflyBSD

FreeBSD is officially supported and tested by the community.

## Plan 9

Plan 9 is included as part of the automated built tests however no functional
tests have been run.

If you do happen to run into any such bugs then I do welcome pull requests.

Feature wise, job control isn't supported in Plan 9 because Plan 9 doesn't
support all of the required signals. All other functions are expected to work.

## Other CPU architectures

Several CPU architectures are supported:

* 386   (x86 32bit)
* AMD64 (x86 64bit)
* ARMv7 (32bit)
* ARMv8 (64bit)

## See Also

* [Download](/DOWNLOAD.md):
  Murex download links
* [Install](/INSTALL.md):
  Installation instructions

<hr/>

This document was generated from [gen/root/supported-platforms_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/root/supported-platforms_doc.yaml).