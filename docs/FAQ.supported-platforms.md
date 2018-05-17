# Supported Operating Systems

The following is a list of platforms _murex_ has been tested on and the
level of support it had:

## Linux

**ArchLinux:**

This is my primary development platform so all features should work.

**Debian:**

Debian is automatically tested via Docker as part of the build tests.

**Ubuntu:**

_murex_ has been tested and works well.

**CentOS:**

CentOS is automatically tested via Docker as part of the build tests.

## OS X / Darwin

Darwin is included as part of the automatic build tests however there has been
no formal usability tests. This means _murex_ will compile and execute however I
cannot make any guarantees about if all the features will function.

If you do encounter any problems then please raise an issue on Github.

## FreeBSD

_murex_ has been tested inside a 10.3-RELEASE AMD64 jail

FreeBSD support is considered very good.

## OpenBSD

Tested on an earlier release of _murex_. Regression tests cannot be run because
of `timeout` dependency however OpenBSD is included as part of the automated
build tests.

OpenBSD support is expected to be good but, as always, please log an issue via
Github if you do encounter problems.

## NetBSD

NetBSD is part of the automated build tests so _murex_ will compile on NetBSD.
However no functional testing has been conducted on that particular platform.

## Windows

Windows is support and part of the automated build tests so will _murex_ will
compile for that platform. However I have not had access to a Windows machine
for some time so there maybe some new bugs introduced in recent versions. There
is also the caveat that without a broad range of command line utilities (eg GNU
coreutils) the usefulness of _murex_ is seriously diminished. There is some work
underway to replicate some of the basics of coreutils as _murex_ builtins but
that level of work is massive, thankless, and so obviously a low priority. Thus
my recommendation is to run _murex_ inside WSL (Windows Subsystem for Linux) on
Windows 10. However if native Windows really is your preference then _murex_
*should* function.

## Plan 9

Not currently supported. There are a few differences in the `syscall` package
which would lead me to believe that _murex_ will not even compile. That's not to
say that I wont ever support Plan 9 in the future however it's not a feature I'm
giving any priority to at present.

# Other CPU architectures

While there isn't any CPU specific code in _murex_ below is a breakdown
of state of support and testing on alternative architectures:

## 386

Untested but should function the same as AMD64. 64 bit data types are rarely
used so performance should be roughly equal. If you do still use 32 bit
platforms and notice any issues then please raise an issue ticket and I will
investigate.

## ARM

Only tested on Linux but full compatibility was present.

## PPC

Untested however the Go compiler does support various PPC architectures so
_murex_ might work.

## MIPS

Untested however the Go compiler does support various MIPS architectures so
_murex_ might work.

## SPARC

Unsupported.
