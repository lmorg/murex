# Supported Operating Systems

The following is a list of platforms _murex_ has been tested on and the
level of support it had:

## Linux

This is one of the primary development platforms; all features should work. The
shell has been extensively tested across a number of distributions and there
are no known distribution specific issues.

## OS X / Darwin

This is one of the primary development platforms; all features work aside alt-
hotkeys.

## Windows

Windows is supported and part of the automated build tests so _murex_ will
compile for that platform however there have been very little in the way of
functional tests on recent versions.

There are also few known bugs / lack of support due to the way how Windows
internals are built. These cannot be easily worked around:

* Windows doesn't decouple the terminal emulator and the shell Which means you
  cannot rely upon STDIN working as expected (eg some commands don't read input
  from STDIN but instead poll the terminal emulator directly)

* Windows sends parameters as a single string rather than an array of string.
  This is to retain backwards compatibility with DOS but it breaks the way how
  quotation marks and escaping works. _murex_ will compile an array of
  parameters based on the quotation strings (there are 3 different types of
  quotations in _murex_), infixed variables, subshells, etc. These would not be
  honoured by any Windows commands because every Windows application then has
  to handle how the one long string of parameters is chopped up into different
  arguments; how quotation marks are handles, spaces, escaping, etc. This means
  there is no standard so one command might handle spaces correctly but another
  wouldn't.

* Job control (`bg`, `^z`, `fg`, etc) isn't supported because Windows doesn't
  have an equivalent of the SIGSTSP (etc) POSIX signal. 

* There is also the caveat that without a broad range of command line utilities
  (eg GNU coreutils) the usefulness of _murex_ is seriously diminished. There
  is some work underway to replicate some of the basics of coreutils as _murex_
  builtins but that level of work is massive, thankless, and targeting a niche
  audience; and so obviously a very low priority.
  
Taking these points into account, the recommendation is to run _murex_ inside a
POSIX compatability layer such as WSL (Windows Subsystem for Linux) on Windows
10 and 11, or Cygwin. However if native Windows is your preference then _murex_
*should* function.

## FreeBSD

An older version was extensively tested inside a 10.3-RELEASE AMD64 jail.

FreeBSD support is considered very good but, as always, please log an issue via
Github if you do encounter problems.

## OpenBSD

Tested on an earlier version of _murex_.

OpenBSD support is expected to be good but, as always, please log an issue via
Github if you do encounter problems.

## NetBSD

NetBSD is part of the automated build tests so _murex_ will compile on NetBSD.
However no functional testing has been conducted on that particular platform.

## DragonflyBSD

DragonflyBSD is part of the automated build tests so _murex_ will compile on
DragonflyBSD. However no functional testing has been conducted on that
particular platform.

## Plan 9

Plan 9 is included as part of the automated built tests however no functional
tests have been run.

If you do happen to run into any such bugs then I do welcome pull requests.

Feature wise, job control isn't supported in Plan 9 because Plan 9 doesn't
support all of the required signals. All other functions are expected to work.

## Other CPU architectures

_murex_ is developed on AMD64 and that is also the architecture which runs all
of the unit tests; however there is nothing CPU specific in _murex_'s source
and the CI pipeline does compile binaries for 386, AMD64, ARMv7 (32bit) and
ARMv8 (64bit) so the shell should be compatible for those architectures.

If you do happen to run into any issues then please report them on the Github
issue tracker.
